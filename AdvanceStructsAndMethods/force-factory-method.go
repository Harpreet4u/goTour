package main

import "github.com/harpreet4u/goTour/AdvanceStructsAndMethods/matrix"
import f "fmt"

func main() {
    // wrong := new(matrix.matrix) // will not compile (matrix is private)
    right := matrix.NewMatrix(2)
    f.Println(right)
}
