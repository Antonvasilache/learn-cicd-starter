package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			name:          "No authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: true,
		},
		{
			name: "Invalid format",
			headers: http.Header{
				"Authorization": []string{"Invalid format"},
			},
			expectedKey:   "",
			expectedError: true,
		},
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey key123"},
			},
			expectedKey:   "key123",
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := auth.GetAPIKey(tt.headers)
			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
			if key != tt.expectedKey {
				t.Errorf("Expected error: %s, got %s", tt.expectedKey, key)
			}
		})
	}
}
