/**********************************************************************
    (Calling with variabel itself)
    Methods on variables do not distinguish between values or pointers.

Interface methods (Calling with interface):
    Value-receiver methods can be called with values and pointers both
        (pointers can be dereferenced)
**  Pointer-receiver methods *cannot* be called with values, because
        the value stored inside an interface has no address.

*************************************************************************/
package main

import f "fmt"

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender
	// List does not implement Appender (Append method requires pointer)
	//CountInto(lst, 1, 10)

	if LongEnough(lst) {
		f.Printf("lst is long enough")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	// Valid: as pointer plst can be deferenced as *plst to value
	// as Lener is implemented on value.
	if LongEnough(plst) {
		f.Printf("lst is long enough")
	}
}
