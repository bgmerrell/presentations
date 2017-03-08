package main

import "fmt"

// START OMIT
func main() {
	var i, j int = 1, 2
	k := 3.2
	a, b, c, d := true, false, "no!", '氣'
	var e int

	fmt.Println(i, j, k, a, b, c, d)
	fmt.Printf("%T, %T, %T, %T, %T, %T, %T\n", i, j, k, a, b, c, d)
	fmt.Println(e)
}

// END OMIT
