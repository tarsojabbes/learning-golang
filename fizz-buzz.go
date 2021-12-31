package main

import "fmt"

// Training flux control
func main() {

	var n int
	fmt.Scanln(&n)
	result_fizz_buzz := fizz_buzz(n)
	fmt.Println(result_fizz_buzz)

	var m int
	fmt.Scanln(&m)
	fizz_buzz_for_loop(m)
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

func fizz_buzz_for_loop(m int) {
	for i := 0; i <= m; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz!")
		} else if i%3 == 0 {
			fmt.Println("Fizz!")
		} else if i%5 == 0 {
			fmt.Println("Buzz!")
		} else {
			fmt.Println(i)
		}
	}
}
