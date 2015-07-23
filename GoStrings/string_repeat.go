/*
   strings.Repeat(s string, count int) string
   Returns new string repeated count times.
*/
package main

import (
	f "fmt"
	strs "strings"
)

func main() {
	orgS := "Hi there!"
	var newS string

	newS = strs.Repeat(orgS, 5)
	f.Printf("Repeated string: %s", newS)
}
