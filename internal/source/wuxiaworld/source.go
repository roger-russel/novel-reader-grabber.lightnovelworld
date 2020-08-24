package wuxiaworld

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/source/wuxiaworld/site"
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
	log.Infof("Title: %v, Author: %v", n.Title, n.Author)

	n.Cover = site.Cover(novelSlug, doc)
	log.Infof("Cover: %v", n.Cover)

	n.Volumes = site.VolumesList(doc)
	log.Infof("Volumes found: %v", len(n.Volumes))

	for _, v := range n.Volumes {
		log.Debugf("%s, found: %v chapters", v.Title, len(*v.Chapters))
	}

	log.Infof("Getting Chapters Content")
	site.Chapters(n)
	log.Infof("Getting Chapters Content Done!")

}
