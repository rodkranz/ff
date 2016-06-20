package update

import (
	"net/http"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
)

type Config struct {
	Version string
	Url     string
	Regex   string
}

type Update struct {
	config Config
}

func NewUpdate(conf Config) *Update {
	return &Update{config: conf}
}

func (u *Update) Check() (string, string, bool) {
	res, err := http.Get(u.config.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	regexEnd := strings.Replace(u.config.Regex, "[ANY]", `(.*)`, -1)
	reg := regexp.MustCompile(regexEnd)
	tag := reg.Find(bytes)

	if len(tag) == 0 {
		return "", "", false
	}

	split := strings.Split(u.config.Regex, "[ANY]")

	if len(tag) < len(split[0]) {
		return "", "", false
	}

	tag        = tag[len(split[0]):len(tag) - 1]
	tagString := string(tag)

	return tagString, u.config.Version, (tagString != u.config.Version)
}

func (u *Update) Update() {
	//fmt.Printf("update!")
}