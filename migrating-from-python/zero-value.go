package main

import "fmt"

// START OMIT
type Foo struct {
	Bar int
	Baz string
	Qux *string
}

func main() {
	p := Foo{}
	fmt.Printf("%#v", p)
	p.Bar = nil
}
// END OMIT
