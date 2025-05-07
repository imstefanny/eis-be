package helpers

import (
	"reflect"
)

func IsEmptyField(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
