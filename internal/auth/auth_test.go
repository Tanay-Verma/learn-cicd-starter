package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name    string
		input   http.Header
		want    string
		wantErr bool
	}

	tests := []test{
		{
			name: "Correct Authorization header",
			input: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			want:    "abc123",
			wantErr: false,
		},
		{
			name:    "No Authorization header",
			input:   http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "Incomplete Authorization header",
			input: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Wrong Authorization header",
			input: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)

			if (err != nil) != tc.wantErr {
				t.Fatalf("GetAPIKey():\nwantErr: %v,\ngot: %v,", tc.wantErr, err)
			}
			if got != tc.want {
				t.Fatalf("GetAPIKey():\nwant: %v,\ngot: %v,", tc.want, got)
			}
		})
	}
}
