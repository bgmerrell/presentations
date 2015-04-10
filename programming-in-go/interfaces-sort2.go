package main

import (
	"fmt"
	"sort"
)

// START OMIT
type lenSortSlice []string

func (s lenSortSlice) Len() int {
	return len(s)
}

func (s lenSortSlice) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s lenSortSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	langs := lenSortSlice{"Carl", "Sally", "Bob", "Cy", "Elizabeth", "Randal"}
	sort.Sort(langs)
	for _, lang := range langs {
		fmt.Println(lang)
	}
}
// END OMIT
