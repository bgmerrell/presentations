package main

import "fmt"

type myUint8 uint8

//method
func (n myUint8) double() uint16 {
	return uint16(n) * 2
}

// function
func triple(n myUint8) uint16 {
	return uint16(n) * 3
}

func main() {
    var n myUint8 = 200;
    fmt.Println(n.double())
    fmt.Println(triple(200))
}
