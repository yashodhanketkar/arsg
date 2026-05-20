package api

import (
	"fmt"

	"github.com/yashodhanketkar/arsg/src/util"
)

// returns string value or empty string
func extractStringValue(m map[string]interface{}, key string) string {
	val, ok := m[key]

	if !ok {
		return ""
	}

	switch v := val.(type) {
	case string:
		return v
	case float32, float64:
		return fmt.Sprintf("%.1f", v)
	default:
		return ""
	}
}

// returns float32 value or 0.0
func extractFloatValue(m map[string]interface{}, key string) float32 {
	val, ok := m[key]

	if !ok {
		return 0.0
	}

	switch v := val.(type) {
	case string:
		return util.FloatParser(v)
	case float32:
		return v
	case float64:
		return float32(v)
	default:
		return 0.0
	}
}
