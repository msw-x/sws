package main

import (
	"fmt"

	"github.com/msw-x/moon/fs"
	"github.com/msw-x/moon/ustring"
)

type Rule struct {
	Source string
	Target string
}

func (o Rule) Ok() bool {
	return o.Source != "" && o.Target != ""
}

func (o Rule) String() string {
	return fmt.Sprintf("%s => %s", o.Source, o.Target)
}

func loadRules(file string) (rules []Rule, err error) {
	var lines []string
	lines, err = fs.ReadLines(file)
	if err == nil {
		for _, line := range lines {
			line = ustring.TrimWhitespaces(line)
			if line != "" {
				var rule Rule
				rule.Source, rule.Target = ustring.SplitPair(line, ":")
				if rule.Ok() {
					rules = append(rules, rule)
				}
			}
		}
	}
	return
}
