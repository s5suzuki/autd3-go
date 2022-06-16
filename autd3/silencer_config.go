/*
 * File: silencer_config.go
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

type SilencerConfig struct {
	ptr unsafe.Pointer
}

func NewSilencerConfig(param ...int) *SilencerConfig {
	step := 10
	freqDiv := 4096
	if len(param) > 0 {
		step = param[0]
	}
	if len(param) > 1 {
		freqDiv = param[1]
	}

	config := new(SilencerConfig)
	config.ptr = unsafe.Pointer(nil)

	C.AUTDCreateSilencer(&config.ptr, C.ushort(step), C.ushort(freqDiv))

	return config
}

func NewSilencerConfigNone() *SilencerConfig {
	return NewSilencerConfig(0xFFFF, 4096)
}

func (c *SilencerConfig) Delete() {
	if c.ptr == unsafe.Pointer(nil) {
		return
	}
	C.AUTDDeleteSilencer(c.ptr)
	c.ptr = unsafe.Pointer(nil)
}

func (g *SilencerConfig) Ptr() unsafe.Pointer {
	return g.ptr
}

func (g *SilencerConfig) IsHeader() bool {
	return true
}
