package main

import "fmt"

//type func declaration
type passwordCheck func(string)bool

func getPassword(password string, passwordCheck passwordCheck)  {
	if passwordCheck(password) {
		fmt.Println("Welcome back")
	} else {
		fmt.Println("You're password is wrong, please insert again")
	}
}

func main() {
	passwordCheck := func(password string) bool {
		return password == "admin"
	}

	getPassword("popoci", passwordCheck)

	getPassword("xxxx", func(password string) bool  {
		return password == "xxxx"
	})
}
