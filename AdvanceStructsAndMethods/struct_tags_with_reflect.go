package main

import (
	f "fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	t := TagType{true, "Yo YO", 1}

	for i := 0; i < 3; i++ {
		refTag(t, i)
	}
}

func refTag(t TagType, i int) {
	tType := reflect.TypeOf(t)
	iField := tType.Field(i)
	f.Println(iField.Tag)
}
