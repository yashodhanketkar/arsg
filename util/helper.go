package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

type ParamType map[string]int

type ConfigType struct {
	Parameters []ParamType `json:"parameters"`
}

var defaultConfigParams = []ParamType{
	{"Art/Animation": 25},
	{"Character/Cast": 30},
	{"Plot": 35},
	{"Bias": 10},
}

func GetParams(config *ConfigType) ([]string, []int) {
	if len(config.Parameters) == 0 {
		config.Parameters = defaultConfigParams
	}

	return paramLists(config.Parameters)
}

func paramLists(args []ParamType) ([]string, []int) {
	var params = make([]string, 0)
	var weights = make([]int, 0)

	for _, p := range args {
		for k, v := range p {
			params = append(params, k)
			weights = append(weights, v)
		}
	}

	return params, weights
}

func CapitalizeFirstLetter(args ...string) (string, error) {
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

func LoadConfig(config *ConfigType) {
	var params = make([]string, 0)
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "arsg", "config.json")

	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("could not open config file:", err)
		os.Exit(1)
	}

	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println("could not unmarshal config file:", err)
	}

	if len(config.Parameters) == 0 {
		config.Parameters = defaultConfigParams
		return
	}

	for _, p := range config.Parameters {
		for k := range p {
			params = append(params, k)
		}
	}
}
