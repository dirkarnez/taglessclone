package taglessclone

import (
	"reflect"
	"errors"
)

func SetField(toPtr interface{}, value interface{}) error {
	reflectValue := reflect.ValueOf(value)
	toReflectValue := reflect.ValueOf(toPtr)

	if toReflectValue.Kind() != reflect.Ptr {
		return errors.New("should be a pointer")
	}

	if reflectValue.Kind() == reflect.Ptr {
		return errors.New("should not be a pointer")
	}

	toReflectValue = toReflectValue.Elem()

	if !toReflectValue.CanSet() {
		return errors.New("cannot set value")
	}

	if toReflectValue.Type() != reflectValue.Type() {
		return errors.New("type does not match")
	}

	toReflectValue.Set(reflectValue)

	return nil
}
