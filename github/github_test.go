package github_test

import (
	"testing"

	"github.com/ilesinge/usercheck/github"
)

func TestUsernameTooLong(t *testing.T) {
	var github github.Github
	output := github.Validate("0123456789012345678901234567890123456789")
	if output {
		t.Errorf("Too long username accepted")
	}
}

func TestUsernameShortEnough(t *testing.T) {
	var github github.Github
	output := github.Validate("012345678901234567890123456789012345678")
	if !output {
		t.Errorf("Valid username refused")
	}
}

func TestUsernameTooShort(t *testing.T) {
	var github github.Github
	output := github.Validate("")
	if output {
		t.Errorf("Too short username accepted")
	}
}

func TestUsernameLongEnough(t *testing.T) {
	var github github.Github
	output := github.Validate("o")
	if !output {
		t.Errorf("Valid username refused")
	}
}

func TestUsernameContainsDoubleDash(t *testing.T) {
	var github github.Github
	output := github.Validate("user--user")
	if output {
		t.Errorf("username containing double dash accepted")
	}
}

func TestUsernameContainsDashAtBeginning(t *testing.T) {
	var github github.Github
	output := github.Validate("-user")
	if output {
		t.Errorf("username beginning with dash accepted")
	}
}

func TestUsernameContainsDashAtEnd(t *testing.T) {
	var github github.Github
	output := github.Validate("user-")
	if output {
		t.Errorf("username ending with dash accepted")
	}
}

func TestUsernameContainsSpecialchar(t *testing.T) {
	var github github.Github
	output := github.Validate("us*er")
	if output {
		t.Errorf("username beginning with special char accepted")
	}
}

func TestUsernameIsValidWithDash(t *testing.T) {
	var github github.Github
	output := github.Validate("ile-singe")
	if !output {
		t.Errorf("Valid username not accepted")
	}
}
