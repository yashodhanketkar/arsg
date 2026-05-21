package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/yashodhanketkar/arsg/src/util"
)

var (
	// INFO: Will replace util.DefaultParams in future updates
	apiParams         = []string{"Art", "Cast", "Plot", "Bias"}
	validContentTypes = map[string]bool{"anime": true, "manga": true, "lightnovel": true}
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

// check if content type is valid
func validateContentType(content_type string) (string, error) {
	if content_type == "" {
		return "", errors.New("No content type provided")
	}

	if !validContentTypes[content_type] {
		return "", errors.New("Invalid content type: " + content_type)
	}

	return content_type, nil
}

func autoCleanBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer func() {
				io.Copy(io.Discard, r.Body)
				r.Body.Close()
			}()

			next.ServeHTTP(w, r)
		}
	})
}
