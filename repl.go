package main

import "strings"

func cleanInput(text string) []string {
	clean := []string{}

	words := strings.Fields(text)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		clean = append(clean, lowerWord)
	}

	return clean
}
