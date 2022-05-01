package entry

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestKind_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		name     string
		data     []byte
		wantKind Kind
		wantErr  bool
	}{
		{
			name:     "nil data",
			data:     nil,
			wantKind: Unknown,
			wantErr:  true,
		},
		{
			name:     "empty data",
			data:     []byte(""),
			wantKind: Unknown,
			wantErr:  true,
		},
		{
			name:     "malformed JSON",
			data:     []byte(`malformed"`),
			wantKind: Unknown,
			wantErr:  true,
		},
		{
			name:     "unknown condition kind",
			data:     []byte(`"unknown kind"`),
			wantKind: Unknown,
			wantErr:  false,
		},
		{
			name:     "integer condition kind",
			data:     []byte(`"integer"`),
			wantKind: Integer,
			wantErr:  false,
		},
		{
			name:     "float condition kind case insensitive",
			data:     []byte(`"fLoAt"`),
			wantKind: Float,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kind := Unknown
			if err := kind.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("Kind.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.wantKind.String(), kind.String())
		})
	}
}
