package assert

import (
	"errors"
	"fmt"
	"github.com/gaminggroup/gobdd/types"
	"github.com/gaminggroup/gobdd/util"
	"reflect"
)

func ToEqual(expected interface{}) types.Matcher {
	return &EqualMatcher{
		Expected: expected,
	}
}

type EqualMatcher struct {
	types.Matcher
	Expected interface{}
}

func (m *EqualMatcher) Match(actual interface{}) (bool, error) {
	if m.Expected == nil {
		return false, errors.New("refused to compare to nil for expected value")
	}
	return reflect.DeepEqual(actual, m.Expected), nil
}

func (m *EqualMatcher) Message(actual interface{}) string {
	return fmt.Sprintf("%s should equal %s", util.Pretty(actual), util.Pretty(m.Expected))
}
