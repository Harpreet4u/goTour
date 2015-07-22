package main

import "fmt"

func main() {
    s := []int{1,2,3,4,5,6}
    fmt.Println("s == ", s)
    fmt.Println("s[1:4] == ", s[1:4])

    fmt.Println("s[:3] == ", s[:3])
    
    fmt.Println("s[4:] == ", s[4:])
}
    
    
    
