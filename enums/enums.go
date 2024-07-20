package enums

// TYPES
const (
	NIL byte = iota
	BOOL
	INT
	UINT
	INT8
	UINT8
	INT16
	UINT16
	INT32
	UINT32
	INT64
	UINT64
	FLOAT32
	FLOAT64
	COMPLEX64
	COMPLEX128
	// UINTPTR
	ARRAY
	STRUCT
	STRING
	SLICE
	MAP
	PTR
	ENDOFTYPE
)

// MAX
const (
	MaxByte      byte = 0xFF
	MaxByteByInt int  = 0xFF
)

//LENGTH
const (
	EncodeHeaderLen     int = 2
	ArraySliceHeaderLen int = 1
)
