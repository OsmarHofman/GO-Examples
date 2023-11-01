package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["vertice"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
	fmt.Println(m["vertice"])

	anotherVertexMap := map[string]Vertex{
		"teste": {-57.32, 65.743},
	}

	fmt.Println((anotherVertexMap))
}
