package main

import "fmt"

// START OMIT
func main() {
	m := make(map[int]struct{})
	m[42] = struct{}{}
	m[24] = struct{}{}
	fmt.Println(m)
	v, ok := m[42]
	fmt.Printf("%v: %v\n", v, ok)
	v, ok = m[1234]
	fmt.Printf("%v: %v\n", v, ok)
}
// END OMIT
