/*
    String Concatenation:
    
    In terms of efficiency: + < strings.Join() < byte-buffer
    
    for ASCII string looping use for-loop
    for Unicode string looping use for-range loop (without range loops over byte not rune)
*/
package main

import f "fmt"

func main() {

    s := "Hel" + "lo"
    f.Println(s)
    
    s += " World!"
    f.Println(s)

    str := "This is in multi"+
            " line yo"
    f.Println(str)
}
