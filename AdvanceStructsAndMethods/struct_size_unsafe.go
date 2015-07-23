package main

import (
    f "fmt"
    us "unsafe"
)

type Empty struct{}
type Filled struct{
    a int
}

func main() {
    f.Println(us.Sizeof(Empty{}))
    f.Println(us.Sizeof(Filled{2}))
}
