package encoder

import (
	"bytes"
	"encoding/binary"
	"goserial/enums"
	"unsafe"
)

func encodeInt(v int) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	intSize := unsafe.Sizeof(v)
	if intSize == 4 {
		err = binary.Write(buf, binary.LittleEndian, int32(v))
	} else {
		err = binary.Write(buf, binary.LittleEndian, int64(v))
	}
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.INT, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeInt8(v int8) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.INT8, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeInt16(v int16) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.INT16, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeInt32(v int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.INT32, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeInt64(v int64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.INT64, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}
