package parser

import (
	"github.com/PuerkitoBio/goquery"
)

//Cover parse where the cover is
func Cover(doc *goquery.Document) (src string, found bool) {

	src, found = doc.Find("header.novel-header figure.cover img").First().Attr("data-src")

	return src, found
}
