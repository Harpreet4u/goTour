/*
    strconv.IntSize // int size on platform
    type T-> string will always succeed but not vice versa

    strconv.Itoa(i int) string
    strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string

    strconv.Atoi(s string) (int, error)
    strconv.parseFloat(s string, bitSize int) (float64, error)
    
*/
package main

import (
    f "fmt"
    sc "strconv"
)

func main() {
    var orig string = "666"
    var a int
    var str string
    
    f.Printf("The size of int is: %d\n", sc.IntSize)
    
    a, err := sc.Atoi(orig)
    if err == nil {
        f.Printf("%d\n", a)
    }

    a += 5
    str = sc.Itoa(a)
    f.Printf("%s\n", str)
}
