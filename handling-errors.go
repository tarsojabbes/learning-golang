package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := select_value()

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Print(err.Error())
	}

}

// This function can either return a string or an error
func select_value() (string, error) {

	fmt.Println("Select a value between 1 and 10")

	var number int32
	fmt.Scanln(&number)

	if 1 <= number && number <= 10 {
		return "You chose a valid number" + string(number), nil
	} else {
		return "", errors.New("You chose an invalid number")
	}

}
