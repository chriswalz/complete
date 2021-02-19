package complete

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bitAutoCompleteTree = &AutoCompleteCLI{
	Desc: "bit command",
	Sub: map[string]*AutoCompleteCLI{
		"checkout": {
			Desc: "checkout changes branches",
			Flags: map[string]*AutoCompleteCLI{
				"quiet": {Desc: "quiet flag desc"},
			},
			Args: map[string]*AutoCompleteCLI{
				"master": {Desc: ".branch description for master."},
				"another-branch": {Desc: "some mildly interesting desc"},
			},
		},
		"remote": {
			Desc: "manage set of tracked repositories",
			Args: map[string]*AutoCompleteCLI{
				"add": {
					Desc: "add a new remote",
					Args: map[string]*AutoCompleteCLI{
						"origin": {Desc: ""},
						"upstream": {Desc: ""},
					},
					Flags: map[string]*AutoCompleteCLI{
						"fetch": {
							Desc: "run git fetch on new remote after it has been created",
						},
						"f": {Desc: "run git fetch on new remote after it has been created"},
					},
				},
				"get-url": {Desc: "retrieves the URLs for a remote"},
			},
		},
		"status": {
			Desc: "show working-tree status",
			Flags: map[string]*AutoCompleteCLI{
				"porcelain": {
					Desc: "produce machine-readable output",
					Args: map[string]*AutoCompleteCLI{
						"v1": {Desc: "v1 porcelain"},
						"v2": {Desc: "v2 porcelain"},
					},
				},
			},
		},

	},
}

func TestNewAutoComplete(t *testing.T) {
	var tests = []struct {
		text string
		wants []Suggestion
	}{
		{"git checkout --quiet ", []Suggestion{
			{"master", ".branch description for master."},
			{"another-branch", "some mildly interesting desc"},
		}},
		{"git checkout --quiet non-existant", []Suggestion{}},
		{"git checkout --qui", []Suggestion{
			{"quiet", "quiet flag desc"},
		}},
		{"git ", []Suggestion{
			{"checkout", "checkout changes branches"},
			{"remote", "manage set of tracked repositories"},
		}},
		{"git remot", []Suggestion{
			{"remote", "manage set of tracked repositories"},
		}},
		{"git remote  ", []Suggestion{
			{"add", "add a new remote"},
			{"get-url", "retrieves the URLs for a remote"},
		}},
		{"git remote a", []Suggestion{
			{"add", "add a new remote"},
		}},
		{"git remote add ", []Suggestion{
			{"origin", ""},
			{"upstream", ""},
		}},
		{"git remote add --fetch ", []Suggestion{
			{"origin", ""},
			{"upstream", ""},
		}},
		{"git status --porcela", []Suggestion{
			{"porcelain", "produce machine-readable output"},
		}},
		{"git status --porcelain=", []Suggestion{
			{"v1", "v1 porcelain"},
			{"v2", "v2 porcelain"},
		}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.text)
		t.Run(testname, func(t *testing.T) {
			got := AutoComplete(tt.text, bitAutoCompleteTree)
			for _, want := range tt.wants{
				assert.Contains(t, got, want)
			}
			if len(got) != 0 && len(tt.wants) == 0 {
				assert.FailNow(t, "yikes there should 0 suggestions in this case")
			}
		})

	}
}