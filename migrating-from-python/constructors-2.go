package main

import (
	"fmt"
)

// START OMIT
type P struct { X int; Y int }
func main() {
	p := P{X: 9}
	fmt.Printf("%#v\n", p)
}
// END OMIT
