package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	inputString := "selamat malam"
	expectedWordCount := map[string]int{
		" ":1,
		"a":4, 
		"e":1, 
		"l":2,
		"m":3,
		"s":1,
		"t":1,
	}

	wordCount := countWords(inputString)

	// compare between wordCount (result) with expected we want to test
	// if it;s not the same then show the message error 
	if !reflect.DeepEqual(wordCount, expectedWordCount) {
		t.Errorf("Word count is incorrect, expected: %v, results: %v", expectedWordCount, wordCount)
	}
}
