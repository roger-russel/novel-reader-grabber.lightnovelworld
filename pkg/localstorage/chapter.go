package localstorage

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

const chaptersFolder string = "/chapters"
const ext = ".html"

//ReadChapter content locally
//It logs if an error different from
//ToDO check if file is a dir or a file
func ReadChapter(source string, novelSlug string, chapterOriginalNumber string) (content []byte, found bool) {
	var err error
	filePath, err := getNovelPath(source, novelSlug)

	if err != nil {
		log.Warningf("Could not read source: %v, novel: %v, chapter: %v, error: %v", source, novelSlug, chapterOriginalNumber, err)
		return []byte{}, false
	}

	filePath += chaptersFolder

	content, err = ioutil.ReadFile(filePath + "/" + chapterOriginalNumber + ext)

	if err != nil && !os.IsNotExist(err) {
		log.Warningf("Could not read source: %v, novel: %v, chapter: %v, error: %v", source, novelSlug, chapterOriginalNumber, err)
	}

	found = true
	if err != nil {
		found = false
	}

	return content, found

}

//WriteChapter locally
func WriteChapter(source string, novelSlug string, chapterOriginalNumber string, content string) (err error) {

	filePath, err := getNovelPath(source, novelSlug)

	if err != nil {
		return err
	}

	filePath += chaptersFolder

	f, err := os.Create(filePath + "/" + chapterOriginalNumber + ext)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	return err

}
