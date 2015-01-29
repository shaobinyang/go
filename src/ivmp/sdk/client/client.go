package main

import (
	"C"
	"bufio"
	sdk "ivmp/sdk/sdkWrapper"
	"log"
	"os"
	"strings"
	"unsafe"
)

func messageCallback(handle int, msg string, user unsafe.Pointer) {
	log.Println("handle:", handle)
	log.Println("callback:", msg)
}

func queryCatalog(handle int, deviceID string) {
	catalog, rt := sdk.RequestConfigure(handle, deviceID, 0)
	if rt == 0 {
		log.Println(catalog)
	} else {
		log.Println("Catalog error:", rt)
	}
}

func queryDeviceStatus(handle int, deviceID string) {
	status, rt := sdk.QueryDeviceStatus(handle, deviceID)
	if rt == 0 {
		log.Println(status)
	} else {
		log.Println("Device status error:", rt)
	}
}

func main() {
	sdk.Init()
	hd := sdk.Login("http://192.168.1.86:59000", sdk.User{"admin", "admin"})
	if hd == 0 {
		log.Println("login error")
		return
	}
	log.Println("Login success, handle is ", hd)
	userInfo, _ := sdk.QueryUserInfo(hd)
	log.Println("User name: ", userInfo.Name)
	platformID, _ := sdk.QueryPlatformID(hd)
	log.Println("Platform id: ", platformID)
	platformName, _ := sdk.QueryPlatformName(hd)
	log.Println("Platform name: ", platformName)

	sdk.Subscribe(hd, "37020086002000000000", "Catalog", 600)
	sdk.Subscribe(hd, "37020086002000000000", "Alarm", 600)
	log.Println("Subscribe over")

	userData := 0

	sdk.SetMessageCallback(hd, messageCallback, unsafe.Pointer(&userData))
	log.Println("Set message callback")

	inputReader := bufio.NewReader(os.Stdin)
LOOP:
	for {
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		switch trimInput {
		case "q":
			break LOOP
		case "catalogPlat":
			queryCatalog(hd, "37020086002000000000")
		case "catalogDomain":
			queryCatalog(hd, "37020086002160000001")
		case "catalogBusiness":
			queryCatalog(hd, "37020086002150000001")
		case "catalogDevice":
			queryCatalog(hd, "37020086001180001001")
		case "statusPlat":
			queryDeviceStatus(hd, "37020086002000000000")
		case "statusDomain":
			queryDeviceStatus(hd, "37020086002160000001")
		case "statusBusiness":
			queryDeviceStatus(hd, "37020086002150000001")
		case "statusDevice":
			queryDeviceStatus(hd, "37020086001180001001")
		}
	}
	log.Println("Logout")
	sdk.Logout(hd)
	log.Println("Release")
	sdk.Release()
	log.Println("Quit")
}
