package main

import (
	"fmt"
)

// START OMIT
type Entity struct {
	id string
	size int
}

func (e *Entity) SizeDescription() string {
	switch {
	case e.size < 100: return "Small"
	case e.size < 200: return "Medium"
	default: return "Large"
	}
}

type Cow struct {
	Entity
	isFemale bool
}

func main() {
	c := Cow{}
	c.size = 150
	fmt.Println(c.SizeDescription())
}
// END OMIT
