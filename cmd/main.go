package main

import (
	"bitbucket.org/simrab/project-trsimrab/pkg/copy"
	"bitbucket.org/simrab/project-trsimrab/pkg/fpicker"
	"bitbucket.org/simrab/project-trsimrab/pkg/translate"
)

// TODO: Organize functions into separate modules

func main() {
	files := fpicker.getFilesWithTranslations("./", ".json")
	copy.copyFiles(files, []string{"de", "fr", "it"})
	translate.fakeTranslateText("it", "home")
}
