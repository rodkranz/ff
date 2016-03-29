package replace

import (
	"os"
	"bufio"
	"bytes"
	"strings"
	"io/ioutil"

	"github.com/rodkranz/ff/core"
)

type Config struct {
	WithReplace bool
	Text        string
	Force 		bool
}

type Replace struct {
	config Config
}

func NewReplace(conf Config) *Replace {
	return &Replace{config: conf}
}

func (r *Replace) GetConf() Config {
	return r.config
}

func (r *Replace) ReplaceLines(elm ff.Element, lines []int) error {
	file, err := os.Open(elm.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	var buffWriter bytes.Buffer

	numLine := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numLine++

		if contains(lines, numLine) {
			line = strings.Replace(line, elm.Text[numLine], r.config.Text, -1)
		}

		buffWriter.WriteString(line + "\n")
	}

	return ioutil.WriteFile(elm.Path, buffWriter.Bytes(), 0644)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}