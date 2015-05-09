package main

import (
	"fmt"
	"time"
)

// START OMIT
type TimestampErr struct {
	When time.Time
	What string
}

func (t *TimestampErr) Error() string {
	return fmt.Sprintf("at %s, %s",
		t.When, t.What)
}

func run() error {
	return &TimestampErr{time.Now(), "it didn't work"}
}

func main() {
	err := run()
	switch err.(type) {
	case *TimestampErr:
		fmt.Println("Got TimestampErr")
	default:
		fmt.Println("Got non-TimestampErr")
	}
	fmt.Println(err)
}
// END OMIT
