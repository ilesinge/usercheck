package github

import "github.com/ilesinge/usercheck/rules"

func getRules() []func(string) bool {
	var rulesSlice []func(string) bool
	rulesSlice = append(rulesSlice, rules.IsLongEnoughGenerator(1))
	rulesSlice = append(rulesSlice, rules.IsShortEnoughGenerator(39))
	rulesSlice = append(rulesSlice, rules.DoesNotContainSubstringGenerator("--"))
	rulesSlice = append(rulesSlice, rules.ContainsValidCharactersGenerator("^[0-9A-Za-z-]+$"))
	rulesSlice = append(rulesSlice, rules.DoesNotStartWithGenerator("-"))
	rulesSlice = append(rulesSlice, rules.DoesNotEndWithGenerator("-"))
	return rulesSlice
}

// Validate check if is github valid
func Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}
