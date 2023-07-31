package main

import (
	"bitbucket.org/simrab/simrab-translate/internal/pkg/copy"
	"bitbucket.org/simrab/simrab-translate/internal/pkg/fpicker"
	"bitbucket.org/simrab/simrab-translate/internal/pkg/translate"
)

// TODO: Organize functions into separate modules

func main() {
	files := fpicker.getFilesWithTranslations("./", ".json")
	copy.copyFiles(files, []string{"de", "fr", "it"})
	translate.fakeTranslateText("it", "home")
}
