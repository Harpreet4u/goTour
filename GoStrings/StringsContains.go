/*
    strings.Contains(s, substr string) bool
*/
package main

import (
    f "fmt"
    strs "strings"
)

func main() {
    str := "HakunaMatata"
    f.Printf("T/F? Does the string \"%s\" contains substring \"%s\"", str, "Matata")
    f.Printf(" %t\n", strs.Contains(str, "Matata"))
}
