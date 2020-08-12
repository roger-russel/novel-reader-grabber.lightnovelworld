package parser

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/tests/_fixtures/helpers"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func TestInfo(t *testing.T) {

	type args struct {
		doc *goquery.Document
	}
	tests := []struct {
		name       string
		args       args
		wantTitle  string
		wantAuthor string
	}{
		{
			name: "get author",
			args: args{
				doc: helpers.GetFixtureDoc("wuxiaworld", "info.html"),
			},
			wantTitle:  "Against the Gods",
			wantAuthor: "Mars Gravity",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTitle, gotAuthor := Info(tt.args.doc)
			if gotTitle != tt.wantTitle {
				t.Errorf("Info() gotTitle = %v, want %v", gotTitle, tt.wantTitle)
			}
			if gotAuthor != tt.wantAuthor {
				t.Errorf("Info() gotAuthor = %v, want %v", gotAuthor, tt.wantAuthor)
			}
		})
	}
}
