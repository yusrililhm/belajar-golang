package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Cilegon", "Banten", "Indonesia"}

	//pass by value
	//address2 := address1
	//address2.City = "Serang"

	//fmt.Println(address1) //address 1 tidak berubah
	//fmt.Println(address2)

	//pass by reference
	address3 := &address1
	address3.City = "Rangkasbitung"

	fmt.Println(address1)
	fmt.Println(address3)

	//operator *
	//address3 = &Address{"Jakarta", "DKI", "Indonesia"}

	//fmt.Println(address1) //address 1 tidak berubah
	//fmt.Println(address3)

	*address3 = Address{"Malang", "Jatim", "Indonesia"}

	fmt.Println(address3)
	fmt.Println(address1)

	//func new
	address4 := new(Address)
	address5 := address4

	address5.Province = "Jawa Timur"

	fmt.Println(address4)
	fmt.Println(address5)
}