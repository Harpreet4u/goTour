/*
   2 ways:
   1. Aggregation (or composition): include a named field of the type of the wanted functionality
   2. Embedding: Embed (anonymously) the type of ewanted functionality
*/

// Aggregation.
package main

import (
	f "fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {

	c := new(Customer)
	c.Name = "Happy"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter:
	c = &Customer{"Happy", &Log{"1 - Yes we can!"}}
	// f.Println(c)
	c.Log().Add("2 - After me the world will be a better place!")
	//f.Println(c.log)
	f.Println(c.Log())
}
func (l *Log) Add(s string) {

	l.msg += "\n" + s
}

func (l *Log) String() string {

	return l.msg
}

func (c *Customer) Log() *Log {

	return c.log
}
