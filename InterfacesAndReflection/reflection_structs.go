package main
import(
    "fmt"
    "reflect"
)

type NotknownType struct {
    a, b string
}

func (n NotknownType) String() string {
    return n.a + "-" + n.b
}

var secret interface{} = NotknownType{"Python", "GO"}

func main() {
    value := reflect.ValueOf(secret)
    typ := reflect.TypeOf(secret)
    
    fmt.Println(typ) // main.NotknownType
    knd := value.Kind() // struct
    fmt.Println(knd)

    for i:=0;i<value.NumField();i++ {
        fmt.Printf("Field %d: %v\n", i, value.Field(i))
        // Error non-exported struct type.
        //value.Field(i).SetString("abc")
    }
    results := value.Method(0).Call(nil)
    fmt.Println(results)
}
