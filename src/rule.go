package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/msw-x/moon/ufs"
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
	add := func(line string) {
		line = ustring.TrimWhitespaces(line)
		if line != "" {
			var rule Rule
			rule.Source, rule.Target = ustring.SplitPair(line, ":")
			if rule.Ok() {
				rules = append(rules, rule)
			}
		}
	}
	lines, err = ufs.ReadLines(file)
	if err == nil {
		for _, line := range lines {
			add(line)
		}
	}
	if s, ok := os.LookupEnv("ADD_ROUTES"); ok {
		for _, v := range strings.Split(s, ",") {
			add(v)
		}
	}
	return
}
