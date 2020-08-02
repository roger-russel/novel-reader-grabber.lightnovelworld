package normalizers

import (
	"fmt"
	"os"
	"testing"
)

func Test_normalizeDirFlag(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				dir: "",
			},
			want: "./",
		},
		{
			name: "do nothing",
			args: args{
				dir: "./",
			},
			want: "./",
		},
		{
			name: "add / at end",
			args: args{
				dir: ".",
			},
			want: "./",
		},
		{
			name: "resolve ~",
			args: args{
				dir: "~/",
			},
			want: func() string {
				return os.Getenv("HOME") + "/"
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeDirFlag(tt.args.dir); got != tt.want {
				t.Errorf("normalizeDirFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeFormatType(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		recover func(t *testing.T)
	}{
		{
			name: "simple",
			args: args{
				format: "mobi",
			},
			want:    "mobi",
			recover: func(t *testing.T) {},
		},
		{
			name: "lower",
			args: args{
				format: "MObi",
			},
			want:    "mobi",
			recover: func(t *testing.T) {},
		},
		{
			name: "panic",
			args: args{
				format: "panic",
			},
			want: "",
			recover: func(t *testing.T) {
				if r := recover(); r != nil {
					want := "Error format type not suported:"
					got := fmt.Sprint(r)[:31]
					if want != got {
						t.Errorf("Error: got %v, want: %v", got, want)
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.recover(t)
			if got := normalizeFormatType(tt.args.format); got != tt.want {
				t.Errorf("normalizeFormatType() = %v, want %v", got, tt.want)
			}
		})
	}
}
