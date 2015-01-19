package main

import "fmt"

var sum = 1

func loop1() {
    for ; sum < 1000; {    // "while" loop
        sum += sum
    }
    fmt.Println(sum)
}

func loop2() {
    for sum < 2000 {    // "while" loop, simplified
        sum += sum
    }
    fmt.Println(sum)
}

func main() {
    loop1()
    loop2()
}
