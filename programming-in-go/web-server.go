package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":8088"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%#v\n", r)
}

func main() {
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.HandleFunc("/", Handle)
	http.ListenAndServe(port, nil)
}
