package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ReduceLogSum", "TestReduceLogSum", NewTestReduceLogSum)
}

/*
&ir.ModelProto{
    IrVersion:   3,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:6},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"data"},
                Output:    {"reduced"},
                Name:      "",
                OpType:    "ReduceLogSum",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "axes",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          7,
                        F:             0,
                        I:             0,
                        S:             nil,
                        T:             (*ir.TensorProto)(nil),
                        G:             (*ir.GraphProto)(nil),
                        SparseTensor:  (*ir.SparseTensorProto)(nil),
                        Floats:        nil,
                        Ints:          {2, 1},
                        Strings:       nil,
                        Tensors:       nil,
                        Graphs:        nil,
                        SparseTensors: nil,
                    },
                    &ir.AttributeProto{
                        Name:          "keepdims",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             0,
                        S:             nil,
                        T:             (*ir.TensorProto)(nil),
                        G:             (*ir.GraphProto)(nil),
                        SparseTensor:  (*ir.SparseTensorProto)(nil),
                        Floats:        nil,
                        Ints:          nil,
                        Strings:       nil,
                        Tensors:       nil,
                        Graphs:        nil,
                        SparseTensors: nil,
                    },
                },
                DocString: "",
            },
        },
        Name:              "test_reduce_log_sum",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "data",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        Output: {
            &ir.ValueInfoProto{
                Name: "reduced",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestReduceLogSum version: 3.
func NewTestReduceLogSum() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ReduceLogSum",
		Title:  "TestReduceLogSum",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x87, 0x1, 0xa, 0x3d, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x22, 0xc, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x75, 0x6d, 0x2a, 0xd, 0xa, 0x4, 0x61, 0x78, 0x65, 0x73, 0x40, 0x2, 0x40, 0x1, 0xa0, 0x1, 0x7, 0x2a, 0xf, 0xa, 0x8, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x18, 0x0, 0xa0, 0x1, 0x2, 0x12, 0x13, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x75, 0x6d, 0x5a, 0x1a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x15, 0xa, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x6},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data"},
		     Output:    []string{"reduced"},
		     Name:      "",
		     OpType:    "ReduceLogSum",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc0002dc000)(name:"axes" type:INTS ints:2 ints:1 ),
		    (*ir.AttributeProto)(0xc0002dc100)(name:"keepdims" type:INT )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292, 0.11827443, 0.639921, 0.14335328, 0.9446689, 0.5218483, 0.41466194, 0.2645556, 0.7742337, 0.45615032, 0.56843394, 0.0187898, 0.6176355, 0.6120957, 0.616934, 0.94374806, 0.6818203, 0.3595079, 0.43703195, 0.6976312, 0.06022547, 0.6667667, 0.67063785, 0.21038257, 0.12892629, 0.31542835, 0.36371076, 0.57019675, 0.43860152, 0.9883738, 0.10204481, 0.20887676, 0.16130951, 0.6531083, 0.2532916, 0.46631077, 0.2444256}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{2.4536822, 2.4298263, 2.07904}),
			),
		},
	}
}
