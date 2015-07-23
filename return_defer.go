package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1 // defer executes after return 1, so ans is 2.
}

func main() {
	fmt.Println(f())
}
