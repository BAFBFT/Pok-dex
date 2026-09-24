package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string

	trimmed := strings.TrimSpace(text)
	i := strings.Index(trimmed, " ")
	for i > -1 {
		result = append(result, strings.ToLower(trimmed[:i]))
		trimmed = trimmed[i+1:]
		i = strings.Index(trimmed, " ")
	}
	return append(result, strings.ToLower(trimmed))
}
