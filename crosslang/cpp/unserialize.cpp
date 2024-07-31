#include "unserialize.h"
#include <exception>
#include <string>

const char *UnserializeException::what() const noexcept
{
    return message.c_str();
}

std::string UnserializeException::type2Str(int type) const
{
    std::string ret;
    switch (type)
    {
    case 0:
    {
        ret = "Unspecified Error";
        break;
    }
    case 1:
    {
        ret = "UnserializeFromNull";
        break;
    }
    case 2:
    {
        ret = "UnserializeToNull";
        break;
    }
    case 3:
    {
        ret = "UnserializeFromUncomplete";
        break;
    }
    case 4:
    {
        ret = "UnserializeFromWrongForm";
        break;
    }
    case 5:
    {
        ret = "UnserializeFromUnsupportType";
        break;
    }
    default:
    {
        ret = "WrongErrorType";
        break;
    }
    }
    return ret;
}

namespace enums
{
    enum Types
    {
        NIL,
        BOOL,
        INT,
        UINT,
        INT8,
        UINT8,
        INT16,
        UINT16,
        INT32,
        UINT32,
        INT64,
        UINT64,
        FLOAT32,
        FLOAT64,
        COMPLEX64,
        COMPLEX128,
        // UINTPTR
        ARRAY,
        STRUCT,
        STRING,
        SLICE,
        MAP,
        PTR,
        ENDOFTYPE
    };
    const int EncodeHeaderLen = 2;
}

void unserialize(byte *data, int len, void *result)
{
    if (data == nullptr)
    {
        throw UnserializeException(1);
    }
    else if (data == nullptr)
    {
        throw UnserializeException(2);
    }
    if (len < 2)
    {
        throw UnserializeException(3);
    }
    switch (data[0])
    {
    case enums::BOOL:
    {
        bool res = readBool(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::INT:
    {
        int64_t res = readInt(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::INT8:
    {
        int8_t res = readInt8(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::INT16:
    {
        int16_t res = readInt16(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::INT32:
    {
        int32_t res = readInt32(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::INT64:
    {
        int64_t res = readInt(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::UINT:
    {
        uint64_t res = readUint(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::UINT8:
    {
        uint8_t res = readUint8(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::UINT16:
    {
        uint16_t res = readUint16(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::UINT32:
    {
        uint32_t res = readUint32(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::UINT64:
    {
        uint64_t res = readUint(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::FLOAT32:
    {
        float res = readFloat32(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::FLOAT64:
    {
        double res = readFloat64(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::COMPLEX64:
    {
        std::complex<float> res = readComplex64(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    case enums::COMPLEX128:
    {
        std::complex<double> res = readComplex128(data, len);
        memcpy(result, &res, sizeof(res));
        break;
    }
    default:
    {
        throw UnserializeException(5);
    }
    }
}

bool readBool(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 1 || (data[2] != 0 && data[2] != 1))
    {
        throw UnserializeException(4);
    }
    return data[2] == 1;
}

int64_t readInt(byte *data, int len)
{
    return (int64_t)readUint(data, len);
}

int8_t readInt8(byte *data, int len)
{
    return (int8_t)readUint8(data, len);
}

int16_t readInt16(byte *data, int len)
{
    return (int16_t)readUint16(data, len);
}

int32_t readInt32(byte *data, int len)
{
    return (int32_t)readUint32(data, len);
}

int64_t readInt64(byte *data, int len)
{
    return (int64_t)readUint64(data, len);
}

uint64_t readUint(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 8)
    {
        throw UnserializeException(4);
    }
    uint64_t ret = 0;
    for (int i = 0; i < 8; ++i)
    {
        ret = (ret << 8) + data[2 + i];
    }
    return ret;
}

uint8_t readUint8(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 1)
    {
        throw UnserializeException(4);
    }
    return data[2];
}

uint16_t readUint16(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 2)
    {
        throw UnserializeException(4);
    }
    uint16_t ret = 0;
    for (int i = 0; i < 2; ++i)
    {
        ret = (ret << 8) + data[2 + i];
    }
    return ret;
}

uint32_t readUint32(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 4)
    {
        throw UnserializeException(4);
    }
    uint32_t ret = 0;
    for (int i = 0; i < 4; ++i)
    {
        ret = (ret << 8) + data[2 + i];
    }
    return ret;
}

uint64_t readUint64(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 8)
    {
        throw UnserializeException(4);
    }
    uint64_t ret = 0;
    for (int i = 0; i < 8; ++i)
    {
        ret = (ret << 8) + data[2 + i];
    }
    return ret;
}

float readFloat32(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 4)
    {
        throw UnserializeException(4);
    }
    byte value[4];
    memcpy(value, data + 2, 4);
    return *((float *)value);
}

double readFloat64(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 8)
    {
        throw UnserializeException(4);
    }
    byte value[8];
    memcpy(value, data + 2, 8);
    return *((double *)value);
}

std::complex<float> readComplex64(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 8)
    {
        throw UnserializeException(4);
    }
    byte value[4];
    memcpy(value, data + 2, 4);
    float real = *((float *)value);
    memcpy(value, data + 6, 4);
    float imag = *((float *)value);
    return std::complex<float>(real, imag);
}

std::complex<double> readComplex128(byte *data, int len)
{
    if (len != (int)(data[1]) || len - enums::EncodeHeaderLen != 16)
    {
        throw UnserializeException(4);
    }
    byte value[8];
    memcpy(value, data + 2, 8);
    double real = *((double *)value);
    memcpy(value, data + 10, 8);
    double imag = *((double *)value);
    return std::complex<double>(real, imag);
}