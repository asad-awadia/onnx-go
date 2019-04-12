package graph

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	"github.com/chewxy/hm"
	"github.com/pkg/errors"
	"gorgonia.org/gorgonia/debugger"
	"gorgonia.org/gorgonia/distro"
	"gorgonia.org/gorgonia/internal/constructor"
	"gorgonia.org/gorgonia/internal/execution"
	"gorgonia.org/gorgonia/internal/value"
	"gorgonia.org/gorgonia/ops"
	"gorgonia.org/tensor"
)

// A Node is a node in the computation graph
type Node struct {
	// metadata of the node
	t     hm.Type // pruned types only plz
	shape tensor.Shape

	// this node is the result of applying the op to the children
	op       ops.Op
	children Nodes // shortcut, instead of having to go through the graph. It is set by the compiler

	// For nicely grouping stuff in graphviz.
	// TODO: Should this be in *Node?
	name  string
	group string

	g *ExprGraph // this node belongs in this graph

	// value bondage
	// inputs are bound to values directly
	boundTo value.Value
	dataOn  execution.Device // where is the data on

	// to track derivations
	derivOf Nodes
	deriv   *Node

	// for hashing nodes
	id   int64 // id is the ID at which the node is added to the graph
	hash uint32

	hashed        bool
	inferredShape bool // is shape inferred?
	unchanged     bool // has this node been modified
	isStmt        bool // is this a statement node
	ofInterest    bool // is this node of particular interest? (for debugging)
}

// SetName of the node
func (n *Node) SetName(name string) {
	n.name = name
}

// GetName of the node
func (n *Node) GetName() string {
	return n.name
}

// NodeConsOpt is a function that provides construction options for any Node.
type NodeConsOpt func(*Node)

// WithType is a node construction option to set a node to the specified type.
// Types in *Node are immutable once set. If the type has already been specified in the node,
// a check will be made to see if the both types are the same. If it isn't, it will panic.
func WithType(t hm.Type) NodeConsOpt {
	f := func(n *Node) {
		if n.t == nil {
			n.t = t
		} else if !n.t.Eq(t) {
			panic(fmt.Sprintf("Node's type is %v. Asking to construct a Node with %v", n.t, t))
		}
	}
	return f
}

// WithChildren sets the children of a node to the specified chidren.
// This construction option does NOT check if existing children exists, and will overwrite the existing children.
func WithChildren(children Nodes) NodeConsOpt {
	f := func(n *Node) {
		n.children = children
	}
	return f
}

// WithOp is a node construction option to set a node's Op to the specified Op.
// `Op`s in `*Node`s are immutable once set and cannot be changed. If the node already has an Op specified
// a check will be made to see if the provided Op and the one already specified in the `*Node` is the same -
// do note that comparison of Ops is done using the `Hashcode()` method of Ops, and hash collisions MAY occur -
// If both ops are different, this function will panic.
func WithOp(op Op) NodeConsOpt {
	f := func(n *Node) {
		if n.op != nil {
			if op.Hashcode() != n.op.Hashcode() {
				panic(fmt.Sprintf("Node Ops are immutable. Cannot set op %v", op))
			}
			return
		}
		n.op = op
		if _, ok := op.(stmtOp); ok {
			n.isStmt = true
		}
	}
	return f
}

// WithName is a node construction option that gives the *Node the provided name. This is especially useful in debugging graphs.
func WithName(name string) NodeConsOpt {
	f := func(n *Node) {
		n.name = name
	}
	return f
}

// WithValue is a node construction option that binds the value to the *Node. This function may panic if:
//	- Gorgonia was unable to convert interface{} into a value.Value.
//	- The type of the value.Value does not match the type of the nodes.
func WithValue(any interface{}) NodeConsOpt {
	v, t, _, err := value.AnyToValue(any)
	if err != nil {
		panic(err)
	}

	f := func(n *Node) {
		if n.t == nil {
			n.t = t
		} else if !n.t.Eq(t) {
			panic(fmt.Sprintf("TypeError: Want %v, Got %v instead (%T %T)", n.t, t, n.t, t)) // yes this is a runtime error
		}

		n.bind(v)
		if n.shape == nil {
			n.shape = v.Shape()
		}
	}
	return f
}

