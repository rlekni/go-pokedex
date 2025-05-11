package main

import "strings"

func ClearInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
