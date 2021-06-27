package constant

import "golang.org/x/text/language"

// Constant
var (
	EnvDevelop    = "develop"
	EnvStaging    = "staging"
	EnvProduction = "production"

	LangVi   = language.Vietnamese.String()
	LangEn   = language.English.String()
	LangList = []interface{}{
		LangVi, LangEn,
	}

	OSNameIOS     = "iOS"
	OSNameAndroid = "android"
	OSNameList    = []interface{}{
		OSNameIOS, OSNameAndroid,
	}

	PhoneCountryCodeVietNam = "+84"
	PhoneCountryCodeList    = []interface{}{
		PhoneCountryCodeVietNam,
	}

	EchoContextKeyStaffID = "staffID"
	EchoContextKeyUserID  = "userID"

	Limit10 = 10
	Limit20 = 20

	DateLayoutFull     = "2006-01-02T15:04:05.000Z"
	DateLayoutYYYYMMDD = "2006-01-02"
)
