package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play := "s"
	for play == "s" {
		fmt.Printf("Deseja jogar? (s/n)")
		fmt.Scan(&play)
		if play == "n" {
			fmt.Printf("Obrigado por jogar")
			break
		}
		var current int
		num := rand.Intn(100)
		fmt.Println(num)
		found := 0
		try := 0
		for found == 0 {
			try += 1
			fmt.Printf("Insira valor: ")
			fmt.Scan(&current)
			if current == num {
				fmt.Printf("Parabens, você achou o numero!\n")
				fmt.Printf("Tentativas: %d", try)
				found = 1
			} else if current > num {
				fmt.Printf("Menor")
			} else if current < num {
				fmt.Printf("Maior")
			}
		}

	}
}
