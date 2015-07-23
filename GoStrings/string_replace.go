/*
   strings.Replace(str, old, new string, n int) string    // replaces n occurrences in str of old with new
    For all occurrences use n = -1

    Returns a copy of a string.
*/
package main

import (
	f "fmt"
	strs "strings"
)

func main() {
	str := "This is sample test. This is a sample test"

	f.Printf("Replacing all \"is\" with \"YO\"\n")
	f.Println(strs.Replace(str, "is", "YO", -1))

	f.Printf("Replacing 2 \"is\" with \"YO\"\n")
	f.Println(strs.Replace(str, "is", "YO", 2))
}
