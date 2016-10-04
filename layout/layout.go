package layout

import (
    "fmt"

    "github.com/rodkranz/ff/terminal"
    "github.com/rodkranz/ff/core"
    "github.com/rodkranz/ff/replace"
    "flag"
    "os"
    "path"
)

type Layout struct {
    bash *terminal.Bash
}

func NewLayout(b *terminal.Bash) *Layout {
    return &Layout{b}
}

func (l *Layout) GetUsage(app, desc, ver string) func() {
    return func() {
        fmt.Printf("NAME: \t %v - %v\n", l.bash.Blue(app), l.bash.White(desc))
        //fmt.Printf("%v", l.bash.Blue(fmt.Sprintf("\n\t\t### %v - %v ###\n\n", app, desc)))
        fmt.Printf("USAGE:\t %v %v\n", l.bash.Blue(path.Base(os.Args[0])), l.bash.White("[arguments...]"))
        fmt.Printf("VER:  \t %v\n", l.bash.Blue(ver))

        fmt.Println("COMMANDS:")
        flag.VisitAll(func(f *flag.Flag) {
            extra := (func(extra string) string {
                if len(extra) == 0 {
                    return ""
                }

                return fmt.Sprintf("[default: %v]", l.bash.White(extra))
            })(f.DefValue)

            space := "\t%v: \t %v %v \n"
            if len(f.Name) < 7 {
                space = "\t%v: \t\t %v %v \n"
            }else if len(f.Name) < 3 {
                space = "\t%v: +\t\t\t %v %v \n"
            }

            fmt.Printf(space, l.bash.Green(f.Name), l.bash.Yellow(f.Usage), extra)
        })
    }
}

func (l *Layout) Draw(elms []ff.Element) {
    for _, elm := range elms {
        if !elm.Enabled {
            continue
        }

        l.showFileName(elm)

        if len(elm.Line) > 0 {
            l.showComment(elm)
            l.newLine()
        }
    }
}

func (l *Layout) DrawWithReplace(rpl replace.Replace, elms []ff.Element, force bool) {
    for _, elm := range elms {
        if !elm.Enabled {
            continue
        }

        l.showFileName(elm)
        l.askToReplaceComment(elm, rpl, force)
        l.newLine()
    }
}

func (l *Layout) DrawError(err error) {

}

func (l *Layout) showFileName(elm ff.Element) {
    if len(elm.Line) == 0 {
        elm.File.Name()

        typeFile := "file"
        if elm.File.IsDir() {
            typeFile = "dir"
        }

        fmt.Printf("[%v] %v\n",
            l.bash.White(typeFile),
            l.bash.Yellow(elm.Path),
        )

        return
    }

    fmt.Printf("[%v] %v (lines: %v)\n",
        l.bash.White("file"),
        l.bash.Yellow(elm.Path),
        l.bash.Blue(len(elm.Line)),
    )
}

func (l *Layout) showComment(elm ff.Element) {
    for _, k := range elm.GetComment() {
        fmt.Printf("[%v] \t%v \n", l.bash.Blue(k), l.bash.ColorWord(elm.Text[k], elm.Line[k]))
    }
}

func (l *Layout) askToReplaceComment(elm ff.Element, rpl replace.Replace, force bool) {
    var lines []int
    for _, k := range elm.GetComment() {
        line := fmt.Sprintf("[%v] \t%v", l.bash.Blue(k), l.bash.ColorWord(elm.Text[k], elm.Line[k]))

        if force {
            fmt.Printf("%v \t [%v %v %v]\n",
                line,
                l.bash.Blue(elm.Text[k]),
                l.bash.Green("->"),
                l.bash.Blue(rpl.GetConf().Text),
            )

            lines = append(lines, k)
            continue
        }
        c := l.bash.Confirm(fmt.Sprintf("%v [Y/N]", line))
        if (c) {
            lines = append(lines, k)
        }
    }

    if len(lines) > 0 {
        rpl.ReplaceLines(elm, lines)
    }
}

func (l *Layout) Clear() {
    fmt.Printf("\r                                  \r")
}
func (l *Layout) ShowWait() {
    fmt.Printf("Searching please wait... \r")
}

func (l *Layout) newLine() {
    fmt.Printf("\n")
}

func (l *Layout) ShowVersion(ver string) {
    fmt.Printf("The current version of FF is %v\n", ver)
}

func (l *Layout) ShowUpdate(verNew, verOld, link string, has bool) bool {

    if (!has) {
        fmt.Printf("The application already updated with latest version.\n")
        return false
    }

    fmt.Printf(`
The [FF] application has update available.
   Version Current: %v
 Version Available: %v

Check Link: %s
`, verOld, verNew, link)


    //return l.bash.Confirm("Would you like to update? [Y/N]");
    return false
}
