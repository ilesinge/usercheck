package twitter_test

import (
	"testing"

	"github.com/ilesinge/usercheck/twitter"
)

func TestUsernameTooLong(t *testing.T) {
	output := twitter.Validate("0123456789012345")
	if output {
		t.Fatalf("Too long username accepted")
	}
}

func TestUsernameShortEnough(t *testing.T) {
	output := twitter.Validate("012345678901234")
	if !output {
		t.Fatalf("Valid username refused")
	}
}

func TestUsernameTooShort(t *testing.T) {
	output := twitter.Validate("")
	if output {
		t.Fatalf("Too short username accepted")
	}
}

func TestUsernameLongEnough(t *testing.T) {
	output := twitter.Validate("o")
	if !output {
		t.Fatalf("Valid username refused")
	}
}

func TestUsernameContainsUppercaseTwitter(t *testing.T) {
	output := twitter.Validate("TWITTERuser")
	if output {
		t.Fatalf("username containing uppercase twitter accepted")
	}
}

func TestUsernameContainsLowercaseTwitter(t *testing.T) {
	output := twitter.Validate("twitterUser")
	if output {
		t.Fatalf("username containing lowercase twitter accepted")
	}
}

func TestUsernameContainsSpecialchar(t *testing.T) {
	output := twitter.Validate("us*er")
	if output {
		t.Fatalf("username containing special char accepted")
	}
}

func TestUsernameIsValidWithUnderscore(t *testing.T) {
	output := twitter.Validate("ile_singe")
	if !output {
		t.Fatalf("Valid username not accepted")
	}
}

/*
func TestIsLongEnough(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"long enough": {
			input:  "o",
			output: true,
		},
		"not long enough": {
			input:  "",
			output: false,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output := isLongEnough(test.input)
		if test.output != output {
			t.Fatalf("expected output to be %v, but got %v", test.output, output)
		}
	}
}

func TestIsShortEnough(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"short enough": {
			input:  "012345678901234",
			output: true,
		},
		"not short enough": {
			input:  "0123456789012345",
			output: false,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output := isShortEnough(test.input)
		if test.output != output {
			t.Fatalf("expected output to be %v, but got %v", test.output, output)
		}
	}
}

func TestDoesNotContainTwitter(t *testing.T) {
	tests := map[string]struct {
		input  string
		output bool
	}{
		"short enough": {
			input:  "coucou",
			output: true,
		},
		"not short enough": {
			input:  "coucouTwitter",
			output: false,
		},
		"not short enough": {
			input:  "coucouTwitter",
			output: false,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output := doesNotContainTwitter(test.input)
		if test.output != output {
			t.Fatalf("expected output to be %v, but got %v", test.output, output)
		}
	}
}
*/
