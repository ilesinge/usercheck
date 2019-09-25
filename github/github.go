package github

import (
	"net/http"

	"github.com/ilesinge/usercheck/rules"
)

// Github type
type Github struct{}

// ErrNetworkFailure error
type ErrNetworkFailure struct {
	Cause error
}

func (e *ErrNetworkFailure) Error() string {
	return "An error occured during Github username availability check"
}

// Unwrap returns the underlying error
func (e *ErrNetworkFailure) Unwrap() error {
	return e.Cause
}

// Validate check if is github valid
func (g *Github) Validate(username string) bool {
	return rules.ValidateRules(getRules(), username)
}

// IsAvailable checks if the username is available on Github
func (g *Github) IsAvailable(username string) (available bool, err error) {
	res, err := http.Get("https://github.com/" + username)
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
		rules.IsShortEnough(39),
		rules.DoesNotContainSubstring("--"),
		rules.ContainsValidCharacters("^[0-9A-Za-z-]+$"),
		rules.DoesNotStartWith("-"),
		rules.DoesNotEndWith("-"),
	}
	return localRules
}
