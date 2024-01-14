package utils

import (
	"fmt"
	"strings"
)

func ExtractValue(data map[string]interface{}, search string) (interface{}, error) {
	keys := strings.Split(search, ".")
	if len(keys) == 1 {
		return data[search], nil
	}

	target := keys[len(keys)-1]
	keys = keys[:len(keys)-1]

	var current = data
	for _, key := range keys {
		if val, ok := current[key].(map[string]interface{}); ok {
			current = val
		} else {
			return nil, fmt.Errorf("Key %s not found", key)
		}
	}

	if val, ok := current[target]; ok {
		return val, nil
	} else {
		return nil, fmt.Errorf("Key %s not found", search)
	}
}
