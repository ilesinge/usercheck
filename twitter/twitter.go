package twitter

import (
	"github.com/ilesinge/usercheck/rules"
)

// Twitter type
type Twitter struct{}

// Validate check if is twitter valid
func (t *Twitter) Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}

func getRules() []func(string) bool {
	localRules := []func(string) bool{
		rules.IsLongEnough(1),
		rules.IsShortEnough(15),
		rules.DoesNotContainSubstring("twitter"),
		rules.ContainsValidCharacters("^[0-9A-Z_a-z]+$"),
	}
	return localRules
}
