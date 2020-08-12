package parser

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/tests/_fixtures/helpers"
)

func TestChaptersList(t *testing.T) {
	type args struct {
		doc *goquery.Document
	}
	tests := []struct {
		name               string
		args               args
		wantChaptersNumber int
	}{
		{
			name: "get a real novel list",
			args: args{
				doc: helpers.GetFixtureDoc("wuxiaworld", "info-sample-list.html"),
			},
			wantChaptersNumber: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChapters := ChaptersList(tt.args.doc)

			if len(gotChapters) != tt.wantChaptersNumber {
				t.Errorf("ChaptersList() gotChapters number = %v, want %v", gotChapters, tt.wantChaptersNumber)
			}
		})
	}
}
