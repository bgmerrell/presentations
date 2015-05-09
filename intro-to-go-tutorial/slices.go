package main

import "fmt"

// START OMIT
func main() {
	var letters []string
	letters = []string{"a", "b", "c", "d"}
	fmt.Println(letters)
	letters = append(letters, "e")
	more := []string{"f", "g", "h"}
	letters = append(letters, more...)  // append is a variadic function
	fmt.Println(letters)
}

// END OMIT
