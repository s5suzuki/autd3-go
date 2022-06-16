/*
 * File: modulation.go
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

// #cgo CFLAGS: -I${SRCDIR}/../../include/
// #cgo LDFLAGS: -L${SRCDIR}/../../ -lautd3capi
//
// #include "autd3_c_api.h"
import "C"

import (
	"unsafe"
)

type Modulation struct {
	P unsafe.Pointer
}

func (m *Modulation) Ptr() unsafe.Pointer {
	return m.P
}

func (m *Modulation) IsHeader() bool {
	return true
}

func (m *Modulation) Delete() {
	if m.P == unsafe.Pointer(nil) {
		return
	}
	C.AUTDDeleteModulation(m.P)
	m.P = unsafe.Pointer(nil)
}

func (m *Modulation) GetSamplingFrequencyDivision() uint32 {
	return uint32(C.AUTDModulationSamplingFrequencyDivision(m.P))
}

func (m *Modulation) SetSamplingFrequencyDivision(value uint32) {
	C.AUTDModulationSetSamplingFrequencyDivision(m.P, C.uint(value))
}

func (m *Modulation) GetSamplingFrequency() float64 {
	return float64(C.AUTDModulationSamplingFrequency(m.P))
}

type Sine struct {
	Modulation
}

func NewSine(freq int, param ...float64) *Sine {
	m := new(Sine)
	m.P = unsafe.Pointer(nil)
	var amp float64 = 1.0
	var offset float64 = 0.5
	if len(param) > 0 {
		amp = param[0]
	}

	if len(param) > 1 {
		offset = param[1]
	}

	C.AUTDModulationSine(&m.P, C.int(freq), C.double(amp), C.double(offset))
	return m
}

type Static struct {
	Modulation
}

func NewStatic(param ...float64) *Static {
	m := new(Static)
	m.P = unsafe.Pointer(nil)
	var amp float64 = 1.0
	if len(param) > 0 {
		amp = param[0]
	}

	C.AUTDModulationStatic(&m.P, C.double(amp))
	return m
}

type SineSquared struct {
	Modulation
}

func NewSineSquared(freq int, param ...float64) *SineSquared {
	m := new(SineSquared)
	m.P = unsafe.Pointer(nil)
	var amp float64 = 1.0
	var offset float64 = 0.5
	if len(param) > 0 {
		amp = param[0]
	}

	if len(param) > 1 {
		offset = param[1]
	}

	C.AUTDModulationSineSquared(&m.P, C.int(freq), C.double(amp), C.double(offset))
	return m
}

type SineLegacy struct {
	Modulation
}

func NewSineLegacy(freq float64, param ...float64) *SineLegacy {
	m := new(SineLegacy)
	m.P = unsafe.Pointer(nil)
	var amp float64 = 1.0
	var offset float64 = 0.5
	if len(param) > 0 {
		amp = param[0]
	}

	if len(param) > 1 {
		offset = param[1]
	}

	C.AUTDModulationSineLegacy(&m.P, C.double(freq), C.double(amp), C.double(offset))
	return m
}

type Square struct {
	Modulation
}

func NewSquare(freq float64, param ...float64) *Square {
	m := new(Square)
	m.P = unsafe.Pointer(nil)
	var low float64 = 0.0
	var high float64 = 1.0
	var duty float64 = 0.5
	if len(param) > 0 {
		low = param[0]
	}
	if len(param) > 1 {
		high = param[1]
	}
	if len(param) > 2 {
		duty = param[2]
	}

	C.AUTDModulationSquare(&m.P, C.int(freq), C.double(low), C.double(high), C.double(duty))
	return m
}

type CustomModulation struct {
	Modulation
}

func NewCustomModulation(buf []uint8, freqDiv uint32) *CustomModulation {
	m := new(CustomModulation)
	m.P = unsafe.Pointer(nil)
	n := len(buf)
	C.AUTDModulationCustom(&m.P, (*C.uchar)(unsafe.Pointer(&buf[0])), C.ulonglong(n), C.uint(freqDiv))
	return m
}
