package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Info get novel info
func Info(doc *goquery.Document) (title, author string) {

	info := doc.Find(".novel-info")

	title = info.Find(".novel-title").Text()
	author = info.Find(".property-item").Text()

	return strings.TrimSpace(title),
		strings.TrimSpace(author)

}
