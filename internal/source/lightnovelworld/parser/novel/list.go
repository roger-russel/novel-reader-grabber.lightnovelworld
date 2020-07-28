package novel

//import "github.com/gosimple/slug"

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"novel-grabber/internal/helpers"
	"novel-grabber/pkg/structs/novel"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

//parseChaptersList generator the list of chapters
func parseChaptersList(page io.Reader) (nextURL string, chapters novel.Chapters) {

	doc, err := goquery.NewDocumentFromReader(page)

	helpers.Must(err)

	doc.Find(".chapter-list li").Each(func(i int, s *goquery.Selection) {

		var err error

		title := strings.TrimSpace(s.Find(".chapter-title").Text())

		sNumber := strings.TrimSpace(s.Find(".chapter-no").Text())
		number, err := strconv.Atoi(sNumber)

		if err != nil {
			log.WithFields(log.Fields{
				"title":  title,
				"number": sNumber,
			}).Warning(fmt.Sprintf("Chapter with title \"%v\" has an invalid numer: %v", title, sNumber))
		}

		updated, found := s.Find(".chapter-update").Attr("datetime")

		if !found {
			log.WithFields(log.Fields{
				"title":  title,
				"number": sNumber,
			}).Warning(fmt.Sprintf("Could not found the release date from chapter \"%v\"", sNumber))
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
			Number:  number,
			Title:   title,
			URL:     url,
			Updated: updated,
		})

	})

	nextPage := doc.Find(".PagedList-skipToNext a").First()

	if nextPage != nil {
		nextURL = nextPage.AttrOr("href", "")
	}

	return nextURL, chapters

}
