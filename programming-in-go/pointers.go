package main

import "fmt"

// START OMIT
type Coordinate struct {
	X    int
	Y    int
	Z    int
	name string
}

func NewCoordinate() *Coordinate {
	return &Coordinate{1, 2, 3, "foo"}
}

func main() {
	var cp *Coordinate = NewCoordinate()
	fmt.Println(cp)
	fmt.Printf("%T\n", cp)
	fmt.Println(*cp)
	fmt.Printf("%T\n", *cp)
	fmt.Println(&cp)
	fmt.Printf("%T\n", &cp)
	// fmt.Println(cp->X) // Doesn't exist
	fmt.Println((*cp).X) // Don't use
	fmt.Println(cp.X)    // Use this instead
}

// END OMIT
