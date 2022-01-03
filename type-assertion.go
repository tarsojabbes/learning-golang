package main

import "fmt"

// Type assertion can be use to check the type of an interface{}
func main() {
	var variable_1 interface{} = "this is a string"
	var variable_2 interface{} = 1984

	is_string := variable_1.(string) // Firstly, the interface and then the type
	is_float := variable_2.(float32)

	fmt.Print(is_string)  // Doesn't trhow an error because variable_1 is a string
	fmt.Println(is_float) // Throws a panic error because variable_2 is int not a float
}
