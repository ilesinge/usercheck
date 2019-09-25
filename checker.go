package usercheck

import "github.com/ilesinge/usercheck/client"

// Checker interface for social network username checking
type Checker interface {
	IsAvailable(string, client.HTTPClient) (available bool, err error)
	Validate(string) bool
}
