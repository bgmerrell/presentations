package main

import "fmt"

type Coordinate struct {
    X int
    Y int
    Z int
    name string
}

func getCoordPtr() *Coordinate {
	return &Coordinate{1, 2, 3, "foo"}
}

func main() {
	var cp *Coordinate = getCoordPtr()
	fmt.Println(*cp)
	fmt.Println(cp)
	fmt.Println(&cp)
}
