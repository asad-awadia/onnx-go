package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Softmax", "TestSoftmaxNegativeAxis", NewTestSoftmaxNegativeAxis)
}

/*
&ir.ModelProto{
    IrVersion:   6,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:11},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"x"},
                Output:    {"y"},
                Name:      "",
                OpType:    "Softmax",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "axis",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          2,
                        F:             0,
                        I:             -1,
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
        Name:              "test_softmax_negative_axis",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "x",
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
                Name: "y",
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestSoftmaxNegativeAxis version: 6.
func NewTestSoftmaxNegativeAxis() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Softmax",
		Title:  "TestSoftmaxNegativeAxis",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x75, 0xa, 0x25, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x53, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x2a, 0x14, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Softmax",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000216e00)(name:"axis" type:INT i:-1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, 0.9772779, 0.95008844, 0.1513572, 0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, 0.20515826, 0.3130677, 0.85409576, 2.5529897, 0.6536186, 0.8644362, 0.742165, 2.2697546, 1.4543657, 0.045758516, 0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, 0.88778573, 1.9807965, 0.34791216, 0.15634897, 1.2302907, 1.2023798, 0.3873268, 0.30230275, 1.048553, 1.420018, 1.7062702, 1.9507754, 0.5096522, 0.4380743, 1.2527953, 0.7774904, 1.6138978, 0.21274029, 0.89546657, 0.3869025, 0.51080513, 1.1806322, 0.028182229, 0.42833188, 0.06651722, 0.3024719, 0.6343221, 0.36274117}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{0.22564866, 0.057690013, 0.102890454, 0.36351478, 0.2502561, 0.29449275, 0.28659353, 0.12893826, 0.1228784, 0.16709709, 0.11251291, 0.41708824, 0.20852564, 0.110024095, 0.15184915, 0.12931651, 0.41267803, 0.113720864, 0.12667899, 0.21760562, 0.44418558, 0.066478014, 0.08207955, 0.07263289, 0.33462396, 0.27603087, 0.067485005, 0.07773696, 0.29854664, 0.28020057, 0.0850921, 0.10637273, 0.17707486, 0.52825713, 0.10320311, 0.10881158, 0.31847993, 0.30971378, 0.13708425, 0.12591052, 0.134613, 0.19516963, 0.25985438, 0.33183125, 0.07853177, 0.11491964, 0.25955123, 0.16136181, 0.37243277, 0.09173455, 0.24800895, 0.14914228, 0.16881499, 0.3298475, 0.10418628, 0.21081902, 0.14681678, 0.1858872, 0.259042, 0.19743498}),
			),
		},
	}
}
