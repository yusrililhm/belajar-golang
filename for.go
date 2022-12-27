package main

import "fmt"

func main() {
    start := 1

    for start <= 10 {
        fmt.Println(start)
        start++
    }

    // for with statement
    for start := 1; start <= 10; start += 2 {
        fmt.Println("Perulangan Ke-",start)
    }

    //for range
    slice := []string{
        "timeng",
        "xiaoching",
        "popoci",
    }

    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }

    //for index/key, name/value := range slice {fmt.Println("index or key : ", index/key, "name or values : ", name/values)}
    person := map[string]string{
        "name": "xiao ching",
        "address": "cilegon",
    }

    for key, value := range person {
        fmt.Println(key, ":", value)
    }
}
