package decoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goserial/enums"
	"goserial/errorlist"
	"reflect"
)

func decodeFloatx(valBytes []byte, v reflect.Value) error {
	switch valBytes[1] {
	case enums.FLOAT32:
		return decodeFloat32(valBytes, v)
	case enums.FLOAT64:
		return decodeFloat64(valBytes, v)
	default:
		return errorlist.ErrUnserializeFromUnknownType
	}
}

func decodeFloat32(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) || length-enums.EncodeHeaderLen != 4 {
		return errorlist.ErrUnserializeFromWrongForm
	}
	buf := bytes.NewReader(valBytes[2:])
	var num float32
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetFloat(float64(num))
	return nil
}

func decodeFloat64(valBytes []byte, v reflect.Value) (e error) {
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
	var num float64
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetFloat(num)
	return nil
}
