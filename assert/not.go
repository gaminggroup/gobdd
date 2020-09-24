package assert

import (
	"fmt"
	"github.com/gaminggroup/gobdd/types"
)

func Not(matcher types.Matcher) types.Matcher {
	return &NotMatcher{
		matcher: matcher,
	}
}

type NotMatcher struct {
	types.Matcher
	matcher types.Matcher
}

func (m *NotMatcher) Match(actual interface{}) (bool, error) {
	success, err := m.matcher.Match(actual)
	return !success, err
}

func (m *NotMatcher) Message(actual interface{}) string {
	return fmt.Sprintf("Not(%s)", m.matcher.Message(actual))
}

