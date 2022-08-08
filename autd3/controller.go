/*
 * File: controller.go
 * Project: controller
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 08/08/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package autd3

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi -lautd3capi-gain-holo
//
// #include "autd3_c_api.h"
// #include "holo_gain.h"
import "C"

import (
	"strings"
	"unsafe"
)

type Controller struct {
	Ptr unsafe.Pointer
}

func NewController() *Controller {
	cnt := new(Controller)
	cnt.Ptr = unsafe.Pointer(nil)
	C.AUTDCreateController(&cnt.Ptr)
	return cnt
}

func (cnt *Controller) Delete() {
	if cnt.Ptr == unsafe.Pointer(nil) {
		return
	}
	C.AUTDFreeController(cnt.Ptr)
	cnt.Ptr = unsafe.Pointer(nil)
}

func (cnt *Controller) ToLegacy() {
	C.AUTDSetMode(cnt.Ptr, 0)
}

func (cnt *Controller) ToNormal() {
	C.AUTDSetMode(cnt.Ptr, 1)
}

func (cnt *Controller) ToNormalPhase() {
	C.AUTDSetMode(cnt.Ptr, 2)
}

func (cnt *Controller) AddDevice(pos [3]float64, rot [3]float64) int {
	return int(C.AUTDAddDevice(cnt.Ptr, C.double(pos[0]), C.double(pos[1]), C.double(pos[2]), C.double(rot[0]), C.double(rot[1]), C.double(rot[2])))
}

func (cnt *Controller) AddDeviceQuaternion(pos [3]float64, rot [4]float64) int {
	return int(C.AUTDAddDeviceQuaternion(cnt.Ptr, C.double(pos[0]), C.double(pos[1]), C.double(pos[2]), C.double(rot[0]), C.double(rot[1]), C.double(rot[2]), C.double(rot[3])))
}

func (cnt *Controller) Open(link *Link) bool {
	return bool(C.AUTDOpenController(cnt.Ptr, link.Ptr))
}

func (cnt *Controller) Clear() int {
	return int(C.AUTDClear(cnt.Ptr))
}

func (cnt *Controller) Stop() int {
	return int(C.AUTDStop(cnt.Ptr))
}

func (cnt *Controller) Close() int {
	return int(C.AUTDClose(cnt.Ptr))
}

func (cnt *Controller) Synchronize() int {
	return int(C.AUTDSynchronize(cnt.Ptr))
}

func (cnt *Controller) UpdateFlags() int {
	return int(C.AUTDUpdateFlags(cnt.Ptr))
}

func (cnt *Controller) SetModDelay(devId int, transIdx int, value uint16) {
	C.AUTDSetModDelay(cnt.Ptr, C.int(devId), C.int(transIdx), C.ushort(value))
}

func (cnt *Controller) Send(param ...Sendable) int {
	if len(param) == 0 {
		return -1
	}

	if len(param) == 1 {
		if param[0].IsHeader() {
			return int(C.AUTDSend(cnt.Ptr, param[0].Ptr(), nil))
		} else {
			return int(C.AUTDSend(cnt.Ptr, nil, param[0].Ptr()))
		}
	}

	if len(param) == 2 {
		if param[0].IsHeader() && !param[1].IsHeader() {
			return int(C.AUTDSend(cnt.Ptr, param[0].Ptr(), param[1].Ptr()))
		} else if !param[0].IsHeader() && param[1].IsHeader() {
			return int(C.AUTDSend(cnt.Ptr, param[1].Ptr(), param[0].Ptr()))
		} else {
			return -1
		}
	}

	return -1
}

func (cnt *Controller) IsOpen() bool {
	return bool(C.AUTDIsOpen(cnt.Ptr))
}

func (cnt *Controller) GetForceFan() bool {
	return bool(C.AUTDGetForceFan(cnt.Ptr))
}

func (cnt *Controller) SetForceFan(value bool) {
	C.AUTDSetForceFan(cnt.Ptr, C.bool(value))
}

func (cnt *Controller) GetCheckTrials() int {
	return int(C.AUTDGetCheckTrials(cnt.Ptr))
}

func (cnt *Controller) SetCheckTrials(value int) {
	C.AUTDSetCheckTrials(cnt.Ptr, C.int(value))
}

func (cnt *Controller) GetSendInterval() int {
	return int(C.AUTDGetSendInterval(cnt.Ptr))
}

func (cnt *Controller) SetSendInterval(value int) {
	C.AUTDSetSendInterval(cnt.Ptr, C.int(value))
}

func (cnt *Controller) GetReadsFPGAInfo() bool {
	return bool(C.AUTDGetReadsFPGAInfo(cnt.Ptr))
}

func (cnt *Controller) SetReadsFPGAInfo(value bool) {
	C.AUTDSetReadsFPGAInfo(cnt.Ptr, C.bool(value))
}

func (cnt *Controller) GetSoundSpeed() float64 {
	return float64(C.AUTDGetSoundSpeed(cnt.Ptr))
}

func (cnt *Controller) SetSoundSpeed(value float64) {
	C.AUTDSetSoundSpeed(cnt.Ptr, C.double(value))
}

func (cnt *Controller) GetAttenuation() float64 {
	return float64(C.AUTDGetAttenuation(cnt.Ptr))
}

func (cnt *Controller) SetAttenuation(value float64) {
	C.AUTDSetAttenuation(cnt.Ptr, C.double(value))
}

func (cnt *Controller) GetTransFrequency(devId int, transIdx int) float64 {
	return float64(C.AUTDGetTransFrequency(cnt.Ptr, C.int(devId), C.int(transIdx)))
}

func (cnt *Controller) SetTransFrequency(devId int, transIdx int, value float64) {
	C.AUTDSetTransFrequency(cnt.Ptr, C.int(devId), C.int(transIdx), C.double(value))
}

func (cnt *Controller) GetTransCycle(devId int, transIdx int) uint16 {
	return uint16(C.AUTDGetTransCycle(cnt.Ptr, C.int(devId), C.int(transIdx)))
}

func (cnt *Controller) SetTransCycle(devId int, transIdx int, value uint16) {
	C.AUTDSetTransCycle(cnt.Ptr, C.int(devId), C.int(transIdx), C.ushort(value))
}

func (cnt *Controller) GetWavelength(devId int, transIdx int) float64 {
	return float64(C.AUTDGetWavelength(cnt.Ptr, C.int(devId), C.int(transIdx)))
}

func (cnt *Controller) NumDevices() int {
	return int(C.AUTDNumDevices(cnt.Ptr))
}

func (cnt *Controller) GetFPGAInfo() []uint8 {
	info := make([]uint8, cnt.NumDevices())
	C.AUTDGetFPGAInfo(cnt.Ptr, (*C.uchar)(unsafe.Pointer(&info[0])))
	return info
}

func (cnt *Controller) GetTransPosition(devId int, transIdx int) [3]float64 {
	x := C.double(0)
	y := C.double(0)
	z := C.double(0)
	C.AUTDTransPosition(cnt.Ptr, C.int(devId), C.int(transIdx), &x, &y, &z)
	return [3]float64{float64(x), float64(y), float64(z)}
}

func (cnt *Controller) GetTransDirectionX(devId int, transIdx int) [3]float64 {
	x := C.double(0)
	y := C.double(0)
	z := C.double(0)
	C.AUTDTransXDirection(cnt.Ptr, C.int(devId), C.int(transIdx), &x, &y, &z)
	return [3]float64{float64(x), float64(y), float64(z)}
}

func (cnt *Controller) GetTransDirectionY(devId int, transIdx int) [3]float64 {
	x := C.double(0)
	y := C.double(0)
	z := C.double(0)
	C.AUTDTransYDirection(cnt.Ptr, C.int(devId), C.int(transIdx), &x, &y, &z)
	return [3]float64{float64(x), float64(y), float64(z)}
}

func (cnt *Controller) GetTransDirectionZ(devId int, transIdx int) [3]float64 {
	x := C.double(0)
	y := C.double(0)
	z := C.double(0)
	C.AUTDTransZDirection(cnt.Ptr, C.int(devId), C.int(transIdx), &x, &y, &z)
	return [3]float64{float64(x), float64(y), float64(z)}
}

func (cnt *Controller) FirmwareInfoList() []string {
	p := unsafe.Pointer(nil)
	n := int(C.AUTDGetFirmwareInfoListPointer(cnt.Ptr, &p))
	list := make([]string, n)
	for i := 0; i < n; i++ {
		info := strings.Repeat("\x00", 256)
		cinfo := C.CString(info)
		C.AUTDGetFirmwareInfo(p, C.int(i), cinfo)
		list[i] = C.GoString(cinfo)
	}
	C.AUTDFreeFirmwareInfoListPointer(p)
	return list
}

func GetLastError() string {
	var n = C.AUTDGetLastError(nil)
	err := strings.Repeat("\x00", int(n))
	cerr := C.CString(err)
	C.AUTDGetLastError(cerr)
	return C.GoString(cerr)
}
