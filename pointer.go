package main

import "fmt"

func main() {
    i, j := 42, 2701

    // NOTE: No pointer Arithmetic in Go like C...
    p := &i //pointing to i
    fmt.Println(*p)
    *p = 21
    fmt.Println(*p)
    
    p = &j
    *p = *p / 37
    fmt.Println(j)
}
