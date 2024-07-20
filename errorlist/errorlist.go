package errorlist

import "errors"

var (
	ErrMaxLengthExceed    = errors.New("max length exceeded")
	ErrUnsupportType      = errors.New("unsupported type")
	ErrUnknownIntType     = errors.New("unknown int type")
	ErrUnknownUintType    = errors.New("unknown uint type")
	ErrUnknownFloatType   = errors.New("unknown float type")
	ErrUnknownComplexType = errors.New("unknown complex type")
)
