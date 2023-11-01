package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	v2 := Vertex{X: 5}
	fmt.Println(v.X, v2.X, v2.Y)
}
