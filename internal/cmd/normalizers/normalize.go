package normalizers

import (
	"fmt"
	"os"
	"strings"

	"github.com/roger-russel/novel-grabber/pkg/structs/cmd"
)

//NormalizeFlags normalize flags given
func NormalizeFlags(flags *cmd.Flags) {
	flags.Dir = normalizeDirFlag(flags.Dir)
	flags.FormatType = normalizeFormatType(flags.FormatType)
}

//normalizeDirFlag normalize dir flag
func normalizeDirFlag(dir string) string {

	if dir == "" {
		return "./"
	}

	if dir[:1] == "~" {
		dir = os.Getenv("HOME") + dir[1:]
	}

	if dir[len(dir)-1:] != "/" {
		dir += "/"
	}

	return dir

}

//normalizeFormatType
func normalizeFormatType(format string) string {
	ok := false
	allowed := []string{"mobi"}

	format = strings.ToLower(format)

	for _, a := range allowed {
		if format == a {
			ok = true
			break
		}
	}

	if ok != true {
		panic(
			fmt.Errorf(
				"Error format type not suported: %v, suported list: %v",
				format, allowed,
			),
		)
	}

	return format

}
