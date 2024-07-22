package main

import (
	"bufio"
	"exercicio/exercicios/collatz_conjecture"
	"exercicio/exercicios/diffsquares"
	"exercicio/exercicios/hamming"
	"exercicio/exercicios/leap"
	"exercicio/exercicios/luhn"
	"exercicio/exercicios/scrabble"
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
		fmt.Println("4. Scrabble")
		fmt.Println("5. IsLeapYear")
		fmt.Println("6. DifferenceOfSquares")
		fmt.Println("7. Luhn")
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
		case 4:
			executeScrabble()
		case 5:
			executeIsLeapYear()
		case 6:
			executeDifferenceOfSquares()
		case 7:
			executeLuhn()
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}

func executeLuhn() {
	validCard := ""
	invalidCard := ""

	fmt.Println("Luhn")
	fmt.Println("1. Gerar cartão válido")
	fmt.Println("2. Gerar cartão inválido")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for {

		if (input != "1") && (input != "2") {
			continue
		}

		switch input {
		case "1":
			validCard = luhn.GenerateRandomCreditCard(true)
			fmt.Printf("Valid Card: %s\n", validCard)

		case "2":
			invalidCard = luhn.GenerateRandomCreditCard(false)
			fmt.Printf("Invalid Card: %s\n", invalidCard)
		}

		if (input == "1") || (input == "2") {
			break
		}

		fmt.Println("1. Gerar cartão válido")
		fmt.Println("2. Gerar cartão inválido")
		reader = bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

	}

	fmt.Println("Input.. ", input)
	fmt.Println("Validando.")
	time.Sleep(1 * time.Second)
	fmt.Println("Validando..")
	time.Sleep(2 * time.Second)

	switch input {
	case "1":
		fmt.Printf("Is Valid Card valid? %v\n", luhn.IsValidLuhn(validCard))
	case "2":
		fmt.Printf("Is Invalid Card valid? %v\n", luhn.IsValidLuhn(invalidCard))
	}

}

func executeDifferenceOfSquares() {
	fmt.Println("Exercício 6: Diferença de Quadrados²")
	fmt.Println("Informe um número inteiro:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Entrada inválida. Por favor, informe um número inteiro.")
		return
	}

	fmt.Println("O resultado é:", diffsquares.Difference(n))

}

func executeIsLeapYear() {
	fmt.Println("IsLeapYear")
	fmt.Println("Escolha o ano que quer verificar se é bissexto")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	year, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, escolha uma opção válida")
		return
	}

	if leap.IsLeapYear(year) {
		fmt.Printf("%d é um ano bissexto", year)

	} else {
		fmt.Printf("%d não é um ano bissexto", year)
	}

}

func executeScrabble() {
	fmt.Println("Scrabble")
	fmt.Println("Sua tarefa é calcular a pontuação Scrabble de uma palavra somando os valores de suas letras.")
	fmt.Println("As letras são avaliadas da seguinte forma:")
	fmt.Println("A, E, I, O, U, L, N, R, S, T = 1 ponto")
	fmt.Println("D, G = 2 ponto")
	fmt.Println("B, C, M, P = 3 ponto")
	fmt.Println("F, H, V, W, Y = 4 ponto")
	fmt.Println("K = 5 ponto")
	fmt.Println("J, X = 8 ponto")
	fmt.Println("Q, Z = 10 ponto")
	fmt.Println("Por exemplo, a palavra \"abracadabra\" tem uma pontuação de ")
	fmt.Println("1 + 1 + 4 + 1 + 3 + 1 +1 + 1 + 1 + 1 + 1 + 1 + 1 = 17 pontos.")
	fmt.Println("")
	fmt.Println("Digite a palavra para calcular a pontuação:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("!! Scrabble!!")
	fmt.Println("")
	fmt.Printf(" '%s' Sua pontuação é de %d!\n", input, scrabble.Score(input))

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
