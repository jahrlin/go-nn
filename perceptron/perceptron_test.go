package gonn

import (
	"github.com/nsf/termbox-go"
	"math/rand"
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

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	draw()
}

func draw() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault,
				termbox.Attribute(rand.Int()%8)+1)
		}
	}
	termbox.Flush()
}
