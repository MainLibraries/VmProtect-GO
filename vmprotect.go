package vmprotect

/*
#include <stdbool.h>
#include <stdlib.h>
#include "VMProtectSDK.h"
#cgo amd64 LDFLAGS: -L${SRCDIR}/bin -lVMProtectSDK64
#cgo 386 LDFLAGS: -L${SRCDIR}/bin -lVMProtectSDK32
*/
import "C"
import (
	"unsafe"
)

type SerialStateFlags int

const (
	SerialStateSuccess             SerialStateFlags = 0
	SerialStateFlagCorrupted       SerialStateFlags = 0x00000001
	SerialStateFlagInvalid         SerialStateFlags = 0x00000002
	SerialStateFlagBlacklisted     SerialStateFlags = 0x00000004
	SerialStateFlagDateExpired     SerialStateFlags = 0x00000008
	SerialStateFlagRunningTimeOver SerialStateFlags = 0x00000010
	SerialStateFlagBadHWID         SerialStateFlags = 0x00000020
	SerialStateFlagMaxBuildExpired SerialStateFlags = 0x00000040
)

type ActivationFlags int

const (
	ActivationOK            ActivationFlags = 0
	ActivationSmallBuffer   ActivationFlags = 1
	ActivationNoConnection  ActivationFlags = 2
	ActivationBadReply      ActivationFlags = 3
	ActivationBanned        ActivationFlags = 4
	ActivationCorrupted     ActivationFlags = 5
	ActivationBadCode       ActivationFlags = 6
	ActivationAlreadyUsed   ActivationFlags = 7
	ActivationSerialUnknown ActivationFlags = 8
	ActivationExpired       ActivationFlags = 9
	ActivationNotAvailable  ActivationFlags = 10
)

type Date struct {
	Year  uint16
	Month uint8
	Day   uint8
}

type SerialNumberData struct {
	State        SerialStateFlags
	UserName     string
	Email        string
	ExpireDate   Date
	MaxBuildDate Date
	RunningTime  int
	UserData     []byte
}

func Begin(markerName string) {
	cstr := C.CString(markerName)
	defer C.free(unsafe.Pointer(cstr))
	C.VMProtectBegin(cstr)
}

func BeginVirtualization(markerName string) {
	cstr := C.CString(markerName)
	defer C.free(unsafe.Pointer(cstr))
	C.VMProtectBeginVirtualization(cstr)
}

func BeginMutation(markerName string) {
	cstr := C.CString(markerName)
	defer C.free(unsafe.Pointer(cstr))
	C.VMProtectBeginMutation(cstr)
}

func BeginUltra(markerName string) {
	cstr := C.CString(markerName)
	defer C.free(unsafe.Pointer(cstr))
	C.VMProtectBeginUltra(cstr)
}

func BeginVirtualizationLockByKey(markerName string) {
	cstr := C.CString(markerName)
	defer C.free(unsafe.Pointer(cstr))
	C.VMProtectBeginVirtualizationLockByKey(cstr)
}

func BeginUltraLockByKey(markerName string) {
	cstr := C.CString(markerName)
	defer C.free(unsafe.Pointer(cstr))
	C.VMProtectBeginUltraLockByKey(cstr)
}

func End() {
	C.VMProtectEnd()
}

func IsProtected() bool {
	return bool(C.VMProtectIsProtected())
}

func IsDebuggerPresent(kernel bool) bool {
	return bool(C.VMProtectIsDebuggerPresent(C.bool(kernel)))
}

func IsVirtualMachinePresent() bool {
	return bool(C.VMProtectIsVirtualMachinePresent())
}

func IsValidImageCRC() bool {
	return bool(C.VMProtectIsValidImageCRC())
}

func DecryptString(value string) string {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	result := C.GoString(C.VMProtectDecryptStringA(cstr))
	return result
}

func FreeString(value string) bool {
	cstr := C.CString(value)
	defer C.free(unsafe.Pointer(cstr))
	return bool(C.VMProtectFreeString(unsafe.Pointer(cstr)))
}

func SetSerialNumber(serial string) int {
	cstr := C.CString(serial)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.VMProtectSetSerialNumber(cstr))
}

func GetSerialNumberState() SerialStateFlags {
	return SerialStateFlags(C.VMProtectGetSerialNumberState())
}

func GetSerialNumberData() (*SerialNumberData, bool) {
	var data C.VMProtectSerialNumberData
	result := bool(C.VMProtectGetSerialNumberData(&data, C.int(unsafe.Sizeof(data))))

	if !result {
		return nil, false
	}

	goData := &SerialNumberData{
		State:        SerialStateFlags(data.nState),
		UserName:     C.GoString((*C.char)(unsafe.Pointer(&data.wUserName[0]))),
		Email:        C.GoString((*C.char)(unsafe.Pointer(&data.wEMail[0]))),
		ExpireDate:   Date{uint16(data.dtExpire.wYear), uint8(data.dtExpire.bMonth), uint8(data.dtExpire.bDay)},
		MaxBuildDate: Date{uint16(data.dtMaxBuild.wYear), uint8(data.dtMaxBuild.bMonth), uint8(data.dtMaxBuild.bDay)},
		RunningTime:  int(data.bRunningTime),
	}

	if data.nUserDataLength > 0 {
		goData.UserData = C.GoBytes(unsafe.Pointer(&data.bUserData[0]), C.int(data.nUserDataLength))
	}

	return goData, true
}

func GetCurrentHWID() string {
	buf := make([]C.char, 256)
	C.VMProtectGetCurrentHWID(&buf[0], 256)
	return C.GoString(&buf[0])
}

func ActivateLicense(code string) (string, ActivationFlags) {
	cstr := C.CString(code)
	defer C.free(unsafe.Pointer(cstr))

	serialBuf := make([]C.char, 256)
	result := ActivationFlags(C.VMProtectActivateLicense(cstr, &serialBuf[0], 256))

	if result != ActivationOK {
		return "", result
	}

	return C.GoString(&serialBuf[0]), result
}

func DeactivateLicense(serial string) ActivationFlags {
	cstr := C.CString(serial)
	defer C.free(unsafe.Pointer(cstr))
	return ActivationFlags(C.VMProtectDeactivateLicense(cstr))
}

func GetOfflineActivationString(code string) (string, ActivationFlags) {
	cstr := C.CString(code)
	defer C.free(unsafe.Pointer(cstr))

	buf := make([]C.char, 1024)
	result := ActivationFlags(C.VMProtectGetOfflineActivationString(cstr, &buf[0], 1024))

	if result != ActivationOK {
		return "", result
	}

	return C.GoString(&buf[0]), result
}

func GetOfflineDeactivationString(serial string) (string, ActivationFlags) {
	cstr := C.CString(serial)
	defer C.free(unsafe.Pointer(cstr))

	buf := make([]C.char, 1024)
	result := ActivationFlags(C.VMProtectGetOfflineDeactivationString(cstr, &buf[0], 1024))

	if result != ActivationOK {
		return "", result
	}

	return C.GoString(&buf[0]), result
}
