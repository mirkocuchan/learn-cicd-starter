package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		shouldError bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 123456"},
			},
			expectedKey: "123456",
			shouldError: false,
		},
		{
			name:        "missing API key",
			headers:     http.Header{},
			expectedKey: "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}
