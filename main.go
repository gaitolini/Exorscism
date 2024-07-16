package main

import (
	"bufio"
	"exercicio/exercicios/collatz_conjecture"
	"exercicio/exercicios/hamming"
	"exercicio/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	for {
		// Exibe o menu
		fmt.Println("\nExercícios em GO:")
		fmt.Println("Escolha o exercício que quer executar")
		fmt.Println("1. Conjectura de Collatz")
		fmt.Println("2. Giga Segundos")
		fmt.Println("3. Hamming")
		fmt.Println("Enter para encerrar..")

		// Lê a entrada do usuário
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		// Converte a entrada para um número
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, escolha uma opção válida.")
			continue
		}

		// Executa o exercício escolhido
		switch choice {
		case 1:
			executeCollatzConjecture()
		case 2:
			executeAddGigaSegundos()
		case 3:
			executeHamming()
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}

func executeHamming() {

	fmt.Println("Exercício 3: Hamming")
	fmt.Println("Escreva uma função que receba duas string que exemplifique pedaços de DNA's")
	fields := []utils.InputField{
		{Prompt: "DNA 1", Type: "string"},
		{Prompt: "DNA 2\n", Type: "string"},
	}

	err := utils.ReadInputs(fields)
	if err != nil {
		fmt.Println("Erro ao ler os dados: ", err)
		return
	}

	dna1, ok1 := fields[0].Value.(string)
	dna2, ok2 := fields[1].Value.(string)

	if !ok1 || !ok2 {
		fmt.Println("Erro ao converter valores para string de DNA1 e ou DNA2")
		return
	}

	hamming, ok3 := hamming.Distance(dna1, dna2)
	if ok3 != nil {
		fmt.Println("Erro ao calcular a distância Hamming")
		return
	}

	fmt.Printf("O valor da distancia de hamming é %d\n", hamming)

}

func executeCollatzConjecture() {

	fields := []utils.InputField{
		{Prompt: "Digite um número inteiro para Conjectura de Collatz", Type: "int"},
	}

	err := utils.ReadInputs(fields)
	if err != nil {
		fmt.Println("Erro ao ler as entradas:", err)
		return
	}

	value := fields[0].Value.(int)
	result, err := collatz_conjecture.Conjecture(value)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(fmt.Sprintf("Conjecture(%d) = %d", value, result))
	}
}

func executeAddGigaSegundos() {
	fields := []utils.InputField{
		{Prompt: "Digite uma data(dd/mm/aaaa) calcular 1 GigaSegundo", Type: "date"},
	}
	utils.ReadInputs(fields)
	date, err := time.Parse("02/01/2006", fields[0].Value.(string))
	if err != nil {
		fmt.Println("Erro:", err)
		return
	} else {
		fmt.Println(fmt.Sprintf("Data: %s", date.Add(time.Second*100000000)))
	}
}
