package util

import (
	"fmt"
	"unicode"
)

func HelperHandler(handler int, args ...string) any {
	switch handler {
	case 1:
		return getParams(args...)
	case 2:
		val, err := capitalizeFirstLetter(args...)
		if err != nil {
			return err
		}
		return val
	default:
		return nil
	}
}

func getParams(args ...string) []string {
	if len(args) == 0 {
		return []string{"Art/Animation", "Character/Cast", "Plot", "Bias"}
	}

	return args
}

func capitalizeFirstLetter(args ...string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("No arguments provided")
	}

	str := args[0]
	if len(str) == 0 {
		return str, nil
	}

	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes), nil
}
