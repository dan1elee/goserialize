package encoder

import (
	"errors"
	"fmt"
	"goserial/enums"
	"reflect"
)

func Encode(v interface{}) ([]byte, error) {
	refVal := reflect.ValueOf(v)
	kind := refVal.Kind()
	switch kind {
	case reflect.Invalid:
		if v == nil {
			return []byte{enums.NIL, 2}, nil
		} else {
			return nil, errors.New("invalid value")
		}

	case reflect.Bool:
		return encodeBool(refVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return encodeIntx(refVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return encodeUintx(refVal)
	case reflect.Float32, reflect.Float64:
		return encodeFloatx(refVal)
	case reflect.Complex64, reflect.Complex128:
		return encodeComplex(refVal)
	case reflect.Array:
		return encodeArray(refVal)
	case reflect.Struct:
		return encodeStruct(refVal)
	case reflect.String:
		return encodeString(refVal)
	case reflect.Slice:
		return encodeSlice(refVal)
	case reflect.Map:
		return encodeMap(refVal)
	case reflect.Ptr:
		return encodePtr(refVal)
	default:
		return nil, errors.New("unsupported type")
	}
}

func encodeBool(v reflect.Value) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			bytes = nil
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	val := v.Bool()
	if val {
		return []byte{enums.BOOL, 3, 1}, nil
	} else {
		return []byte{enums.BOOL, 3, 0}, nil
	}
}

func encodeIntx(v reflect.Value) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			bytes = nil
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	val := v.Int()
	switch v.Kind() {
	case reflect.Int:
		return encodeInt(int(val))
	case reflect.Int8:
		return encodeInt8(int8(val))
	case reflect.Int16:
		return encodeInt16(int16(val))
	case reflect.Int32:
		return encodeInt32(int32(val))
	case reflect.Int64:
		return encodeInt64(int64(val))
	default:
		return nil, errors.New("unknown int type")
	}
}

func encodeUintx(v reflect.Value) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			bytes = nil
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	val := v.Uint()
	switch v.Kind() {
	case reflect.Int:
		return encodeUint(uint(val))
	case reflect.Int8:
		return encodeUint8(uint8(val))
	case reflect.Int16:
		return encodeUint16(uint16(val))
	case reflect.Int32:
		return encodeUint32(uint32(val))
	case reflect.Int64:
		return encodeUint64(uint64(val))
	default:
		return nil, errors.New("unknown uint type")
	}
}

func encodeFloatx(v reflect.Value) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			bytes = nil
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	val := v.Float()
	switch v.Kind() {
	case reflect.Float32:
		return encodeFloat32(float32(val))
	case reflect.Float64:
		return encodeFloat64(float64(val))
	default:
		return nil, errors.New("unknown float type")
	}
}

func encodeComplex(v reflect.Value) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			bytes = nil
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	val := v.Complex()
	switch v.Kind() {
	case reflect.Complex64:
		return encodeComplex64(complex64(val))
	case reflect.Complex128:
		return encodeComplex128(complex128(val))
	default:
		return nil, errors.New("unknown complex type")
	}
}

func encodeArray(v reflect.Value) (bytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}

func encodeStruct(v reflect.Value) (bytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}

func encodeString(v reflect.Value) (bytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}

func encodeSlice(v reflect.Value) (bytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}

func encodeMap(v reflect.Value) (bytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}

func encodePtr(v reflect.Value) (bytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}
