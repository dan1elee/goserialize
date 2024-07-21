from enum import Enum, auto
import struct


class Type(Enum):
    NIL = 0
    BOOL = auto()
    INT = auto()
    UINT = auto()
    INT8 = auto()
    UINT8 = auto()
    INT16 = auto()
    UINT16 = auto()
    INT32 = auto()
    UINT32 = auto()
    INT64 = auto()
    UINT64 = auto()
    FLOAT32 = auto()
    FLOAT64 = auto()
    COMPLEX64 = auto()
    COMPLEX128 = auto()
    # UINTPTR=auto()
    ARRAY = auto()
    STRUCT = auto()
    STRING = auto()
    SLICE = auto()
    MAP = auto()
    PTR = auto()
    ENDOFTYPE = auto()


def decode(data: bytes):
    length = len(data)
    print(length)
    if length < 2:
        return None
    elif data[1] != length:
        return None
    if data[0] == Type.INT32.value:
        return struct.unpack('<i', data[2:])[0]

