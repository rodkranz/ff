package terminal

import (
	"github.com/fatih/color"
	"strings"
	"fmt"
	"bufio"
	"os"
)

type Bash struct {
	White  func(a ...interface{}) string
	Yellow  func(a ...interface{}) string
	Blue    func(a ...interface{}) string
	Green   func(a ...interface{}) string
	HiWrite func(a ...interface{}) string
}

func NewBash() *Bash {
	return &Bash{
		White:   color.New(color.FgWhite).SprintFunc(),
		Yellow:  color.New(color.FgYellow).SprintFunc(),
		Green:   color.New(color.FgGreen).SprintFunc(),
		Blue:    color.New(color.FgBlue).SprintFunc(),
		HiWrite: color.New(color.FgBlack, color.BgGreen).SprintFunc(),
	}
}

func (b *Bash) ColorWord(word, line string) string {
	index := strings.Index(line, word)
	if index == -1 {
		return line
	}

	startLine := line[0:index]
	endLine := line[index + len(word):]

	return fmt.Sprintf("%v%v%v", startLine, b.HiWrite(word), endLine)
}

func (b *Bash) Confirm(question string) bool {
	reader      := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s", question);
		variable, _ := reader.ReadString('\n')

		if len(variable) > 0 {
			variable = strings.ToLower(variable[0:1])
		}

		switch variable {
		case "n" :
			return false
		case "y" :
			return true
		}
	}
}