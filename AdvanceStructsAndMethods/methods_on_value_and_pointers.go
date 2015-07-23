/*
   methods can be called on pointers as well as value.
*/
package main

import f "fmt"

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return f.Sprint(b) }

func main() {
	var a B // value
	a.change()
	f.Println(a.write())

	b := new(B) // pointer
	b.change()
	f.Println(b.write())
}
