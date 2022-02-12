package challenge

import (
	"reflect"
	"testing"

	"github.com/BranDebs/challenge-bot/internal/domain/condition"
)

func TestNewConditionTypes(t *testing.T) {
	tests := []struct {
		name string

		data []byte

		want ConditionTypes
	}{
		{
			name: "empty condition types",
			data: nil,
			want: nil,
		}, {
			name: "invalid condition types",
			data: []byte(`{malformed stuff,`),
			want: nil,
		}, {
			name: "one valid condition type",
			data: []byte(`{"age": "integer"}`),
			want: ConditionTypes{
				"age": condition.Integer,
			},
		}, {
			name: "more valid condition type",
			data: []byte(`{"age": "integer", "money": "float", "did work": "boolean"}`),
			want: ConditionTypes{
				"age":      condition.Integer,
				"money":    condition.Float,
				"did work": condition.Boolean,
			},
		}, {
			name: "more valid condition type with some invalid types",
			data: []byte(`{"age": "integer", "money": "float", "did work": "boolean", "invalid": "blah blah"}`),
			want: ConditionTypes{
				"age":      condition.Integer,
				"money":    condition.Float,
				"did work": condition.Boolean,
				"invalid":  condition.Unknown,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := NewConditionTypes(tt.data)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want: %v got: %v", tt.want, got)
				return
			}
		})
	}
}
