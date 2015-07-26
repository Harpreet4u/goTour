/*
   Allows program to regain control of panicking goroutine,
   stopping the terminating sequence and resuming normal execution.

   recover is only useful inside a **deferred function**
   It retrives the error value passed through the call of panic.

   NOTE: panic unwinds stack until deferred recover() is found
           or program terminates.

   analogous to catch block
   log.Fatal --> calls os.Exit(1) after writing log message

*/
package main

import f "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		// Gets err returned from panic.
		if e := recover(); e != nil {
			f.Printf("Panicking %s\n", e)
		}
	}()
	badCall()
	f.Printf("After bad call\n")
}

func main() {
	f.Printf("Calling test\n")
	test()
	f.Printf("Test completed\n")
}
