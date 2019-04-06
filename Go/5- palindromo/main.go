package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	fmt.Print("Ingresa: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, " ", "", -1)
	text = strings.ToLower(text)
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	fmt.Println(checkPal(text))
}

func checkPal(s string) string {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "No es palindromo"
		}
	}
	return "Si es palindromo"
}