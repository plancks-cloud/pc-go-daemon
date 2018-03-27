package util

import (
	"reflect"
)

func GetType(obj interface{}) (res string) {
	t := reflect.TypeOf(obj)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
