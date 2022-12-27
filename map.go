package main

import "fmt"

func main() {
    //nameMap := map[key]{value}

    person := map[string]string{
        "name": "xiao ching",
        "address": "cilegon",
    }

    fmt.Println(person)
    fmt.Println(person["name"])
    fmt.Println(person["address"])
    fmt.Println(len(person))

    catType := make(map[string]string)
    catType["name"] = "timeng"
    catType["ras"] = "anggora"

    fmt.Println(catType)

    delete(catType, "ras")
    fmt.Println(catType)
}
