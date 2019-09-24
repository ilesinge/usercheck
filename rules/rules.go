package rules

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// IsLongEnoughGenerator creates min length checker
func IsLongEnoughGenerator(i int) func(string) bool {
	return func(username string) bool {
		length := utf8.RuneCountInString(username)
		return length >= i
	}
}

// IsShortEnoughGenerator creates max length checker
func IsShortEnoughGenerator(i int) func(string) bool {
	return func(username string) bool {
		length := utf8.RuneCountInString(username)
		return length <= i
	}
}

// DoesNotContainSubstringGenerator creates substring checker
func DoesNotContainSubstringGenerator(s string) func(string) bool {
	return func(username string) bool {
		return !strings.Contains(strings.ToLower(username), s)
	}
}

// ContainsValidCharactersGenerator creates pattern checker
func ContainsValidCharactersGenerator(pattern string) func(string) bool {
	var usernameMask = regexp.MustCompile(pattern)
	return func(username string) bool {
		return usernameMask.MatchString(username)
	}
}

// DoesNotStartWithGenerator creates a begin checker
func DoesNotStartWithGenerator(s string) func(string) bool {
	return func(username string) bool {
		return !strings.HasPrefix(username, "-")
	}
}

// DoesNotEndWithGenerator creates an end checker
func DoesNotEndWithGenerator(s string) func(string) bool {
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
