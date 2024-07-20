package encoder

import (
	"errors"
	"fmt"
	"goserial/enums"
	"goserial/errorlist"
	"reflect"
)

func Encode(refVal reflect.Value) ([]byte, error) {
	kind := refVal.Kind()
	switch kind {
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
		return nil, errorlist.ErrUnsupportType
	}
}

func encodeBool(v reflect.Value) (valBytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			valBytes = nil
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	val := v.Bool()
	if val {
		return []byte{enums.BOOL, byte(enums.EncodeHeaderLen + 1), 1}, nil
	} else {
		return []byte{enums.BOOL, byte(enums.EncodeHeaderLen + 1), 0}, nil
	}
}

func encodeIntx(v reflect.Value) (valBytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			valBytes = nil
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
		return nil, errorlist.ErrUnknownIntType
	}
}

func encodeUintx(v reflect.Value) (valBytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			valBytes = nil
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
		return nil, errorlist.ErrUnknownUintType
	}
}

func encodeFloatx(v reflect.Value) (valBytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			valBytes = nil
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
		return nil, errorlist.ErrUnknownFloatType
	}
}

func encodeComplex(v reflect.Value) (valBytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			valBytes = nil
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
		return nil, errorlist.ErrUnknownComplexType
	}
}

func encodeArray(v reflect.Value) (valBytes []byte, err error) {
	length := v.Len()
	if length <= 0 {
		return []byte{enums.ARRAY, byte(enums.EncodeHeaderLen)}, nil
	} else if length > enums.MaxByteByInt {
		return nil, errorlist.ErrMaxLengthExceed
	}
	var content []byte
	var contentLength int = 0
	for i := 0; i < length; i++ {
		valI := v.Index(i)
		ret, err := Encode(valI)
		if err != nil {
			return nil, err
		}
		content = append(content, ret...)
		contentLength += len(content)
		if contentLength > enums.MaxByteByInt-enums.EncodeHeaderLen-enums.ArraySliceHeaderLen {
			return nil, errorlist.ErrMaxLengthExceed
		}
	}
	return append([]byte{enums.ARRAY, byte(contentLength + 3), byte(length)}, content...), nil
}

func encodeStruct(v reflect.Value) (valBytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}

func encodeString(v reflect.Value) (valBytes []byte, err error) {
	val := v.String()
	val2Bytes := []byte(val)
	length := len(val2Bytes)
	if length > enums.MaxByteByInt-enums.EncodeHeaderLen {
		return nil, errorlist.ErrMaxLengthExceed
	}
	return append([]byte{enums.STRING, byte(length + enums.EncodeHeaderLen)}, val2Bytes...), nil
}

func encodeSlice(v reflect.Value) (valBytes []byte, err error) {
	length := v.Len()
	if length <= 0 {
		return []byte{enums.SLICE, byte(enums.EncodeHeaderLen)}, nil
	} else if length > enums.MaxByteByInt {
		return nil, errorlist.ErrMaxLengthExceed
	}
	var content []byte
	var contentLength int = 0
	for i := 0; i < length; i++ {
		valI := v.Index(i)
		ret, err := Encode(valI)
		if err != nil {
			return nil, err
		}
		content = append(content, ret...)
		contentLength += len(content)
		if contentLength > enums.MaxByteByInt-enums.EncodeHeaderLen-enums.ArraySliceHeaderLen {
			return nil, errorlist.ErrMaxLengthExceed
		}
	}
	return append([]byte{enums.SLICE, byte(contentLength + enums.EncodeHeaderLen + enums.ArraySliceHeaderLen), byte(length)}, content...), nil
}

func encodeMap(v reflect.Value) (valBytes []byte, err error) {
	// todo
	return nil, errors.New("TODO")
}

func encodePtr(v reflect.Value) (valBytes []byte, err error) {
	if v.IsNil() {
		return []byte{enums.PTR, 2}, nil
	}
	elem, e := Encode(v.Elem())
	if e != nil {
		return nil, e
	}
	elemLength := len(elem)
	if elemLength > enums.MaxByteByInt-enums.EncodeHeaderLen {
		return nil, errorlist.ErrMaxLengthExceed
	}
	return append([]byte{enums.PTR, byte(elemLength + enums.EncodeHeaderLen)}, elem...), nil
}
