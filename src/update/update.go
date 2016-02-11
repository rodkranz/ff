package update

import (
	"net/http"
	"log"
	"io/ioutil"
	"regexp"
)

type Update struct {
	Version string
	Url     string
}

func NewUpdate(url string, version string) (string, bool) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	tagLink := "href=\"/rodkranz/ff/releases/tag/"
	reg := regexp.MustCompile(tagLink + "(.*)\"")
	tag := reg.Find(bytes)

	if len(tag) == 0 {
		return "", false
	}

	tag = tag[len(tagLink):len(tag)-1]
	tagString := string(tag)
	hasUpdate := (tagString != version)

	return tagString, hasUpdate
}