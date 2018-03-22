package main

import (
	"fmt"
	"testing"
)

func TestAnchorId(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"'rune' is a weird name", "rune-is-a-weird-name"},
		{"`new` and `make` instead of one", "new-and-make-instead-of-one"},
		{"no language interoperability (only C)", "no-language-interoperability-only-c"},
		{"no function/operator overloading", "no-function-operator-overloading"},
		{`string containing "double quotes"`, "string-containing-double-quotes"},
		{"ignores C#", "ignores-csharp"},
		{"no exceptions", "no-exceptions"},
		{"stuck in '70s", "stuck-in-70s"},
		{"`fork()` is either wrong or impossible", "fork-is-either-wrong-or-impossible"},
		{"no `map`/`reduce`/`filter`", "no-map-reduce-filter"},
		{":= is weird", "colon-equals-is-weird"},
		{"c-style", "c-style"},
	}
	for _, test := range tests {
		if result := anchorID(test.input); result != test.want {
			t.Errorf(`anchorID(%q) = %q, want %q`, test.input, result, test.want)
		}
	}
}

func ExampleAnchorID() {
	fmt.Println(anchorID("'rune' is a weird name"))
	// Output:
	// rune-is-a-weird-name
}
