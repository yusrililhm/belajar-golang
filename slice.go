package main

import "fmt"

func main() {
    //var nameSlice

    var nameArray = [12]string{
        "januari",
        "februari",
        "maret",
        "april",
        "mei",
        "juni",
        "juli",
        "agustus",
        "september",
        "oktober",
        "november",
        "desember",
    }

    var slice = nameArray[4:7]

    fmt.Println(slice)
    fmt.Println(len(slice))
    fmt.Println(cap(slice))

    //append(slice, data)

    var slice2 = append(slice, "libur", "gajian", "lebaran")
    fmt.Println(slice2)
    fmt.Println(slice)

    //make([]types, len, cap)

    days := make([]string, 3, 7)
    days[0] = "Sunday"
    days[1] = "Monday"
    days[2] = "Tuesday"

    fmt.Println(days)
    fmt.Println(cap(days))
    fmt.Println(len(days))

    //copy(destination, source)

    fromSlice := days[:]
    toSlice := make([]string, len(fromSlice), cap(fromSlice))

    copy(toSlice, fromSlice)

    fmt.Println(toSlice)
}
