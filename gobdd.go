package gobdd

import (
	"fmt"
	"github.com/gaminggroup/gobdd/types"
	"runtime"
	"testing"
)

var current *step
var suites []*suite

type suite struct {
	body func()
}

type step struct {
	name     string
	text     string
	parent   *step
	children []*step
	specs    []*spec
	body     func()
}

type spec struct {
	run         func() (bool, error)
	description func() string
	trace       string
}

type fail struct {
	message string
	trace string
}

func IsRoot() bool {
	return current == nil
}

func NewRoot(name string, text string, body func()) {
	current = &step{
		name: name,
		text: text,
		parent: nil,
		children: []*step{},
		specs: []*spec{},
		body: body,
	}
}

func AddStep(name string, text string, body func()) {
	current.children = append(current.children, &step{
		name: name,
		text: text,
		parent: current,
		children: []*step{},
		specs: []*spec{},
		body: body,
	})
}

func RunSuites(t *testing.T) {
	for _, suite := range suites {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic: ", r)
					t.Fail()
					_, fn, line, _ := runtime.Caller(3)
					fmt.Println(fmt.Sprintf("%s:%d", fn, line))
				}
			}()
			current = nil
			suite.body()
			_, fails := iterateSteps(t)
			if len(fails) > 0 {
				//fmt.Println(fmt.Sprintf("%d of %d tests failed", len(fails), tests))
				for _, fail := range fails {
					fmt.Println(*getDescription(current))
					fmt.Println("    FAIL: " + fail.message)
					fmt.Println("        " + fail.trace)
				}
				//fmt.Println("")
				t.Fail()
			}
		}()
	}
	suites = nil
}

func iterateSteps(t *testing.T) (int, []*fail) {
	tests := 0
	var fails []*fail
	current.body()
	for _, spec := range current.specs {
		tests++
		success, err := spec.run()
		if err != nil {
			fails = append(fails, &fail{
				message: "Unable to run spec: " + err.Error(),
				trace: spec.trace,
			})
		} else if !success {
			fails = append(fails, &fail{
				message: spec.description(),
				trace: spec.trace,
			})
		}
	}
	for _, step := range current.children {
		current = step
		t, f := iterateSteps(t)
		tests += t
		fails = append(fails, f...)
	}
	return tests, fails
}

func getDescription(step *step) *string {
	thisStep := step
	description := ""
	for {
		if thisStep == nil {
			break
		}
		description = thisStep.name + " " + thisStep.text + " " + description
		thisStep = thisStep.parent
	}
	return &description
}

func Suite(body func()) {
	suites = append(suites, &suite{
		body: body,
	})
}

func Expect(actual interface{}, matcher types.Matcher) {
	_, fn, line, _ := runtime.Caller(1)
	current.specs = append(current.specs, &spec{
		run: func() (bool, error) {
			success, err := matcher.Match(actual)
			return success, err
		},
		description: func() string {
			description := matcher.Message(actual)
			return description
		},
		trace: fmt.Sprintf("%s:%d", fn, line),
	})
}
