package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// RemoveDiacritics ...
func RemoveDiacritics(s string) string {
	if s != "" {
		s = strings.ToLower(s)
		s = replaceStringWithRegex(s, `đ`, "d")
		t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
		result, _, _ := transform.String(t, s)
		result = replaceStringWithRegex(result, `[^a-zA-Z0-9\s]`, "")

		return result
	}
	return ""
}

// replaceStringWithRegex ...
func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}

// FormatPhoneFull ...
func FormatPhoneFull(countryCode, number string) string {
	// Country code always format as: +{country_code}, ex: +84
	// Convert number to 9 chars, without "0" at start
	numberLength := len(number)

	// Remove char 0
	if numberLength == 10 && string(number[0]) == "0" {
		number = number[1:]
	}

	// Join and return
	return countryCode + number
}
