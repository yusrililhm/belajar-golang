package main

import "fmt"

func getFullName() (string, string) {
    return "xiao", "ching"
}

func main() {
    firstName, _ := getFullName()
    fmt.Println(firstName)
    //fmt.Println(lastName)
    //fmt.Println(firstName, lastName)
}
