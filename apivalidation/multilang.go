package apivalidation

import (
	"encoding/json"
	"errors"

	"github.com/Caris-Project/caris-modules/constant"
	"github.com/Caris-Project/caris-modules/model"
)

// MultilangRequired ...
func MultilangRequired(v interface{}) (err error) {
	err = errors.New("invalid multiple languages data")

	value, ok := v.(model.MultiLang)
	if !ok {
		return
	}

	// Get byte data
	byteData, _ := json.Marshal(value)

	// Convert to map
	var mapData map[string]string
	if e := json.Unmarshal(byteData, &mapData); e != nil {
		return
	}

	// Check language
	hasInvalidLang := false
	for _, langInterface := range constant.LangList {
		lang := langInterface.(string)
		if mapData[lang] == "" {
			hasInvalidLang = true
			break
		}
	}

	// Set error
	if !hasInvalidLang {
		err = nil
	}

	return
}
