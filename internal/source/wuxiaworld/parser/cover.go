package parser

import (
	"github.com/PuerkitoBio/goquery"
)

//Cover parse where the cover is
func Cover(doc *goquery.Document) (src string, found bool) {

	src, found = doc.Find(".novel-index img").First().Attr("src")

	return src, found
}
