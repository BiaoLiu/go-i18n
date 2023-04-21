package main

import (
	"log"
	"path/filepath"

	"github.com/BiaoLiu/go-i18n/cmd/i18n/internal/generator"
	"github.com/BiaoLiu/go-i18n/cmd/i18n/internal/utils"
	"github.com/spf13/cobra"
)

var (
	output string
)

func init() {
	generatorCmd.Flags().StringVarP(&output, "output", "", "i18n", "output dir name; default i18n")
}

var generatorCmd = &cobra.Command{
	Use:   "gen $path",
	Short: "create a scaffold project",
	Long:  "create a scaffold project",
	Run:   runGeneratorCmd,
}

func runGeneratorCmd(_ *cobra.Command, args []string) {
	var dir string
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}
	if len(args) == 1 && utils.IsDirectory(args[0]) {
		dir = args[0]
	} else {
		dir = filepath.Dir(args[0])
	}
	err := generator.Extract(dir, output)
	if err != nil {
		log.Fatal(err)
	}
	err = generator.Generator(output)
	if err != nil {
		log.Fatal(err)
	}
}
