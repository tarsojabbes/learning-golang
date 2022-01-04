package main

import (
	"fmt"
	"time"
)

func main() {

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// This one is a normal function an the algoritghm only goes one when this is finished
	fmt.Println("Time elapsed to calculate all values")
	calculate_double(values)

	// This one is a GoRoutine, the rest of the algorithm is going to run even if it's not finished
	fmt.Println("Time elapsed to calculate all values")
	// Different from the above, this function will not log any result
	// It's possible to access its value using channels
	go calculate_double(values)

}

// Calculates the time of a function to run
func elapsedTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

// Simple function that calculates the double of each value in a array
func calculate_double(values []int) {
	elapsedTime(time.Now(), "Calculate double")

	var doubled_values []int
	for _, value := range values {
		doubled_values = append(doubled_values, value*2)
	}
}
