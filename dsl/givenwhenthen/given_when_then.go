package givenwhenthen

import (
	"github.com/gaminggroup/gobdd"
)

const (
	feature    = "FEATURE"
	background = "BACKGROUND"
	scenario   = "SCENARIO"
	given      = "GIVEN"
	andGiven   = "AND GIVEN"
	when       = "WHEN"
	andWhen    = "AND WHEN"
	then       = "THEN"
	andThen    = "AND THEN"
)

func Feature(text string, body func()) {
	if !gobdd.IsRoot() {
		panic("Feature must be added as the first element in a Suite")
	}
	gobdd.NewRoot(feature, text + ":", body)
}

func Background(text string, body func()) {
	if gobdd.IsRoot() {
		panic("Background cannot be added as the first element in a Suite")
	}
	gobdd.AddStep(background, text + ":", body)
}

func Scenario(text string, body func()) {
	if gobdd.IsRoot() {
		panic("Scenario cannot be added as the first element in a Suite")
	}
	gobdd.AddStep(scenario, text + ":", body)
}

func Given(text string, body func()) {
	if gobdd.IsRoot() {
		panic("Given cannot be added as the first element in a Suite")
	}
	gobdd.AddStep(given, text, body)
}

func AndGiven(text string, body func()) {
	if gobdd.IsRoot() {
		panic("AndGiven cannot be added as the first element in a Suite")
	}
	Given(text, body)
}

func When(text string, body func()) {
	if gobdd.IsRoot() {
		panic("When cannot be added as the first element in a Suite")
	}
	gobdd.AddStep(when, text, body)
}

func AndWhen(text string, body func()) {
	if gobdd.IsRoot() {
		panic("AndWhen cannot be added as the first element in a Suite")
	}
	gobdd.AddStep(andWhen, text, body)
}

func Then(text string, body func()) {
	if gobdd.IsRoot() {
		panic("Then cannot be added as the first element in a Suite")
	}
	gobdd.AddStep(then, text, body)
}

func AndThen(text string, body func()) {
	Then(text, body)
}

