package retrieve

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// MissingTranslationKeys represents the JSON structure for missing translation keys
type MissingTranslationKeys struct {
	MissingTranslationKeys map[string]bool `json:"missingTranslationKeys"`
}

// isTargetFile checks if the file has .ts or .html extension
func isTargetFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".ts" || ext == ".html"
}

func GetMissingKeys(searchDir, referenceFile, outputFile string) {
	// Compile the regex patterns for translation keys
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`'([^']*)'[ ]*\|[ ]*translate`),
		regexp.MustCompile(`translate[ ]*\([ ]*'([^']*)'[ ]*\)`),
	}

	// Map to store unique translation keys
	foundKeys := make(map[string]bool)

	// Walk through all files in the search directory recursively
	err := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, hidden files, and non-target files
		if info.IsDir() || strings.HasPrefix(filepath.Base(path), ".") || !isTargetFile(path) {
			return nil
		}

		// Open and scan file
		file, err := os.Open(path)
		if err != nil {
			return nil // Skip files we can't open
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// Check for all patterns
			for _, pattern := range patterns {
				matches := pattern.FindAllStringSubmatch(line, -1)
				for _, match := range matches {
					if len(match) > 1 {
						foundKeys[match[1]] = true
					}
				}
			}
		}

		return scanner.Err()
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d total translation keys in .ts and .html files.\n", len(foundKeys))

	// Read reference file
	referenceData, err := os.ReadFile(referenceFile)
	if err != nil {
		fmt.Printf("Error reading reference file: %v\n", err)
		os.Exit(1)
	}

	// Parse reference JSON
	var referenceJSON map[string]interface{}
	err = json.Unmarshal(referenceData, &referenceJSON)
	if err != nil {
		fmt.Printf("Error parsing reference JSON: %v\n", err)
		os.Exit(1)
	}

	// Extract all keys from the reference JSON
	referenceKeys := make(map[string]bool)
	extractKeys(referenceJSON, "", referenceKeys)

	// Find keys that are not in the reference file
	missingKeys := make(map[string]bool)
	for key := range foundKeys {
		if !referenceKeys[key] {
			missingKeys[key] = true
		}
	}

	fmt.Printf("Found %d translation keys that are not in the reference file.\n", len(missingKeys))

	// Create JSON structure for missing keys
	missingKeysStruct := MissingTranslationKeys{
		MissingTranslationKeys: missingKeys,
	}

	// Write missing keys to output file
	missingKeysJSON, err := json.MarshalIndent(missingKeysStruct, "", "  ")
	if err != nil {
		fmt.Printf("Error creating missing keys JSON: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFile, missingKeysJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Missing translation keys saved to %s\n", outputFile)
	fmt.Print("Processing completed.")
}

// extractKeys recursively extracts all keys from a nested JSON structure
func extractKeys(data interface{}, prefix string, result map[string]bool) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			newKey := key
			if prefix != "" {
				newKey = prefix + "." + key
			}
			
			// Add this key
			result[newKey] = true
			
			// Process nested structures
			extractKeys(value, newKey, result)
		}
	case []interface{}:
		for i, item := range v {
			newKey := fmt.Sprintf("%s[%d]", prefix, i)
			extractKeys(item, newKey, result)
		}
	default:
		// For primitive values, just add the key
		if prefix != "" {
			result[prefix] = true
		}
	}
}
