package json

import "encoding/json"

func ToJson(value interface{}) string {
	jsonStr, _ := json.Marshal(value)
	return string(jsonStr)
}
