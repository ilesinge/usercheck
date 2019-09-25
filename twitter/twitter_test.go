package twitter_test

import (
	"testing"

	"github.com/ilesinge/usercheck/twitter"
)

func TestUsernameTooLong(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("0123456789012345")
	if output {
		t.Errorf("Too long username accepted")
	}
}

func TestUsernameShortEnough(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("012345678901234")
	if !output {
		t.Errorf("Valid username refused")
	}
}

func TestUsernameTooShort(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("")
	if output {
		t.Errorf("Too short username accepted")
	}
}

func TestUsernameLongEnough(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("o")
	if !output {
		t.Errorf("Valid username refused")
	}
}

func TestUsernameContainsUppercaseTwitter(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("TWITTERuser")
	if output {
		t.Errorf("username containing uppercase twitter accepted")
	}
}

func TestUsernameContainsLowercaseTwitter(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("twitterUser")
	if output {
		t.Errorf("username containing lowercase twitter accepted")
	}
}

func TestUsernameContainsSpecialchar(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("us*er")
	if output {
		t.Errorf("username containing special char accepted")
	}
}

func TestUsernameIsValidWithUnderscore(t *testing.T) {
	var twitter twitter.Twitter
	output := twitter.Validate("ile_singe")
	if !output {
		t.Errorf("Valid username not accepted")
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
			t.Errorf("expected output to be %v, but got %v", test.output, output)
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
			t.Errorf("expected output to be %v, but got %v", test.output, output)
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
			t.Errorf("expected output to be %v, but got %v", test.output, output)
		}
	}
}
*/
