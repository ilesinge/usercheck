package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	twitterRules := getTwitterRules()
	fmt.Println(validateRules(twitterRules, "coucou"))
	fmt.Println(validateRules(twitterRules, ""))
	fmt.Println(validateRules(twitterRules, "coucoucoucoucoucoucoucoucoucou"))
	fmt.Println(validateRules(twitterRules, "coucouTwitter"))
	fmt.Println(validateRules(twitterRules, "twiTtercoucou"))
	fmt.Println(validateRules(twitterRules, "coucou*"))
	fmt.Println(validateRules(twitterRules, "%coucou"))
}

func getTwitterRules() []func(string) bool {
	var twitterRules []func(string) bool
	twitterRules = append(twitterRules, isLongEnoughGenerator(1))
	twitterRules = append(twitterRules, isShortEnoughGenerator(15))
	twitterRules = append(twitterRules, doesNotContainSubstringGenerator("twitter"))
	twitterRules = append(twitterRules, containsValidCharactersGenerator("^[0-9A-Z_a-z]+$"))
	return twitterRules
}

func validateRules(rules []func(string) bool, username string) bool {
	valid := true
	for _, rule := range rules {
		valid = valid && rule(username)
	}
	return valid
}

func isLongEnoughGenerator(i int) func(string) bool {
	return func(username string) bool {
		length := utf8.RuneCountInString(username)
		return length >= i
	}
}

func isShortEnoughGenerator(i int) func(string) bool {
	return func(username string) bool {
		length := utf8.RuneCountInString(username)
		return length <= i
	}
}

func doesNotContainSubstringGenerator(s string) func(string) bool {
	return func(username string) bool {
		return !strings.Contains(strings.ToLower(username), s)
	}
}

func containsValidCharactersGenerator(pattern string) func(string) bool {
	var usernameMask = regexp.MustCompile(pattern)
	return func(username string) bool {
		return usernameMask.MatchString(username)
	}
}
