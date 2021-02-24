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
		name         string
		args         args
		wantTitle    string
		wantAuthor   string
		wantChapters int
	}{
		{
			name: "simple",
			args: args{
				doc: helpers.GetFixtureDoc("lightnovelworld", "2021-07-02-info.html"),
			},
			wantTitle:    "I Alone Level-Up (Solo Leveling)",
			wantAuthor:   "Chugong",
			wantChapters: 271,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTitle, gotAuthor, gotChapters := Info(tt.args.doc)
			if gotTitle != tt.wantTitle {
				t.Errorf("Info() gotTitle = %v, want %v", gotTitle, tt.wantTitle)
			}
			if gotAuthor != tt.wantAuthor {
				t.Errorf("Info() gotAuthor = %v, want %v", gotAuthor, tt.wantAuthor)
			}
			if gotChapters != tt.wantChapters {
				t.Errorf("Info() gotChapters = %d, want %d", gotChapters, tt.wantChapters)
			}
		})
	}
}
