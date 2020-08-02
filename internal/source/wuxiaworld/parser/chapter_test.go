package parser

import (
	"io"
	"testing"

	"github.com/roger-russel/novel-grabber/tests/_fixtures/helpers"
)

func TestChapter(t *testing.T) {
	type args struct {
		page io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sample",
			args: args{
				page: helpers.GetFixtureFile("wuxiaworld", "content-sample.html"),
			},
			want: helpers.GetFixtureString("wuxiaworld", "content-sample-parsed.html"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chapter(tt.args.page); got != tt.want {
				t.Errorf("Chapter() =\n%v\n\n\nwant=\n%v", got, tt.want)
			}
		})
	}
}
