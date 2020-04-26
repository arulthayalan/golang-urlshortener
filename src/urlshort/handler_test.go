package urlshort

import (
	"testing"

)

type Case struct {
	key   string
	value string
}


func TestBuildMap(t *testing.T) {
	pathUrls := []pathUrl{
		{
			Path: "/urlshort",
			URL: "https://github.com/gophercises/urlshort",
		},
		{
			"/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution",
		},
	}
	cases := []Case{
		{"/urlshort", "https://github.com/gophercises/urlshort",},
		{"/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution",},
	}
	pathToUrls := buildMap(pathUrls)

	for _, testcase := range cases {
		got := pathToUrls[testcase.key]

		if got == "" || got != testcase.value {
			t.Errorf("buildMap key: %q == %v, want %v", testcase.key, got, testcase.value)
		}
	}
}