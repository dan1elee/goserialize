package decoder

import (
	"errors"
	"fmt"
	"goserial/enums"
	"goserial/errorlist"
	"reflect"
)

func Decode(valBytes []byte, refVal reflect.Value) error {
	if valBytes == nil {
		return errorlist.ErrUnserializeFromNil
	}
	if len(valBytes) < 2 {
		return errorlist.ErrUnserializeFromUncomplete
	}
	switch valBytes[0] {
	case enums.BOOL:
		return decodeBool(valBytes, refVal)
	case enums.INT, enums.INT8, enums.INT16, enums.INT32, enums.INT64:
		return decodeIntx(valBytes, refVal)
	case enums.UINT, enums.UINT8, enums.UINT16, enums.UINT32, enums.UINT64:
		return decodeUintx(valBytes, refVal)
	case enums.FLOAT32, enums.FLOAT64:
		return decodeFloatx(valBytes, refVal)
	case enums.COMPLEX64, enums.COMPLEX128:
		return decodeComplex(valBytes, refVal)
	case enums.ARRAY:
		return decodeArray(valBytes, refVal)
	case enums.STRUCT:
		return decodeStruct(valBytes, refVal)
	case enums.STRING:
		return decodeString(valBytes, refVal)
	case enums.SLICE:
		return decodeSlice(valBytes, refVal)
	case enums.MAP:
		return decodeMap(valBytes, refVal)
	case enums.PTR:
		return decodePtr(valBytes, refVal)
	default:
		return errorlist.ErrUnserializeFromUnknownType
	}
}

func decodeBool(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) || length-enums.EncodeHeaderLen != 1 || (valBytes[2] != 0 && valBytes[2] != 1) {
		return errorlist.ErrUnserializeFromWrongForm
	}
	v.SetBool(valBytes[2] == 1)
	return nil
}

func decodeArray(valBytes []byte, v reflect.Value) error {
	return errors.New("todo")
}

func decodeStruct(valBytes []byte, v reflect.Value) error {
	return errors.New("todo")
}

func decodeString(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) {
		return errorlist.ErrUnserializeFromWrongForm
	}
	if length == 2 {
		v.SetString("")
	} else {
		v.SetString(string(valBytes[2:]))
	}
	return nil
}

func decodeSlice(valBytes []byte, v reflect.Value) error {
	return errors.New("todo")
}

func decodeMap(valBytes []byte, v reflect.Value) error {
	return errors.New("todo")
}

func decodePtr(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) {
		return errorlist.ErrUnserializeFromWrongForm
	}
	if length == 2 {
		v.SetPointer(nil)
	} else {
		Decode(valBytes[2:], v)
	}
	return errors.New("todo")
}
