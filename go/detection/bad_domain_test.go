package main

import (
	"testing"
)

func TestFmtRegex(t *testing.T) {
	tc := struct {
		iocs []string
		want string
	}{
		iocs: []string{
			"example.com",
			"not.a.domain.example",
			"test.google.com",
		},
		want: ".*(example.com|not.a.domain.example|test.google.com)$",
	}
	if got := fmtRegex(tc.iocs); got != tc.want {
		t.Errorf("fmtRegex() returned unexpected regex string;\nwant: %s\ngot:%s\n", tc.want, got)
	}
}

func TestBadDomainDetection(t *testing.T) {
	// <TODO: Implement me!>
	// If you break up your code into subfunctions (e.g. one for filter(), or for aggregate()),
	// feel free to unit test your subfunctions, rather than run().
	// E.g., for filter() and aggregate(), you could have:
	// def TestBadDomainFilter(self):
	//     [...]
	//
	// def TestBadDomainAggregate(self):
	//     [...]
	// Hint #1: `cmp` and `protocmp` packages are your friends.
	// Hint #2: Testing large functions can be hard - consider breaking your code into subfunctions to make unit testing easier.
	// Hint #3: Use the test above as an example of how to set up a table driven test.
}
