package main

import "fmt"
import "math"

type Coordinate struct {
    X float64
    Y float64
    Z float64
    name string
}

func (c Coordinate) getOutlier() string {
    max := math.Abs(c.X)
    o := "X"
    if max < math.Abs(c.Y) { max = c.Y; o = "Y" }
    if max < math.Abs(c.Z) { max = c.Z; o = "Z" }
    return o
}

func main() {
    c := Coordinate{5, 2, 3, "bogey"}
    fmt.Println(c.getOutlier())
}
