package complete

import (
	"strings"
)

type AutoCompleteCLI struct {
	Flags map[string]*AutoCompleteCLI
	Args  map[string]*AutoCompleteCLI
	Sub   map[string]*AutoCompleteCLI
	Desc  string
}

type Suggestion struct {
	name string
	Desc string
}

func AutoComplete(text string, completionTree *AutoCompleteCLI) []Suggestion {
	var prev *AutoCompleteCLI = nil
	curr := completionTree
	curr = prev
	curr = completionTree
	fields := strings.Fields(strings.TrimSpace(text))
	last := fields[len(fields)-1]
	for _, field := range fields {
		if strings.HasPrefix(field, "-") {
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
	if strings.HasSuffix(text, " ") {
		if curr.Args != nil {
			for k, v := range curr.Args {
				s = append(s, Suggestion{
					name: k,
					Desc: v.Desc,
				})
			}
		}
		if curr.Sub != nil {
			for k, v := range curr.Sub {
				s = append(s, Suggestion{
					name: k,
					Desc: v.Desc,
				})
			}
		}

	} else if strings.HasPrefix(last, "-") { // ends with flag
		for k, v := range curr.Flags {
			s = append(s, Suggestion{
				name: k,
				Desc: v.Desc,
			})
		}
	}
	if !strings.HasPrefix(last, "-") && !strings.HasSuffix(text, " ") {
		for k, v := range curr.Args {
			if strings.HasPrefix(k, last) {
				s = append(s, Suggestion{
					name: k,
					Desc: v.Desc,
				})
			}
		}
		for k, v := range curr.Sub {
			if strings.HasPrefix(k, last) {
				s = append(s, Suggestion{
					name: k,
					Desc: v.Desc,
				})
			}
		}
	}
	// ends with flag & space
	// ends with sub

	return s
}