// WithGrad is a node construction option that binds the value to the *Node. This function may panic if:
//	- There isn't already a value associated with the node (.boundTo == nil)
//	- The type of the value.Value does not match the value of the node.
func WithGrad(any interface{}) NodeConsOpt {
	v, t, _, err := value.AnyToValue(any)
	if err != nil {
		panic(err)
	}
	f := func(n *Node) {
		if n.boundTo == nil {
			panic("No value already bound to node")
		}
		if !value.TypeOf(n.boundTo).Eq(t) {
			panic("Different types ")
		}

		if dv, ok := n.boundTo.(*value.DualValue); !ok {
			if err := n.bind(&value.DualValue{Value: n.boundTo, D: v}); err != nil {
				panic(err)
			}
		} else {
			dv.D = v
		}
	}
	return f
}

// WithInit is a node construction option to initialize a *Node with the InitWFn provided.
func WithInit(fn distro.InitWFn) NodeConsOpt {
	f := func(n *Node) {
		dt, err := dtypeOf(n.t)
		if err != nil {
			panic(err)
		}

		var v value.Value
		v = tensor.New(tensor.WithShape(n.shape...), tensor.WithBacking(fn(dt, n.shape...)))
		WithValue(v)(n)
	}
	return f
}

// WithShape is a node construction option to initialize a *Node with a particular shape.
// This function panics if the shape's dimensions do not match the specified dimensions of the *Node.
func WithShape(shp ...int) NodeConsOpt {
	s := tensor.Shape(tensor.BorrowInts(len(shp)))
	copy(s, shp)
	f := func(n *Node) {
		nd := n.Dims()
		// if nd == 1 && s.IsVector() {
		// 	goto safe
		// }
		isVec := s.IsColVec() || s.IsRowVec()
		acceptVec := (isVec && (nd == 1))
		sameDims := nd == s.Dims()
		acceptScalar := nd == 0 && scalarEquiv(s)

		if !acceptVec && !sameDims && !acceptScalar {
			panic(fmt.Sprintf("Node %v, has %d dimensions(Shape: %v). Input shape is %v, which has %d dimensions", n, n.Dims(), n.shape, s, s.Dims()))
		}
		// safe:
		n.shape = s
	}
	return f
}

// WithGroupName is a node construction option to group a *Node within a particular group. This option is useful for debugging with graphs.
func WithGroupName(name string) NodeConsOpt {
	f := func(n *Node) {
		if n.group == "" {
			n.group = name
		}
	}
	return f
}

/*
func newNode(opts ...NodeConsOpt) *Node {
	n := borrowNode()
	n.dataOn = CPU
	n.id = -1

	for _, opt := range opts {
		opt(n)
	}
	n.fix()

	incrNN()
	return n
}
*/

// ID returns the ID of the node. This satisfies the gonum/graph.Node interface
func (n *Node) ID() int64 { return n.id }

// helper functions to help compilation process
func (n *Node) isArg() bool      { return n.op == nil }
func (n *Node) isInput() bool    { return (n.isArg() || n.isRandom()) && !n.isStmt }
func (n *Node) isMutable() bool  { return !n.isInput() && n.op.ReturnsPtr() }
func (n *Node) isConstant() bool { _, ok := n.op.(constant); return ok }
func (n *Node) isRandom() bool   { _, ok := n.op.(randomOp); return ok }

func (n *Node) isRoot() bool {
	if n.g == nil {
		return true
	}
	return len(n.g.to[n]) == 0
}

// IsVar returns true if  the node represents a differentiable variable (i.e. it's an argument to the function that is not a statement)
func (n *Node) IsVar() bool { return n.isArg() && !n.isStmt && !n.isConstant() }

// type related isX() helper methods

// IsScalar indicates if a node represents a a scalar value. This is based on the type of the node, not the actual value associated with the node
func (n *Node) IsScalar() bool { _, ok := n.t.(tensor.Dtype); return ok }

