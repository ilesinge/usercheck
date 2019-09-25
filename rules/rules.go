package rules

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// IsLongEnough creates min length checker
func IsLongEnough(i int) func(string) bool {
	return func(username string) bool {
		length := utf8.RuneCountInString(username)
		return length >= i
	}
}

// IsShortEnough creates max length checker
func IsShortEnough(i int) func(string) bool {
	return func(username string) bool {
		length := utf8.RuneCountInString(username)
		return length <= i
	}
}

// DoesNotContainSubstring creates substring checker
func DoesNotContainSubstring(s string) func(string) bool {
	return func(username string) bool {
		return !strings.Contains(strings.ToLower(username), s)
	}
}

// ContainsValidCharacters creates pattern checker
func ContainsValidCharacters(pattern string) func(string) bool {
	var usernameMask = regexp.MustCompile(pattern)
	return func(username string) bool {
		return usernameMask.MatchString(username)
	}
}

// DoesNotStartWith creates a begin checker
func DoesNotStartWith(s string) func(string) bool {
	return func(username string) bool {
		return !strings.HasPrefix(username, "-")
	}
}

// DoesNotEndWith creates an end checker
func DoesNotEndWith(s string) func(string) bool {
	return func(username string) bool {
		return !strings.HasSuffix(username, "-")
	}
}

// ValidateRules checks if a set of rules is verified
func ValidateRules(rules []func(string) bool, username string) bool {
	valid := true
	for _, rule := range rules {
		valid = valid && rule(username)
	}
	return valid
}
