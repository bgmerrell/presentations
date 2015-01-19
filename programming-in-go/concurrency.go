package main

import (
    "fmt"
    "math/rand"
    "time"
)

// START OMIT
func waitAndPrint(s string, c chan int) {
	d := time.Duration(rand.Intn(5000))
	time.Sleep(d * time.Millisecond)
	fmt.Printf("%s", s)
	c <- int(d)
}

func main() {
	rand.Seed(7)
	c1 := make(chan int)
	c2 := make(chan int)
	go waitAndPrint("over", c1)
	go waitAndPrint("hang", c2)
}
// END OMIT
