package main

import (
    "os"
    "runtime"
    "fmt"
    "flag"
    "strings"
    "regexp"

    "github.com/fatih/color"

    "github.com/rodkranz/ff/core"
    "github.com/rodkranz/ff/update"
    "github.com/rodkranz/ff/terminal"
    "github.com/rodkranz/ff/layout"
    "github.com/rodkranz/ff/replace"
)

const (
    APP string = "FF"
    DESC string = "Find Text Or File."
    VER string = "1.2.0"
    GITHUB string = "https://github.com/rodkranz/ff/releases"
)

var conf = &configuration{}

func init() {
    // configuration of update
    conf.upConfig = update.Config{
        VER,
        GITHUB,
        `href=\"/rodkranz/ff/releases/tag/[ANY]\"`,
    }

    // General
    flag.IntVar(&conf.CPUNum, "cpu", runtime.NumCPU(), fmt.Sprintf("Number of CPU you have %d available", runtime.NumCPU()))
    flag.BoolVar(&conf.force, "force", false, "Replace all result without ask.")
    //flag.IntVar(&conf.reach, 	"r", 10, "Range around of the word")

    // FF Config
    flag.StringVar(&conf.ffConfig.Text, "t", "", "Text that you are looking for")
    flag.StringVar(&conf.ffConfig.FilterFile, "f", "", "Filter by file name")
    flag.StringVar(&conf.ffConfig.Directory, "d", "./", "Set a specify directory")
    flag.BoolVar(&conf.ffConfig.CaseInsensitive, "cis", false, "Search text case insensitive")

    // Replace Config
    flag.StringVar(&conf.rpConfig.Text, "replace", "", "Replace result to text")

    // Update Config
    flag.BoolVar(&conf.showVersion, "ver", false, "Show the version")
    flag.BoolVar(&conf.upCheck, "up", false, "Check update")

    // Output config
    flag.BoolVar(&color.NoColor, "-no-color", false, "Disable color output")

    // FF with regex
    var withRegex bool
    flag.BoolVar(&withRegex, "reg", false, "Search by this Regex")

    // exclude files or folder with this name by default.
    exclude := *flag.String("-exclude-dir", ".bzr,CVS,.git,.hg,.svn", "Exclude dir from reader")

    // Allow use the CPUs available
    runtime.GOMAXPROCS(conf.CPUNum)

    //New Output mode
    bash := terminal.NewBash()

    //Layout that will draw the contents
    conf.draw = layout.NewLayout(bash)

    // Show Usage Pretty
    flag.Usage = conf.draw.GetUsage(APP, DESC, VER)

    // Parse parameters from console.
    flag.Parse()

    if len(exclude) > 0 {
        conf.ffConfig.ExcludeFiles = strings.Split(exclude, ",")
    }

    if withRegex {
        // if has regex
        conf.ffConfig.Regex = regexp.MustCompile(conf.ffConfig.Text)
    }

    // if has replace
    conf.rpConfig.WithReplace = false
    if (len(conf.rpConfig.Text) != 0) {
        conf.rpConfig.WithReplace = true
    }

    if conf.CPUNum > runtime.NumCPU() {
        conf.CPUNum = runtime.NumCPU()
    }

    // Check if has parameters
    if flag.NArg() == 1 {
        if par := flag.Arg(0); par[0] != 45 {
            conf.ffConfig.Text = flag.Arg(0)
        }
    }
}

func main() {
    draw := conf.draw

    if conf.showVersion || conf.upCheck {

        // show the current version of program
        if conf.showVersion {
            draw.ShowVersion(APP, VER, GITHUB)
            os.Exit(0)
        }

        // create a instance of update
        up := update.NewUpdate(conf.upConfig)

        // Check if has update and inform layout to do what next.
        if draw.ShowUpdate(up.Check()) {
            // if has update.
            up.Update();
        }

        // finalize the application correctly
        os.Exit(0)
    }

    // Show wait text...
    draw.ShowWait()

    // create a virtual storage
    finder := ff.NewFinder(conf.ffConfig)

    // search the files/text
    finder.FindFiles(conf.ffConfig.Directory)
    finder.FindText(conf.ffConfig.Text)

    result, err := finder.GetResult()
    if err != nil {
        // Draw error
        draw.DrawError(err)
        // finalize the application with error.
        os.Exit(1)
    }

    draw.Clear()
    if conf.rpConfig.WithReplace {
        // Draw result with Replace
        rpl := replace.NewReplace(conf.rpConfig)
        draw.DrawWithReplace(*rpl, result, conf.force)
    } else {
        // Draw results
        draw.Draw(result)
    }

    // finalize the application correctly
    os.Exit(0)
}

type configuration struct {
    rpConfig    replace.Config
    ffConfig    ff.Config
    upConfig    update.Config
    draw        *layout.Layout

    force       bool
    showVersion bool
    upCheck     bool
    CPUNum      int
}
