/*
 * File: gain.go
 * Project: gain
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

type Gain struct {
	P unsafe.Pointer
}

func (g *Gain) Ptr() unsafe.Pointer {
	return g.P
}

func (g *Gain) IsHeader() bool {
	return false
}

func (g *Gain) Delete() {
	if g.P == unsafe.Pointer(nil) {
		return
	}
	C.AUTDDeleteGain(g.P)
	g.P = unsafe.Pointer(nil)
}

type IGain interface {
	Ptr() unsafe.Pointer
}

type Focus struct {
	Gain
}

func NewFocus(pos [3]float64, amp ...float64) *Focus {
	g := new(Focus)
	g.P = unsafe.Pointer(nil)
	var amp_ float64 = 1.0
	if len(amp) != 0 {
		amp_ = amp[0]
	}
	C.AUTDGainFocus(&g.P, C.double(pos[0]), C.double(pos[1]), C.double(pos[2]), C.double(amp_))
	return g
}

type Null struct {
	Gain
}

func NewNull() *Null {
	g := new(Null)
	g.P = unsafe.Pointer(nil)
	C.AUTDGainNull(&g.P)
	return g
}

type Grouped struct {
	Gain
}

func NewGrouped(cnt *Controller) *Grouped {
	g := new(Grouped)
	g.P = unsafe.Pointer(nil)
	C.AUTDGainGrouped(&g.P, cnt.Ptr)
	return g
}

func (g *Grouped) Add(devID int, gain IGain) {
	C.AUTDGainGroupedAdd(g.P, C.int(devID), gain.Ptr())
}

type BesselBeam struct {
	Gain
}

func NewBesselBeam(pos [3]float64, dir [3]float64, theta float64, amp ...float64) *BesselBeam {
	g := new(BesselBeam)
	g.P = unsafe.Pointer(nil)
	var amp_ float64 = 1.0
	if len(amp) != 0 {
		amp_ = amp[0]
	}
	C.AUTDGainBesselBeam(&g.P, C.double(pos[0]), C.double(pos[1]), C.double(pos[2]), C.double(dir[0]), C.double(dir[1]), C.double(dir[2]), C.double(theta), C.double(amp_))
	return g
}

type PlaneWave struct {
	Gain
}

func NewPlaneWave(dir [3]float64, amp ...float64) *PlaneWave {
	g := new(PlaneWave)
	g.P = unsafe.Pointer(nil)
	var amp_ float64 = 1.0
	if len(amp) != 0 {
		amp_ = amp[0]
	}
	C.AUTDGainPlaneWave(&g.P, C.double(dir[0]), C.double(dir[1]), C.double(dir[2]), C.double(amp_))
	return g
}

type CustomGain struct {
	Gain
}

func NewCustomGain(amps []float64, phases []float64) *CustomGain {
	g := new(CustomGain)
	g.P = unsafe.Pointer(nil)
	n := len(amps)
	C.AUTDGainCustom(&g.P, (*C.double)(unsafe.Pointer(&amps[0])), (*C.double)(unsafe.Pointer(&phases[0])), C.ulonglong(n))
	return g
}
