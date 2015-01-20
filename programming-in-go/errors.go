package main

import (
	"fmt"
	"time"
)

// START OMIT
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %s, %s",
		e.When, e.What)
}

func run() error {
	return MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	err := run()
	switch err.(type) {
		case MyError: fmt.Printf("MyError: %s", err)
		default: fmt.Println("Other error: %s", err)
	}
}
// END OMIT
