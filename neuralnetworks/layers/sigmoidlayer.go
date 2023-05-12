package layers

import (
	"fmt"
	"lossless/utils"
	"math"

	"gonum.org/v1/gonum/mat"
)

type SigmoidLayer struct {
	n_inputs int
}

func (layer *SigmoidLayer) Initialize(n_inputs int) {
	layer.n_inputs = n_inputs
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func (layer *SigmoidLayer) Pass(input mat.Matrix) mat.Matrix {
	input.(*mat.Dense).Apply(func(i int, j int, v float64) float64 { return sigmoid(v) }, input)
	return input
}

func (layer *SigmoidLayer) Back(inputs mat.Matrix, outputs mat.Matrix, forwardGradients mat.Matrix) (ShiftType, mat.Matrix) {
	outputSlice := utils.GetSlice(outputs)
	_, c := forwardGradients.Dims()
	forwardGradients.(*mat.Dense).Apply(func(i, j int, v float64) float64 {
		val := outputSlice[i*c+j]
		return v * val * (1 - val)
	}, forwardGradients)

	return &NilShift{}, forwardGradients
}

func (layer *SigmoidLayer) GetShape() mat.Matrix { return nil }

func (layer *SigmoidLayer) NumOutputs() int {
	return layer.n_inputs
}

func (layer *SigmoidLayer) ToBytes() []byte {
	return make([]byte, 0)
}

func (layer *SigmoidLayer) FromBytes(bytes []byte) {}

func (layer *SigmoidLayer) PrettyPrint() string {
	return fmt.Sprintln("Sigmoid Activation")
}
