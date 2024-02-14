package utils

import (
	"encoding/json"
)

func ToJson(data map[string]string) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
