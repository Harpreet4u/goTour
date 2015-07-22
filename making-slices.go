package main

import "fmt"

func main() {
    a := make([]int, 5) // 0 initialized array with slice pointing to it.
    printSlice("a", a)
    
    b := make([]int, 0, 5) // len 0 and cap is 5
    printSlice("b", b)
    
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    // %v default value printing.
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
