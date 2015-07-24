// Reflection is also called metaprogramming.

// converts to empty interface and then derive type and value...
// func TypeOf(i interface{}) Type
// func ValueOf(i interface{}) Value

package main
import(
    f "fmt" 
    "reflect"
)

func main() {
    var x float64 = 3.4
    f.Println("type:", reflect.TypeOf(x))
    v := reflect.ValueOf(x)
    f.Println("value:", v)
    f.Println("type:", v.Type())
    f.Println("kind:", v.Kind())
    f.Println("value:", v.Float())
    f.Println(v.Interface())
    
    y := v.Interface().(float64)
    f.Println(y)
}
    
