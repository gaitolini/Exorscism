package luhn

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateRandomCreditCard gera um número de cartão de crédito aleatório
// Se o parâmetro valid for true, gera um número válido, caso contrário gera um número inválido
func GenerateRandomCreditCard(valid bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Gera os primeiros 15 dígitos aleatórios
	cardNumber := make([]int, 15)
	for i := 0; i < 15; i++ {
		cardNumber[i] = r.Intn(10)
	}

	// Calcula o dígito verificador usando o algoritmo de Luhn
	checkDigit := calculateLuhnCheckDigit(cardNumber)

	// Se queremos um número válido, usamos o dígito verificador calculado
	// Se queremos um número inválido, ajustamos o dígito verificador para garantir a invalidade
	if valid {
		cardNumber = append(cardNumber, checkDigit)
	} else {
		invalidDigit := (checkDigit + 1) % 10
		cardNumber = append(cardNumber, invalidDigit)
		for IsValidLuhn(ConvertToString(cardNumber)) {
			invalidDigit = (invalidDigit + 1) % 10
			cardNumber[15] = invalidDigit
		}
	}

	// Converte o número do cartão para uma string
	return ConvertToString(cardNumber)
}

// calculateLuhnCheckDigit calcula o dígito verificador de um número de cartão de crédito usando o algoritmo de Luhn
func calculateLuhnCheckDigit(cardNumber []int) int {
	sum := 0
	doubleDigit := true

	// Itera sobre os dígitos do número do cartão de trás para frente
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := cardNumber[i]
		if doubleDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		doubleDigit = !doubleDigit
	}

	// O dígito verificador é o valor que torna a soma um múltiplo de 10
	checkDigit := (10 - (sum % 10)) % 10
	return checkDigit
}

// IsValidLuhn verifica se um número de cartão de crédito é válido usando o algoritmo de Luhn
func IsValidLuhn(cardNumber string) bool {
	CardValid := strings.Trim(cardNumber, " ")

	cardNumberInts := make([]int, len(CardValid))
	for i, digit := range cardNumber {
		cardNumberInts[i] = int(digit - '0')
	}

	sum := 0
	doubleDigit := true

	// Itera sobre os dígitos do número do cartão de trás para frente
	for i := len(cardNumberInts) - 1; i >= 0; i-- {
		digit := cardNumberInts[i]
		if doubleDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		doubleDigit = !doubleDigit
	}

	return sum%10 == 0
}

// ConvertToString converte um slice de inteiros em uma string
func ConvertToString(cardNumber []int) string {
	cardNumberStr := ""
	for _, digit := range cardNumber {
		cardNumberStr += strconv.Itoa(digit)
	}
	return cardNumberStr
}
