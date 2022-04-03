package operator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperator_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		data         []byte
		wantOperator Operator
		wantErr      bool
	}{
		{
			name:         "nil data",
			data:         nil,
			wantOperator: Unknown,
			wantErr:      true,
		},
		{
			name:         "empty data",
			data:         []byte(""),
			wantOperator: Unknown,
			wantErr:      true,
		},
		{
			name:         "malformed data",
			data:         []byte(`malformed"`),
			wantOperator: Unknown,
			wantErr:      true,
		},
		{
			name:         "unknown condition kind",
			data:         []byte(`"unknown kind"`),
			wantOperator: Unknown,
			wantErr:      false,
		},
		{
			name:         "equal operator",
			data:         []byte(`"eq"`),
			wantOperator: Equal,
			wantErr:      false,
		},
		{
			name:         "less than operator case insensitive",
			data:         []byte(`"lT"`),
			wantOperator: LessThan,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operator := Unknown
			if err := operator.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("Operator.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.wantOperator.String(), operator.String())
		})
	}
}
