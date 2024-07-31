#ifndef _UNSERIALIZE_H
#define _UNSERIALIZE_H

#include <cstdint>
#include <complex>

typedef uint8_t byte;
typedef std::complex<float> complex64;
typedef std::complex<double> complex128;

class UnserializeException : public std::exception
{
public:
    UnserializeException() : type(0), message(type2Str(0)) {}
    UnserializeException(int type) : type(type), message(type2Str(type)) {}
    ~UnserializeException() throw()
    {
    }
    virtual const char *what() const noexcept override;

private:
    int type;
    std::string message;
    std::string type2Str(int type) const;
};

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

typedef union unionval UnionVal;

typedef struct
{
    int len;
    UnionVal *valList;
} Arr;

typedef union unionval
{
    bool boolVal;
    int8_t int8Val;
    int16_t int16Val;
    int32_t int32Val;
    int64_t int64Val;
    uint8_t uint8Val;
    uint16_t uint16Val;
    uint32_t uint32Val;
    uint64_t uint64Val;
    float float32Val;
    double float64Val;
    complex64 complex64Val;
    complex128 complex128Val;
    Arr arrVal;
    // TODO
} UnionVal;

typedef struct
{
    enums::Types type;
    UnionVal val;
} typeAndVal;

enums::Types getDataType(byte *data, int len);

void unserialize(byte *data, int len, void *result);

bool readBool(byte *data, int len);

int64_t readInt(byte *data, int len);
int8_t readInt8(byte *data, int len);
int16_t readInt16(byte *data, int len);
int32_t readInt32(byte *data, int len);
int64_t readInt64(byte *data, int len);

uint64_t readUint(byte *data, int len);
uint8_t readUint8(byte *data, int len);
uint16_t readUint16(byte *data, int len);
uint32_t readUint32(byte *data, int len);
uint64_t readUint64(byte *data, int len);

float readFloat32(byte *data, int len);
double readFloat64(byte *data, int len);

complex64 readComplex64(byte *data, int len);
complex128 readComplex128(byte *data, int len);

#endif
