package gonn

import (
	"testing"
)

func Test_perceptron(t *testing.T) {

	p1 := perceptron{inputs: []int{1, 4, 9, -2, 7, 58}, weights: []int{-2, -2, 1, 3, 7, -1}, bias: 1}

	expected1 := 0
	output1 := p1.feedForward()

	if output1 != expected1 {
		t.Errorf("\nResult	%v\nExpected %v", output1, expected1)
	}

	p2 := perceptron{inputs: []int{5, -2, -9, -52, 10, 10, 11, 32}, weights: []int{1, -2, 2, 4, -5, 1, 1, 2}, bias: -1}

	expected2 := 0
	output2 := p2.feedForward()

	if output2 != expected2 {
		t.Errorf("\nResult	%v\nExpected %v", output2, expected2)
	}
}
