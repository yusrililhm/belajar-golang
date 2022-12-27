package main

import "fmt"

func sumAll(numbers ...int) int {
    total := 0

    for _, number := range numbers {
        total += number
    }

    return total
}

func main() {
    total := sumAll(1, 3, 5, 7, 9, 11)
    fmt.Println(total)

    //slice param
    angka := []int{10,20,30}
    total = sumAll(angka...)
    fmt.Println(total)
}
