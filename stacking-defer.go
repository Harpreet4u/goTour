package main

import "fmt"

func main() {
    fmt.Println("Stacking defers.")
    
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("start defers in LIFO")
}
