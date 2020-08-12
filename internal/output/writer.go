package output

import (
	"fmt"

	"github.com/roger-russel/novel-grabber/pkg/mobi"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	log "github.com/sirupsen/logrus"
)

//Writer is the output handler that will call the selected output format type
func Writer(n *novel.Novel, format string, dir string) {
	switch format {
	case "mobi":
		log.Info("Writing mobi file")
		mobi.Write(n, dir)
	default:
		panic(fmt.Errorf("Unkonw output format: %v", format))
	}
}
