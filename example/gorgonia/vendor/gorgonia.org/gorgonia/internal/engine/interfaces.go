package engine

import (
	"hash"

	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/node"
	"gorgonia.org/gorgonia/ops"
)

type hashWriter interface {
	WriteHash(hash.Hash)
}

type arityer interface {
	Arity() int
}

// START_OPERATION OMIT

// Operation is any op that can be applied to a node n of the graph g
//type Operation func(g graph.WeightedDirected, n node.Node) (ops.Op, error)
type Operation func(g graph.WeightedDirected, n node.Node) (ops.Op, error)

// END_OPERATION OMIT
