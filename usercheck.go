package main

import (
	"fmt"

	"github.com/ilesinge/usercheck/rules"
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
	twitterRules = append(twitterRules, rules.IsLongEnoughGenerator(1))
	twitterRules = append(twitterRules, rules.IsShortEnoughGenerator(15))
	twitterRules = append(twitterRules, rules.DoesNotContainSubstringGenerator("twitter"))
	twitterRules = append(twitterRules, rules.ContainsValidCharactersGenerator("^[0-9A-Z_a-z]+$"))
	return twitterRules
}

func validateRules(rules []func(string) bool, username string) bool {
	valid := true
	for _, rule := range rules {
		valid = valid && rule(username)
	}
	return valid
}
