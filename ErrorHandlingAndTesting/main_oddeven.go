package main
import (
    f "fmt"
    "github.com/harpreet4u/goTour/ErrorHandlingAndTesting/even"
)

func main() {
    for i := 0; i<=100; i++ {
        f.Printf("Is the integer %d even? %v\n", i, even.Even(i))
    }
}
