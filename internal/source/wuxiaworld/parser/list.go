package parser

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

//ChaptersList generator the list of chapters
func ChaptersList(doc *goquery.Document) (chapters novel.Chapters) {

	root := doc.Find(".panel-group")
	root.Find(".panel.panel-default .chapter-item a").Each(
		func(number int, s *goquery.Selection) {

			var err error

			title := strings.TrimSpace(s.Text())

			sNumber := title

			if err != nil {
				log.WithFields(log.Fields{
					"title":  title,
					"number": number,
				}).Warning(fmt.Sprintf("Chapter with title \"%v\" has an invalid number: %v", title, sNumber))
				return
			}

			url, found := s.Attr("href")

			if !found {
				log.WithFields(log.Fields{
					"title":  title,
					"number": sNumber,
				}).Warning(fmt.Sprintf("Could not found the url from chapter \"%v\"", sNumber))
			}

			chapters = append(chapters, novel.Chapter{
				Number:         number,
				OriginalNumber: normalizer(title),
				Title:          title,
				URL:            url,
				Updated:        "n/a",
			})
		},
	)

	return chapters

}

func normalizer(title string) string {
	title = slug.Make(title)

	if title == "prologue" {
		title = "chapter-0-prologue"
	}

	return title

}
