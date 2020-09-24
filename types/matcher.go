package types

type Matcher interface {
	Match(actual interface{}) (bool, error)
	Message(actual interface{}) (message string)
}
