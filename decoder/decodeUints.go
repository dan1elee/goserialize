package decoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goserial/enums"
	"goserial/errorlist"
	"reflect"
)

func decodeUintx(valBytes []byte, v reflect.Value) error {
	switch valBytes[1] {
	case enums.UINT, enums.UINT64:
		return decodeUint(valBytes, v)
	case enums.UINT8:
		return decodeUint8(valBytes, v)
	case enums.UINT16:
		return decodeUint16(valBytes, v)
	case enums.UINT32:
		return decodeUint32(valBytes, v)
	default:
		return errorlist.ErrUnserializeFromUnknownType
	}
}

func decodeUint(valBytes []byte, v reflect.Value) (e error) {
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
	var num uint64
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetUint(num)
	return nil
}

func decodeUint8(valBytes []byte, v reflect.Value) (e error) {
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
	var num uint8
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetUint(uint64(num))
	return nil
}

func decodeUint16(valBytes []byte, v reflect.Value) (e error) {
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
	var num uint16
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetUint(uint64(num))
	return nil
}

func decodeUint32(valBytes []byte, v reflect.Value) (e error) {
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
	var num uint32
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		return err
	}
	v.SetUint(uint64(num))
	return nil
}
