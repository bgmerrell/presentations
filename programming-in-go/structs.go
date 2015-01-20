package main

import "fmt"

type Coordinate struct {
    X int
    Y int
    Z int
    name string
}

func main() {
    c := Coordinate{X: 1, Z: 2}
    fmt.Printf("%v\n", c) // %v, the value in a default format
    fmt.Println(c)
    fmt.Printf("%#v\n", c) // %#v, a Go-syntax representation of the value
    fmt.Printf("%T\n", c) // %T, a Go-syntax representation of the type of the value
    fmt.Printf("X: %d\n", c.X) // %d, a number (base 10)
    fmt.Printf("name: %s\n", c.name) // %s, a string (or some bytes)
}
