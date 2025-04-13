package translate

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

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

func FakeTranslateText(targetLanguage, text string) (string, error) {
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
  // TODO Add this as optional to enable transaltion from google cloud
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
