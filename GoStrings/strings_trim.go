/*
   strings.TrimSpace(s string) string

   General Trim:
   strings.Trim(s, chars string) string
   strings.TrimLeft(s, chars string) string
   strings.TrimRight(s, chars string) string
*/

package main

import (
	f "fmt"
	strs "strings"
)

func main() {
	str := " This is a string "

	f.Printf("Trimmed spaces in string: %s\n", strs.TrimSpace(str))
	f.Printf("Trimmed string: %s\n", strs.Trim(str, " Tg"))

}
