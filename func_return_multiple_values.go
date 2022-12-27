package main

import "fmt"

/**func nameFunction() (types 1, types 2, ..., types n) {
    ---block code---
    return value 1, value 2, ... value n
}*/

func getFulName() (string, string) {
    return "xiao", "ching"
}

func main() {
    firstName, _ := getFulName()
    fmt.Println(firstName)
    //fmt.Println(lastName)
    //fmt.Println(firstName, lastName)

    a, b := getFulName()
    fmt.Println(a, b)
}
