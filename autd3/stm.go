/*
 * File: stm.go
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

type STM struct {
	ptr unsafe.Pointer
}

func (g *STM) Ptr() unsafe.Pointer {
	return g.ptr
}

func (g *STM) IsHeader() bool {
	return false
}

func (g *STM) Delete() {
	if g.ptr == unsafe.Pointer(nil) {
		return
	}
	C.AUTDDeleteSTM(g.ptr)
	g.ptr = unsafe.Pointer(nil)
}

func (m *STM) GetFrequency() float64 {
	return float64(C.AUTDSTMFrequency(m.ptr))
}

func (m *STM) SetFrequency(value float64) float64 {
	return float64(C.AUTDSTMSetFrequency(m.ptr, C.double(value)))
}

func (m *STM) GetSamplingFrequencyDivision() uint32 {
	return uint32(C.AUTDSTMSamplingFrequencyDivision(m.ptr))
}

func (m *STM) SetSamplingFrequencyDivision(value uint32) {
	C.AUTDSTMSetSamplingFrequencyDivision(m.ptr, C.uint(value))
}

func (m *STM) GetSamplingFrequency() float64 {
	return float64(C.AUTDSTMSamplingFrequency(m.ptr))
}

type PointSTM struct {
	STM
}

func NewPointSTM() *PointSTM {
	stm := new(PointSTM)
	stm.ptr = unsafe.Pointer(nil)
	C.AUTDPointSTM(&stm.ptr)
	return stm
}

func (stm *PointSTM) Add(pos [3]float64, shift ...uint8) {
	var shift_ uint8 = 0
	if len(shift) != 0 {
		shift_ = shift[0]
	}
	C.AUTDPointSTMAdd(stm.ptr, C.double(pos[0]), C.double(pos[1]), C.double(pos[2]), C.uchar(shift_))
}

type GainSTM struct {
	STM
}

func NewGainSTM(cnt *Controller) *GainSTM {
	stm := new(GainSTM)
	stm.ptr = unsafe.Pointer(nil)
	C.AUTDGainSTM(&stm.ptr, cnt.Ptr)
	return stm
}

func (stm *GainSTM) Add(gain IGain) {
	C.AUTDGainSTMAdd(stm.ptr, gain.Ptr())
}

type Mode uint16

const (
	PhaseDutyFull Mode = 1
	PhaseFull     Mode = 2
	PhaseHalf     Mode = 4
)

func (stm *GainSTM) GetMode() Mode {
	return Mode(C.AUTDGetGainSTMMode(stm.ptr))
}

func (stm *GainSTM) SetMode(mode Mode) {
	C.AUTDSetGainSTMMode(stm.ptr, C.ushort(mode))
}
