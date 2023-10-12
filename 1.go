package main
import "fmt"

func main() {
    cont := 0
    aux := 1
    x := 3 
    n := 4
    for cont < n {
        aux *= x
        cont += 1
    }
    fmt.Println(aux)
}
