package main
import "fmt"

func main() {
    var n int
    fmt.Print("Type a number: ")
    fmt.Scan(&n)
    aux := 1
    cont := 1
    for cont <= n {
        fmt.Println(aux)
        aux += 2
        cont += 1
    }
}
