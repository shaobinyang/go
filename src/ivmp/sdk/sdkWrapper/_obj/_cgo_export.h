/* Created by cgo - DO NOT EDIT. */

#line 3 "/home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go"



#include <stdlib.h>
#include <string.h>
#include <hmt_net_sdk.h>

extern void messageCallback1(int handle, void* msg, void* user);
static inline void messageCallback_cgo(int handle, void* msg, void* user)
{
messageCallback1(handle, msg, user);
}




typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef __complex float GoComplex64;
typedef __complex double GoComplex128;

typedef struct { char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;


extern void messageCallback1(int p0, void* p1, void* p2);
