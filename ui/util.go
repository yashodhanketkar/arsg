package ui

import "strings"

func numericInput(str string) string {
	var inputBuilder strings.Builder

	for _, r := range str {
		if strings.ContainsRune("0123456789.", r) {
			inputBuilder.WriteRune(r)
		}
	}
	return inputBuilder.String()
}
