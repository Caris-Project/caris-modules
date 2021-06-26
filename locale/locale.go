package locale

import (
	"github.com/thoas/go-funk"
)

// Locale ...
type Locale struct {
	Key     string
	Message string
	Code    int
}

var notFoundKey = Locale{
	Key:     CommonKeyNotFound,
	Message: "không tìm thấy",
	Code:    -1,
}

// GetByKey give key and receive message + code
func GetByKey(key string) Locale {
	item := funk.Find(list, func(item Locale) bool {
		return item.Key == key
	})

	if item == nil {
		return notFoundKey
	}
	return item.(Locale)
}

var list = make([]Locale, 0)

// Init ...
func Init() {
	// 1-199
	list = append(list, common...)
	// 200-299
	list = append(list, user...)
	// 300-399
	list = append(list, staff...)

	// 500 - 599
	list = append(list, post...)
}
