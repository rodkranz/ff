package main

import (
	"fmt"
	"html/template"
	"bytes"
	"bufio"
	"log"
	"io/ioutil"
	"os/exec"
	//"os"
	"encoding/json"
	"os"
)

const sourcePath = "/home/rlopes/Codes/go/src/github.com/rodkranz/ff"

const tmpBash = "_build.sh"

const buildScript = `#!/usr/bin/env bash

cd "{{.SourcePath}}"
mkdir "{{.SourcePath}}/release"

env GOOS={{.GOOS}} GOARCH={{.GOARCH}} go build -o {{.FileName}} main.go

tar -zcf "{{.SourcePath}}/release/{{.GOOS}}_{{.GOARCH}}.tar.gz" {{.FileName}}

rm ff
echo "done"`

type Release struct {
	GOOS       string `json:"os"`
	GOARCH     string `json:"arch"`
	FileName   string `json:"filename"`
	SourcePath string
}

func main() {
	jsonBlob, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", sourcePath, ".buildrc"))
	if err != nil {
		fmt.Println("error:", err)
	}
	var releases []Release

	if err = json.Unmarshal(jsonBlob, &releases); err != nil {
		fmt.Println("error:", err)
	}

	for _, r := range releases {
		r.SourcePath = sourcePath

		Compiler(r)
	}
}

func Compiler(r Release) {
	bashFile := fmt.Sprintf("%s/%s", sourcePath, tmpBash)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	t := template.Must(template.New("buildScript").Parse(buildScript))
	if err := t.Execute(w, r); err != nil {
		log.Println("executing template:", err)
	}
	w.Flush()

	if err := ioutil.WriteFile(bashFile, b.Bytes(), 0777); err != nil {
		log.Println("executing template:", err)
	}

	_, err := exec.Command(bashFile).Output()
	if err != nil {
		log.Fatal(err)
	}

	os.Remove(bashFile)
	fmt.Printf("Compilation [arch=%s] for [os=%s] with [name=%s] finished\n", r.GOARCH, r.GOOS, r.FileName)
}