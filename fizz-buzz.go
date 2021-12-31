package main

import "fmt"

// Training flux control
func main() {
	var n int
	fmt.Scanln(&n)

	result_fizz_buzz := fizz_buzz(n)

	fmt.Println(result_fizz_buzz)
}

func fizz_buzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz!"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return "Neither Fizz nor Buzz"
	}
}
