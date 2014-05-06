package main

import (
	"bytes"
	"fmt"
	"os"
)

// START OMIT
func main() {
	// Named something like "New" by convention
	b1 := bytes.NewBufferString("Hello ") // HL
	fmt.Fprintf(b1, "world!\n")
	fmt.Println(b1)

	var b2 bytes.Buffer // no init needed
	b2.Write([]byte("Hello again "))
	fmt.Fprintf(&b2, "world!\n\n")
	b2.WriteTo(os.Stdout)

	b3 := bytes.Buffer{}
	b3.WriteString("Goodbye\n")
	b3.WriteTo(os.Stdout)
}
// END OMIT
