/*
    Marshaling or Encoding
    Unmarshaling or decoding

    Marshaling is used to convert in memory data to the special format (data -> string), and vice versa for unmarshaling (string -> data structure).

    Encoding does the same thing but the output is a stream of data (implementing io.Writer); decoding starts from a stream of data (implementing io.Reader) and populates a data structure.

*/
package main

import (
    "fmt"
    "encoding/json"
    "log"
    "os"
)

type Address struct {
    Type string
    City string
    Country string
}

type VCard struct {
    FirstName string
    LastName string
    Addresses []*Address
    Remark string
}

func main() {   
    pa := &Address{"private", "sdfsf", "Belgium"}
    wa := &Address{"work", "sdfs", "Belgium"}
    vc := VCard{"Happy", "Singh", []*Address{pa, wa}, "none"}

    fmt.Println(vc)
    
    // converts to string
    js, _ := json.Marshal(vc)
    fmt.Printf("JSON format: %s\n", js)

// encoder
    file, _ := os.OpenFile("vcard.json", os.O_CREATE | os.O_WRONLY, 0)
    defer file.Close()

    // converts to stream Writer
    enc := json.NewEncoder(file)
    err := enc.Encode(vc)
    if err != nil {
        log.Println("Error in encoding...")
    }
}
    
