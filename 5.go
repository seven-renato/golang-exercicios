package main

import "fmt"

func main() {
    str := "ANA E MARIANA GOSTAM DE BANANA"
    subStr := "ANA"
    cont := 0
    for i := 0; i < len(str); i++ {
        if string(str[i]) == string(subStr[0]) {
            teste := true
            for j := 1; j < len(subStr); j++ {
                if i+j >= len(str) || string(str[i+j]) != string(subStr[j]) {
                    teste = false
                    break
                }
            }
            if teste {
                cont++
            }
        }
    }
    fmt.Println("Number of substrings found:", cont)
}
