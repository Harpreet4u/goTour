// read input from the console
package main
import f "fmt"

var (
    firstName, lastName, s string
    i int
    fi float32
    input = "56.33 / 2134 / Go"
    format = "%f / %d / %s"
)

func main() {
    f.Println("Enter full name:")
    f.Scanln(&firstName, &lastName) // scans till newline with space deliminter
// f.Scanf("%s %s", &firstName, &lastName)

    f.Printf("Hi %s %s!\n", firstName, lastName)
    f.Sscanf(input, format, &fi, &i, &s) // Reads from string "input"
    f.Println("From string we read: ", fi, i, s)
    
} 