// IsVector indicates if a node represents a vector value. This is based on the type of the node, not the actual value associated with the node
func (n *Node) IsVector() bool {
	if t, ok := n.t.(constructor.TensorType); ok {
		return t.Dims == 1
	}

	return false
}

// IsColVec indicates if a node represents a Column Vector. This is based on the type of the node, not the actual value associated with the node
func (n *Node) IsColVec() bool {
	if _, ok := n.t.(constructor.TensorType); ok {
		if n.shape != nil {
			return n.shape.IsColVec()
		}
	}
	return false
}

// IsRowVec indicates if a node represents a Row Vector. This is based on the type of the node, not the actual value associated with the node
func (n *Node) IsRowVec() bool {
	if _, ok := n.t.(constructor.TensorType); ok {
		if n.shape != nil {
			return n.shape.IsRowVec()
		}
	}
	return false
}

// IsMatrix indicates if a node represents a matrix. This is based on the type of the node, not the actual value associated with the node
func (n *Node) IsMatrix() bool {
	if _, ok := n.t.(constructor.TensorType); ok {
		return n.shape.Dims() == 2
	}
	return false
}

// methods

// Graph returns the graph of the node
func (n *Node) Graph() *ExprGraph { return n.g }

/*
// CloneTo clones the node into a new graph. If CloneTo() is called on the same graph as the n, it will return n. The reason this is done is because
// at any given time, every node  should be unique in the *ExprGraph.
//
//TODO: clone children as well (this means that CloneTo() is only currently suitable fo input nodes)
func (n *Node) CloneTo(g *ExprGraph) *Node {
	if n.g != nil && g == n.g {
		return n
	}
	n2 := n.Clone().(*Node)
	n2.g = g
	n2 = g.AddNode(n2)
	return n2
}
*/

/*
// Clone clones the node. There are some caveats:
//		- the graph is not copied over - the node essentially does not belong to a collection
//		- there is no ID
// 		- the children are not cloned
func (n *Node) Clone() (retVal interface{}) {
	n2 := newNode(In(n.g), WithOp(n.op), WithName(n.name), WithType(n.t))
	if n.shape != nil {
		n2.shape = n.shape.Clone()
		n2.inferredShape = n.inferredShape
	}

	if n.boundTo != nil {
		var err error
		if n2.boundTo, err = CloneValue(n.boundTo); err != nil {
			log.Printf("Unable to clone %v\n%T\n%v", n, n.boundTo, n.boundTo)
			panic(err)
		}
	}

	// reset
	n2.g = nil

	// other things
	n2.name = n.name
	n2.group = n.group
	n2.dataOn = n.dataOn
	n2.hash = n.hash

	n2.hashed = n.hashed
	n2.inferredShape = n.inferredShape
	n2.unchanged = n.unchanged
	n2.isStmt = n.isStmt
	n2.ofInterest = n.ofInterest
	return n2
}
*/

// Value returns the valuse bound to the node. May return nil
func (n *Node) Value() value.Value {
	if n.isConstant() {
		return n.op.(constant).Value()
	}
	if dv, ok := n.boundTo.(*value.DualValue); ok {
		return dv.Value
	}
	return n.boundTo
}

// Grad returns the gradient if there is one.
func (n *Node) Grad() (value.Value, error) {
	if dv, ok := n.boundTo.(*value.DualValue); ok {
		return dv.D, nil
	}
	if n.deriv != nil {
		return n.deriv.Value(), nil
	}

	return nil, errors.Errorf("No Gradient node/value found for %T", n)
}

// Dims indicates how many dimensions the node's result has
func (n *Node) Dims() int {
	if n.shape != nil {
		return n.shape.Dims()
	}
	switch nt := n.t.(type) {
	case constructor.TensorType:
		return nt.Dims
	case tensor.Dtype:
		return 0
	default:
		panic(fmt.Sprintf("Dims undefined for %v(%T)", nt, nt))
	}
}

// Type returns the type of the node
func (n *Node) Type() hm.Type { return n.t }

// Dtype returns the dtype of the node
func (n *Node) Dtype() tensor.Dtype {
	dt, err := dtypeOf(n.t)
	if err != nil {
		panic(err)
	}
	return dt
}

