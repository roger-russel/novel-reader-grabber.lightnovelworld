package site

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"

	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/source/wuxiaworld/parser"
	"github.com/roger-russel/novel-grabber/pkg/localstorage"
)

//Cover get novel cover
func Cover(novelSlug string, doc *goquery.Document) (imgPath string) {

	var err error

	cover, found := parser.Cover(doc)

	if !found {
		log.Warningf("Fail parsing image cover url from source: %v, novel: %v, error: %v", SOURCE, novelSlug, err)
		return ""
	}

	img, err := helpers.Download(cover)

	if err != nil {
		log.Warningf("Fail download image cover to source: %v, novel: %v, error: %v", SOURCE, novelSlug, err)
		return ""
	}

	imgPath, err = localstorage.WriteCover(SOURCE, novelSlug, "jpg", img)

	if err != nil {
		log.Warningf("Fail writen into localstorage the cover from source: %v, novel: %v, error: %v", SOURCE, novelSlug, err)
	}

	return imgPath

}
