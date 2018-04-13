package osin

import (
	"fmt"
)

func makeKey(namespace, id string) string {
	return fmt.Sprintf("%s:%s", namespace, id)
}

func assertToString(in interface{}) (string, error) {
	var ok bool
	var data string
	if in == nil {
		return "", nil
	} else if data, ok = in.(string); ok {
		return data, nil
	} else if str, ok := in.(fmt.Stringer); ok {
		return str.String(), nil
	}
	return "", AssertStringError
}
