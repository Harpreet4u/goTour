/*
   strings.ToLower(s string) string
   strings.ToUpper  (s string) string

*/
package main

import (
	f "fmt"
	strs "strings"
)

func main() {
	orgS := "Hi there! Happy"
	var lower string
	var upper string

	lower = strs.ToLower(orgS)
	f.Printf("Lower Case String: %s\n", lower)

	upper = strs.ToUpper(orgS)
	f.Printf("Upper Case String: %s\n", upper)
}
