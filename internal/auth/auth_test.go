package auth

import (
    "net/http"
    "testing"
    "errors"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name string
        headers http.Header
        expectedKey string
        expectedErr error
    }{
        {
            name: "No Authorization Header",
            headers: http.Header{},
            expectedKey: "",
            expectedErr: ErrNoAuthHeaderIncluded,
        },
        {
            name: "Malformed Authorization Header - No ApiKey Prefix",
            headers: http.Header{
                "Authorization": []string{"Bearer 123456"},
            },
            expectedKey: "",
            expectedErr: errors.New("malformed authorization header"),
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            apiKey, err := GetAPIKey(tt.headers)

            if apiKey != tt.expectedKey {
                t.Errorf("expected key %v, got %v", tt.expectedKey, apiKey)
            }

            if err != nil && tt.expectedErr != nil {
                if err.Error() != tt.expectedErr.Error() {
                    t.Errorf("expected key %v, got %v", tt.expectedErr, err)
                }
            } else if err != tt.expectedErr {
                t.Errorf("expected key %v, got %v", tt.expectedErr, err)
            }
        })
    }
}
