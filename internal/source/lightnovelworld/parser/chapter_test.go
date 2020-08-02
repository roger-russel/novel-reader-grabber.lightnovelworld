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
			name: "content",
			args: args{
				page: helpers.GetFixtureFile("lightnovelworld", "content.html"),
			},
			want: helpers.GetFixtureString("lightnovelworld", "content-parsed.html"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chapter(tt.args.page); got != tt.want {
				t.Errorf("Chapter() =\n%v,\n\n\nwant=\n%v", got, tt.want)
			}
		})
	}
}
