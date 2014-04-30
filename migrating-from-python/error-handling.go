package main

import (
	"fmt"
	"time"
)

// START OMIT
type FooError struct { When time.Time; What string }
func (e *FooError) Error() string {
	return fmt.Sprintf("at %s, %s", e.When, e.What)
}

type BarError struct{}
func (e *BarError) Error() string { return "BarError" }

func run() error {
	return &FooError{ time.Now(), "it didn't work", }
}

func main() {
	var err error
	err = run()
	if _, ok := err.(*BarError); ok { fmt.Println(err) }
	fmt.Println("...")
	if _, ok := err.(*FooError); ok { fmt.Println(err) }
}
// END OMIT
