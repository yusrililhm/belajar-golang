package main

import (
    "fmt"
    "errors"
)

func pembagian(pembilang float32, penyebut float32) (float32, error) {
    if penyebut == 0 {
        return 0, errors.New("Can't devided by zero")
    } else {
        return pembilang / penyebut, nil
    }
}

func main() {
    result, err := pembagian(5, 0)
    if err == nil {
        fmt.Println("Result is :", result)
    } else {
        fmt.Println("Error :", err.Error())
    }
}
