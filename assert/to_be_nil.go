package assert

import (
	"fmt"
	"github.com/gaminggroup/gobdd/types"
	"github.com/gaminggroup/gobdd/util"
)

func ToBeNil() types.Matcher {
	return &NilMatcher{}
}

type NilMatcher struct {
	types.Matcher
}

func (m *NilMatcher) Match(actual interface{}) (bool, error) {
	return actual == nil, nil
}

func (m *NilMatcher) Message(actual interface{}) string {
	return fmt.Sprintf("%s should be nil", util.Pretty(actual))
}