// Shape returns the shape of the node
func (n *Node) Shape() tensor.Shape { return n.shape.Clone() }

// Strides returns the strides of the value of the node
func (n *Node) Strides() []int {
	if n.boundTo != nil {
		switch v := n.boundTo.(type) {
		case *value.DualValue:
			return v.Value.(tensor.Tensor).Strides()
		case tensor.Tensor:
			return v.Strides()
		default:
			log.Printf("Unhandled type for Strides(): %T. Using fallback method and assuming dense tensor types", n.boundTo)
		}
	}
	return n.shape.CalcStrides()
}

// Device returns the device the data will be on
func (n *Node) Device() execution.Device { return n.dataOn }

// Op returns the Op of the node
func (n *Node) Op() Op { return n.op }

// IsVec returns whether this node is a vector
func (n *Node) IsVec() bool { return n.IsVector() }

// Name returns the name of the node. If a name was specified and it is too long,
// the short name will be used instead (except in inputs)
//
// The short name is typically of the form: OpName(%1, %2 ...), making it read more like a function call
func (n *Node) Name() string {
	if n.name != "" {
		return n.name
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s(", n.op)
	for i, child := range n.children {
		fmt.Fprintf(&buf, "%%%x", child.id)
		if i < len(n.children)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")
	return buf.String()
}

// WriteHash writes the hash to the provided Hash32.
func (n *Node) WriteHash(h hash.Hash32) {
	fmt.Fprintf(h, "%v%v", n.t, n.shape)

	if n.isInput() {
		h.Write([]byte(n.name))
	} else {

		n.op.WriteHash(h)
	}

	// if len(n.children) == 0 {
	// 	binary.Write(h, binary.LittleEndian, byte(0))
	// }

	binary.Write(h, binary.LittleEndian, byte(len(n.children)))
	for _, child := range n.children {
		binary.Write(h, binary.LittleEndian, child.Hashcode())
	}

}

// Hashcode provides the hash for the tree, assuming that the node is the root of the tree.
// Original implementation was here by Vatine (who's apparently 80 years old and using SO!?!):
//		http://stackoverflow.com/questions/1988665/hashing-a-tree-structure
func (n *Node) Hashcode() uint32 {
	if n.hashed {
		return n.hash
	}
	h := fnv.New32a()
	n.WriteHash(h)
	n.hash = h.Sum32()
	n.hashed = true
	return n.hash
}

// String() implements the fmt.Stringer interface
func (n *Node) String() string {
	var buf bytes.Buffer
	if n.Name() != "" {
		fmt.Fprintf(&buf, "%s :: ", n.Name())
	} else {
		fmt.Fprintf(&buf, "%s :: ", n.op)
	}
	if c, ok := n.op.(constant); ok {
		fmt.Fprintf(&buf, "%v{%v}", n.t, c.Value())
	} else {
		fmt.Fprintf(&buf, "%v", n.t)
	}
	return buf.String()
}

// private methods

// TODO: check type, check shape, check if needsGrad -> promote to value.DualValue
func (n *Node) bind(v value.Value) error {
	// pc, _, _, _ := runtime.Caller(1)
	// log.Printf("binding to %p. Called by %v", n, runtime.FuncForPC(pc).Name())

	if n.boundTo == nil {
		n.boundTo = v
		return nil
	}

	if dv, ok := n.boundTo.(*value.DualValue); ok {
		if vdv, ok := v.(*value.DualValue); ok {
			if vdv == dv {
				return nil
			}
			if n.isRandom() {
				// then simply replace the value in it
				dv.Value = vdv.Value
				return nil
			}
			// n.boundTo = vdv
			// return nil
			log.Printf("n %p", n)
			panic("Undefined behaviour") // no seriously there literally is no defined behaviour of what should the right thing be. I'll come back to this TODO.
		}
		dv.Value = v
		return nil
	}

	n.boundTo = v

	return nil
}

// bindCopy copies the value if to the bound value.
func (n *Node) bindCopy(v value.Value) (err error) {
	if n.boundTo == nil {
		var cloned value.Value
		if cloned, err = value.CloneValue(v); err != nil {
			return
		}
		n.boundTo = cloned
		return nil
	}

	var copied value.Value
	if dv, ok := n.boundTo.(*value.DualValue); ok {

		if vdv, ok := v.(*value.DualValue); ok {
			if vdv == dv {
				return nil // no need to copy!
			}

			if n.isRandom() {
				// returnValue(dv.Value)
				dv.Value = vdv.Value
				return nil
			}

			return errors.Errorf("Cannot yet handle bindCopy() of *value.DualValue into *value.DualValue") // TODO FIX
		}
		if copied, err = value.Copy(dv.Value, v); err != nil {
			return errors.Wrapf(err, "Failed to copy while binding to node with *value.DualValue")
		}
		dv.Value = copied // in case they're scalars
		return nil
	}
	if copied, err = value.Copy(n.boundTo, v); err != nil {
		return errors.Wrapf(err, "Failed to copy while binding to node")
	}
	n.boundTo = copied // in case it's a scalar
	return nil
}

// unbind releases the values back to the pool
func (n *Node) unbind() {
	if n.boundTo == nil {
		return
	}

	if dv, ok := n.boundTo.(*value.DualValue); ok {
		value.ReturnDV(dv)
	}

	if t, ok := n.boundTo.(tensor.Tensor); ok {
		returnTensor(t)
	}
	n.boundTo = nil
}

func (n *Node) dotCluster() string {
	var group string
	var isConst bool
	var isInput = n.isInput()

	if n.op != nil {
		_, isConst = n.op.(constant)
	}

	switch {
	case isConst:
		group = constantsClust
	case isInput:
		group = inputsClust
	case n.group == "":
		group = exprgraphClust
	default:
		group = n.group
	}
	return group
}

func (n *Node) fix() {
	if n.IsScalar() {
		n.shape = scalarShape
	}

	if n.isConstant() {
		return
	}

	if n.g == nil {
		panic(fmt.Sprintf("no graph supplied %v", n))
	}
}

/*
func (n *Node) fixEdges() {
	if n.g == nil {
		return
	}

	if len(n.children) > 0 {
		for _, child := range n.children {
			e := edge{from: n, to: child}
			n.g.SetEdge(e)
		}
	} else {
		n.g.leaves = append(n.g.leaves, n)
	}
}
*/

func (n *Node) setShape(s tensor.Shape, inferred bool) {
	n.shape = s
	n.inferredShape = inferred
}

func (n *Node) setGroup(grp string) {
	n.group = grp
}

/*
func (n *Node) clone(opts ...NodeConsOpt) *Node {
	if n.isInput() {
		return n
	}

	nn := newNode(WithChildren(n.children),
		WithType(n.t),
		WithOp(n.op),
		WithName(n.name),
		In(n.g),
	)

	for _, opt := range opts {
		opt(nn)
	}

	// if the shape is already known...
	if n.shape != nil {
		nn.shape = n.shape
		nn.inferredShape = n.inferredShape
	}

	return nn
}
*/

func (n *Node) diffWRT() []bool {
	if sdop, ok := n.op.(SDOp); ok {
		return sdop.DiffWRT(len(n.children))
	}
	return nil
}

// dfs but does not use channels. useful for extracting paths. used particularly in test
func (n *Node) seqWalk() Nodes {
	retVal := Nodes{n}
	for _, child := range n.children {
		retVal = append(retVal, child.seqWalk()...)
	}
	return retVal
}

// Group the node belongs to (useful for graphviz representation)
func (n *Node) Group() debugger.GroupID {
	var groupID debugger.GroupID
	var isConst bool
	var isInput = n.isInput()

	if n.op != nil {
		_, isConst = n.op.(constant)
	}
	switch {
	case isConst:
		groupID = debugger.ConstantCluster
	case isInput:
		groupID = debugger.InputCluster
	case n.group == "":
		groupID = debugger.ExprGraphCluster
	default:
		groupID = debugger.UndefinedCluster
	}
	return groupID
}
