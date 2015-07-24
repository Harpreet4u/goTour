package main

import f "fmt"

type myType int

func main() {
    var dataSlice = []myType{1,2,3}
    // compile error, as we can't copy slices of different types
    // empty interface can take any type but
    // memory layout is different for both
    // interface have 2*N (receiver+table_of_methods_ptrs) words 
    // memory layout w.r.t N word layout of myType slice

// http://code.google.com/p/go-wiki/wiki/InterfaceSlice
    //**//var interfaceSlice []interface{} = dataSlice
    
    var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
    for ix, d := range dataSlice {
        interfaceSlice[ix] = d
    }
    f.Println(interfaceSlice)
}
