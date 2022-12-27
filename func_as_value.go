package main

import "fmt"

func getHelloName(name string) string{
    return "hello " + name
}

func main() {
    sayHelloName := getHelloName
    fmt.Println(sayHelloName("xiao ching"))
}
