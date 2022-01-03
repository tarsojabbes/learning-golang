package main

import (
	"bufio"
	"fmt"
	"os"
)

// Slices are resizable arrays
func main() {

	result := add_names_to_list(5)

	fmt.Println(result)

}

// Appending items to a Slice
func add_names_to_list(n int) []string {
	var my_slice = make([]string, n)

	reader := bufio.NewReader(os.Stdin)

	for n > 0 {

		name, _ := reader.ReadString('\n')
		my_slice = append(my_slice, name)
		n--
	}

	return my_slice
}
