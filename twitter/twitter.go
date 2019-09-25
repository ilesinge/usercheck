package twitter

import (
	"net/http"

	"github.com/ilesinge/usercheck/rules"
)

// Twitter type
type Twitter struct{}

// ErrNetworkFailure error
type ErrNetworkFailure struct {
	Cause error
}

func (e *ErrNetworkFailure) Error() string {
	return "An error occured during Twitter username availability check"
}

// Unwrap returns the underlying error
func (e *ErrNetworkFailure) Unwrap() error {
	return e.Cause
}

// Validate check if is twitter valid
func (t *Twitter) Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}

// IsAvailable checks if the username is available on Twitter
func (t *Twitter) IsAvailable(username string) (available bool, err error) {
	res, err := http.Get("https://twitter.com/" + username)
	if err != nil {
		// Need a pointer because it's the pointer that satisfies the interface (pointer receptor)
		err = &ErrNetworkFailure{Cause: err}
		return
	}
	defer res.Body.Close()
	available = (res.StatusCode == http.StatusNotFound)
	return
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
