package lightnovelworld

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/source/lightnovelworld/site"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

//New Novel
func New(n *novel.Novel, novelSlug string) {
	infoPage := site.InfoPage(novelSlug)

	doc, err := goquery.NewDocumentFromReader(infoPage)
	helpers.Must(err)

	n.Slug = novelSlug
	n.Title, n.Author = site.Info(doc)
	log.Infof("Title: %v", n.Title)

	n.Cover = site.Cover(novelSlug, doc)
	log.Infof("Cover: %v", n.Cover)

	n.Chapters = site.ChaptersList(doc)
	log.Infof("Chapters found:  %v", len(n.Chapters))

	log.Info("Getting Chapters Content")
	site.Chapters(n)

}
