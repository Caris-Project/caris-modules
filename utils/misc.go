package utils

import (
	"encoding/json"

	"gorm.io/datatypes"
)

// ConvertToMap ...
func ConvertToMap(data interface{}) (result map[string]interface{}) {
	dataByte, _ := json.Marshal(data)
	json.Unmarshal(dataByte, &result)
	return result
}

// ConvertToJSONString ...
func ConvertToJSONString(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

// ConvertToDataTypesJSON ...
func ConvertToDataTypesJSON(d interface{}) datatypes.JSON {
	bytes, _ := json.Marshal(d)

	// JSON data
	jsonData := datatypes.JSON(bytes)
	return jsonData
}

// GetValueOfLangFromJSON ...
func GetValueOfLangFromJSON(data datatypes.JSON, lang string) string {
	byteData, err := data.MarshalJSON()
	if err != nil {
		return ""
	}
	var mapData map[string]string
	if err = json.Unmarshal(byteData, &mapData); err != nil {
		return ""
	}
	return mapData[lang]
}
