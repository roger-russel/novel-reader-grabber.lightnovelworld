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
		wantVolumeNumber   int
		wantChaptersNumber int
	}{
		{
			name: "get a real novel list",
			args: args{
				doc: helpers.GetFixtureDoc("wuxiaworld", "info-sample-list.html"),
			},
			wantVolumeNumber:   2,
			wantChaptersNumber: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVolumes := ChaptersList(tt.args.doc)

			if len(gotVolumes) != tt.wantVolumeNumber {
				t.Errorf("ChaptersList() gotVolumes number = %v, want %v", gotVolumes, tt.wantVolumeNumber)
			}

			var gotChapters int

			for _, v := range gotVolumes {
				gotChapters += len(*v.Chapters)
			}

			if gotChapters != tt.wantChaptersNumber {
				t.Errorf("ChaptersList() gotChapters number = %v, want %v", gotChapters, tt.wantChaptersNumber)
			}
		})
	}
}

func Test_normalizer(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantFloat float32
		wantErr   bool
	}{
		{
			name: "zero",
			args: args{
				url: "/novel/against-the-gods/atg-chapter-0",
			},
			want:      "0",
			wantFloat: 0,
			wantErr:   false,
		},
		{
			name: "Devil number",
			args: args{
				url: "/novel/against-the-gods/atg-chapter-666",
			},
			want:      "666",
			wantFloat: 666,
			wantErr:   false,
		},
		{
			name: "Float Number",
			args: args{
				url: "/novel/against-the-gods/atg-chapter-917-05",
			},
			want:      "917.05",
			wantFloat: 917,
			wantErr:   false,
		},
		{
			name: "Error",
			args: args{
				url: "/novel/against-the-gods/atg-chapter-917s",
			},
			want:      "",
			wantFloat: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := normalizer(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("normalizer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("normalizer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantFloat {
				t.Errorf("normalizer() got1 = %v, want %v", got1, tt.wantFloat)
			}
		})
	}
}
