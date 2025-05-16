package utils

import (
	"encoding/json"
	"gorm.io/datatypes"
	"reflect"
	"sort"
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

func ParseToJSON(value interface{}) datatypes.JSON {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil
	}

	return datatypes.JSON(bytes)
}

func ToStringSlice(input interface{}) ([]string, bool) {
	arr, ok := input.([]interface{})
	if !ok {
		return nil, false
	}
	result := make([]string, len(arr))
	for i, val := range arr {
		str, ok := val.(string)
		if !ok {
			return nil, false
		}
		result[i] = str
	}
	return result, true
}

func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}
