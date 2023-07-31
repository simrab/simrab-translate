package main

import (
	"bitbucket.org/simrab/simrab-translate/internal/pkg/copy"
	"bitbucket.org/simrab/simrab-translate/internal/pkg/fpicker"
	"bitbucket.org/simrab/simrab-translate/internal/pkg/translate"
)

// TODO: Organize functions into separate modules

func main() {
	files := fpicker.GetFilesWithTranslations("./", ".json")
	copy.CopyFiles(files, []string{"de", "fr", "it"})
	translate.FakeTranslateText("it", "home")
}
