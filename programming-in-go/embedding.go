package main

import (
	"fmt"
)

// START OMIT
type Food struct {
	Calories int
}

func (f *Food) String() string {
	return "I'm a generic food"
}

type Pizza struct {
	Food
	Sauce string
}

func main() {
	p := Pizza{}
	p.Sauce = "tomato"
	fmt.Printf("%s\n", p.String())
}
// END OMIT
