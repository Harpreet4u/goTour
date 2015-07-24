/*
   An interface can contain a value of any type.
   To detect dynamic type of interface use:
       v, ok := varI.(T)
       // varI must be interface type otherwise compile-error
       v is converted value of varI to type T
       ok true if can be  converted else false
*/
package main

import (
	f "fmt"
	m "math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var shaperI Shaper
	sq1 := new(Square)
	sq1.side = 5

	shaperI = sq1
	// Is Square the type of shaperI ?
	// Removing * from *Square will give compile-error as Square
	// doesn't implements Shaper interface
	if t, ok := shaperI.(*Square); ok {
		f.Printf("The type of shaperI is %T\n", t)
	}

	if u, ok := shaperI.(*Circle); ok {
		f.Printf("The type of shaperI is: %T\n", u)
	} else {
		f.Printf("shaperI does not contain a variable of type Circle\n")
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * m.Pi
}
