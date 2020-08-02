package parser

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/tests/_fixtures/helpers"
)

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
			name: "simple",
			args: args{
				doc: helpers.GetFixtureDoc("lightnovelworld", "info.html"),
			},
			wantTitle:  "I Alone Level-Up (Solo Leveling) Web Novel",
			wantAuthor: "Chugong",
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
