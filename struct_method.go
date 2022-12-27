package main

import "fmt"

type CustomerPriority struct {
    Name, Address, Hobby string
    Age int
}

func (customer CustomerPriority) sayIntroduce(name string)  {
	fmt.Println("Hello", name, "and My name is", customer.Name)
}

func main() {
	xiao := CustomerPriority{Name: "Xiao Ching", Age: 19}
	xiao.sayIntroduce("timeng")
}
