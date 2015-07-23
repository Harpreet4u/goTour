package main

import (
    f "fmt"
}

// Example of structs with TAGS.

type TagType struct{ // tags
    field1 bool "An important answer"
    field2 string "The name of the thing"
    field3 int "how much there are"
}
