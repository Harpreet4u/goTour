package main

type Foo map[string]string
type Bar struct{
    a string
    b int
}

func main() {
    // OK:
    y := new(Bar)
    (*y).a = "Hello"
    y.b = 3
    
    // not OK:
    //z := make(Bar) // compile error
    //z.a = "Hello"
    //z.b = 3

    // OK:
    x := make(Foo)
    x["x"] = "Hola"
    x["y"] = "world"
    // not OK:
    u := new(Foo) // NOTE: compile with nil pointer / Bad thing.
    
    (*u)["x"] = "hola" // !! panic !!: runtime error: assignment to entry in nil map as no allocated memory
}
