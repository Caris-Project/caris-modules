package translation

import (
	"fmt"
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thoas/go-funk"
	"golang.org/x/text/language"

	"github.com/Caris-Project/caris-modules/constant"
)

type languageData struct {
	Lang      string
	Localizer *i18n.Localizer
}

var listOther = make([]languageData, 0)

// Init ...
func Init() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load files
	pwd := getPwd()
	loadOther(bundle, pwd)
}

func getPwd() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}

func loadOther(bundle *i18n.Bundle, pwd string) {
	for _, langInterface := range constant.LangList {
		lang := langInterface.(string)

		// Add language
		path := fmt.Sprintf("%s/other/active.%s.toml", pwd, lang)
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
		listOther = append(listOther, item)
	}
}

// GetOtherText ...
func GetOtherText(lang, messageID string) string {
	itemInterface := funk.Find(listOther, func(i languageData) bool {
		return i.Lang == lang
	})
	if itemInterface == nil {
		return ""
	}

	item := itemInterface.(languageData)
	msg, _ := item.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	return msg
}
