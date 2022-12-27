package main

import "fmt"

//func nameFunction(parameter types) {blok code}

func sayHello(name string) {
    fmt.Println("Hello " + name)
}

func penjumlahan(angka1 int, angka2 int) {
    fmt.Println(angka1 + angka2)
}

func main() {
    sayHello("xiao ching")
    penjumlahan(10, 90)
}
