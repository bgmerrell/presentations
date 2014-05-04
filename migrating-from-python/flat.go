package main

import (
    "fmt"
    "os"
    "log"
)

// START OMIT
type S struct {
    filename string
}

func (s S) printfile() {
    b := make([]byte, 1024)
    f, err := os.Open(s.filename)
    defer f.Close()
    if err != nil {
        log.Fatal("Failed to open file")
    }
    f.Read(b)
    fmt.Println(string(b))
}

func main() {
    s := S{filename: "/var/tmp/file.txt"}
    s.printfile()
}
// END OMIT
