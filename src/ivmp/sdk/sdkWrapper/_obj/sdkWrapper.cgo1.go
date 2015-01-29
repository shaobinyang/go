// Created by cgo - DO NOT EDIT

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:1
package sdkWrapper
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:19

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:18
import "unsafe"
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:21

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:20
type User struct {
	Name		string
	Password	string
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:26

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:25
type UserAuthority struct {
	Ctrl		int
	Preset		int
	Video		int
	Download	int
	Monitor		int
	Config		int
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:35

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:34
type UserInfo struct {
	Name		string
	UserType	int
	Authority	UserAuthority
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:41

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:40
type MessageCallback func(int, string, unsafe.Pointer)
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:43

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:42
func Init() {
	_Cfunc_hmtNetInit()
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:47

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:46
func Release() {
	_Cfunc_hmtNetRelease()
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:51

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:50
func Login(url string, user User) int {
									curl := _Cfunc_CString(url)
									defer _Cfunc_free(unsafe.Pointer(curl))
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:55

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:54
	var cuser _Ctype_struct__HmtNetUser
									name := _Cfunc_CString(user.Name)
									defer _Cfunc_free(unsafe.Pointer(name))
									_Cfunc_strcpy(&cuser.szUserName[0], name)
									password := _Cfunc_CString(user.Password)
									defer _Cfunc_free(unsafe.Pointer(password))
									_Cfunc_strcpy(&cuser.szPassWord[0], password)
									return int(_Cfunc_hmtNetLogin(curl, &cuser))
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:65

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:64
func Logout(handle int) {
	_Cfunc_hmtNetLogout(_Ctype_HMT_NET_HANDLE(handle))
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:69

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:68
func QueryUserInfo(handle int) (UserInfo, int) {
	var info UserInfo
	var cUserInfo _Ctype_struct__HmtNetUserInfo
	rt := _Cfunc_hmtNetQueryUserInfo(_Ctype_HMT_NET_HANDLE(handle), &cUserInfo)
	if rt != 0 {
		return info, int(rt)
	}
	info.Name = _Cfunc_GoString(&cUserInfo.szUserName[0])
	info.UserType = int(cUserInfo.userType)
	info.Authority.Ctrl = int(cUserInfo.userAuthority.ctrl)
	info.Authority.Preset = int(cUserInfo.userAuthority.preset)
	info.Authority.Video = int(cUserInfo.userAuthority.video)
	info.Authority.Download = int(cUserInfo.userAuthority.download)
	info.Authority.Monitor = int(cUserInfo.userAuthority.monitor)
	info.Authority.Config = int(cUserInfo.userAuthority.config)
	return info, 0
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:87

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:86
func QueryPlatformID(handle int) (string, int) {
	var data string
	var bufSize _Ctype_int = 1024
	cData := _Cfunc__CMalloc(_Ctype_size_t(bufSize))
	defer _Cfunc_free(unsafe.Pointer(cData))
	rt := _Cfunc_hmtNetQueryPlatformID(_Ctype_HMT_NET_HANDLE(handle),
		(*_Ctype_char)(cData),
		&(bufSize))
	if rt != 0 {
		return data, int(rt)
	}
	data = _Cfunc_GoString((*_Ctype_char)(cData))
	return data, 0
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:102

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:101
func QueryPlatformName(handle int) (string, int) {
	var data string
	var bufSize _Ctype_int = 1024
	cData := _Cfunc__CMalloc(_Ctype_size_t(bufSize))
	defer _Cfunc_free(unsafe.Pointer(cData))
	rt := _Cfunc_hmtNetQueryPlatformName(_Ctype_HMT_NET_HANDLE(handle),
		(*_Ctype_char)(cData),
		&(bufSize))
	if rt != 0 {
		return data, int(rt)
	}
	data = _Cfunc_GoString((*_Ctype_char)(cData))
	return data, 0
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:117

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:116
func Subscribe(handle int, deviceID string, subType string, expires int) int {
	cDeviceID := _Cfunc_CString(deviceID)
	defer _Cfunc_free(unsafe.Pointer(cDeviceID))
	cSubType := _Cfunc_CString(subType)
	defer _Cfunc_free(unsafe.Pointer(cSubType))
	rt := _Cfunc_hmtNetSubscribe(_Ctype_HMT_NET_HANDLE(handle),
		cDeviceID,
		cSubType,
		(_Ctype_int)(expires))
	return int(rt)
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:129

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:128
var gMessageCallback MessageCallback
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:132

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:131
func messageCallback1(handle _Ctype_int, msg unsafe.Pointer, user unsafe.Pointer) {
	gMessageCallback(int(handle), _Cfunc_GoString((*_Ctype_char)(msg)), user)
}
//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:136

//line /home/ysb/work/go/src/ivmp/sdk/sdkWrapper/sdkWrapper.go:135
func SetMessageCallback(handle int,
	callback MessageCallback,
	user unsafe.Pointer) {
	gMessageCallback = callback
	_Cfunc_hmtNetSetMessageCallback(_Ctype_HMT_NET_HANDLE(handle),
		(_Ctype_HmtNetMessageProc)(unsafe.Pointer(_Cgo_ptr(_Cfpvar_fp_messageCallback_cgo))),
		user)
}
