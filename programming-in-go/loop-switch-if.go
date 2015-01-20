package main

import "fmt"

const MAX = 10

func main() {
    for i := 0; i < MAX; i++ {
        switch i {
        case 5:
            fmt.Println("Half way")
        case 9:
            fmt.Println("Almost there!")
        default:
            if i%2 == 0 {
                fmt.Println(i, "(even)")
            } else {
                fmt.Println(i, "(odd)")
            }
        }
    }
}
