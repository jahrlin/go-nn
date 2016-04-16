package gonn

type perceptron struct {
	inputs  []int
	weights []int
	bias    int
}

func (p *perceptron) feedForward() int {
	sum := 0

	for i, value := range p.inputs {
		sum += value * p.weights[i]
	}

	sum += p.bias

	return activate(sum)
}

func activate(sum int) int {
	if sum <= 0 {
		return 0
	}

	return 1
}
