package main

import "fmt"

func main() {

    name := "Timeng"

    switch name {

        case "xiao ching":
            fmt.Println("Heloooo Xiao Ching")

        case "Popoci":
            fmt.Println("Heloooo Popoci")

        default:
            fmt.Println("kamu siapaaaaa?")

    }

    //short statement

    nilaiUjian := 76

    switch ujian := 80; ujian > nilaiUjian {

        case true:
            fmt.Println("Selamat")
            fmt.Println("Kamu Lulussss")

        case false:
            fmt.Println("coba lagi")
            fmt.Println("tetap semangat")

    }

    //tanpa kondisi

    ujian := 75

    switch {

        case ujian > 89:
            fmt.Println("A")

        case ujian > 80:
            fmt.Println("B")

        case ujian > 70:
            fmt.Println("C")

        default:
            fmt.Println("D")

    }
}
