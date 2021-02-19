package complete

import (
	"github.com/google/shlex"
	"strings"
)

type AutoCompleteCLI struct {
	Flags map[string]*AutoCompleteCLI
	Args  map[string]*AutoCompleteCLI
	Sub   map[string]*AutoCompleteCLI
	Dynamic func(prefix string) []AutoCompleteCLI
	Desc  string
}

type Suggestion struct {
	Name string
	Desc string
}

func AutoComplete(text string, completionTree *AutoCompleteCLI) ([]Suggestion, error) {
	var prev *AutoCompleteCLI = nil
	curr := completionTree
	curr = prev
	curr = completionTree

	fields, err := shlex.Split(strings.TrimSpace(text))
	if err != nil {
		return nil, err
	}
	last := fields[len(fields)-1]
	// i.e. in "git checkout --fie", remove "git"
	fields = fields[1:]
	for i, field := range fields {
		if strings.HasPrefix(field, "-") {
			continue
		}
		if i == len(fields) - 1 && !strings.HasSuffix(text, " ") {
			continue
		}
		if curr.Sub[field] != nil {
			prev = curr
			curr = curr.Sub[field]
		} else if curr.Args[field] != nil {
			prev = curr
			curr = curr.Args[field]
		}
	}
	s := make([]Suggestion, 0, 20)
	// ends with sub & space
	if curr == nil {
		return s, nil
	} else if strings.HasSuffix(text, " ") {
		if curr.Args != nil {
			for k, v := range curr.Args {
				s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
			}
		}
		if curr.Sub != nil {
			for k, v := range curr.Sub {
				s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
			}
		}

	} else if strings.HasPrefix(last, "-") { // ends with flag
		hasTwoDashes := strings.HasPrefix(last, "--")
		searchTerm := strings.TrimPrefix(last, "-")
		searchTerm = strings.TrimPrefix(searchTerm, "-")
		for k, v := range curr.Flags {
			if hasTwoDashes && len(k) > 1 && (strings.HasPrefix(k, searchTerm) || searchTerm == ""){
					s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
			}
			if !hasTwoDashes && len(k) == 1 {
				s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
			}
		}
	}
	if !strings.HasPrefix(last, "-") && !strings.HasSuffix(text, " ") {
		for k, v := range curr.Args {
			if strings.HasPrefix(k, last) {
				s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
			}
		}
		for k, v := range curr.Sub {
			if strings.HasPrefix(k, last) {
				s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
			}
		}
	}
	// ends with = (suggest flag args)
	if strings.HasPrefix(last, "-") && strings.HasSuffix(last, "=") {
		flagName := strings.TrimPrefix(last, "-")
		flagName = strings.TrimPrefix(flagName, "-")
		flagName = strings.TrimSuffix(flagName, "=")
		flagCompleter := curr.Flags[flagName]
		for k, v := range flagCompleter.Args {
				s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
		}
		for k, v := range flagCompleter.Sub {
				s = append(s, Suggestion{
					Name: k,
					Desc: v.Desc,
				})
		}
	}

	// ends with flag & space
	// ends with sub

	return s, nil
}
