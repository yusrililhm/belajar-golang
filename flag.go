package main

import (
	"flag"
	"fmt"
)

func main() {
	host :=flag.String("host","localhost", "put your database host")
	user :=flag.String("user","root", "put your database user")
	password :=flag.String("password","root", "put your database password")

	flag.Parse()

	fmt.Println(*host, *user, *password)
}