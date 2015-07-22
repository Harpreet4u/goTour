package main

import "fmt"

// Split function with named return.
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    // No need to return vars. It's auto return.
    return 
}

func main() {
    fmt.Println(split(17))
}
