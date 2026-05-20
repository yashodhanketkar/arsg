package util

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParams(t *testing.T) {
	t.Run("should pass for default parameters", func(t *testing.T) {
		tests := []struct {
			name        string
			args        []string
			wantParams  []string
			wantWeights []int
			config      ConfigType
		}{
			{
				name:        "Get params default",
				wantParams:  DefaultParams,
				wantWeights: []int{25, 30, 35, 10},
				config:      ConfigType{},
			},
			{
				name:        "Get params via args",
				wantParams:  []string{"studio", "genre"},
				wantWeights: []int{1, 1},
				config: ConfigType{
					Parameters: []ParamType{{"studio": 1}, {"genre": 1}},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotParams, gotWeights := GetParams(&tt.config)
				assert.Equal(t, tt.wantParams, gotParams)
				assert.Equal(t, tt.wantWeights, gotWeights)
			})
		}
	})

	t.Run("test readParams - custom parameters", func(t *testing.T) {
		var config ConfigType
		LoadConfig(&config)

		expectedParam := []ParamType{
			{"Art/Animation": 25},
			{"Character/Cast": 30},
			{"Plot": 35},
			{"Bias": 10},
		}

		expectedParamList := DefaultParams
		actualParamList, _ := GetParams(&config)

		assert.Equal(t, expectedParam, config.Parameters)
		assert.Equal(t, expectedParamList, actualParamList)
	})

	t.Run("test readParams - default parameters", func(t *testing.T) {
		var config ConfigType = ConfigType{}

		expectedParamList := DefaultParams
		actualParamList, _ := GetParams(&config)

		assert.Equal(t, defaultConfigParams, config.Parameters)
		assert.Equal(t, expectedParamList, actualParamList)
	})
}

func TestCapitalizeFirstLetter(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want any
		err  error
	}{
		{
			name: "no args provided",
			args: []string{},
			want: "",
			err:  fmt.Errorf("No arguments provided"),
		},
		{
			name: "empty string",
			args: []string{""},
			want: "",
			err:  nil,
		},
		{
			name: "single character string",
			args: []string{"a"},
			want: "A",
			err:  nil,
		},
		{
			name: "mulit-character string",
			args: []string{"abcd"},
			want: "Abcd",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CapitalizeFirstLetter(tt.args...)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHandleParameters(t *testing.T) {
	tests := []struct {
		name     string
		input    *ConfigType
		expected []ParamType
	}{
		{
			name:     "nil parameters",
			input:    &ConfigType{Parameters: nil},
			expected: defaultConfigParams,
		},
		{
			name:     "empty parameters",
			input:    &ConfigType{Parameters: []ParamType{}},
			expected: defaultConfigParams,
		},
		{
			name:     "single parameters",
			input:    &ConfigType{Parameters: []ParamType{{"testPara": 34}}},
			expected: []ParamType{{"testPara": 34}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleParameters(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestExportPath(t *testing.T) {
	home, err := os.UserHomeDir()

	if err != nil {
		t.Fatalf("Failed to fetch user home directory %s", err)
	}

	tests := []struct {
		name   string
		config ConfigType
		want   string
	}{
		{
			name:   "empty path",
			config: ConfigType{ExportPath: ""},
			want:   filepath.Join(home, "temp", "export.json"),
		},
		{
			name:   "relative path",
			config: ConfigType{ExportPath: "tmp/export.json"},
			want:   filepath.Join(home, "tmp", "export.json"),
		},
		{
			name:   "absolute path",
			config: ConfigType{ExportPath: "/tmp/export.json"},
			want:   "/tmp/export.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleExportPath(&tt.config)
			assert.Equal(t, tt.want, tt.config.ExportPath)
		})
	}
}
