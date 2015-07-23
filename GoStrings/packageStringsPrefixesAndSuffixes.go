/*
    strings.HasPrefix(s, prefix string) bool
    strings.HasSuffix(s, suffix string) bool
*/

package main

import (
    f "fmt"
    strs "strings"
)

func main() {
    var str string = "This is sample string"
    f.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")
    
    f.Printf(" %t\n", strs.HasPrefix(str, "Th"))
    f.Printf("T/F? Does the string \"%s\" have suffix %s?", str, "ring")
    
    f.Printf(" %t\n", strs.HasSuffix(str, "ring"))
} 
