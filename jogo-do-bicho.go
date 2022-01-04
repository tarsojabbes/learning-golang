package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	cartela := map[int]string{
		1: "Avestruz", 2: "Águia", 3: "Burro", 4: "Borboleta", 5: "Cachorro",
		6: "Cabra", 7: "Carneiro", 8: "Camelo", 9: "Cobra", 10: "Coelho",
		11: "Cavalo", 12: "Elefante", 13: "Galo", 14: "Gato", 15: "Jacaré",
		16: "Leão", 17: "Macaco", 18: "Porco", 19: "Pavão", 20: "Peru",
		21: "Touro", 22: "Tigre", 23: "Urso", 24: "Veado", 25: "Vaca",
	}

	numero := numero_aleatorio()

	resultado := jogo(numero, cartela)

	fmt.Println(resultado)

}

func numero_aleatorio() int {
	// Muda o número cada vez que chamamos a função
	rand.Seed(time.Now().UnixNano())

	numero := rand.Intn(24) + 1 // rand.Intn(max-min) + min

	return numero
}

func jogo(number int, cartela map[int]string) string {
	resultado := cartela[number]
	return "O número foi " + strconv.Itoa(number) + " e o animal sorteado foi " + resultado
}
