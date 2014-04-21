package main

import fmt

func loop1() {
        sum := 0
        for ; sum < MAX; {
                sum += i
        }
        fmt.Println(sum)
}

func loop2() {
        sum := 0
        for sum < MAX {
                sum += i
        }
        fmt.Println(sum)
}

func main() {
	loop1()
	loop2()
}
