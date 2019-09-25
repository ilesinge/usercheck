package github

import (
	"net/http"

	"github.com/ilesinge/usercheck/client"
	"github.com/ilesinge/usercheck/rules"
)

// Github type
type Github struct{}

// Validate check if is github valid
func (g *Github) Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}

// IsAvailable checks if the username is available on Github
func (g *Github) IsAvailable(username string, cli client.HTTPClient) (available bool, err error) {
	req, err := http.NewRequest("GET", "https://github.com/"+username, nil)
	if err != nil {
		return
	}
	res, err := (cli).Do(req)
	if err != nil {
		// Need a pointer because it's the pointer that satisfies the interface (pointer receptor)
		err = &client.ErrNetworkFailure{Cause: err}
		return
	}
	defer res.Body.Close()
	available = (res.StatusCode == http.StatusNotFound)
	return
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
