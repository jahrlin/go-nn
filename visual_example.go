package main

import "./perceptron"
import "math/rand"
import "time"
import "net/http"
import "encoding/json"

func f(x float64) float64 {
	return 0.4*x + float64(1)
}

func random(min, max int) float64 {
	return float64(rand.Intn(max-min) + min)
}

var count = 0
var ptron = gonn.NewPerceptron(3)
var dots = []Dot{}
var lines = []Line{}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//n sets of training data
	trainers := make([]*gonn.Trainer, 2500)

	//create some random inputs, calculate the answer beforehand for supervised learning
	for i := 0; i < len(trainers); i++ {
		x := random(-400, 400)
		y := random(-100, 100)

		answer := float64(1)

		if float64(y) < f(float64(x)) {
			answer = -1
		}

		trainers[i] = gonn.NewTrainer(x, y, answer)
	}

	//everything is set up, start looping through traniners and keep sending input to our ptron
	count = len(trainers)
	for i := 0; i < count; i++ {
		ptron.Train(trainers[i].Inputs, trainers[i].Answer)
		guess := ptron.FeedForward(trainers[i].Inputs)
		fill := true
		if guess > 0 {
			fill = false
		}

		dots = append(dots, Dot{trainers[i].Inputs[0], trainers[i].Inputs[1], fill})

		y1 := (-ptron.Weights[2] - ptron.Weights[0]*-400) / ptron.Weights[1]
		y2 := (-ptron.Weights[2] - ptron.Weights[0]*400) / ptron.Weights[1]

		lines = append(lines, Line{-400, y1, 400, y2})
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/data/", data)

	http.ListenAndServe(":3000", nil)
}

type Dot struct {
	X    float64
	Y    float64
	Fill bool
}

type Line struct {
	X1 int
	Y1 float64
	X2 int
	Y2 float64
}

type Data struct {
	Dots  []Dot
	Lines []Line
}

func data(w http.ResponseWriter, r *http.Request) {
	payload := Data{dots, lines}

	js, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
