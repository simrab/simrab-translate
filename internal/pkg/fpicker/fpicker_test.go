package fpicker

import (
	"testing"
)

func TestGetFilesWithTranslations(t *testing.T) {
	results := GetFilesWithTranslations("./", ".json")
	if results == nil {
		t.Error("Test GetFilesWithTranslations FAILED: no files returned")
		return
	}
	if len(results) == 1 && results[0] == "en.json" {
		t.Logf("Test GetFilesWithTranslations PASSED: %v", results[0])
	}
}
