package main

import (
	"fmt"
	"time"
)

func repeat(s string, n, frequency int) {
    for i := 0; i < n; i++ {
        time.Sleep(time.Duration(frequency) * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go repeat("fast", 10, 250)
}
