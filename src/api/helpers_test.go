package api

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type trackingCloser struct {
	io.Reader
	closed bool
}

func (tc *trackingCloser) Close() error {
	tc.closed = true
	return nil
}

func TestExtractValue(t *testing.T) {
	val := map[string]interface{}{
		"string_key":     "string_val",
		"string_num_key": "10.0",
		"float32_key":    float32(43.542),
		"float64_key":    float64(343.3343),
		"int32_key":      int32(234),
		"empty_key":      "",
	}

	t.Run("should result in string extract", func(t *testing.T) {
		tests := []struct {
			name string
			key  string
			want string
		}{
			{
				name: "should return correct value - string",
				key:  "string_key",
				want: "string_val",
			},
			{
				name: "should return correct value - numeric string",
				key:  "string_num_key",
				want: "10.0",
			},
			{
				name: "should return correct value - float32",
				key:  "float32_key",
				want: "43.5",
			},
			{
				name: "should return correct value - float64",
				key:  "float64_key",
				want: "343.3",
			},
			{
				name: "should return empty string",
				key:  "empty_key",
				want: "",
			},
			{
				name: "should return empty string for int32",
				key:  "int32_key",
				want: "",
			},
			{
				name: "should return empty string for invalid key",
				key:  "invalid_key",
				want: "",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := extractStringValue(val, tt.key)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("should result in float32 extract", func(t *testing.T) {
		tests := []struct {
			name string
			key  string
			want float32
		}{
			{
				name: "should return correct value - string",
				key:  "string_key",
				want: 0.0,
			},
			{
				name: "should return correct value - numeric string",
				key:  "string_num_key",
				want: 10.0,
			},
			{
				name: "should return correct value - float32",
				key:  "float32_key",
				want: 43.542,
			},
			{
				name: "should return correct value - float64",
				key:  "float64_key",
				want: 343.3343,
			},
			{
				name: "should return 0.0 for empty key",
				key:  "empty_key",
				want: 0.0,
			},
			{
				name: "should return 0.0 for int32 key",
				key:  "int32_key",
				want: 0.0,
			},
			{
				name: "should return 0.0 for invalid key",
				key:  "invalid_key",
				want: 0.0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := extractFloatValue(val, tt.key)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}

func TestValidateContentType(t *testing.T) {
	t.Run("should result in valid content type", func(t *testing.T) {
		tests := []struct {
			name     string
			content  string
			expected string
		}{
			{
				name:     "should return valid content type",
				content:  "anime",
				expected: "anime",
			},
			{
				name:     "should return valid content type",
				content:  "manga",
				expected: "manga",
			},
			{
				name:     "should return valid content type",
				content:  "lightnovel",
				expected: "lightnovel",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := validateContentType(tt.content)
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			})
		}
	})

	t.Run("should result in invalid content type", func(t *testing.T) {
		tests := []struct {
			name        string
			content     string
			expectedErr string
		}{
			{
				name:        "should return invalid content type",
				content:     "movies",
				expectedErr: "Invalid content type: movies",
			},
			{
				name:        "should return invalid content type",
				content:     "",
				expectedErr: "No content type provided",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := validateContentType(tt.content)
				assert.Error(t, err)
				assert.Equal(t, "", got)
			})
		}
	})
}

func TestAutoCleanBody(t *testing.T) {
	handlerCalled := false
	nextHandler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handlerCalled = true
			w.WriteHeader(http.StatusOK)
		})

	middlewareToTest := autoCleanBody(nextHandler)

	t.Run("should clean and close request body automatically", func(t *testing.T) {
		bodyByes := []byte("some payload data")
		mockBody := &trackingCloser{Reader: bytes.NewReader(bodyByes)}

		req := httptest.NewRequest(http.MethodPost, "/", mockBody)
		w := httptest.NewRecorder()

		middlewareToTest.ServeHTTP(w, req)

		assert.True(t, handlerCalled, "Next handler was not called")
		assert.True(t, mockBody.closed, "Body was not closed")
	})
}
