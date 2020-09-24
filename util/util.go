package util

import "fmt"

func Pretty(object interface{}) string {
	if str, ok := object.(string); ok {
		return str
	}
	return fmt.Sprintf("%+v", object)
}
