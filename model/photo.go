package model

import (
	"encoding/json"

	"gorm.io/datatypes"
)

const (
	SizeSmallPrefix  = "sm_"
	SizeMediumPrefix = "md_"
)

// Photo ...
type Photo struct {
	ID     string    `gorm:"primaryKey" json:"id"`
	Name   string    `json:"name"`
	Small  PhotoSize `json:"small"`
	Medium PhotoSize `json:"medium"`
}

// PhotoSize ...
type PhotoSize struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// ParsePhotoFromJSON ...
func ParsePhotoFromJSON(jsonData datatypes.JSON) (result *Photo) {
	if jsonData == nil {
		return
	}
	bytes, _ := json.Marshal(jsonData)
	if err := json.Unmarshal(bytes, &result); err != nil {
		return
	}

	result.Small.URL = c.Host + SizeSmallPrefix + result.Name
	result.Medium.URL = c.Host + SizeMediumPrefix + result.Name
	return
}

// ParsePhotosFromJSON ...
func ParsePhotosFromJSON(jsonData datatypes.JSON) (result []Photo) {
	result = make([]Photo, 0)

	if jsonData == nil {
		return
	}

	bytes, _ := json.Marshal(jsonData)
	if err := json.Unmarshal(bytes, &result); err != nil {
		return
	}

	for i, _ := range result {
		result[i].Small.URL = c.Host + SizeSmallPrefix + result[i].Name
		result[i].Medium.URL = c.Host + SizeMediumPrefix + result[i].Name
	}
	return
}
