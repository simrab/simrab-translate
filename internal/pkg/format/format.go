package format

import (
	"encoding/json"
	"fmt"
)

func IterateValues(val []byte, lang string) []byte {
	m := make(map[string]interface{})
	if err := json.Unmarshal(val, &m); err != nil {
		return nil
	}
	for i, v := range m {
		m[i] = fmt.Sprintf("%v-%v", v, lang)
	}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return jsonBytes
}
