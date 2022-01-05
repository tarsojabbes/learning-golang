package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

func main() {

	point, _ := getPoint()
	findQuarter(point)

}

func getPoint() (Point, error) {
	reader := bufio.NewReader(os.Stdin)

	point, _ := reader.ReadString('\n')

	array_x_y := strings.Fields(point)

	x, _ := strconv.ParseFloat(array_x_y[0], 32)
	y, _ := strconv.ParseFloat(array_x_y[1], 32)

	result := Point{x, y}

	return result, nil
}

func findQuarter(point Point) {

	if point.x == 0 && point.y == 0 {
		fmt.Println("Origem")
	} else if point.x > 0 && point.y > 0 {
		fmt.Println("Q1")
	} else if point.x < 0 && point.y > 0 {
		fmt.Println("Q2")
	} else if point.x > 0 && point.y < 0 {
		fmt.Println("Q4")
	} else if point.x < 0 && point.y < 0 {
		fmt.Println("Q3")
	} else if point.x == 0 && point.y != 0 {
		fmt.Println("Eixo Y")
	} else if point.x != 0 && point.y == 0 {
		fmt.Println("Eixo X")
	}
}
