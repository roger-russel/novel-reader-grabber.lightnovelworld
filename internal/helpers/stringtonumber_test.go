package helpers

import "testing"

func TestStringToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "working",
			args: args{
				s: "100",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "working with multiples",
			args: args{
				s: "100-101",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "working",
			args: args{
				s: "100-101",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "faling",
			args: args{
				s: "a-101",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToNumber(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
