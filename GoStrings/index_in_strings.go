/*
   strings.Index(s, str string) int    // return first occurrence
   strings.LastIndex(s, str string) int // returns last occurrence

   For non-ASCII
   strings.IndexRune(s string, ch int) int
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Happy, Hi."
	fmt.Printf("The position of \"Happy\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Happy"))
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))
	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
}
