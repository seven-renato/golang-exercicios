package main
import "fmt"

func main() {
    var n int
    fmt.Print("Type a temperature: ")
    fmt.Scan(&n)
    aux:= ((n-32)*5)/9
    fmt.Println(aux)
}
