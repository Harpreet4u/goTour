package main

import "fmt"

func main() {
    var a []int
    printSlice("a", a)
    
    // func append(s []T, vs ...T) []T
    // ... unpacking values. e.g. []int{1,2,4} --> 1, 2, 4
    a = append(a, 0)
    printSlice("a", a)
    
    a = append(a, 1)
    printSlice("a", a)
    
    a = append(a, 2, 3, 5)
    printSlice("a", a)
}

func printSlice(s string, x []int) {
        fmt.Printf("string: %s, values: %v len: %d cap: %d\n", s, x, len(x), cap(x))
}
