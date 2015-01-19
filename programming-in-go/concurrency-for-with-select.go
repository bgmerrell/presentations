package main

import (
    "fmt"
    "time"
)

func main() {
    var tick chan time.Time = time.Tick(1 * time.Second)
    boom := time.After(5 * time.Second)
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM!")
            return
        default:
            fmt.Println("    .")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
