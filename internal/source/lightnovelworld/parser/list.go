package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

//ChaptersList generator the list of chapters
func ChaptersList(doc *goquery.Document) (nextURL string, chapters novel.Chapters) {

	doc.Find("ul.chapter-list li").Each(func(i int, s *goquery.Selection) {

		var err error

		title := strings.TrimSpace(s.Find(".chapter-title").Text())

		sNumber := strings.TrimSpace(s.Find(".chapter-no").Text())

		number, err := strconv.ParseFloat(sNumber, 32)

		if err != nil {
			log.WithFields(log.Fields{
				"title":  title,
				"number": sNumber,
			}).Warning(fmt.Sprintf("Chapter with title \"%v\" has an invalid number: %v", title, sNumber))
			return
		}

		updated, found := s.Find(".chapter-update").Attr("datetime")

		if !found {
			updated = "n/a"
		}

		url, found := s.Find("a").Attr("href")

		if !found {
			log.WithFields(log.Fields{
				"title":   title,
				"number":  sNumber,
				"updated": updated,
			}).Warning(fmt.Sprintf("Could not found the url from chapter \"%v\"", sNumber))
		}

		chapters = append(chapters, novel.Chapter{
			Number:         float32(number),
			OriginalNumber: sNumber,
			Title:          title,
			URL:            url,
			Updated:        updated,
		})

	})

	nextPage := doc.Find(".PagedList-skipToNext a").First()

	if nextPage != nil {
		nextURL = nextPage.AttrOr("href", "")
	}

	return nextURL, chapters

}
