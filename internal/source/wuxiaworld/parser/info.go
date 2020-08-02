package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Info get novel info
func Info(doc *goquery.Document) (title, author string) {

	info := doc.Find(".novel-index")

	title = info.Find("h2").Text()

	info.Find("dd").Each(func(index int, s *goquery.Selection) {
		if index == 1 {
			author = s.Text()
		}
	})

	return strings.TrimSpace(title),
		strings.TrimSpace(author)

}
