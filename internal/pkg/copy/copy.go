package copy

import (
	"fmt"
	"os"
	"strings"

	"bitbucket.org/simrab/simrab-translate/internal/pkg/format"
)

func CopyFiles(files []string, langs []string) {
	for _, file := range files {
		// Copy file names
		for _, lang := range langs {
			copyFile(file, file[0:strings.LastIndex(file, "-")+1]+lang+".json", lang)
		}
	}
}

func copyFile(src string, dst string, lang string) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !sourceFileStat.Mode().IsRegular() {
		fmt.Println("not a regular file")
	}
	file, err := os.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := format.IterateValues(file, lang)
	if data == nil {
		return
	}
	// Copy copyFile into the new file
	err = ioutil.WriteFile(dst, data, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// (wip) Open and read data in original file
	fileData, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileData.Sync()
	defer fileData.Close()
}
