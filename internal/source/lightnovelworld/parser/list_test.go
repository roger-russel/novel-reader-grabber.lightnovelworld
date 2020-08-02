package parser

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/tests/_fixtures/helpers"
)

func TestChaptersList(t *testing.T) {

	type args struct {
		page *goquery.Document
	}

	tests := []struct {
		name           string
		args           args
		wantNext       string
		wantListNumber int
	}{
		{
			name: "get a real novel list",
			args: args{
				page: helpers.GetFixtureDoc("lightnovelworld", "info.html"),
			},
			wantNext:       "/novel/i-alone-level-up-solo-leveling-web-novel/2",
			wantListNumber: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNext, gotList := ChaptersList(tt.args.page)

			if tt.wantNext != gotNext {
				t.Errorf("Expected Next: %v got %v", tt.wantNext, gotNext)
			}

			gotListNumber := len(gotList)

			if tt.wantListNumber != gotListNumber {
				t.Errorf("Expected number of chapters: %v got %v", tt.wantListNumber, gotListNumber)
			}
		})
	}
}
