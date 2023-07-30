package format

import (
	"encoding/json"
	"fmt"
)

func iterateValues(val []byte, lang string) []byte {
	m := make(map[string]interface{})
	if err := json.Unmarshal(val, &m); err != nil {
		return nil
	}
	for i, v := range m {
		m[i] = fmt.Sprintf("%v-%v", v, lang)
		// fmt.Println(v)
	}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return jsonBytes
}
