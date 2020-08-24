package site

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/pool"
	"github.com/roger-russel/novel-grabber/internal/source/wuxiaworld/parser"
	"github.com/roger-russel/novel-grabber/pkg/localstorage"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

const pullSize int = 10

//VolumesList get the novel info Page
func VolumesList(doc *goquery.Document) (volumes novel.Volumes) {

	vol := parser.ChaptersList(doc)
	volumes = append(volumes, vol...)

	return volumes

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

	}

	content = strings.TrimSpace(content)

	localstorage.WriteChapter(SOURCE, novelSlug, originalNumber, content)

	return content

}

type chChapter struct {
	Index   int
	Content string
	Volume  int
}

//Chapters get all Chapter
func Chapters(n *novel.Novel) {

	ch := make(chan chChapter, 1)
	pl := pool.New(pullSize)

	go updateContent(ch, n)

	for volIndex, v := range n.Volumes {
		for chIndex := range *v.Chapters {
			pl.Add(
				func(volIndex, chIndex int, slug string, originalNumber string, URL string) func() {
					return func() {
						content := Chapter(slug, originalNumber, URL)
						log.Debug(pl.Status())
						ch <- chChapter{
							Index:   chIndex,
							Content: content,
							Volume:  volIndex,
						}
					}
				}(volIndex, chIndex, n.Slug, (*v.Chapters)[chIndex].OriginalNumber, (*v.Chapters)[chIndex].URL),
			)
		}
	}

	pl.Run()

}

func updateContent(ch chan chChapter, n *novel.Novel) {
	for c := range ch {
		(*n.Volumes[c.Volume].Chapters)[c.Index].Content = c.Content
	}
}
