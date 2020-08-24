package mobi

import (
	"fmt"
	"sort"

	"github.com/766b/mobi"
	"github.com/gosimple/slug"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
)

//Write mobi file
func Write(n *novel.Novel, dir string) {

	if len(n.Volumes) < 1 {

		fullFileName := dir + n.Slug + ".mobi"
		writeVolume(fullFileName, n.Title, n.Cover, n.Author, &n.Chapters)

	} else {

		for _, v := range n.Volumes {

			if len(*v.Chapters) < 1 {
				continue
			}

			title := n.Title + " - " + v.Title

			slugTitle := slug.Make(title)
			fullFileName := dir + slugTitle + ".mobi"

			writeVolume(fullFileName, title, n.Cover, n.Author, v.Chapters)

		}

	}

}

func writeVolume(fullFileName, title, cover string, author string, chapters *novel.Chapters) {

	m, err := mobi.NewWriter(fullFileName)
	helpers.Must(err)

	m.Title(title)
	//m.Compression(mobi.CompressionNone)
	m.Compression(mobi.CompressionPalmDoc)
	//ToDO add thumbnail instead of second cover

	m.AddCover(cover, cover)

	m.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	m.NewExthRecord(mobi.EXTH_AUTHOR, author)

	sort.SliceStable(*chapters, func(i, j int) bool {
		return (*chapters)[i].Number < (*chapters)[j].Number
	})

	for _, chapter := range *chapters {
		m.NewChapter(chapter.Title, []byte(chapter.Content))
	}

	m.Write()
	fmt.Printf("Writen book at: %v\n", fullFileName)
}
