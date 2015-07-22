package main

import "fmt"

func main() {
    // More efficient than array, array is unlike c
    // array doesn't points to 1st element.
    // whole array is copied while passing in func args
    s := []int{1, 2, 3}

    fmt.Println("s==", s)

    for i := 0; i < len(s); i++ {
        fmt.Printf("s[%d] == %d\n", i, s[i])
    }
}   
