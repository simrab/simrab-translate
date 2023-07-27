package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/translate"

	"golang.org/x/text/language"
)

// TODO: Organize functions into separate modules

func main() {
	files := getFilesWithTranslations("./", ".json")
	copyFiles(files, []string{"it", "fr", "de"})
	fakeTranslateText("it", "home")
}

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

func copyFiles(files []string, langs []string) {
	for _, file := range files {
		// Copy file names
		for _, lang := range langs {
			copyFile(file, file[0:strings.LastIndex(file, "-")+1]+lang+".json")
		}
	}
}

func copyFile(src string, dst string) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !sourceFileStat.Mode().IsRegular() {
		fmt.Println("not a regular file")
	}
	file, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Copy copyFile into the new file
	err = ioutil.WriteFile(dst, file, 0o644)
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
	vals, err := io.ReadAll(fileData)
	if err != nil {
		fmt.Println(err)
		return
	}
	iterateValues(vals)
}

func iterateValues(val []byte) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(val, &m); err != nil {
		fmt.Println(err)
		return
	}
	// for _, v := range m {
	// 	fmt.Print(v)
	// }
}

// func translateText(targetLanguage, text string) (string, error) {
// 	ctx := context.Background()

// 	lang, err := language.Parse(targetLanguage)
// 	if err != nil {
// 		return "", fmt.Errorf("language.Parse: %w", err)
// 	}

// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer client.Close()

// 	resp, err := client.Translate(ctx, []string{text}, lang, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("Translate: %w", err)
// 	}
// 	if len(resp) == 0 {
// 		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
// 	}
// 	fmt.Println(resp[0].Text)
// 	return resp[0].Text, nil
// }

func fakeTranslateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	fmt.Println(lang)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %w", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// resp, err := client.Translate(ctx, []string{text}, lang, nil)
	// if err != nil {
	// 	return "", fmt.Errorf("Translate: %w", err)
	// }
	// if len(resp) == 0 {
	// 	return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	// }
	// fmt.Println(resp[0].Text)
	return text, nil
}
