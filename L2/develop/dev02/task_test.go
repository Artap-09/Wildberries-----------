package main

import "testing"

type testInOut struct{
	in string
	out string
}

var tests = []testInOut {
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd" , "abcd"},
	{"45" , ""},
	{"", ""},
	{"qwe\\4\\5" , "qwe45"},
	{"qwe\\45" , "qwe44444"},
	{"qwe\\\\5" , "qwe\\\\\\\\\\"},
}

func TestUnpacking(t *testing.T)  {
	for _, inOut := range tests {
		out,_:=Unpacking(inOut.in)
		if out != inOut.out {
			t.Error(
				"For", inOut.in,
				"expected", inOut.out,
				"got", out,
			)
		}
	}
}