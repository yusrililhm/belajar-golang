package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func ChangeProvinceToWestJava(alamat *Alamat)  {
	alamat.Province = "Jawa Barat"
}

func main() {
	alamat := Alamat{"Bekasi", "", "Indonesia"}

	ChangeProvinceToWestJava(&alamat)

	fmt.Println(alamat)
}