package main

import "fmt"

type Point struct {
    X int
    Y int
}

func (p Point) DoubleValues() {
	p.X *= 2
	p.Y *= 2
}

func main() {
	p := Point{X: 2, Y: 4}
	p.DoubleValues()
	fmt.Println(p) // What do you expect here?
}
