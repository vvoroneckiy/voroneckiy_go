package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}

	firstChar := word[0]

	rest := word[1:]

	runes := []rune(rest) 
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedRest := string(runes)

	return string(firstChar) + reversedRest
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)

	encryptedWords := make([]string, len(words))

	for i, word := range words {
		encryptedWords[i] = encryptWord(word)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	testPhrases := []string{
		"Pepe Schnele is a legend",
		"Hello world",
		"a",
		"go is awesome",
		"Крипто Пепе", 
		"",
	}

	fmt.Println("=== Шифратор фраз ===\n")

	for _, phrase := range testPhrases {
		encrypted := encryptPhrase(phrase)
		fmt.Printf("Исходная фраза: \"%s\"\n", phrase)
		fmt.Printf("Зашифровано:    \"%s\"\n\n", encrypted)
	}
}