package fpicker

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func getFilesWithTranslations(root string, ext string) []string {
	var files []string
	// Get all files with .en extension
	e := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println(e)
			return e
		}
		exts := filepath.Ext(d.Name())
		if exts == ext && strings.Contains(d.Name(), "en") {
			files = append(files, s)
		}
		return nil
	})

	if e != nil {
		fmt.Println(e)
	}

	return files
}
