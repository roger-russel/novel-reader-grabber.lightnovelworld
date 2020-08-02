package localstorage

import (
	"bytes"
	"io"
	"os"
)

//WriteCover on disk
//coverExt type are jpg, png, gif ...
func WriteCover(source string, novelSlug string, coverExt string, data io.Reader) (imgPath string, err error) {

	folder, err := getNovelPath(source, novelSlug)

	if err != nil {
		return "", err
	}

	imgPath = folder + "/cover." + coverExt

	f, err := os.Create(imgPath)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	imgBytes := buf.Bytes()
	f.Write(imgBytes)
	f.Close()

	return imgPath, err
}
