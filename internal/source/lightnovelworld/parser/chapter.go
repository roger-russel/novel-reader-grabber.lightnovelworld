package parser

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
)

//Chapter parse the chapter content
func Chapter(page io.Reader) string {

	doc, err := goquery.NewDocumentFromReader(page)

	helpers.Must(err)

	content := doc.Find(".chapter-content").First()

	if content != nil {
		c, _ := content.Html()
		return strings.TrimSpace(c)
	}

	return ""
}
