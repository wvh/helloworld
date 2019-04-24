package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{
			in:  "",
			out: "Hello World!",
		},
		{
			in:  "Jack",
			out: "Hello Jack!",
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			res := HelloWorld(test.in)
			if res != test.out {
				t.Errorf("got %q, want %q", res, test.out)
			}
		})
	}
}
