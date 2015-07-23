/*
   strings.Join(sl []string, sep string) string
*/

package main

import (
	f "fmt"
	strs "strings"
)

func main() {
	str := "This is a string"
	sl := strs.Fields(str)

	f.Printf("Joined by : string %s\n", strs.Join(sl, ":"))
}
