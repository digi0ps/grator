package utils

import (
	"testing"
)

func TestExtractValue(t *testing.T) {
	t.Run("should return value for a simple key", func(t *testing.T) {
		data := map[string]interface{}{"key": "value"}
		value, err := ExtractValue(data, "key")

		if err != nil {
			t.Errorf("Expected no error, got %s", err.Error())
		}

		if value != "value" {
			t.Errorf("Expected value to be %s, got %s", "value", value)
		}
	})

	t.Run("should return value for a nested key", func(t *testing.T) {
		data := map[string]interface{}{"key": map[string]interface{}{"nested": "value"}}
		value, err := ExtractValue(data, "key.nested")

		if err != nil {
			t.Errorf("Expected no error, got %s", err.Error())
			return
		}

		if value != "value" {
			t.Errorf("Expected value to be %s, got %s", "value", value)
			return
		}
	})
}
