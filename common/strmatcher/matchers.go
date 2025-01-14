package strmatcher

import (
	"github.com/dlclark/regexp2"
	"strings"
)

type fullMatcher string

func (m fullMatcher) Match(s string) bool {
	return string(m) == s
}

func (m fullMatcher) String() string {
	return "full:" + string(m)
}

type substrMatcher string

func (m substrMatcher) Match(s string) bool {
	return strings.Contains(s, string(m))
}

func (m substrMatcher) String() string {
	return "keyword:" + string(m)
}

type domainMatcher string

func (m domainMatcher) Match(s string) bool {
	pattern := string(m)
	if !strings.HasSuffix(s, pattern) {
		return false
	}
	return len(s) == len(pattern) || s[len(s)-len(pattern)-1] == '.'
}

func (m domainMatcher) String() string {
	return "domain:" + string(m)
}

type regexMatcher struct {
	pattern *regexp2.Regexp
}

func (m *regexMatcher) Match(s string) bool {
	r, e := m.pattern.MatchString(s)
	if e != nil {
		return false
	}
	return r
}

func (m *regexMatcher) String() string {
	return "regexp:" + m.pattern.String()
}
