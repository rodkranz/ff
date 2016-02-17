package update

import (
	"net/http"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
)

type UpJob struct {
	Version string
	Url     string
	regex   string
}

func NewUpdate(version, url, regex string) *UpJob {
	return &UpJob{version, url, regex}
}

func (u *UpJob) Check() (string, bool) {
	res, err := http.Get(u.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(strings.Replace(u.regex, "[ANY]", `(.*)">`, -1))
	tag := reg.Find(bytes)

	if len(tag) == 0 {
		return "", false
	}

	tag = tag[len(u.regex):len(tag)-1]
	tagString := string(tag)
	hasUpdate := (tagString != u.Version)

	return tagString, hasUpdate
}