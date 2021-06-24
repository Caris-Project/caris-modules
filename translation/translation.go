package translation

import (
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type languageData struct {
	Lang      string
	Localizer *i18n.Localizer
}

// Init ...
func Init() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load files
	pwd := getPwd()
	loadOther(bundle, pwd)
	loadNotificationTitle(bundle, pwd)
	loadNotificationContent(bundle, pwd)
}

func getPwd() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}
