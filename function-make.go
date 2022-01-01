package main

import "fmt"

func main() {

	// Initalize a array of string with the size of 5
	a := make([]string, 5)
	fmt.Println(a)

	// Make() can be use to initialize maps
	// Maps are just like in Python, where you store key-values pairs
	first_map := make(map[string]int)

	// Populating
	first_map["Tarso"] = 13
	first_map["Jabbes"] = 21

	fmt.Println(first_map)

	// Another way to initialize maps with the Make() function
	second_map := map[string]int{
		"Paraíba": 83,
		"Roraima": 95,
	}

	fmt.Println(second_map)

}
