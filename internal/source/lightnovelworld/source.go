package lightnovelworld

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/source/lightnovelworld/site"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

type infoDTO struct {
	Slug     string
	Title    string
	Author   string
	Chapters novel.Chapters
	Cover    string
}

func info(novelSlug string) (i infoDTO) {

	infoPage := site.InfoPage(novelSlug)
	doc, err := goquery.NewDocumentFromReader(infoPage)
	helpers.Must(err)

	var chap int

	i.Slug = novelSlug
	i.Title, i.Author, chap = site.Info(doc)
	fmt.Printf("Title: %s\nAuthor: %s\nChapters: %d\n", i.Title, i.Author, chap)

	i.Chapters = site.ChaptersList(doc)

	return i

}

//New Novel
func New(n *novel.Novel, novelSlug string) {

	dto := info(novelSlug)

	n.Slug = dto.Slug
	n.Title = dto.Title
	n.Author = dto.Author
	n.Cover = dto.Cover
	n.Chapters = dto.Chapters

	log.Info("Getting Chapters Content")
	site.Chapters(n)

}

//Info retrieves intormation about the novel
func Info(novelSlug string) {
	_ = info(novelSlug)
}
