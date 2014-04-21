package main

import "fmt"

func main() {
	i := 0
	for {
		i += 1
		break // Would be infinite loop if not for me!
	}
	fmt.Println(i)
}
