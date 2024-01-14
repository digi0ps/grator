package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`%([a-zA-Z0-9_-]+)%`)

func ParseTemplate(target string, data map[string]interface{}) (string, error) {
	matches := regex.FindAllStringSubmatch(target, -1)

	result := target

	for _, matchGroup := range matches {
		match := matchGroup[0]
		key := matchGroup[1]
		if value, ok := data[key]; ok {
			result = strings.Replace(result, match, fmt.Sprintf("%v", value), -1)
		} else {
			return "", fmt.Errorf("key %s not found", key)
		}
	}

	return result, nil
}
