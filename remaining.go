package main

import "fmt"

func main() {

	var initial int
	var final int

	fmt.Scanln(&initial)
	fmt.Scanln(&final)

	if initial > final {
		initial, final = final, initial
	}

	for i := initial; i < final; i++ {
		if i%5 == 2 || i%5 == 3 {
			fmt.Println(i)
		}
	}

}
