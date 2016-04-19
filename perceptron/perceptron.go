package gonn

import "math/rand"

type perceptron struct {
	Inputs  []float64
	Weights []float64
	bias    float64
	c       float64
}

func (p *perceptron) GetInputs() []float64 {
	return p.Inputs
}

func NewPerceptron(n int) *perceptron {
	p := new(perceptron)
	p.Weights = make([]float64, n)

	for i := 0; i < n; i++ {
		p.Weights[i] = (rand.Float64()*2 - 1) * 1
	}

	return p
}

func (p *perceptron) Train(inputs []float64, desired float64) {
	guess := p.FeedForward(inputs)
	error := desired - guess
	for i := 0; i < len(p.Weights); i++ {
		p.Weights[i] += float64(0.00001) * float64(error) * inputs[i]
	}
}

func (p *perceptron) FeedForward(inputs []float64) float64 {
	sum := float64(0)

	for i, value := range inputs {
		sum += value * p.Weights[i]
	}

	sum += p.bias

	return activate(sum)
}

func activate(sum float64) float64 {
	if sum <= 0 {
		return -1
	}

	return 1
}

type Trainer struct {
	Inputs []float64
	Answer float64
}

func NewTrainer(x, y, a float64) *Trainer {
	t := new(Trainer)
	t.Inputs = []float64{x, y, 1}
	t.Answer = a

	return t
}
