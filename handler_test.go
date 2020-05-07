package urlshort

import (
	"testing"
   "github.com/google/go-cmp/cmp"
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

func TestParseYml(t *testing.T) {
	wantPathUrls :=  []pathUrl{
		{
			Path: "/urlshort",
			URL: "https://github.com/gophercises/urlshort",
		},
	}
	inYaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
`

	gotPathUrls, err := 	parseYaml([]byte(inYaml))

	if (err != nil) {
		t.Errorf("parseYml %v", err)
	}

	for i, pu := range gotPathUrls {
		if diff := cmp.Diff(wantPathUrls[i], pu); diff != "" {
			t.Errorf("parseYaml() mismatch (-want +got):\n%s", diff)
		}
	}

}