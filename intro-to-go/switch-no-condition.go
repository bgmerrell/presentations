package main

import "fmt"

func main() {
	x := 0
	switch {
	case x < 0:
		fmt.Println("Negative")
	case x > 0:
		fmt.Println("Positive")
	default:
		fmt.Println("Zero")
	}
}
