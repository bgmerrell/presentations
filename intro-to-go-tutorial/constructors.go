package main

import "fmt"

type Foo struct {
	i int
	s string
	m map[int]bool
}

func NewFoo() *Foo {
	return &Foo{i: 1, s: "bar", m: map[int]bool{1: false, 2: true, 3: true}}
}

func main() {
	f := NewFoo()
	fmt.Printf("%#v\n", f)
}
