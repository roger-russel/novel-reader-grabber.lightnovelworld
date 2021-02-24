package epub

import (
	"fmt"
	"sort"

	"github.com/bmaupin/go-epub"

	"github.com/gosimple/slug"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
)

const ext = "epub"

func Write(n *novel.Novel, dir string) {

	if len(n.Volumes) < 1 {

		fullFileName := dir + n.Slug + "." + ext

		writeVolume(fullFileName, n.Title, n.Cover, n.Author, &n.Chapters)

	} else {

		for _, v := range n.Volumes {

			if len(*v.Chapters) < 1 {
				continue
			}

			title := n.Title + " - " + v.Title

			slugTitle := slug.Make(title)
			fullFileName := dir + slugTitle + "." + ext

			writeVolume(fullFileName, title, n.Cover, n.Author, v.Chapters)

		}
	}
}

func writeVolume(fullFileName, title, cover string, author string, chapters *novel.Chapters) {

	e := epub.NewEpub(title)

	e.SetCover(cover, "")
	e.SetLang("en")

	sort.SliceStable(*chapters, func(i, j int) bool {
		return (*chapters)[i].Number < (*chapters)[j].Number
	})

	for _, chapter := range *chapters {
		_, err := e.AddSection(chapter.Content, chapter.Title, chapter.Title, "")
		if err != nil {
			fmt.Printf("Error adding chapter: %s, due error: %v", chapter.Title, err)
		}
	}

	fmt.Printf("Writing book: %v\n", fullFileName)
	err := e.Write(fullFileName)

	if err != nil {
		fmt.Printf("Fail writing book: %s, due error: %v", fullFileName, err)
	}

}
