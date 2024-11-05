package vietqr

import (
	"testing"
)

func Test_hashCrc(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "text data",
			args: args{
				"hello word",
			},
			want: "6646",
		},
		{
			name: "vietqr data",
			args: args{
				"00020101021238490010A000000727011900069704160105135790208QRIBFTTA530370454061200005802VN6304",
			},
			want: "9A71",
		},
		{
			name: "empty data",
			args: args{
				"",
			},
			want: "FFFF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashCrc(tt.args.s); got != tt.want {
				t.Errorf("hashCrc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validCrcContent(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid data",
			args: args{
				"hello word nocrc",
			},
			want: false,
		},
		{
			name: "valid data",
			args: args{
				"hello word6646",
			},
			want: true,
		},
		{
			name: "valid data lowercase",
			args: args{
				"hello word lowercasecade",
			},
			want: true,
		},
		{
			name: "empty valid data",
			args: args{
				"FFFF",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validCrcContent(tt.args.s); got != tt.want {
				t.Errorf("validCrcContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
