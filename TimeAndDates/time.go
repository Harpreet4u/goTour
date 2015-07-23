/*
   Datatype: time.Time

   current time: time.Now()
   Day: time.Now().Day()/.Month()/Year()

   Type: Duration --> time elapsed between 2 instants as an int64 nanosecond count.
   Type: Location --> maps time instants to the Zone, UTC

   func (t Time) Format(layout string) string --> Formatting time

*/

package main

import (
	f "fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()

	f.Println(t)
	f.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	f.Println(t)
	f.Println(time.Now())

	// Duration type
	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(week)
	f.Println(week_from_now)

	// formatting...
	f.Println(time.Now().Format("02 Jan 2006 15:04"))
	f.Println(t.Format("20060102"))
}
