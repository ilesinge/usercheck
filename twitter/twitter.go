package twitter

import "github.com/ilesinge/usercheck/rules"

func getRules() []func(string) bool {
	var rulesSlice []func(string) bool
	rulesSlice = append(rulesSlice, rules.IsLongEnoughGenerator(1))
	rulesSlice = append(rulesSlice, rules.IsShortEnoughGenerator(15))
	rulesSlice = append(rulesSlice, rules.DoesNotContainSubstringGenerator("twitter"))
	rulesSlice = append(rulesSlice, rules.ContainsValidCharactersGenerator("^[0-9A-Z_a-z]+$"))
	return rulesSlice
}

// Validate check if is twitter valid
func Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}
