package main

import "time"
import "fmt"

func main() {
	start := time.Now()
	for i := 0; i < 1e9; i++ {
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Time elapsed: %s\n", delta)
}
