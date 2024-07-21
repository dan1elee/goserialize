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
    if length < 2:
        return None
    elif data[1] != length:
        return None
    if data[0]==Type.BOOL.value:
        return struct.unpack('<?', data[2:])[0]
    elif data[0] == Type.INT.value:
        return struct.unpack('<q', data[2:])[0]
    elif data[0] == Type.INT8.value:
        return struct.unpack('<b', data[2:])[0]
    elif data[0] == Type.INT16.value:
        return struct.unpack('<h', data[2:])[0]
    elif data[0] == Type.INT32.value:
        return struct.unpack('<i', data[2:])[0]
    elif data[0] == Type.INT64.value:
        return struct.unpack('<q', data[2:])[0]
    elif data[0] == Type.UINT.value:
        return struct.unpack('<Q', data[2:])[0]
    elif data[0] == Type.UINT8.value:
        return struct.unpack('<B', data[2:])[0]
    elif data[0] == Type.UINT16.value:
        return struct.unpack('<H', data[2:])[0]
    elif data[0] == Type.UINT32.value:
        return struct.unpack('<I', data[2:])[0]
    elif data[0] == Type.UINT64.value:
        return struct.unpack('<Q', data[2:])[0]
    elif data[0] == Type.FLOAT32.value:
        return struct.unpack('<f', data[2:])[0]
    elif data[0] == Type.FLOAT64.value:
        return struct.unpack('<d', data[2:])[0]
    else:
        return None
