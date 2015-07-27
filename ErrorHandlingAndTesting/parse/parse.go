package parse

import (
	f "fmt"
	sc "strconv"
	strs "strings"
)

type ParseError struct {
	Index int // index into space-separated list of words
	Word  string
	err   error
}

func (e *ParseError) String() string {
	return f.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

func Parse(input string) (numbers []int, err error) {
	defer func() {
		// Recover receives error from panic.
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = f.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strs.Fields(input)
	numbers = fields2numbers(fields) // panic here
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := sc.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
