/*
Interfaces
An interface type is defined by a set of methods.

A value of interface type can hold any value that implements those methods.

*/

package main

import (
    "fmt"
    "math"
)

type Abser interface{
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    
    a = f // interface can hold any value that implements its methods.
    
    fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
