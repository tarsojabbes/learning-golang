package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	converted_input := convert_input(input)

	result := swaps(converted_input)

	if result == 0 {
		fmt.Println("Carlos")
	} else if result%2 == 0 {
		fmt.Println("Carlos")
	} else if result%2 != 0 {
		fmt.Println("Marcelo")
	}
}

func convert_input(input string) []float32 {

	input_array := strings.Fields(input)

	var float_numbers_array []float32

	for _, value := range input_array {
		float_value, _ := strconv.ParseFloat(value, 32)
		float_numbers_array = append(float_numbers_array, float32(float_value))
	}

	return float_numbers_array
}

func swaps(array []float32) int {
	var swap int
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				aux_var := array[j]
				array[j] = array[j+1]
				array[j+1] = aux_var
				swap++
			}
		}
	}

	return swap
}
