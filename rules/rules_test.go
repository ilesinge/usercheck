package rules_test

import (
	"testing"

	"github.com/ilesinge/usercheck/rules"
)

func TestIsNotLongEnough(t *testing.T) {
	rule := rules.IsLongEnough(5)
	output := rule("1234")
	if output {
		t.Errorf("Invalid username accepted")
	}
}

func TestIsLongEnough(t *testing.T) {
	rule := rules.IsLongEnough(5)
	output := rule("12345")
	if !output {
		t.Errorf("Valid username rejected")
	}
}

func TestIsNotShortEnough(t *testing.T) {
	rule := rules.IsShortEnough(5)
	output := rule("123456")
	if output {
		t.Errorf("Invalid username accepted")
	}
}

func TestIsShortEnough(t *testing.T) {
	rule := rules.IsShortEnough(5)
	output := rule("12345")
	if !output {
		t.Errorf("Valid username rejected")
	}
}

func TestDoesNotContainSubstring(t *testing.T) {
	rule := rules.DoesNotContainSubstring("blob")
	output := rule("hallo")
	if !output {
		t.Errorf("Valid username rejected")
	}
}

func TestDoesContainSubstring(t *testing.T) {
	rule := rules.DoesNotContainSubstring("blob")
	output := rule("halloblob")
	if output {
		t.Errorf("Invalid username accepted")
	}
}

func TestContainsValidCharacters(t *testing.T) {
	rule := rules.ContainsValidCharacters("^[a-z]+$")
	output := rule("hallo")
	if !output {
		t.Errorf("Valid username rejected")
	}
}

func TestDoesNotContainValidCharacters(t *testing.T) {
	rule := rules.ContainsValidCharacters("^[a-z]+$")
	output := rule("hal$lo")
	if output {
		t.Errorf("Invalid username accepted")
	}
}

func TestDoesNotStartWith(t *testing.T) {
	rule := rules.DoesNotStartWith("abc")
	output := rule("hallo")
	if !output {
		t.Errorf("Valid username rejected")
	}
}

func TestDoesStartWith(t *testing.T) {
	rule := rules.DoesNotStartWith("abc")
	output := rule("abchallo")
	if !output {
		t.Errorf("Invalid username accepted")
	}
}

func TestDoesNotEndWith(t *testing.T) {
	rule := rules.DoesNotEndWith("abc")
	output := rule("hallo")
	if !output {
		t.Errorf("Valid username rejected")
	}
}

func TestDoesEndWith(t *testing.T) {
	rule := rules.DoesNotEndWith("abc")
	output := rule("halloabc")
	if !output {
		t.Errorf("Invalid username accepted")
	}
}

/*
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
*/
