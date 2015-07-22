// You can declare a method on any type that is declared in your package, not just struct types.

package main

import (
    "fmt"
    "math"
)

type MyFloat float64 // check above comment

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
