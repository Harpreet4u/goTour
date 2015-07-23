package main

import (
	f "fmt"
	"github.com/harpreet4u/goTour/AdvanceStructsAndMethods/person"
)

func main() {
	p := new(person.Person)
	// p.firstName // error visibility private
	p.SetFirstName("Happy")
	f.Println(p.FirstName())
}
