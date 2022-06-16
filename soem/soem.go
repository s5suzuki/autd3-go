/*
 * File: soem.go
 * Project: soem
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
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
)

type Adapter struct {
	Desc string
	Name string
}

type SOEM struct {
	ptr unsafe.Pointer
}

var _callback *func(string)

//export onLostCallback
func onLostCallback(msg *C.char) {
	if _callback != nil {
		(*_callback)(C.GoString(msg))
	}
}

func RegisterOnLostCallback(callback func(string)) {
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

func NewSOEM(ifname string, devNum int, cycleTicks int, highPrecision bool) *SOEM {
	l := new(SOEM)
	l.ptr = unsafe.Pointer(nil)
	callback := unsafe.Pointer(nil)
	C.AUTDLinkSOEMGetCallback(&callback)
	C.AUTDLinkSOEM(&l.ptr, C.CString(ifname), C.int(devNum), C.ushort(cycleTicks), callback, C.bool(highPrecision))
	return l
}

func (l *SOEM) Ptr() unsafe.Pointer {
	return l.ptr
}
