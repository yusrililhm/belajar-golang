package main

import "fmt"

func main() {

    //array 1
    var nameArray [3]string

    nameArray[0] = "Xiao"
    nameArray[2] = "Popoci"

    fmt.Println(nameArray)
    fmt.Println(nameArray[0])
    fmt.Println(nameArray[2])

    //array 2
    var nameCat = [...]string{
        "timeng",
        "macing",
        "Popoci",
    }

    fmt.Println(nameCat)
    fmt.Println(len(nameCat))
}
