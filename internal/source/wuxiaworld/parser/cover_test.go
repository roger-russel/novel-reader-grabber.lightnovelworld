package parser

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/tests/_fixtures/helpers"
)

func TestCover(t *testing.T) {
	type args struct {
		doc *goquery.Document
	}
	tests := []struct {
		name      string
		args      args
		wantSrc   string
		wantFound bool
	}{
		{
			name: "cover",
			args: args{
				doc: helpers.GetFixtureDoc("wuxiaworld", "info.html"),
			},
			wantSrc:   "https://cdn.wuxiaworld.com/images/covers/atg.jpg?ver=b6377e1043744b345c0bdf83557f8b89a8018e94",
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSrc, gotFound := Cover(tt.args.doc)
			if gotSrc != tt.wantSrc {
				t.Errorf("Cover() gotSrc = %v, want %v", gotSrc, tt.wantSrc)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Cover() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}

}
