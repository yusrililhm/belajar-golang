package main

import "fmt"

func factorial(value int) int  {
	result := 1
	for i := value; i > 0; i-- {
		result *= i		
	}

	return result
}

func factorialRecursive(value int64) int64  {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorialRecursive(5))
}