package decoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goserialize/enums"
	"goserialize/errorlist"
	"reflect"
)

func decodeComplex(valBytes []byte, v reflect.Value) error {
	switch valBytes[1] {
	case enums.COMPLEX64:
		return decodeComplex64(valBytes, v)
	case enums.COMPLEX128:
		return decodeComplex128(valBytes, v)
	default:
		return errorlist.ErrUnserializeFromUnknownType
	}
}

func decodeComplex64(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) || length-enums.EncodeHeaderLen != 8 {
		return errorlist.ErrUnserializeFromWrongForm
	}
	buf := bytes.NewReader(valBytes[2:])
	var real float32
	err := binary.Read(buf, binary.LittleEndian, &real)
	if err != nil {
		return err
	}
	var imag float32
	err = binary.Read(buf, binary.LittleEndian, &imag)
	if err != nil {
		return err
	}
	complexValue := complex(real, imag)
	v.SetComplex(complex128(complexValue))
	return nil
}

func decodeComplex128(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) || length-enums.EncodeHeaderLen != 8 {
		return errorlist.ErrUnserializeFromWrongForm
	}
	buf := bytes.NewReader(valBytes[2:])
	var real float64
	err := binary.Read(buf, binary.LittleEndian, &real)
	if err != nil {
		return err
	}
	var imag float64
	err = binary.Read(buf, binary.LittleEndian, &imag)
	if err != nil {
		return err
	}
	complexValue := complex(real, imag)
	v.SetComplex(complexValue)
	return nil
}
