from enum import Enum, auto
from typing import List, Any
from exceptions.exceptions import WrongFormException, WrongFormErrorType
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
        raise WrongFormException(WrongFormErrorType.LengthTooShort)
    elif data[1] != length:
        raise WrongFormException(WrongFormErrorType.LengthNotEqual)
    if data[0] == Type.BOOL.value:
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
    elif data[0] == Type.COMPLEX64.value:
        real = struct.unpack('<f', data[2:6])[0]
        imag = struct.unpack('<f', data[6:10])[0]
        return complex(real, imag)
    elif data[0] == Type.COMPLEX128.value:
        real = struct.unpack('<d', data[2:10])[0]
        imag = struct.unpack('<d', data[10:18])[0]
        return complex(real, imag)
    elif data[0] == Type.ARRAY.value:
        return decodeArray(data)
    elif data[0] == Type.STRING.value:
        return data[2:].decode('utf-8')
    else:
        raise WrongFormException(WrongFormErrorType.TypeNotSupport)


def decodeArray(data: bytes) -> List[Any]:
    ret = list()
    length = len(data)
    actualLen = data[2]
    if actualLen != 0:
        elemSize = (length - 3) // actualLen
        for i in range(actualLen):
            ret.append(decode(data[3+elemSize*i:3+elemSize*(i+1)]))
    return ret
