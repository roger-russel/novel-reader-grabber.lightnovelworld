package site

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/pool"
	"github.com/roger-russel/novel-grabber/internal/source/lightnovelworld/parser"
	"github.com/roger-russel/novel-grabber/pkg/localstorage"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

const pullSize int = 2

//ChaptersList get the novel info Page
func ChaptersList(doc *goquery.Document) novel.Chapters {
	var chapters novel.Chapters

	next, cha := parser.ChaptersList(doc)
	chapters = append(chapters, cha...)

	for next != "" {
		infoPage, err := helpers.Download(URL + next)
		helpers.Must(err)

		doc, err := goquery.NewDocumentFromReader(infoPage)
		helpers.Must(err)

		next, cha = parser.ChaptersList(doc)
		chapters = append(chapters, cha...)
	}

	return chapters

}

//Chapter get Chapter
func Chapter(novelSlug string, originalNumber string, chapterURL string) string {

	baContent, found := localstorage.ReadChapter(SOURCE, novelSlug, originalNumber)

	content := string(baContent)

	if !found {

		page, err := helpers.Download(URL + chapterURL)

		if err != nil {
			log.Warningf("Error downloading: %v%v, error:%v", URL, chapterURL, err)
			return ""
		}

		content = parser.Chapter(page)
		content = strings.TrimSpace(content)
		err = localstorage.WriteChapter(SOURCE, novelSlug, originalNumber, content)

		if err != nil {
			log.Warningf("Error Writing Chapter: %v%v, error:%v", URL, chapterURL, err)
			return ""
		}
	}

	return content
}

type chChapter struct {
	Index   int
	Content string
}

//Chapters get all Chapter
func Chapters(n *novel.Novel) {

	ch := make(chan chChapter, 1)
	pl := pool.New(pullSize)

	go updateContent(ch, n)

	for index := range n.Chapters {

		pl.Add(
			func(index int, slug string, originalNumber string, URL string) func() {
				return func() {

					content := Chapter(slug, originalNumber, URL)

					//Slow get frequency to try to avoid error 429 from site
					time.Sleep(1 * time.Second)

					ch <- chChapter{
						Index:   index,
						Content: content,
					}
				}
			}(index, n.Slug, n.Chapters[index].OriginalNumber, n.Chapters[index].URL),
		)
	}

	pl.Run()

}

func updateContent(ch chan chChapter, n *novel.Novel) {
	for c := range ch {
		n.Chapters[c.Index].Content = c.Content
	}
}
