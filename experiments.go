// Solving URI-1094 to learn about reading user input with "bufio",
// splitting input with "strings", type conversion and fmt.Printf

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

	var times int
	fmt.Scanln(&times)

	var total_cobaias int
	var total_coelhos int
	var total_ratos int
	var total_sapos int

	for times > 0 {
		experiment, _ := reader.ReadString('\n')

		experiment_array := strings.Fields(experiment)

		quantity, _ := strconv.Atoi(experiment_array[0])
		total_cobaias += quantity

		animal := experiment_array[1]

		if animal == "C" {
			coelhos, _ := strconv.Atoi(experiment_array[0])
			total_coelhos += coelhos
		} else if animal == "S" {
			sapos, _ := strconv.Atoi(experiment_array[0])
			total_sapos += sapos
		} else if animal == "R" {
			ratos, _ := strconv.Atoi(experiment_array[0])
			total_ratos += ratos
		}

		times--
	}

	perc_coelhos := (float32(total_coelhos) / float32(total_cobaias)) * 100
	perc_sapos := (float32(total_sapos) / float32(total_cobaias)) * 100
	perc_ratos := (float32(total_ratos) / float32(total_cobaias)) * 100

	fmt.Printf("Total: %d cobaias\n", total_cobaias)
	fmt.Printf("Total de coelhos: %d\n", total_coelhos)
	fmt.Printf("Total de ratos: %d\n", total_ratos)
	fmt.Printf("Total de sapos: %d\n", total_sapos)
	fmt.Printf("Percentual de coelhos %0.2f \n", perc_coelhos)
	fmt.Printf("Percentual de ratos %0.2f\n", perc_ratos)
	fmt.Printf("Percentual de sapos %0.2f\n", perc_sapos)

}
