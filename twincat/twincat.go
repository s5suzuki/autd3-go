/*
 * File: twincat.go
 * Project: twincat
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package twincat

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi-link-twincat
//
// #include "twincat_link.h"
import "C"
import (
	"unsafe"
)

type TwinCAT struct {
	ptr unsafe.Pointer
}

func NewTwinCAT() *TwinCAT {
	l := new(TwinCAT)
	l.ptr = unsafe.Pointer(nil)
	C.AUTDLinkTwinCAT(&l.ptr)
	return l
}

func (l *TwinCAT) Ptr() unsafe.Pointer {
	return l.ptr
}
