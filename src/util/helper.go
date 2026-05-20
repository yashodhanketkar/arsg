package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

type ParamType map[string]int

type ConfigType struct {
	Parameters []ParamType `json:"parameters"`
	ExportPath string      `json:"export_path"`
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

func LoadConfig(config *ConfigType) error {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "arsg", "config.json")

	file, err := os.ReadFile(configPath)
	if err != nil {
		return errors.New("could not open config file: " + err.Error())
	}

	if err := json.Unmarshal(file, &config); err != nil {
		return errors.New("could not unmarshal config file: " + err.Error())
	}

	handleParameters(config)
	handleExportPath(config)

	return nil
}

func handleParameters(config *ConfigType) []ParamType {
	var params []ParamType

	if len(config.Parameters) == 0 {
		config.Parameters = defaultConfigParams
		return defaultConfigParams
	}

	for _, p := range config.Parameters {
		params = append(params, p)
	}

	return params
}

func handleExportPath(config *ConfigType) {
	home := os.Getenv("HOME")

	if config.ExportPath == "" {
		config.ExportPath = filepath.Join(home, "temp", "export.json")
		return
	}

	if filepath.IsAbs(config.ExportPath) {
		return
	}

	config.ExportPath = filepath.Join(home, config.ExportPath)
}
