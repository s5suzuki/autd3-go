/*
 * File: mod_delay_config.go
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

type ModDelayConfig struct {
	ptr unsafe.Pointer
}

func NewModDelayConfig() *ModDelayConfig {
	config := new(ModDelayConfig)
	config.ptr = unsafe.Pointer(nil)

	C.AUTDCreateModDelayConfig(&config.ptr)

	return config
}

func (c *ModDelayConfig) Delete() {
	if c.ptr == unsafe.Pointer(nil) {
		return
	}
	C.AUTDDeleteModDelayConfig(c.ptr)
	c.ptr = unsafe.Pointer(nil)
}

func (g *ModDelayConfig) Ptr() unsafe.Pointer {
	return g.ptr
}

func (g *ModDelayConfig) IsHeader() bool {
	return true
}
