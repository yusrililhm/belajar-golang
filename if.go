package main

import "fmt"

func main() {
    nilaiUjian := 80
    nilaiAbsen := 80

    ujian := 90
    absen := 70

    if ujian > nilaiUjian && absen > nilaiAbsen {
        fmt.Println("Lulus")
    } else if ujian > nilaiUjian && absen < nilaiAbsen {
        fmt.Println("harus rajin masuk")
    } else {
        fmt.Println("Coba lagi")
    }

    //short statement
    if ujian := 90; ujian > nilaiUjian {
        fmt.Println("Luluuuuus")
    } else {
        fmt.Println("Cobaaaaa lagiiiii")
    }
}
