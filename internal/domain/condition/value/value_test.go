package value

import (
	"testing"

	"github.com/BranDebs/challenge-bot/internal/domain/condition/kind"
)

func TestValue_Valid(t *testing.T) {
	type args struct {
		k kind.Kind
	}
	tests := []struct {
		name string
		v    Value
		args args
		want bool
	}{
		{
			name: "unknown value kind",
			v:    "",
			args: args{
				k: kind.Unknown,
			},
			want: false,
		},
		{
			name: "mismatched value kind",
			v:    "12",
			args: args{
				k: kind.Boolean,
			},
			want: false,
		},
		{
			name: "valid value kind",
			v:    "12.34",
			args: args{
				k: kind.Float,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Valid(tt.args.k); got != tt.want {
				t.Errorf("Value.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				v: Value(""),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "invalid boolean string",
			args: args{
				v: Value("hello"),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "false boolean string",
			args: args{
				v: Value("false"),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "true boolean string",
			args: args{
				v: Value("true"),
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBool(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				v: Value(""),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid integer string",
			args: args{
				v: Value("hello"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "0 integer string",
			args: args{
				v: Value("0"),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "non 0 integer string",
			args: args{
				v: Value("17"),
			},
			want:    17,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				v: Value(""),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "invalid float string",
			args: args{
				v: Value("hello"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "0 float string",
			args: args{
				v: Value("0"),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "non 0 float string",
			args: args{
				v: Value("12.34"),
			},
			want:    12.34,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFloat(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
