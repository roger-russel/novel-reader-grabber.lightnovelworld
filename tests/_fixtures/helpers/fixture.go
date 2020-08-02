package helpers

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/roger-russel/novel-grabber/internal/helpers"
)

//GetFixtureDoc on fixture folder
func GetFixtureDoc(source, file string) (doc *goquery.Document) {

	f := GetFixtureFile(source, file)

	doc, err := goquery.NewDocumentFromReader(f)
	helpers.Must(err)

	return doc
}

//GetFixtureFile on fixture folder
func GetFixtureFile(source, file string) (f io.Reader) {

	goPath := os.Getenv("GOPATH")
	module := "github.com/roger-russel/novel-grabber"

	f, err := os.Open(goPath + "/src/" + module + "/tests/_fixtures/" + source + "/" + file)

	helpers.Must(err)

	return f

}

//GetFixtureString return the content as string
func GetFixtureString(source, file string) string {
	contentBytes := GetFixtureBytes(source, file)
	return strings.TrimSpace(string(contentBytes))
}

//GetFixtureBytes return the content as bytes
func GetFixtureBytes(source, file string) (content []byte) {

	goPath := os.Getenv("GOPATH")
	module := "github.com/roger-russel/novel-grabber"

	content, err := ioutil.ReadFile(goPath + "/src/" + module + "/tests/_fixtures/" + source + "/" + file)

	helpers.Must(err)

	return content

}
