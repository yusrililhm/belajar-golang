package main

import "fmt"


//func nameFunction(param types) types {code program}
func sayHelloTo(name string)string {
    if name == "" {
        return "Hello World"
    } else {
        return "Hello " + name
    }

}

func sayHelloWorld(names string) string {
    return "Hello World from " + names
}

func main(){
    hello := sayHelloTo("xiao ching")
    fmt.Println(hello)

    fmt.Println(sayHelloTo(""))

    fmt.Println(sayHelloWorld("popoci"))
}
