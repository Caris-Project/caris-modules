package model

import (
	"encoding/json"
	"reflect"
	"strings"

	"gorm.io/datatypes"

	"github.com/Caris-Project/caris-modules/utils"
)

// MultiLang ...
type MultiLang struct {
	En string `json:"en"`
	Vi string `json:"vi"`
}

// GenerateJSON ...
func (l MultiLang) GenerateJSON() datatypes.JSON {
	bytes, _ := json.Marshal(l)

	// JSON data
	jsonData := datatypes.JSON(bytes)
	return jsonData
}

// GenerateSearchTokens ...
func (l MultiLang) GenerateSearchTokens() TsVector {
	// Search tokens
	values := make([]string, 0)
	v := reflect.ValueOf(l)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface().(string)
		values = append(values, utils.RemoveDiacritics(value))
	}

	vecValue := strings.Join(values, " ")
	searchTokens := TsVector{Value: vecValue}
	return searchTokens
}

// GenerateJSONAndSearchTokens ...
func (l MultiLang) GenerateJSONAndSearchTokens() (datatypes.JSON, TsVector) {
	bytes, _ := json.Marshal(l)

	// JSON data
	jsonData := datatypes.JSON(bytes)

	// Search tokens
	values := make([]string, 0)
	v := reflect.ValueOf(l)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface().(string)
		values = append(values, utils.RemoveDiacritics(value))
	}

	vecValue := strings.Join(values, " ")
	searchTokens := TsVector{Value: vecValue}
	return jsonData, searchTokens
}

// GetMultiLangFromJSON ...
func GetMultiLangFromJSON(jsonData datatypes.JSON) (data MultiLang) {
	bytes, _ := jsonData.MarshalJSON()
	json.Unmarshal(bytes, &data)
	return
}

// GetArrayValues ...
func (l MultiLang) GetArrayValues() []string {
	values := make([]string, 0)
	v := reflect.ValueOf(l)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface().(string)
		values = append(values, utils.RemoveDiacritics(value))
	}
	return values
}
