package main

import "fmt"

func main()  {
	name := "udin"
	increment := func ()  {
		name := "budi"
		fmt.Println(name)
	}

	increment()
	fmt.Println(name)
}