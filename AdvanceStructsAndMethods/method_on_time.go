/*
   Methods can be defined only on local-types within a package.
*/
package main

import (
	f "fmt"
	"time"
)

type myTime struct {
	time.Time // anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[:3]
}

func main() {
	m := myTime{time.Now()}
	f.Println("Full time now:", m.String())
	f.Println("First 3 chars:", m.first3Chars())
}
