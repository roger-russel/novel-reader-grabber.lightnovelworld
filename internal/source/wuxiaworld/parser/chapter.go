package parser

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	log "github.com/sirupsen/logrus"
)

//Chapter parse the chapter content
func Chapter(page io.Reader) (content string) {

	doc, err := goquery.NewDocumentFromReader(page)

	helpers.Must(err)

	root := doc.Find("#chapter-content").First()

	if root != nil {

		root.Find("p").Each(func(i int, s *goquery.Selection) {
			c, err := s.Html()

			if err != nil {
				log.Warningf("Fail parsing wuxiaworld content, place: %v", s.Text())
			}

			c = strings.TrimSpace(c)

			if c != "" {
				content += "<p>" + c + "</p>"
			}

		})

		return strings.TrimSpace(content)
	}

	return content
}
