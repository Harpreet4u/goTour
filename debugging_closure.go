/*
   runtime.Caller(1)
       or
   log.setFlags(log.Llongfile)
*/
package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	fmt.Println("YO")
	where()

	log.SetFlags(log.Llongfile)
	w := log.Print

	w("")
	fmt.Println("YO")
	w("")
}
