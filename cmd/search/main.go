package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func main() {

	cfg := struct {
		Dir      string
		FileName string
	}{}

	var cmdSearch = &cobra.Command{
		Use:   "search [string to search]",
		Short: "Search anything in files and show in screen",
		Long: `Search text in files help you find what you need.
Simplify your task in replace and find texts`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("File", cfg.FileName)
			fmt.Println("Dir", cfg.Dir)

			err := filepath.Walk(cfg.Dir, search)
			if err != nil {
				fmt.Println("Err filepath.Walk:", err)
				os.Exit(1)
			}

			fmt.Println("Search for: " + strings.Join(args, " "))
		},
	}

	cmdSearch.Flags().StringVarP(&cfg.Dir, "dir", "d", ".", "Directory to search.")
	cmdSearch.Flags().StringVarP(&cfg.FileName, "file_name", "f", "", "Filter files by the name.")

	// replaceValue := cmdTimes.Flags().StringP("replace", "r", "", "Replace the result by text.")
	// replaceForced := cmdTimes.Flags().BoolP()P("replace-forced", "rf", false, "Replace all result without ask.")

	var rootCmd = &cobra.Command{Use: "ff"}
	rootCmd.AddCommand(cmdSearch)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}

func search(path string, info os.FileInfo, err error) error {



}
