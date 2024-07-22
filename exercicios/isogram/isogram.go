package isogram

import "strings"

// Determine se uma palavra ou frase é um isograma.

// Um isograma (também conhecido como "palavra não-padrão")
// é uma palavra ou frase sem uma letra repetitiva,
//no entanto, espaços e hífens podem aparecer várias vezes.

// Exemplos de isogramas:

// Lenhadores
// fundo
// jusante
// criança de seis anos
// A palavra isogramas, no entanto, não é um isograma,
//porque o s se repete.

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	// Remove os espaços e hífens

	letterMap := make(map[rune]bool)

	for _, char := range word {
		if char == ' ' || char == '-' {
			continue
		}

		if letterMap[char] {
			return false
		}

		letterMap[char] = true
	}

	return true
}
