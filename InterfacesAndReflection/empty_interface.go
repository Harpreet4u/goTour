/*
    The empty or minimal interface has no methods.
    type Any interface{}
*/
package main

import f "fmt"
var i = 5
var str = "ABC"

type Person struct {
    name string
    age int
}

type Any interface{}

func main() {
    var val Any

    val = 5
    f.Printf("val has value: %v\n", val)

    val = str
    f.Printf("val has value: %v\n", val)

    person := new(Person)
    person.name = "Happy"
    person.age = 55
    
    val = person
    f.Printf("val has value: %v\n", val)

    switch t := val.(type) {
    case int:
        f.Printf("Type int %T\n", t)
    case string:
        f.Printf("Type string %T\n", t)
    case bool:
        f.Printf("Type bool %T\n", t)
    case *Person:
        f.Printf("Type pointer to Person %T\n", t)
    default:
        f.Printf("Unexpected type %T\n", t)
    }
}
