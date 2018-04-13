package env

import (
	"fmt"
	"reflect"
)

type InvalidInjectError struct {
	Type reflect.Type
}

func (ie *InvalidInjectError) Error() string {
	return fmt.Sprintf("Except for pointer type, type of %s is receive", ie.Type.String())
}

func (ie *InvalidInjectError) RuntimeError() {}
