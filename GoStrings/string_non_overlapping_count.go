/*
   strings.Count(s, str string) int
   Counts non-overlapping occurrences in s of str.
*/
package main

import (
	f "fmt"
	strs "strings"
)

func main() {
	str := "Hello, how is it going, Hugo?"
	manyG := "gggggggggg"

	f.Printf("Number of H's in %s is: ", str)
	f.Printf("%d\n", strs.Count(str, "H"))

	f.Printf("Number of gg's in %s is: ", manyG)
	f.Printf("%d\n", strs.Count(manyG, "gg"))
}
