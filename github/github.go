package github

import "github.com/ilesinge/usercheck/rules"

// Github type
type Github struct{}

// Validate check if is github valid
func (g *Github) Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}

func getRules() []func(string) bool {
	localRules := []func(string) bool{
		rules.IsLongEnough(1),
		rules.IsShortEnough(39),
		rules.DoesNotContainSubstring("--"),
		rules.ContainsValidCharacters("^[0-9A-Za-z-]+$"),
		rules.DoesNotStartWith("-"),
		rules.DoesNotEndWith("-"),
	}
	return localRules
}
