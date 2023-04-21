package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:     "i18n",
	Short:   "",
	Long:    "",
	Version: "1.0",
}

func init() {
	rootCmd.AddCommand(generatorCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
