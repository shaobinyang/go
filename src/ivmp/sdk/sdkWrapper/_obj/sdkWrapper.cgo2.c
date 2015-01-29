#line 3 "/home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go"



#include <stdlib.h>
#include <string.h>
#include <hmt_net_sdk.h>

extern void messageCallback1(int handle, void* msg, void* user);
static inline void messageCallback_cgo(int handle, void* msg, void* user)
{
messageCallback1(handle, msg, user);
}




// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

extern char* _cgo_topofstack(void);

#include <errno.h>
#include <string.h>

void
_cgo_0d4ddd74268d_Cfunc_free(void *v)
{
	struct {
		void* p0;
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	free((void*)a->p0);
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetInit(void *v)
{
	struct {
		int r;
		char __pad4[4];
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetInit();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetLogin(void *v)
{
	struct {
		char* p0;
		HmtNetUser* p1;
		HMT_NET_HANDLE r;
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetLogin((void*)a->p0, (void*)a->p1);
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetLogout(void *v)
{
	struct {
		HMT_NET_HANDLE p0;
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	hmtNetLogout(a->p0);
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformID(void *v)
{
	struct {
		HMT_NET_HANDLE p0;
		char* p1;
		int* p2;
		int r;
		char __pad28[4];
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetQueryPlatformID(a->p0, (void*)a->p1, (void*)a->p2);
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformName(void *v)
{
	struct {
		HMT_NET_HANDLE p0;
		char* p1;
		int* p2;
		int r;
		char __pad28[4];
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetQueryPlatformName(a->p0, (void*)a->p1, (void*)a->p2);
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetQueryUserInfo(void *v)
{
	struct {
		HMT_NET_HANDLE p0;
		HmtNetUserInfo* p1;
		int r;
		char __pad20[4];
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetQueryUserInfo(a->p0, (void*)a->p1);
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetRelease(void *v)
{
	struct {
		int r;
		char __pad4[4];
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetRelease();
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetSetMessageCallback(void *v)
{
	struct {
		HMT_NET_HANDLE p0;
		HmtNetMessageProc p1;
		void* p2;
		int r;
		char __pad28[4];
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetSetMessageCallback(a->p0, (void*)a->p1, (void*)a->p2);
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_hmtNetSubscribe(void *v)
{
	struct {
		HMT_NET_HANDLE p0;
		char* p1;
		char* p2;
		int p3;
		char __pad28[4];
		int r;
		char __pad36[4];
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = hmtNetSubscribe(a->p0, (void*)a->p1, (void*)a->p2, a->p3);
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

void
_cgo_0d4ddd74268d_Cfunc_strcpy(void *v)
{
	struct {
		char* p0;
		char* p1;
		const char* r;
	} __attribute__((__packed__, __gcc_struct__)) *a = v;
	char *stktop = _cgo_topofstack();
	__typeof__(a->r) r = (__typeof__(a->r)) strcpy((void*)a->p0, (void*)a->p1);
	a = (void*)((char*)a + (_cgo_topofstack() - stktop));
	a->r = r;
}

