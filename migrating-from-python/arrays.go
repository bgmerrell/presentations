package main

import "fmt"

// START OMIT
func main() {
	letterSlice := []string{"a", "b", "c", "d"}
	fmt.Printf("%T\n", letterSlice)
	letterArray := [4]string{"a", "b", "c", "d"}
	fmt.Printf("%T\n", letterArray)
}
// END OMIT
