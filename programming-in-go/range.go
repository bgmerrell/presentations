package main

import "fmt"

// START OMIT
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("%d: %d\n", i, v)
    }
}
// END OMIT
