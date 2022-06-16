/*
 * File: emulator.go
 * Project: emulator
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package emulator

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi-link-emulator
//
// #include "emulator_link.h"
import "C"
import (
	"unsafe"

	"github.com/shinolab/autd3-go/v2/autd3"
)

type Emulator struct {
	ptr unsafe.Pointer
}

func NewEmulator(port uint16, cnt *autd3.Controller) *Emulator {
	l := new(Emulator)
	l.ptr = unsafe.Pointer(nil)
	C.AUTDLinkEmulator(&l.ptr, C.ushort(port), cnt.Ptr)
	return l
}

func (l *Emulator) Ptr() unsafe.Pointer {
	return l.ptr
}
