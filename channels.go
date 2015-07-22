/*
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

*/
package main

import "fmt"

func sum(a []int, c chan int) {
    sum := 0
    for _, val := range a {
        sum += val
    }
    c <- sum 
}

func main() {   
    a := []int{7, 2, 8, -9, 4, 0}
    
    c := make(chan int) // channel of type int
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c

    fmt.Println(x, y, x+y)
}
    
