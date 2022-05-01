package condition

import (
	"reflect"
	"testing"

	"github.com/BranDebs/challenge-bot/internal/domain/condition/operator"
	"github.com/BranDebs/challenge-bot/internal/domain/value"
)

func TestFromJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []*Condition
		wantErr bool
	}{
		{
			name:    "nil data",
			data:    nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty data",
			data:    []byte(`""`),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "malformed data",
			data:    []byte(`[{"malformed}]`),
			want:    nil,
			wantErr: true,
		},
		{
			name: "one condition",
			data: []byte(`[{"name":"condition1","kind":"integer","value":"1337","operator":"lt"}]`),
			want: []*Condition{
				{
					Name:     "condition1",
					Kind:     value.Integer,
					Value:    value.Value("1337"),
					Operator: operator.LessThan,
				},
			},
			wantErr: false,
		},
		{
			name: "two conditions",
			data: []byte(`[{"name":"condition1","kind":"integer","value":"1337","operator":"lt"},{"name":"condition2","kind":"float","value":"1.337","operator":"eq"}]`),
			want: []*Condition{
				{
					Name:     "condition1",
					Kind:     value.Integer,
					Value:    value.Value("1337"),
					Operator: operator.LessThan,
				},
				{
					Name:     "condition2",
					Kind:     value.Float,
					Value:    value.Value("1.337"),
					Operator: operator.Equal,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromJSON(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
