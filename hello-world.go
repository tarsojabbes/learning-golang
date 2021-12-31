package main

import "fmt"

func main() {

	var name string

	// Reads a line and assigns the result to name
	fmt.Scanln(&name)

	var year_of_birth int

	// Reads another line and assigns the result to year_of_birth
	fmt.Scanln(&year_of_birth)

	// Calculates the age based on the read line
	age := calculate_age(year_of_birth)

	fmt.Println("Hello " + name)
	// %d is the placeholder to Int values
	fmt.Printf("You're %d years-old", age)

	fmt.Println("----------- Insert the values of X, Y and Z to calculate the volume---------")

	var x float32
	var y float32
	var z float32

	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&z)

	volume := calculate_volume(x, y, z)

	// %f is the placeholder por Float values
	fmt.Printf("The volume is %f", volume)
}

func calculate_age(year int) int {
	age := 2021 - year
	return age
}

func calculate_volume(x float32, y float32, z float32) float32 {
	volume := x * y * z
	return volume
}
