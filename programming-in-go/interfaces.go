package main

import "fmt"

// START OMIT
type Talker interface{
	Talk() string
}

type Cow struct{}

type Fox struct{}

func (f Fox) Talk() string {
	return "Ring-ding-ding-ding-dingeringeding"
}

func ask(n Talker) {
	fmt.Printf("What goes \"%s?\"\n", n.Talk())
}

func main() {
	ask(Cow{})
	ask(Fox{})
}
// END OMIT
