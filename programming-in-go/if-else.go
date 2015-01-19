package main

import "fmt"

func getX() int {
	return 0
}

func main() {
	x := getX() // Statements like this can be combined with the conditional
	if x < 0 {
		fmt.Println("Negative")
	} else if x > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Zero")
	}
}
