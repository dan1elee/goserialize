package decoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goserial/enums"
	"goserial/errorlist"
	"reflect"
)

func decodeIntx(valBytes []byte, v reflect.Value) error {
	switch valBytes[1] {
	case enums.INT, enums.INT64:
		return decodeInt(valBytes, v)
	case enums.INT8:
		return decodeInt8(valBytes, v)
	case enums.INT16:
		return decodeInt16(valBytes, v)
	case enums.INT32:
		return decodeInt32(valBytes, v)
	default:
		return errorlist.ErrUnserializeFromUnknownType
	}
}

func decodeInt(valBytes []byte, v reflect.Value) (e error) {
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
	var num int64
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetInt(num)
	return nil
}

func decodeInt8(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) || length-enums.EncodeHeaderLen != 1 {
		return errorlist.ErrUnserializeFromWrongForm
	}
	buf := bytes.NewReader(valBytes[2:])
	var num int8
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetInt(int64(num))
	return nil
}

func decodeInt16(valBytes []byte, v reflect.Value) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("panic: %v", r)
		}
	}()
	length := len(valBytes)
	if length != int(valBytes[1]) || length-enums.EncodeHeaderLen != 2 {
		return errorlist.ErrUnserializeFromWrongForm
	}
	buf := bytes.NewReader(valBytes[2:])
	var num int16
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetInt(int64(num))
	return nil
}

func decodeInt32(valBytes []byte, v reflect.Value) (e error) {
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
	var num int32
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetInt(int64(num))
	return nil
}
