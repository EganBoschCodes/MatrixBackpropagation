package layers

import (
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

func (layer *SigmoidLayer) Back(inputs mat.Matrix, outputs mat.Matrix, forwardGradients mat.Matrix) (mat.Matrix, mat.Matrix) {
	forwardGradients.(*mat.Dense).Apply(func(i, j int, v float64) float64 {
		val := outputs.At(i, j)
		return v * val * (1 - val)
	}, forwardGradients)

	return nil, forwardGradients
}

func (layer *SigmoidLayer) GetShape() mat.Matrix {
	return nil
}

func (layer *SigmoidLayer) ApplyShift(shift mat.Matrix, scale float64) {}

func (layer *SigmoidLayer) NumOutputs() int {
	return layer.n_inputs
}
