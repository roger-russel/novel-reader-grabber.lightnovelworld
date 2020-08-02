package localstorage

import (
	"os"
)

const rootFolder string = ".novel-grabber"

var fullPathRootFolder string

func init() {
	fullPathRootFolder = os.Getenv("HOME") + "/" + rootFolder
}
