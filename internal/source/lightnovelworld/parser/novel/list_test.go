package novel

import (
	"io"
	"os"
	"testing"
)

func Test_parseChaptersList(t *testing.T) {

	rootPath := "../../../"

	type args struct {
		page io.Reader
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
				page: func() io.Reader {
					f, err := os.Open(rootPath + "tests/fixtures/info.html")
					if err != nil {
						panic(err)
					}
					return f
				}(),
			},
			wantNext:       "/novel/i-alone-level-up-solo-leveling-web-novel/2",
			wantListNumber: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNext, gotList := parseChaptersList(tt.args.page)

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
