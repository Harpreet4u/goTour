/*
To store data in an anonymous field or get access to the data we use the name of the data type:
outer.int; a consequence is that we can only have one anonymous field of each data type in a
struct.

*/
package main

import f "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS // anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 4.3
	outer.int = 74
	// Can access embedded struct fields directly.
	outer.in1 = 3
	outer.in2 = 10

	f.Printf("outer.b is: %d\n", outer.b)
	f.Printf("outer.c is: %f\n", outer.c)
	f.Printf("outer.int is: %d\n", outer.int)
	f.Printf("outer.in1 is: %d\n", outer.in1)
	f.Printf("outer.in2 is: %d\n", outer.in2)

	// with struct literal/constructor
	outer2 := outerS{2, 3.4, 68, innerS{5, 2}}
	f.Println("outer2 is: ", outer2)
}
