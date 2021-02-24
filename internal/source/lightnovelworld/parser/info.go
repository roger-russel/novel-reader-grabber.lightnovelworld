package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Info get novel info
func Info(doc *goquery.Document) (title, author string, chapters int) {

	info := doc.Find(".novel-info")

	title = info.Find(".novel-title").Text()
	author = info.Find(".property-item span[itemprop=author]").Text()

	chp := info.Find("div.header-stats span").First().Find("strong")
	chp.Find("i").Remove()
	chaptersStr := strings.TrimSpace(chp.Text())

	chapters, err := strconv.Atoi(chaptersStr)

	if err != nil {
		fmt.Printf("could not retrieve chapters number due error: %v", err)
	}

	return strings.TrimSpace(title),
		strings.TrimSpace(author),
		chapters

}
