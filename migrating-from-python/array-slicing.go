package main

import "fmt"

// START OMIT
func main() {
	letterArray := [4]string{"a", "b", "c", "d"}
	fmt.Printf("%T\n", letterArray)
	var s [1]string = letterArray[1:2]
	fmt.Println(s)
}
// END OMIT
