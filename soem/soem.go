/*
 * File: soem.go
 * Project: soem
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 14/08/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package soem

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi-link-soem
//
// #include "soem_link.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/shinolab/autd3-go/v2/autd3"
)

type Adapter struct {
	Desc string
	Name string
}

type SOEM struct {
	ifname        string
	devNum        int
	sendCycle     uint16
	sync0Cycle    uint16
	highPrecision bool
	freerun       bool
	callback      func(string)
}

var _callback *func(string)

//export onLostCallback
func onLostCallback(msg *C.char) {
	if _callback != nil {
		(*_callback)(C.GoString(msg))
	}
}

func registerOnLostCallback(callback func(string)) {
	_callback = &callback
}

func EnumerateAdapters() []Adapter {
	p := unsafe.Pointer(nil)
	var n = int(C.AUTDGetAdapterPointer(&p))
	adapters := make([]Adapter, n)
	for i := 0; i < n; i++ {
		desc := strings.Repeat("\x00", 128)
		name := strings.Repeat("\x00", 128)
		c_desc := C.CString(desc)
		c_name := C.CString(name)
		C.AUTDGetAdapter(p, C.int(i), c_desc, c_name)
		adapters[i] = Adapter{C.GoString(c_desc), C.GoString(c_name)}
	}
	C.AUTDFreeAdapterPointer(p)
	return adapters
}

func NewSOEM(devNum int) *SOEM {
	l := new(SOEM)
	l.ifname = ""
	l.devNum = devNum
	l.sendCycle = 1
	l.sync0Cycle = 1
	l.highPrecision = false
	l.freerun = false
	l.callback = nil
	return l
}

func (l *SOEM) Ifname(ifname string) *SOEM {
	l.ifname = ifname
	return l
}

func (l *SOEM) SendCycle(cycle uint16) *SOEM {
	l.sendCycle = cycle
	return l
}

func (l *SOEM) Sync0Cycle(cycle uint16) *SOEM {
	l.sync0Cycle = cycle
	return l
}

func (l *SOEM) HighPrecision(flag bool) *SOEM {
	l.highPrecision = flag
	return l
}

func (l *SOEM) FreeRun(flag bool) *SOEM {
	l.freerun = flag
	return l
}

func (l *SOEM) OnLost(onLost func(string)) *SOEM {
	l.callback = onLost
	return l
}

func (link *SOEM) Build() *autd3.Link {
	l := new(autd3.Link)
	l.Ptr = unsafe.Pointer(nil)

	if link.callback != nil {
		registerOnLostCallback(link.callback)
	}

	callback := unsafe.Pointer(nil)
	C.AUTDLinkSOEMGetCallback(&callback)
	if link.ifname == "" {
		C.AUTDLinkSOEM(&l.Ptr, nil, C.int(link.devNum), C.ushort(link.sync0Cycle), C.ushort(link.sendCycle), C.bool(link.freerun), callback, C.bool(link.highPrecision))
	} else {
		C.AUTDLinkSOEM(&l.Ptr, C.CString(link.ifname), C.int(link.devNum), C.ushort(link.sync0Cycle), C.ushort(link.sendCycle), C.bool(link.freerun), callback, C.bool(link.highPrecision))
	}
	return l
}
