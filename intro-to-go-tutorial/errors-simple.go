package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// START OMIT
var ErrEmptyList error = errors.New("List is empty")

func getRandom(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, ErrEmptyList
	} else {
		return numbers[rand.Intn(len(numbers))], nil
	}
}

func main() {
	if n, err := getRandom([]int{}); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(n)
	}
}
// END OMIT
