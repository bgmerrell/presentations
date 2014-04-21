package main

import "fmt"

func funcCaller(f func(int) int) {
    fmt.Println(f(10))
}

func main() {
    // Function as a variable
    var doubleFunc func(int) int = func(n int) int { return n * 2 }
    fmt.Println(doubleFunc(10))
    // Calling a function literal
    fmt.Println(func(n int) int { return n * 3}(10))
    funcCaller(func(n int) int { return n * 4 })
}
