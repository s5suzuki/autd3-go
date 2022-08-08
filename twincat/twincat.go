/*
 * File: twincat.go
 * Project: twincat
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 08/08/2022
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

	"github.com/shinolab/autd3-go/v2/autd3"
)

type TwinCAT struct{}

func NewTwinCAT() *TwinCAT {
	l := new(TwinCAT)
	return l
}

func (l *TwinCAT) Build() *autd3.Link {
	link := new(autd3.Link)
	link.Ptr = unsafe.Pointer(nil)
	C.AUTDLinkTwinCAT(&link.Ptr)
	return link
}
