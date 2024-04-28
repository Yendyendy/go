package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	key := "golang"

	codedText := vigenere(text, key)
	vigenereDecode(codedText, key)
}

func vigenere(plainText, key string) string {
	plainText = strings.ToUpper(plainText)
	key = strings.ToUpper(key)

	var pivote byte = 'A'
	var codedText []byte

	for i := 0; i < len(plainText); i++ {

		aux := key[i%len(key)] - pivote
		codedText = append(codedText, plainText[i]+aux)

	}
	fmt.Println("Coded: ", string(codedText))

	return string(codedText)
}

func vigenereDecode(codedText, key string) string {
	key = strings.ToUpper(key)

	var pivote byte = 'A'
	var plainText []byte

	for i := 0; i < len(codedText); i++ {
		aux := key[i%len(key)] - pivote
		plainText = append(plainText, codedText[i]+aux)
	}
	fmt.Println("Decoded: ", string(plainText))

	return string(plainText)
}
