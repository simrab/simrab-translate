/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"bitbucket.org/simrab/simrab-translate/internal/pkg/copy"
	"bitbucket.org/simrab/simrab-translate/internal/pkg/fpicker"
	"bitbucket.org/simrab/simrab-translate/internal/pkg/translate"
	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "This command will add translation for the specified languages",
	Long: `This command translate each file in json format with en in his name, and creates a translations file foreach language you specify.
		Add abbr languages names after the command, every language separeted by a space:
		ES: go run main.go translate it en`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("ERROR: You need to specify at least one language")
			return
		}
		files := fpicker.GetFilesWithTranslations("./", ".json")
		copy.CopyFiles(files, args)
		translate.FakeTranslateText("it", "home")
		defer fmt.Println("Translations created successfully")
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// translateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// translateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
