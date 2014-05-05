package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := url.Parse("http://foo.org/index%html")
	fmt.Println(u.Host)
}
