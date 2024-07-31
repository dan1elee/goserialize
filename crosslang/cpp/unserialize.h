#ifndef _UNSERIALIZE_H
#define _UNSERIALIZE_H

#include <cstdint>
#include <complex>

typedef uint8_t byte;

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

std::complex<float> readComplex64(byte *data, int len);
std::complex<double> readComplex128(byte *data, int len);
#endif
