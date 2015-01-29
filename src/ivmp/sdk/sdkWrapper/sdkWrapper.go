package sdkWrapper

/*
#cgo CFLAGS: -I$HOME/include/
#cgo LDFLAGS: -L$HOME/lib/ -lhmtnetsdk
#include <stdlib.h>
#include <string.h>
#include <hmt_net_sdk.h>

void messageCallback_cgo(int handle, void* msg, void* user);

*/
import "C"
import "unsafe"

type User struct {
	Name     string
	Password string
}

type UserAuthority struct {
	Ctrl     int
	Preset   int
	Video    int
	Download int
	Monitor  int
	Config   int
}

type UserInfo struct {
	Name      string
	UserType  int
	Authority UserAuthority
}

type MessageCallback func(int, string, unsafe.Pointer)

func Init() {
	C.hmtNetInit()
}

func Release() {
	C.hmtNetRelease()
}

func Login(url string, user User) int {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))

	var cuser C.HmtNetUser
	name := C.CString(user.Name)
	defer C.free(unsafe.Pointer(name))
	C.strcpy(&cuser.szUserName[0], name)
	password := C.CString(user.Password)
	defer C.free(unsafe.Pointer(password))
	C.strcpy(&cuser.szPassWord[0], password)
	return int(C.hmtNetLogin(curl, &cuser))
}

func Logout(handle int) {
	C.hmtNetLogout(C.HMT_NET_HANDLE(handle))
}

func QueryUserInfo(handle int) (UserInfo, int) {
	var info UserInfo
	var cUserInfo C.HmtNetUserInfo
	rt := C.hmtNetQueryUserInfo(C.HMT_NET_HANDLE(handle), &cUserInfo)
	if rt != 0 {
		return info, int(rt)
	}
	info.Name = C.GoString(&cUserInfo.szUserName[0])
	info.UserType = int(cUserInfo.userType)
	info.Authority.Ctrl = int(cUserInfo.userAuthority.ctrl)
	info.Authority.Preset = int(cUserInfo.userAuthority.preset)
	info.Authority.Video = int(cUserInfo.userAuthority.video)
	info.Authority.Download = int(cUserInfo.userAuthority.download)
	info.Authority.Monitor = int(cUserInfo.userAuthority.monitor)
	info.Authority.Config = int(cUserInfo.userAuthority.config)
	return info, 0
}

func QueryPlatformID(handle int) (string, int) {
	var data string
	var bufSize C.int = 1024
	cData := C.malloc(C.size_t(bufSize))
	defer C.free(unsafe.Pointer(cData))
	rt := C.hmtNetQueryPlatformID(C.HMT_NET_HANDLE(handle),
		(*C.char)(cData),
		&(bufSize))
	if rt != 0 {
		return data, int(rt)
	}
	data = C.GoString((*C.char)(cData))
	return data, 0
}

func QueryPlatformName(handle int) (string, int) {
	var data string
	var bufSize C.int = 1024
	cData := C.malloc(C.size_t(bufSize))
	defer C.free(unsafe.Pointer(cData))
	rt := C.hmtNetQueryPlatformName(C.HMT_NET_HANDLE(handle),
		(*C.char)(cData),
		&(bufSize))
	if rt != 0 {
		return data, int(rt)
	}
	data = C.GoString((*C.char)(cData))
	return data, 0
}

func Subscribe(handle int, deviceID string, subType string, expires int) int {
	cDeviceID := C.CString(deviceID)
	defer C.free(unsafe.Pointer(cDeviceID))
	cSubType := C.CString(subType)
	defer C.free(unsafe.Pointer(cSubType))
	rt := C.hmtNetSubscribe(C.HMT_NET_HANDLE(handle),
		cDeviceID,
		cSubType,
		(C.int)(expires))
	return int(rt)
}

var gMessageCallback MessageCallback

//export messageCallback
func messageCallback(handle C.int, msg unsafe.Pointer, user unsafe.Pointer) {
	gMessageCallback(int(handle), C.GoString((*C.char)(msg)), user)
}

func SetMessageCallback(handle int,
	callback MessageCallback,
	user unsafe.Pointer) {
	gMessageCallback = callback
	C.hmtNetSetMessageCallback(C.HMT_NET_HANDLE(handle),
		(C.HmtNetMessageProc)(unsafe.Pointer(C.messageCallback_cgo)),
		user)
}

func RequestConfigure(handle int, deviceID string, configType int) (string, int) {
	cDeviceID := C.CString(deviceID)
	defer C.free(unsafe.Pointer(cDeviceID))
	var data string
	var bufSize C.int = 102400
	cData := C.malloc(C.size_t(bufSize))
	defer C.free(unsafe.Pointer(cData))

	rt := C.hmtNetRequestConfigure(C.HMT_NET_HANDLE(handle),
		cDeviceID,
		(C.int)(configType),
		(*C.char)(cData),
		&bufSize)
	if rt != 0 {
		return data, int(rt)
	}
	data = C.GoString((*C.char)(cData))
	return data, 0
}

func QueryDeviceStatus(handle int, deviceID string) (string, int) {
	cDeviceID := C.CString(deviceID)
	defer C.free(unsafe.Pointer(cDeviceID))
	var data string
	var bufSize C.int = 10240
	cData := C.malloc(C.size_t(bufSize))
	defer C.free(unsafe.Pointer(cData))
	rt := C.hmtNetQueryDeviceStatus(C.HMT_NET_HANDLE(handle),
		cDeviceID,
		(*C.char)(cData),
		&bufSize)
	if rt != 0 {
		return data, int(rt)
	}
	data = C.GoString((*C.char)(cData))
	return data, 0
}
