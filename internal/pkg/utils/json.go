package utils

import (
	"encoding/json"
	"gorm.io/datatypes"
)

func ParseJSONToMap(data []byte) map[string]interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}

	return result
}

func ParseMapToJSON(value map[string]interface{}) datatypes.JSON {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil
	}

	return datatypes.JSON(bytes)
}
