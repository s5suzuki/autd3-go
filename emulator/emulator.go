/*
 * File: emulator.go
 * Project: emulator
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 14/08/2022
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
	port uint16
}

func NewEmulator() *Emulator {
	l := new(Emulator)
	l.port = 50632
	return l
}

func (l *Emulator) Port(port uint16) *Emulator {
	l.port = port
	return l
}

func (link *Emulator) Build() *autd3.Link {
	l := new(autd3.Link)
	l.Ptr = unsafe.Pointer(nil)
	C.AUTDLinkEmulator(&l.Ptr, C.ushort(link.port))
	return l
}
