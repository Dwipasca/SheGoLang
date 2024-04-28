package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\n", strings.Repeat("-", 20), "\n")
	fmt.Println("Mini challange 2 - Calculate Words")
	fmt.Println("\n", strings.Repeat("-", 20), "\n")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input your words: ")
	inputString, _ := reader.ReadString('\n')

	cleanInputString := strings.TrimSpace(inputString)

	wordCount := make(map[string]int)

	for _, char := range cleanInputString {
		wordCount[string(char)]++
	}

	fmt.Println(wordCount)
}
