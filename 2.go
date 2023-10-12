package main
import "fmt"

func primo(n int) bool {
    if n <= 3 && n >= 1{
        return true
    }
    if n % 2 == 0{ 
        return false
    }
    aux := 3
    for aux*aux <= n {
        if n%aux == 0 {
            return false
        }
        aux += 2
    }
    return true
}
func main() {
    fmt.Println(primo(24))
}