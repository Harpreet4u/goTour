/*
    Strings --> immutable
    ASCII --> 1 byte (uint8)
    UTF-8 --> 2-4 bytes (int32) (rune)

    len terminated and variable width unlike c++, java (fixed width)
    less disk space

    NOTE: Taking the address of a character in a string, like &str[i], is illegal.
*/
package main

import (
    f "fmt" // Package aliasing
)

func main() {
    var str string // Default value is ""

    f.Println("Default value: ", str)
    
    // Interpreted strings. ""
    str = "This is \n in newline"
    f.Println(str)
    f.Println("")
    
    // Raw strings. `` 
    str = `This is \n in same line`
    f.Println(str)
   
    f.Println("")
    f.Printf("Integer: %d\n", str[0])
    f.Printf("Character: %c\n", str[0])
    f.Printf("UTF-8 bytes(int16 or int -- 2 bytes): %X\n", str[0])
    f.Printf("UTF-8 code points(int32 -- 4 bytes): %U\n", str[0])

    // ERROR: cannot take the address of str[0]
    // f.Printf("%v", &str[0])
}


