package output

import (
	"fmt"

	"github.com/roger-russel/novel-grabber/pkg/mobi"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
)

//Writer is the output handler that will call the selected output format type
func Writer(n *novel.Novel, format string, dir string) {
	switch format {
	case "mobi":
		mobi.Write(n, dir)
	default:
		panic(fmt.Errorf("Unkonw output format: %v", format))
	}
}
