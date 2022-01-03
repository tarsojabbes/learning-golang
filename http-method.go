package main

import "fmt"

func main() {
	var method string
	fmt.Scanln(&method)

	result := http_method(method)

	fmt.Println(result)
}

// Using switch method
func http_method(method string) string {
	var result_method string
	switch method {
	case "GET":
		result_method = "GET Method"
		break
	case "POST":
		result_method = "POST Method"
		break
	case "PUT":
		result_method = "PUT Method"
	case "DELETE":
		result_method = "DELETE Method"
	}

	return result_method
}
