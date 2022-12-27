package main

import "fmt"

//function type declaration
type filter func(string) string

func sayHelloFilter(name string, filter filter)  {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string  {
	if name == "babi" {
		return "xxx"
	} else {
		return name
	}
}



func main() {
	sayHelloFilter("babi", spamFilter)
	
	filter := spamFilter
	sayHelloFilter("Xiao Ching", filter)
}
