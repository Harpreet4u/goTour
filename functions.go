package main

import "fmt"  // Formatting package used for printing.

// or func add(x, y int) int 
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
