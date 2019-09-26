package twitter

import (
	"net/http"

	"github.com/ilesinge/usercheck/client"
	"github.com/ilesinge/usercheck/rules"
)

// Twitter type
type Twitter struct{}

// Validate check if is twitter valid
func (t *Twitter) Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}

// IsAvailable checks if the username is available on Twitter
func (t *Twitter) IsAvailable(username string, cli client.HTTPClient) (available bool, err error) {
	req, err := http.NewRequest("GET", "https://twitter.com/"+username, nil)
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

// Name gets social network name
func (t *Twitter) Name() string {
	return "Twitter"
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
