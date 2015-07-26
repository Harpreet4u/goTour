package main

import (
    f "fmt"
    errs "errors"
    m "math"
)

func main() {
    if val, err := Sqrt(-2); err != nil {
        f.Printf("Error: %s\n", err)
    } else {
        f.Printf("Sqrt : %g\n", val)
    }
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errs.New("math - Square root of negative number")
    }
    return m.Sqrt(f), nil
}
