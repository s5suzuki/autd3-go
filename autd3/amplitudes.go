/*
 * File: amplitudes.go
 * Project: autd3
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package autd3

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi
//
// #include "autd3_c_api.h"
import "C"
import "unsafe"

type Amplitudes struct {
	ptr unsafe.Pointer
}

func NewAmplitudes(cnt *Controller, value float64) *Amplitudes {
	config := new(Amplitudes)
	config.ptr = unsafe.Pointer(nil)

	C.AUTDCreateAmplitudes(&config.ptr, cnt.Ptr, C.double(value))

	return config
}

func (c *Amplitudes) Delete() {
	if c.ptr == unsafe.Pointer(nil) {
		return
	}
	C.AUTDDeleteAmplitudes(c.ptr)
	c.ptr = unsafe.Pointer(nil)
}

func (g *Amplitudes) Ptr() unsafe.Pointer {
	return g.ptr
}

func (g *Amplitudes) IsHeader() bool {
	return true
}
