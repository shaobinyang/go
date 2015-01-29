
#include "runtime.h"
#include "cgocall.h"
#include "textflag.h"

#pragma dataflag NOPTR
static void *cgocall_errno = runtime·cgocall_errno;
#pragma dataflag NOPTR
void *·_cgo_runtime_cgocall_errno = &cgocall_errno;

#pragma dataflag NOPTR
static void *runtime_gostring = runtime·gostring;
#pragma dataflag NOPTR
void *·_cgo_runtime_gostring = &runtime_gostring;

#pragma dataflag NOPTR
static void *runtime_gostringn = runtime·gostringn;
#pragma dataflag NOPTR
void *·_cgo_runtime_gostringn = &runtime_gostringn;

#pragma dataflag NOPTR
static void *runtime_gobytes = runtime·gobytes;
#pragma dataflag NOPTR
void *·_cgo_runtime_gobytes = &runtime_gobytes;

#pragma dataflag NOPTR
static void *runtime_cmalloc = runtime·cmalloc;
#pragma dataflag NOPTR
void *·_cgo_runtime_cmalloc = &runtime_cmalloc;

void ·_Cerrno(void*, int32);
#pragma cgo_import_static messageCallback_cgo
extern byte *messageCallback_cgo;
#pragma dataflag NOPTR /* C pointer, not heap pointer */ 
void *·_Cfpvar_fp_messageCallback_cgo = messageCallback_cgo;


#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_free
void _cgo_0d4ddd74268d_Cfunc_free(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_free = _cgo_0d4ddd74268d_Cfunc_free;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetInit
void _cgo_0d4ddd74268d_Cfunc_hmtNetInit(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetInit = _cgo_0d4ddd74268d_Cfunc_hmtNetInit;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetLogin
void _cgo_0d4ddd74268d_Cfunc_hmtNetLogin(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetLogin = _cgo_0d4ddd74268d_Cfunc_hmtNetLogin;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetLogout
void _cgo_0d4ddd74268d_Cfunc_hmtNetLogout(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetLogout = _cgo_0d4ddd74268d_Cfunc_hmtNetLogout;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformID
void _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformID(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformID = _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformID;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformName
void _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformName(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformName = _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformName;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetQueryUserInfo
void _cgo_0d4ddd74268d_Cfunc_hmtNetQueryUserInfo(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetQueryUserInfo = _cgo_0d4ddd74268d_Cfunc_hmtNetQueryUserInfo;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetRelease
void _cgo_0d4ddd74268d_Cfunc_hmtNetRelease(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetRelease = _cgo_0d4ddd74268d_Cfunc_hmtNetRelease;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetSetMessageCallback
void _cgo_0d4ddd74268d_Cfunc_hmtNetSetMessageCallback(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetSetMessageCallback = _cgo_0d4ddd74268d_Cfunc_hmtNetSetMessageCallback;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_hmtNetSubscribe
void _cgo_0d4ddd74268d_Cfunc_hmtNetSubscribe(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_hmtNetSubscribe = _cgo_0d4ddd74268d_Cfunc_hmtNetSubscribe;
#pragma cgo_import_static _cgo_0d4ddd74268d_Cfunc_strcpy
void _cgo_0d4ddd74268d_Cfunc_strcpy(void*);
#pragma dataflag NOPTR
void *·_cgo_0d4ddd74268d_Cfunc_strcpy = _cgo_0d4ddd74268d_Cfunc_strcpy;
#pragma cgo_export_dynamic messageCallback1
extern void ·messageCallback1();

#pragma cgo_export_static _cgoexp_0d4ddd74268d_messageCallback1
#pragma textflag 7
void
_cgoexp_0d4ddd74268d_messageCallback1(void *a, int32 n)
{
	runtime·cgocallback(·messageCallback1, a, n);
}
