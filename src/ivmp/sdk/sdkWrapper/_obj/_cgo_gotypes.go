// Created by cgo - DO NOT EDIT

package sdkWrapper

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

type _Ctype_HMT_NET_HANDLE _Ctype_ulong

type _Ctype_HmtNetMessageProc *[0]byte

type _Ctype_HmtNetUser _Ctype_struct__HmtNetUser

type _Ctype_HmtNetUserInfo _Ctype_struct__HmtNetUserInfo

type _Ctype_UserAuthority _Ctype_struct__UserAuthority

type _Ctype_char int8

type _Ctype_int int32

type _Ctype_size_t _Ctype_ulong

type _Ctype_struct__HmtNetUser struct {
//line :1
	szUserName	[256]_Ctype_char
//line :1
	szPassWord	[256]_Ctype_char
//line :1
}

type _Ctype_struct__HmtNetUserInfo struct {
//line :1
	szUserName	[256]_Ctype_char
//line :1
	userType	_Ctype_int
//line :1
	userAuthority	_Ctype_struct__UserAuthority
//line :1
}

type _Ctype_struct__UserAuthority struct {
//line :1
	ctrl		_Ctype_int
//line :1
	preset		_Ctype_int
//line :1
	video		_Ctype_int
//line :1
	download	_Ctype_int
//line :1
	monitor		_Ctype_int
//line :1
	config		_Ctype_int
//line :1
}

type _Ctype_ulong uint64

type _Ctype_void [0]byte

var _cgo_runtime_cgocall_errno func(unsafe.Pointer, uintptr) int32
var _cgo_runtime_cmalloc func(uintptr) unsafe.Pointer
var _Cfpvar_fp_messageCallback_cgo unsafe.Pointer


func _Cfunc_CString(s string) *_Ctype_char {
	p := _cgo_runtime_cmalloc(uintptr(len(s)+1))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], s)
	pp[len(s)] = 0
	return (*_Ctype_char)(p)
}

var _cgo_runtime_gostring func(*_Ctype_char) string
func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}

func _Cfunc__CMalloc(n _Ctype_size_t) unsafe.Pointer {
	return _cgo_runtime_cmalloc(uintptr(n))
}

var _cgo_0d4ddd74268d_Cfunc_free unsafe.Pointer
func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetInit unsafe.Pointer
func _Cfunc_hmtNetInit() (r1 _Ctype_int) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetInit, uintptr(unsafe.Pointer(&r1)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetLogin unsafe.Pointer
func _Cfunc_hmtNetLogin(p0 *_Ctype_char, p1 *_Ctype_struct__HmtNetUser) (r1 _Ctype_HMT_NET_HANDLE) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetLogin, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetLogout unsafe.Pointer
func _Cfunc_hmtNetLogout(p0 _Ctype_HMT_NET_HANDLE) (r1 _Ctype_void) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetLogout, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformID unsafe.Pointer
func _Cfunc_hmtNetQueryPlatformID(p0 _Ctype_HMT_NET_HANDLE, p1 *_Ctype_char, p2 *_Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformID, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformName unsafe.Pointer
func _Cfunc_hmtNetQueryPlatformName(p0 _Ctype_HMT_NET_HANDLE, p1 *_Ctype_char, p2 *_Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetQueryPlatformName, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetQueryUserInfo unsafe.Pointer
func _Cfunc_hmtNetQueryUserInfo(p0 _Ctype_HMT_NET_HANDLE, p1 *_Ctype_struct__HmtNetUserInfo) (r1 _Ctype_int) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetQueryUserInfo, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetRelease unsafe.Pointer
func _Cfunc_hmtNetRelease() (r1 _Ctype_int) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetRelease, uintptr(unsafe.Pointer(&r1)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetSetMessageCallback unsafe.Pointer
func _Cfunc_hmtNetSetMessageCallback(p0 _Ctype_HMT_NET_HANDLE, p1 *[0]byte, p2 unsafe.Pointer) (r1 _Ctype_int) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetSetMessageCallback, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_hmtNetSubscribe unsafe.Pointer
func _Cfunc_hmtNetSubscribe(p0 _Ctype_HMT_NET_HANDLE, p1 *_Ctype_char, p2 *_Ctype_char, p3 _Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_hmtNetSubscribe, uintptr(unsafe.Pointer(&p0)))
	return
}

var _cgo_0d4ddd74268d_Cfunc_strcpy unsafe.Pointer
func _Cfunc_strcpy(p0 *_Ctype_char, p1 *_Ctype_char) (r1 *_Ctype_char) {
	_cgo_runtime_cgocall_errno(_cgo_0d4ddd74268d_Cfunc_strcpy, uintptr(unsafe.Pointer(&p0)))
	return
}
