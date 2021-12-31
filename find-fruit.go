package main

import "fmt"

func main() {
	fruit := get_fruit()
	array_fruit := get_array_fruit()
	result := find_fruit(fruit, array_fruit)

	if result {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}

func get_fruit() string {
	var fruit string
	fmt.Scanln(&fruit)
	return fruit
}

func get_array_fruit() []string {
	var array_fruit [5]string

	for i := range array_fruit {
		fmt.Scanln(&array_fruit[i])
	}

	return array_fruit[:]

}

func find_fruit(fruit string, array_fruit []string) bool {
	found_fruit := false

	for _, current_fruit := range array_fruit {
		if current_fruit == fruit {
			found_fruit = true
		}
	}

	if found_fruit {
		return true
	} else {
		return false
	}

}
