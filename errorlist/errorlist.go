package errorlist

import "errors"

var (
	// for encoder
	ErrMaxLengthExceed    = errors.New("max length exceeded")
	ErrUnsupportType      = errors.New("unsupported type")
	ErrUnknownIntType     = errors.New("unknown int type")
	ErrUnknownUintType    = errors.New("unknown uint type")
	ErrUnknownFloatType   = errors.New("unknown float type")
	ErrUnknownComplexType = errors.New("unknown complex type")

	// for decoder
	ErrUnserializeToNil           = errors.New("unserialize to nil")
	ErrUnserializeToNotPtr        = errors.New("unserialize to not a ptr")
	ErrUnserializeFromNil         = errors.New("unserialize from nil")
	ErrUnserializeFromUncomplete  = errors.New("unserialize from uncomplete")
	ErrUnserializeFromUnknownType = errors.New("unserialize from unknown type")
	ErrUnserializeFromWrongForm   = errors.New("unserialize from wrong form")
)
