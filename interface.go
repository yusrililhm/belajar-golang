package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHi(hashName HasName)  {
	fmt.Println("Hello", hashName.GetName())
}

func (person Person) GetName()string {
	return person.Name
}

func main() {
	nama := Person{Name: "Xiao"}
	SayHi(nama)
}