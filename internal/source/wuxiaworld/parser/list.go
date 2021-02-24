package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

//ChaptersList generator the list of chapters
func ChaptersList(doc *goquery.Document) (volumes novel.Volumes) {

	root := doc.Find(".panel-group")

	root.Find(".panel.panel-default").Each(
		func(i int, v *goquery.Selection) {
			var number int
			var title string

			sPanel := v.Find(".panel-title").First()
			sNumber := sPanel.Find(".book").First()

			if sNumber != nil {
				var err error

				number, err = strconv.Atoi(sNumber.Text())

				if err != nil {
					log.Warning(err)
				}

			} else {
				number = i
			}

			sTitle := sPanel.Find(".title a").First()

			if sTitle != nil {
				title = sTitle.Text()
			}

			volume := novel.Volume{
				Title:  strings.TrimSpace(title),
				Number: number,
			}

			chapters := &novel.Chapters{}
			v.Find(".chapter-item a").Each(func(n int, c *goquery.Selection) {
				parseChapters(chapters, n, c)
			})

			volume.Chapters = chapters

			volumes = append(volumes, volume)

		},
	)

	return volumes

}

func parseChapters(chapters *novel.Chapters, n int, c *goquery.Selection) {

	var err error

	title := strings.TrimSpace(c.Text())

	url, found := c.Attr("href")

	if !found {
		log.WithFields(log.Fields{
			"title": title,
		}).Warning(fmt.Sprintf("Could not found the url from chapter \"%v\"", url))
	}

	sNumber, number, err := normalizer(url)

	if err != nil {
		log.WithFields(log.Fields{
			"title": title,
			"url":   url,
		}).Warning(fmt.Sprintf("Chapter with title \"%v\" has an invalid url to extract number: %v", title, url))
		return
	}

	*chapters = append(*chapters, novel.Chapter{
		Number:         number,
		OriginalNumber: sNumber,
		Title:          title,
		URL:            url,
		Updated:        "n/a",
	})
}

func normalizer(url string) (string, float32, error) {
	var sNumber string

	re := regexp.MustCompile("([0-9]+)(-([0-9]+))?$")
	m := re.FindStringSubmatch(url)

	if m == nil {
		return "", 0, fmt.Errorf("Could not parse chapter number from url: %v", url)
	}

	sNumber = m[1]

	if m[3] != "" {
		sNumber += "." + m[3]
	}

	number, err := strconv.ParseFloat(m[1], 32)

	return sNumber, float32(number), err

}
