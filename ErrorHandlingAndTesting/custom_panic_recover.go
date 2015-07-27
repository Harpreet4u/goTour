/*
   Best practices for custom packages:
   1. always recover from panic in your package: no explicit panic
       should be allowes to cross package boundary
   2. return errors as error values to the callers of your package.
*/
package main

import (
	f "fmt"
	"github.com/harpreet4u/goTour/ErrorHandlingAndTesting/parse"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		f.Printf("Parsing %q\n", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			f.Println(err)
			continue
		}
		f.Println(nums)
	}
}
