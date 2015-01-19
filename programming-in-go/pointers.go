package main

import "fmt"

// START OMIT
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
	fmt.Println(cp)
	fmt.Printf("%T\n", cp)
	fmt.Println(*cp)
	fmt.Printf("%T\n", *cp)
	fmt.Println(&cp)
	fmt.Printf("%T\n", &cp)
	fmt.Println((*cp).X) // Don't use
	fmt.Println(cp.X) // Use this instead
}
// END OMIT
