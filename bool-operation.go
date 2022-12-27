package main

import "fmt"

func main() {
    // &&, ||, !

    a := true
    b := false
    c := true
    d := false

    fmt.Println(a && b)
    fmt.Println(a && c)

    fmt.Println(a || b)
    fmt.Println(b || d)

    fmt.Println(!b)
}
