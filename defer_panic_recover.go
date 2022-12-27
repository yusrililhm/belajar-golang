package main

import "fmt"

//defer

func logging() {
    fmt.Println("selesai memangil function")
}

func runAplication() {
    defer logging()
    fmt.Println("Run Aplication")
}

func Penjumlahan(angka1 int, angka2 int) {
    defer logging()
    fmt.Println(angka1 / angka2)
}

// Panic

func endApp() {
    fmt.Println("Selesai")

    message := recover()

    if message != nil {
        fmt.Println("Terjadi error : ", message)
    }
    fmt.Println("Aplikasi berjalan")
}

func runApp(error bool) {

    defer endApp()

    if error {
        panic("Aplikasi error")
    }

    fmt.Println("Success")
}

//recover




func main() {
	//runAplication()
    //penjumlahan(10, 10)
    //runApp(false)
    runApp(false)
}
