package encoder

import (
	"bytes"
	"encoding/binary"
	"goserial/enums"
)

func encodeComplex64(v complex64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, real(v))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, imag(v))
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.COMPLEX64, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}

func encodeComplex128(v complex128) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, real(v))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, imag(v))
	if err != nil {
		return nil, err
	}
	valBytes := buf.Bytes()
	return append([]byte{enums.COMPLEX128, byte(len(valBytes) + enums.EncodeHeaderLen)}, valBytes...), nil
}
