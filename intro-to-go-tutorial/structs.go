package main

import "fmt"

type Coordinate struct {
    X int
    Y int
    Z int
    name string
}

func main() {
    c1 := Coordinate{X: 1, Z: 2}
    c2 := Coordinate{1, 2, 3, "foo"}
    fmt.Println(c1)
    fmt.Printf("%#v\n", c1) // %#v, a Go-syntax representation of the value
    fmt.Printf("%T\n", c1) // %T, a Go-syntax representation of the type of the value
    fmt.Printf("X: %d\n", c1.X) // %d, a number (base 10)
    fmt.Printf("name: %s\n", c1.name) // %s, a string (or some bytes)
    fmt.Printf("name: %s\n", c2.name) // %s, a string (or some bytes)
}
