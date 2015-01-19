package main

import "fmt"

// START OMIT
type Noiser interface{
	GetNoise() string
}

type Cow struct{}

type Fox struct{}

func (f Fox) GetNoise() string {
	return "Ring-ding-ding-ding-dingeringeding"
}

func ask(n Noiser) {
	fmt.Printf("What goes \"%s?\"\n", n.GetNoise())
}

func main() {
	ask(Cow{})
	ask(Fox{})
}
// END OMIT
