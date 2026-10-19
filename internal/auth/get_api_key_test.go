package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{headers: http.Header{"Authorization": {"ApiKey test-key"}}, wantKey: "test-key"},
		{headers: http.Header{}, wantErr: true},
		{headers: http.Header{"Authorization": {"ApiKey"}}, wantErr: true},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.headers)
		if (err != nil) != tc.wantErr {
			t.Fatalf("expected error: %v, got: %v", tc.wantErr, err)
		}
		if got != tc.wantKey {
			t.Fatalf("expected key: %q, got: %q", tc.wantKey, got)
		}
	}
}
