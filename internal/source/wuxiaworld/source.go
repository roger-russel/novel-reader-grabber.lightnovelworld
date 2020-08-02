package wuxiaworld

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/internal/source/lightnovelworld/site"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
)

//New Novel
func New(n *novel.Novel, novelSlug string) {
	infoPage := site.InfoPage(novelSlug)

	doc, err := goquery.NewDocumentFromReader(infoPage)
	helpers.Must(err)

	n.Slug = novelSlug
	n.Title, n.Author = site.Info(doc)
	fmt.Println("Title: ", n.Title)

	/*
		n.Cover = site.Cover(novelSlug, doc)
		fmt.Println("Cover: ", n.Cover)

		n.Chapters = site.ChaptersList(doc)
		fmt.Println("Chapters found: ", len(n.Chapters))

		fmt.Println("Getting Chapters Content")
		site.Chapters(n)
	*/
}
