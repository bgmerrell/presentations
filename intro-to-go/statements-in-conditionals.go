package main

import "fmt"

func getX() int {
	return 0;
}

func main() {
	if x := getX(); x < 0 {  // Or, switch := getX(); x {
		fmt.Println("Negative")
	} else if x > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Zero")
	}
	// fmt.Println(x) here would fail to compile, why?
}
