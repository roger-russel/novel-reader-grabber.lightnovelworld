package epub

import (
	"fmt"
	"sort"

	"github.com/bmaupin/go-epub"

	"github.com/gosimple/slug"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
)

const ext = "epub"

func Write(pw progress.Writer, n *novel.Novel, dir string) {
	trkVol := &progress.Tracker{
		Total:   int64(len(n.Volumes)),
		Message: fmt.Sprintf("Volumes: %d", len(n.Volumes)),
		Units:   progress.UnitsDefault,
	}

	pw.AppendTracker(trkVol)

	defer trkVol.MarkAsDone()

	if len(n.Volumes) < 1 {
		fullFileName := dir + n.Slug + "." + ext
		writeVolume(pw, fullFileName, n.Title, n.Cover, n.Author, &n.Chapters)
		trkVol.Increment(1)
		return
	}

	for i, v := range n.Volumes {
		trkVol.Increment(1)
		if len(*v.Chapters) < 1 {
			continue
		}

		title := fmt.Sprintf("%s-volume-%d-%s", n.Title, i, v.Title)

		slugTitle := slug.Make(title)
		fullFileName := dir + slugTitle + "." + ext

		writeVolume(pw, fullFileName, title, n.Cover, n.Author, v.Chapters)
	}
}

func writeVolume(pw progress.Writer, fullFileName, title, cover string, author string, chapters *novel.Chapters) {
	total := len(*chapters)
	trkChapter := &progress.Tracker{
		Total:   int64(total),
		Message: fmt.Sprintf("Volume %s - Chapters: %d", title, total),
		Units:   progress.UnitsDefault,
	}

	pw.AppendTracker(trkChapter)
	defer trkChapter.MarkAsDone()

	e := epub.NewEpub(title)

	e.SetCover(cover, "")
	e.SetLang("en")

	sort.SliceStable(*chapters, func(i, j int) bool {
		return (*chapters)[i].Number < (*chapters)[j].Number
	})

	for i, chapter := range *chapters {
		_, err := e.AddSection(chapter.Content, chapter.Title, fmt.Sprintf("%d-%s", i, chapter.Title), "")
		trkChapter.Increment(1)
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
