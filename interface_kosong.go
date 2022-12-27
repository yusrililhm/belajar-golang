package main

import "fmt"

func ups(i int) interface{} {
    if i == 1 {
        return 1
    } else if i == 5 {
        return 5
    } else {
        return "upssss"
    }
}

func main() {
    data := ups(90)
    fmt.Println(data)
}
