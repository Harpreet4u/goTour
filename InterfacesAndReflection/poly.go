package main

import f "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3} // Rectangle implements Shaper
	// compile error if q is Square value as pointer implements Shaper
	q := &Square{5} // Only Square pointer implements Shaper not Sqaure value

	shapes := []Shaper{r, q}
	for n, _ := range shapes {
		f.Println("Area of shape is: ", shapes[n].Area())
	}
}
