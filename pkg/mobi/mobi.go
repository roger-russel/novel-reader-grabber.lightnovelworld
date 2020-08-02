package mobi

import (
	"fmt"
	"sort"

	"github.com/766b/mobi"
	"github.com/roger-russel/novel-grabber/internal/helpers"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
)

//Write mobi file
func Write(n *novel.Novel, dir string) {

	fullFileName := dir + n.Slug + ".mobi"

	m, err := mobi.NewWriter(fullFileName)
	helpers.Must(err)

	m.Title(n.Title)
	m.Compression(mobi.CompressionNone)
	//ToDO add thumbnail instead of second cover

	m.AddCover(n.Cover, n.Cover)

	m.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	m.NewExthRecord(mobi.EXTH_AUTHOR, n.Author)

	sort.SliceStable(n.Chapters, func(i, j int) bool {
		return n.Chapters[i].Number < n.Chapters[j].Number
	})

	for _, chapter := range n.Chapters {
		m.NewChapter(chapter.Title, []byte(chapter.Content))
	}

	m.Write()
	fmt.Printf("Writen book at: %v\n", fullFileName)
}
