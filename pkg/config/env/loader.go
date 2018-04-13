package env

import (
	"fmt"
	"os"
	"reflect"
)

const tagName = "env"

func InjectWithEnv(v interface{}) error {
	reflectValue := reflect.ValueOf(v)
	reflectType := reflect.TypeOf(v)

	if !validateInterface(reflectType) || reflectValue.IsNil() {
		return &InvalidInjectError{reflectType}
	}

	ele := reflectValue.Elem()
	for i, max := 0, ele.NumField(); i < max; i += 1 {
		field := ele.Field(i)
		fieldTag, hasTag := reflectType.Elem().Field(i).Tag.Lookup(tagName)
		fieldName := reflectType.Elem().Field(i).Name

		if !hasTag {
			continue
		}

		if !field.CanSet() {
			return fmt.Errorf("fail to set field %s", fieldName)
		}

		env := os.Getenv(fieldTag)
		field.Set(reflect.ValueOf(env))
	}
	return nil
}
