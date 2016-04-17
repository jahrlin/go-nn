package gonn

import "math/rand"

type perceptron struct {
	Inputs  []int
	Weights []float32
	bias    int
	c       float32
}

func (p *perceptron) GetInputs() []int {
	return p.Inputs
}

func NewPerceptron(n int) *perceptron {
	p := new(perceptron)
	p.Weights = make([]float32, n)

	for i := 0; i < n; i++ {
		p.Weights[i] = (rand.Float32()*2 - 1) * 1
	}

	return p
}

func (p *perceptron) Train(inputs []int, desired int) {
	guess := p.FeedForward(inputs)
	error := desired - guess
	for i := 0; i < len(p.Weights); i++ {
		p.Weights[i] += float32(0.00001) * float32(error) * float32(inputs[i])
	}
}

func (p *perceptron) FeedForward(inputs []int) int {
	sum := float32(0)

	for i, value := range inputs {
		sum += float32(value) * p.Weights[i]
	}

	sum += float32(p.bias)

	return activate(sum)
}

func activate(sum float32) int {
	if sum <= 0 {
		return -1
	}

	return 1
}

type Trainer struct {
	Inputs []int
	Answer int
}

func NewTrainer(x, y, a int) *Trainer {
	t := new(Trainer)
	t.Inputs = []int{x, y, 1}
	t.Answer = a

	return t
}
