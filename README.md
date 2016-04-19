# go-nn

a perceptron implementation written in Go

includes a demo with a neural network and supervised learning that tries to guess what f(x) is

every point on the graph is a randomly generated training input

green line shows f(x), for reference, and blue line shows the current guess
 
## demo 
`$ go run visual_example.go`

`$ open localhost:3000`

![Browser demo]
(https://raw.githubusercontent.com/jahrlin/go-nn/master/visual_example.gif)

## TODO / issues
- [x] get rid of the casting and use same-size floats instead
- [ ] subscribe to google.visualization.events to avoid melting the client GPU
- [ ] should probably use requestAnimationFrame for that sweet 60 fps
- [ ] make it async with a websocket or something
