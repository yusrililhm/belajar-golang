package main

import "fmt"

func mahaSiswa() (nim string, smt int8, address string) {
    nim = "11121042"
    smt = 3
    address = "cilegon"

    return
}

func main() {
    nim, smt, address := mahaSiswa()
    fmt.Println(nim, smt, address)
}
