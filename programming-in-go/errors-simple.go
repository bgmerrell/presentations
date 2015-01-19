package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// START OMIT
func getRandom(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("Error: list is empty")
	} else {
		return numbers[rand.Intn(len(numbers))], nil
	}
}

func main() {
	if n, err := getRandom([]int{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
// END OMIT
