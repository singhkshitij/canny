package utils

import "regexp"

func Is404Error(err string) bool {
	pattern := "(?i)(?:code) = \\bnotfound\\b"
	isValid, _ := regexp.MatchString(pattern, err)
	return isValid
}
