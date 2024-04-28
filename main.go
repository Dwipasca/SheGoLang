package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords (inputString string) map[string]int {
	// initialization map with key as string and value as int
	wordCount := make(map[string]int)

	cleanInputString := strings.TrimSpace(inputString)

	for _, char := range cleanInputString {
		wordCount[string(char)] += 1
		fmt.Println(string(char))
	}

	return wordCount
}

func main() {
	fmt.Println("\n", strings.Repeat("-", 20))
	fmt.Println("Mini challange 2 - Calculate Words")
	fmt.Println("\n", strings.Repeat("-", 20))

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input your words: ")
	inputString, _ := reader.ReadString('\n')

	wordCount := countWords(inputString)

	fmt.Println(wordCount)
}
