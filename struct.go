package main

import "fmt"

type Customer struct {
    Name, Address, Hobby string
    Age int
}

func main() {
    var xiao Customer
    xiao.Name = "Xiao Ching"
    xiao.Address = "Cilegon"
    xiao.Hobby = "ngoding"
    xiao.Age = 19

    fmt.Println(xiao)

    //literals 2
    timeng := Customer{
        Name: "Timeng",
        Address: "Cilegon",
        Hobby: "ngoding",
        Age: 20,
    }

    fmt.Println(timeng)

    //literals 3 but not recomended
    popoci := Customer{"Popoci", "Cilegon", "ngoding", 19}
    fmt.Println(popoci)
}
