package routemiddleware

import (
	"github.com/Caris-Project/caris-modules/constant"
	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
)

// List allow languages
var languages = []string{constant.LangVi, constant.LangEn}

// Locale set default locale for user, based on header
func Locale(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		lang := c.Request().Header.Get("Accept-Language")

		// Set default language
		if !funk.Contains(languages, lang) {
			lang = constant.LangVi
		}

		// Set locale
		c.Set("lang", lang)
		return next(c)
	}
}
