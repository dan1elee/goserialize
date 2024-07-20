package goserialize

import (
	"goserialize/decoder"
	"goserialize/encoder"
	"goserialize/errorlist"
	"reflect"
)

func Serialize(v interface{}) ([]byte, error) {
	return encoder.Encode(reflect.ValueOf(v))
}

func Unserialize(data []byte, v interface{}) error {
	if v == nil {
		return errorlist.ErrUnserializeToNil
	}
	refVal := reflect.ValueOf(v)
	if refVal.IsNil() {
		return errorlist.ErrUnserializeToNil
	}
	if refVal.Kind() != reflect.Ptr {
		return errorlist.ErrUnserializeToNotPtr
	}
	return decoder.Decode(data, refVal.Elem())
}
