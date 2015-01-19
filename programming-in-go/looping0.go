package main

import "fmt"

const MAX = 10

func main() {
	sum := 0
	// Traditional "for" loop
	for i := 0; i < MAX; i++ {
		sum += i
	}
	fmt.Println(sum)
}
