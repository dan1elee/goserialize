package encoder

import (
	"bytes"
	"encoding/binary"
	"goserial/enums"
)

func encodeFloat32(v float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.FLOAT32, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeFloat64(v float64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.FLOAT64, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}
