/*
   strings.Fields(s string) []string // Splits on whitespace
   strings.Split(s, sep string,) []string // Splits on separator
*/

package main

import (
	f "fmt"
	strs "strings"
)

func main() {
	str := "This is  sample string to split"
	abc := "This-is-to-split"

	f.Printf("Splitted string: %s\n", strs.Fields(str))
	f.Printf("Splitted string: %s\n", strs.Split(abc, "-"))

}
