package twitter

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var usernameMask = regexp.MustCompile("^[0-9A-Z_a-z]+$")

// Validate twitter username
func Validate(username string) bool {
	return isLongEnough(username) &&
		isShortEnough(username) &&
		doesNotContainTwitter(username) &&
		containsValidCharacters(username)
}

func isLongEnough(username string) bool {
	length := utf8.RuneCountInString(username)
	return length >= 1
}

func isShortEnough(username string) bool {
	length := utf8.RuneCountInString(username)
	return length <= 15
}

func doesNotContainTwitter(username string) bool {
	return !strings.Contains(strings.ToLower(username), "twitter")
}

func containsValidCharacters(username string) bool {
	return usernameMask.MatchString(username)
}
