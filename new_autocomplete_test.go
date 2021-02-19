package complete

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var gitAutoCompleteTree = &CompTree{
	Desc: "bit command",
	Sub: map[string]*CompTree{
		"checkout": {
			Desc: "checkout changes branches",
			Flags: map[string]*CompTree{
				"quiet": {Desc: "suppress progress reporting"},
			},
			Args: map[string]*CompTree{
				"master": {Desc: ".branch description for master."},
				"another-branch": {Desc: "some mildly interesting desc"},
			},
		},
		"remote": {
			Desc: "manage set of tracked repositories",
			Args: map[string]*CompTree{
				"add": {
					Desc: "add a new remote",
					Args: map[string]*CompTree{
						"origin": {Desc: ""},
						"upstream": {Desc: ""},
					},
					Flags: map[string]*CompTree{
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
			Flags: map[string]*CompTree{
				"porcelain": {
					Desc: "produce machine-readable output",
					Args: map[string]*CompTree{
						"v1": {Desc: "v1 porcelain"},
						"v2": {Desc: "v2 porcelain"},
					},
				},
			},
		},
		"commit": {
			Desc: "record changes to repository",
			Flags: map[string]*CompTree{
				"a": {
					Desc: "stage all modified and deleted paths",
				},
				"m": {
					Desc: "use the given message as the commit message",
					Args: map[string]*CompTree{
						`"`: {Desc: "your commit message"},
					},
				},
			},
		},

	},
}

func TestGitAutoComplete(t *testing.T) {
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
			{"quiet", "suppress progress reporting"},
		}},
		{"git checkout -qui", []Suggestion{}},
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
		{`git commit -`, []Suggestion{
			{"a", "stage all modified and deleted paths"},
			{"m", "use the given message as the commit message"},
		}},
		{`git commit -a`, []Suggestion{
			{"a", "stage all modified and deleted paths"},
			{"m", "use the given message as the commit message"},
		}},
		{`git commit -am`, []Suggestion{}},
		{`git commit -a "an awesome commit value" `, []Suggestion{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.text)
		t.Run(testname, func(t *testing.T) {
			got, err := AutoComplete(tt.text, gitAutoCompleteTree)
			if err != nil {
				t.Error(err)
			}
			for _, want := range tt.wants{
				assert.Contains(t, got, want)
			}
			if len(got) != 0 && len(tt.wants) == 0 {
				assert.FailNow(t, "yikes there should 0 suggestions in this case")
			}
		})

	}
}

func TestBitCLIAutoComplete(t *testing.T) {
	var tests = []struct {
		text string
		wants []Suggestion
	}{
		{"bit checkout --quiet ", []Suggestion{
			{"master", ".branch description for master."},
			{"another-branch", "some mildly interesting desc"},
		}},
		{"bit checkout --quiet non-existant", []Suggestion{}},
		{"bit checkout --qui", []Suggestion{
			{"quiet", "suppress progress reporting"},
		}},
		{"bit checkout --yik", []Suggestion{
		}},
		// What there is a single dash just show all single dash flags
		{"bit checkout -zzz", []Suggestion{
			{"q", "suppress progress reporting"},
			{"2", "check out stage #2 for unmerged paths"},
		}},
		{"bit ", []Suggestion{
			{"checkout", "Switch branches or restore working tree files"},
			{"remote", "Manage set of tracked repositories"},
		}},
		{"bit remot", []Suggestion{
			{"remote", "Manage set of tracked repositories"},
		}},
		{"bit remote  ", []Suggestion{
			{"add", "add a new remote"},
			{"get-url", "retrieves the URLs for a remote"},
		}},
		{"bit remote a", []Suggestion{
			{"add", "add a new remote"},
		}},
		{"bit remote add ", []Suggestion{
			{"origin", ""},
			{"upstream", ""},
		}},
		{"bit remote add --fetch ", []Suggestion{
			{"origin", ""},
			{"upstream", ""},
		}},
		{"bit status --porcela", []Suggestion{
			{"porcelain", "produce machine-readable output"},
		}},
		{"bit status --porcelain=", []Suggestion{
			{"v1", "v1 porcelain"},
			{"v2", "v2 porcelain"},
		}},
		{`bit co`, []Suggestion{
			{"commit", "Record changes to the repository"},
		}},
		{`bit commit`, []Suggestion{
			{"commit", "Record changes to the repository"},
		}},
		{`bit commit -`, []Suggestion{
			{"a", "stage all modified and deleted paths"},
			{"m", "use the given message as the commit message"},
		}},
		{`bit commit -a`, []Suggestion{
			{"a", "stage all modified and deleted paths"},
			{"m", "use the given message as the commit message"},
		}},
		// What there is a single dash just show all single dash flags
		{`bit commit -am`, []Suggestion{
			{"a", "stage all modified and deleted paths"},
			{"m", "use the given message as the commit message"},
		}},
		{`bit commit -a "an awesome commit value" `, []Suggestion{}},
	}

	bitCompleteTree := getBitTree()

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.text)
		t.Run(testname, func(t *testing.T) {
			got, err := AutoComplete(tt.text, bitCompleteTree)
			if err != nil {
				t.Error(err)
			}
			for _, want := range tt.wants{
				assert.Contains(t, got, want)
			}
			if len(got) != 0 && len(tt.wants) == 0 {
				t.Logf("%v", got)
				t.Logf("%v", tt.wants)
				assert.FailNow(t, "yikes there should 0 suggestions in this case")

			}
		})

	}
}