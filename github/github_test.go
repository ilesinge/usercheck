package github_test

import (
	"testing"

	"github.com/ilesinge/usercheck/github"
)

func TestUsernameTooLong(t *testing.T) {
	var github github.Github
	output := github.Validate("0123456789012345678901234567890123456789")
	if output {
		t.Fatalf("Too long username accepted")
	}
}

func TestUsernameShortEnough(t *testing.T) {
	var github github.Github
	output := github.Validate("012345678901234567890123456789012345678")
	if !output {
		t.Fatalf("Valid username refused")
	}
}

func TestUsernameTooShort(t *testing.T) {
	var github github.Github
	output := github.Validate("")
	if output {
		t.Fatalf("Too short username accepted")
	}
}

func TestUsernameLongEnough(t *testing.T) {
	var github github.Github
	output := github.Validate("o")
	if !output {
		t.Fatalf("Valid username refused")
	}
}

func TestUsernameContainsDoubleDash(t *testing.T) {
	var github github.Github
	output := github.Validate("user--user")
	if output {
		t.Fatalf("username containing double dash accepted")
	}
}

func TestUsernameContainsDashAtBeginning(t *testing.T) {
	var github github.Github
	output := github.Validate("-user")
	if output {
		t.Fatalf("username beginning with dash accepted")
	}
}

func TestUsernameContainsDashAtEnd(t *testing.T) {
	var github github.Github
	output := github.Validate("user-")
	if output {
		t.Fatalf("username ending with dash accepted")
	}
}

func TestUsernameContainsSpecialchar(t *testing.T) {
	var github github.Github
	output := github.Validate("us*er")
	if output {
		t.Fatalf("username beginning with special char accepted")
	}
}

func TestUsernameIsValidWithDash(t *testing.T) {
	var github github.Github
	output := github.Validate("ile-singe")
	if !output {
		t.Fatalf("Valid username not accepted")
	}
}
