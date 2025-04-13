package retrieve

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		prefix   string
		expected map[string]bool
	}{
		{
			name:   "Simple map",
			input:  map[string]interface{}{"key1": "value1", "key2": "value2"},
			prefix: "",
			expected: map[string]bool{
				"key1": true,
				"key2": true,
			},
		},
		{
			name: "Nested map",
			input: map[string]interface{}{
				"parent": map[string]interface{}{
					"child": "value",
				},
			},
			prefix: "",
			expected: map[string]bool{
				"parent":        true,
				"parent.child": true,
			},
		},
		{
			name: "Array in map",
			input: map[string]interface{}{
				"array": []interface{}{"item1", "item2"},
			},
			prefix: "",
			expected: map[string]bool{
				"array":      true,
				"array[0]":   true,
				"array[1]":   true,
			},
		},
		{
			name: "With prefix",
			input: map[string]interface{}{
				"key": "value",
			},
			prefix: "test",
			expected: map[string]bool{
				"test.key": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make(map[string]bool)
			extractKeys(tt.input, tt.prefix, result)

			// Check if all expected keys are present
			for key := range tt.expected {
				if !result[key] {
					t.Errorf("Expected key %s not found in result", key)
				}
			}

			// Check if there are no unexpected keys
			for key := range result {
				if !tt.expected[key] {
					t.Errorf("Unexpected key %s found in result", key)
				}
			}
		})
	}
}

func TestMissingTranslationKeysStruct(t *testing.T) {
	// Test the JSON marshaling of MissingTranslationKeys struct
	missingKeys := map[string]bool{
		"key1": true,
		"key2": true,
	}

	missingKeysStruct := MissingTranslationKeys{
		MissingTranslationKeys: missingKeys,
	}

	jsonData, err := json.Marshal(missingKeysStruct)
	if err != nil {
		t.Fatalf("Failed to marshal MissingTranslationKeys: %v", err)
	}

	// Unmarshal back to verify structure
	var unmarshaled MissingTranslationKeys
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal MissingTranslationKeys: %v", err)
	}

	// Verify the keys
	if len(unmarshaled.MissingTranslationKeys) != len(missingKeys) {
		t.Errorf("Expected %d keys, got %d", len(missingKeys), len(unmarshaled.MissingTranslationKeys))
	}

	for key := range missingKeys {
		if !unmarshaled.MissingTranslationKeys[key] {
			t.Errorf("Expected key %s not found in unmarshaled data", key)
		}
	}
}

func TestGetMissingKeys(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-translate")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	testFiles := []struct {
		path    string
		content string
	}{
		{
			path:    filepath.Join(tempDir, "test.ts"),
			content: "{{ 'test' | translate }}\n{{ 'test2'|translate }}\ntranslate('new')",
		},
		{
			path:    filepath.Join(tempDir, "test.html"),
			content: "translate('value')",
		},
	}

	for _, tf := range testFiles {
		err := os.WriteFile(tf.path, []byte(tf.content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", tf.path, err)
		}
	}

	// Create reference JSON file
	referenceJSON := map[string]interface{}{
		"test":  "translation1",
		"test2": "translation2",
	}
	referenceData, err := json.Marshal(referenceJSON)
	if err != nil {
		t.Fatalf("Failed to marshal reference JSON: %v", err)
	}
	referencePath := filepath.Join(tempDir, "reference.json")
	err = os.WriteFile(referencePath, referenceData, 0644)
	if err != nil {
		t.Fatalf("Failed to write reference file: %v", err)
	}

	// Create output file path
	outputPath := filepath.Join(tempDir, "missing.json")

	// Run GetMissingKeys
	GetMissingKeys(tempDir, referencePath, outputPath)

	// Read and verify output
	outputData, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var result MissingTranslationKeys
	err = json.Unmarshal(outputData, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal output: %v", err)
	}

	// Verify expected missing keys
	expectedMissing := map[string]bool{
		"new":   true,
		"value": true,
	}

	if len(result.MissingTranslationKeys) != len(expectedMissing) {
		t.Errorf("Expected %d missing keys, got %d", len(expectedMissing), len(result.MissingTranslationKeys))
	}

	for key := range expectedMissing {
		if !result.MissingTranslationKeys[key] {
			t.Errorf("Expected missing key %s not found in result", key)
		}
	}
} 