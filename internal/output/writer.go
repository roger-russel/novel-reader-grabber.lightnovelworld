package output

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/roger-russel/novel-grabber/pkg/epub"
	"github.com/roger-russel/novel-grabber/pkg/mobi"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

//Writer is the output handler that will call the selected output format type
func Writer(n *novel.Novel, format string, dir string) {
	pw := progressWriter()
	go pw.Render()
	go func() {
		for pw.IsRenderInProgress() {
			// for manual-stop mode, stop when there are no more active trackers
			if pw.LengthActive() == 0 {
				pw.Stop()
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	switch format {
	case "mobi":
		log.Info("Writing mobi file")
		mobi.Write(n, dir)

	case "epub":
		log.Info("Writing epub file")
		epub.Write(pw, n, dir)

	default:
		panic(fmt.Errorf("Unkonw output format: %v", format))
	}
}

func progressWriter() progress.Writer {
	pw := progress.NewWriter()
	pw.SetAutoStop(true)
	pw.SetTrackerLength(25)
	pw.SetNumTrackersExpected(2)
	pw.ShowOverallTracker(true)
	pw.ShowTime(true)
	pw.ShowTracker(true)
	pw.ShowValue(true)
	pw.SetMessageWidth(24)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Colors = progress.StyleColorsExample
	pw.Style().Options.PercentFormat = "%4.1f%%"

	return pw
}
