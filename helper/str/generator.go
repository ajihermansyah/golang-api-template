package str

import (
	"regexp"
	"strings"
)

func PhoneConvertToAbbv(phone string) string {
	isMatch, _ := regexp.MatchString(`^[0{1}]`, phone)
	if isMatch {
		re := regexp.MustCompile(`^[0{1}]`)
		s := re.ReplaceAllString(phone, `+62`)

		return s
	}

	return phone
}

func Replacer(source string, replacer *strings.Replacer) string {
	return replacer.Replace(source)
}
