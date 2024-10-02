package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Função para gerar números aleatórios em um intervalo específico
func generateRandomNumbers(quantity, min, max int, wg *sync.WaitGroup, numbersChan chan []int) {
	defer wg.Done()

	if min >= max {
		fmt.Println("Erro: O valor mínimo deve ser menor que o valor máximo.")
		return
	}

	numbers := make([]int, quantity)
	for i := 0; i < quantity; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}

	numbersChan <- numbers
}

// Função para salvar os números gerados em um arquivo
func saveToFile(numbers [][]int) error {
	file, err := os.Create("numeros_gerados.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for i, set := range numbers {
		_, err := file.WriteString(fmt.Sprintf("Intervalo %d:\n", i+1))
		if err != nil {
			return err
		}
		for _, num := range set {
			_, err := file.WriteString(fmt.Sprintf("%d ", num))
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var numIntervals, quantity, min, max int

	// Definir uma semente para o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Gerador de Números Aleatórios Avançado")
	fmt.Println("-------------------------------------")

	// Solicitar o número de intervalos de geração
	fmt.Print("Quantos intervalos de números aleatórios você deseja gerar? ")
	fmt.Scanln(&numIntervals)

	numbersChan := make(chan []int, numIntervals)
	var wg sync.WaitGroup
	allNumbers := make([][]int, numIntervals)

	// Para cada intervalo, solicitar os parâmetros e iniciar uma goroutine para gerar os números
	for i := 0; i < numIntervals; i++ {
		fmt.Printf("\nIntervalo %d:\n", i+1)
		fmt.Print("Quantos números aleatórios você deseja gerar? ")
		fmt.Scanln(&quantity)

		fmt.Print("Digite o valor mínimo: ")
		fmt.Scanln(&min)
		fmt.Print("Digite o valor máximo: ")
		fmt.Scanln(&max)

		wg.Add(1)
		go generateRandomNumbers(quantity, min, max, &wg, numbersChan)
	}

	// Aguardar o término de todas as goroutines
	go func() {
		wg.Wait()
		close(numbersChan)
	}()

	// Coletar os números gerados de cada goroutine
	for numbers := range numbersChan {
		allNumbers = append(allNumbers, numbers)
	}

	// Exibir os números gerados
	for i, set := range allNumbers {
		if len(set) > 0 {
			fmt.Printf("\nNúmeros gerados no intervalo %d:\n", i+1)
			for _, num := range set {
				fmt.Printf("%d ", num)
			}
			fmt.Println()
		}
	}

	// Perguntar ao usuário se deseja salvar os números em um arquivo
	var saveChoice string
	fmt.Print("\nDeseja salvar os números gerados em um arquivo? (s/n): ")
	fmt.Scanln(&saveChoice)

	if saveChoice == "s" || saveChoice == "S" {
		err := saveToFile(allNumbers)
		if err != nil {
			fmt.Println("Erro ao salvar os números no arquivo:", err)
		} else {
			fmt.Println("Números salvos com sucesso no arquivo 'numeros_gerados.txt'.")
		}
	}
}
