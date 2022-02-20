package kind

import (
	"testing"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name string

		got string

		want Kind
	}{
		{
			name: "empty string",
			got:  "",
			want: Unknown,
		}, {
			name: "invalid kind",
			got:  "invalid",
			want: Unknown,
		}, {
			name: "caps insensitive",
			got:  "INTegER",
			want: Integer,
		}, {
			name: "boolean",
			got:  "boolean",
			want: Boolean,
		}, {
			name: "integer",
			got:  "integer",
			want: Integer,
		}, {
			name: "float",
			got:  "float",
			want: Float,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := FromString(tt.got)

			if tt.want != got {
				t.Fatalf("want: %v got: %v", tt.want, got)
				return
			}
		})
	}
}
