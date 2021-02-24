package site

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/source/lightnovelworld/parser"
)

//InfoPage get the info page novel
func InfoPage(novelSlug string) io.Reader {

	url := URL + NOVEL + novelSlug

	fmt.Println("URL:", url)

	infoPage, err := helpers.Download(url)
	helpers.Must(err)
	return infoPage
}

//Info get novel information
func Info(doc *goquery.Document) (title, author string, chapters int) {
	return parser.Info(doc)
}
