package main

import (
    "fmt"
    "math/rand"
    "time"
)

// START OMIT
func waitAndPrint(s string, c chan int) {
	d := time.Duration(rand.Intn(1000))
	time.Sleep(d * time.Millisecond)
	fmt.Printf("%s", s)
	c <- int(d)
}

func main() {
	rand.Seed(5)
	c1 := make(chan int)
	c2 := make(chan int)
	go waitAndPrint("gun", c1)
	go waitAndPrint("shot", c2)
}
// END OMIT
