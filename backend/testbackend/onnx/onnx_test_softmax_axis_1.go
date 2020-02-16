package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Softmax", "TestSoftmaxAxis1", NewTestSoftmaxAxis1)
}

/*
&ir.ModelProto{
    IrVersion:   3,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:9},
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
                        I:             1,
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
        Name:              "test_softmax_axis_1",
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

// NewTestSoftmaxAxis1 version: 3.
func NewTestSoftmaxAxis1() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Softmax",
		Title:  "TestSoftmaxAxis1",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x65, 0xa, 0x1c, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x53, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x13, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x31, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Softmax",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000216e00)(name:"axis" type:INT i:1 )
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
				tensor.WithBacking([]float32{0.10431331, 0.026669053, 0.047564402, 0.16804633, 0.11568889, 0.047495, 0.04622104, 0.020794816, 0.0198175, 0.02694897, 0.020643288, 0.0765252, 0.038259212, 0.020186651, 0.027860498, 0.02495367, 0.079632774, 0.021944242, 0.02444472, 0.04199045, 0.18644135, 0.027903317, 0.03445187, 0.030486748, 0.14045423, 0.06214639, 0.015193771, 0.017501924, 0.06721566, 0.063085176, 0.016946724, 0.021184921, 0.035265766, 0.10520633, 0.020553662, 0.01697049, 0.049670823, 0.048303634, 0.021379953, 0.019637281, 0.0550496, 0.07981405, 0.106266685, 0.13570142, 0.032115337, 0.029896934, 0.067523584, 0.041979104, 0.09689029, 0.023865214, 0.047235616, 0.028405538, 0.03215239, 0.062822536, 0.019843249, 0.029607078, 0.020618707, 0.026105694, 0.03637944, 0.02772745}),
			),
		},
	}
}
