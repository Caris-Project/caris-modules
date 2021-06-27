package translation

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
)

var listNotificationTitle = make([]languageData, 0)

func loadNotificationTitle(bundle *i18n.Bundle, pwd string) {
	for _, lang := range langList {

		// Add language
		path := fmt.Sprintf("%s/notification/title/active.%s.toml", pwd, lang)
		_, err := bundle.LoadMessageFile(path)
		if err != nil {
			panic(err)
		}

		// Add localizer
		localizer := i18n.NewLocalizer(bundle, lang)

		// Add item
		item := languageData{
			Lang:      lang,
			Localizer: localizer,
		}
		listNotificationTitle = append(listNotificationTitle, item)
	}
}

// GetNotificationTitle ...
func GetNotificationTitle(lang, messageID string, data interface{}) string {
	itemInterface := funk.Find(listNotificationTitle, func(i languageData) bool {
		return i.Lang == lang
	})
	if itemInterface == nil {
		return ""
	}

	item := itemInterface.(languageData)
	msg, _ := item.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
	return msg
}
