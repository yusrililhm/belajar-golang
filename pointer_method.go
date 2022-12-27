package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Graduate()  {
	man.Name = "Mr. " + man.Name
}

func main()  {
	xiao := Man{"Xiao Ching"}

	xiao.Graduate()

	fmt.Println(xiao.Name)
}