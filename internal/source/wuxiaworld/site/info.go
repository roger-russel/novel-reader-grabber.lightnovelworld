package site

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/source/wuxiaworld/parser"
)

//InfoPage get the info page novel
func InfoPage(novelSlug string) io.Reader {
	infoPage, err := helpers.Download(URL + NOVEL + novelSlug)
	helpers.Must(err)
	return infoPage
}

//Info get novel information
func Info(doc *goquery.Document) (title, author string) {
	return parser.Info(doc)
}
