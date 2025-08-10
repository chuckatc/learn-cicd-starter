package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		want    string
	}{
		"valid header": {
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			want: "valid_api_key",
		},
		"missing header": {
			headers: http.Header{},
			want:    "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if err != nil && err != ErrNoAuthHeaderIncluded {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
