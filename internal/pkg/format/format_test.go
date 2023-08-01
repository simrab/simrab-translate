package format

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestIterateValues(t *testing.T) {
	file, err := ioutil.ReadFile("./mockTranslate.en.json")
	if err != nil {
		t.Errorf("can't find mock file")
		return
	}

	result := IterateValues(file, "it")
	if result == nil {
		t.Errorf("Test IterateValues FAILED: No returned file")
		return
	}
	if fmt.Sprintf("%T", result) != "[]uint8" {
		t.Errorf("%T is not a valid type, []byte is the correct type", result)
		return
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(result, &m); err != nil {
		t.Errorf("Test IterateValues FAILED: Error while Unmarshal")
		return
	}
	if m["test"] == "this is a test-it" {
		t.Logf("IterateValuesValueCheck PASSED: Value is as expected: %v", m["test"])
	}
}
