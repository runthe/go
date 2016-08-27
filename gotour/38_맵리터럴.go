package main

import "fmt"

type Vertex struct {
	X, Y float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.40, -74.40,
	},
	"Google": Vertex{
		30.30, -30.30,
	},
}

func main() {
	fmt.Println(m)
}
