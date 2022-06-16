/*
 * File: holo.go
 * Project: holo
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package holo

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi-gain-holo
//
// #include "holo_gain.h"
import "C"
import (
	"unsafe"

	"github.com/shinolab/autd3-go/v2/autd3"
)

type Backend struct {
	P unsafe.Pointer
}

func (b *Backend) Delete() {
	C.AUTDDeleteBackend(b.P)
}

type IBackend interface {
	Ptr() unsafe.Pointer
}

type BackendEigen struct {
	Backend
}

func NewBackendEigen() *BackendEigen {
	b := new(BackendEigen)
	b.P = unsafe.Pointer(nil)

	C.AUTDEigenBackend(&b.P)
	return b
}

func (b *BackendEigen) Ptr() unsafe.Pointer {
	return b.P
}

type Constraint interface {
	Ty() int
	Param() unsafe.Pointer
}

type DontCare struct {
}

func NewDontCare() *DontCare {
	return new(DontCare)
}

func (c *DontCare) Ty() int {
	return 0
}
func (c *DontCare) Param() unsafe.Pointer {
	return nil
}

type Normalize struct {
}

func NewNormalize() *Normalize {
	return new(Normalize)
}
func (c *Normalize) Ty() int {
	return 1
}
func (c *Normalize) Param() unsafe.Pointer {
	return nil
}

type Uniform struct {
	value float64
	param unsafe.Pointer
}

func NewUniform(value float64) *Uniform {
	c := new(Uniform)
	c.value = value
	c.param = unsafe.Pointer(&c.value)
	return c
}

func (c *Uniform) Ty() int {
	return 2
}

func (c *Uniform) Param() unsafe.Pointer {
	return c.param
}

type Clamp struct {
}

func (c *Clamp) Ty() int {
	return 3
}
func (c *Normalize) Clamp() unsafe.Pointer {
	return nil
}

func NewClamp() *Clamp {
	return new(Clamp)
}

type Holo struct {
	autd3.Gain
}

func (g *Holo) Add(pos [3]float64, amp float64) {
	C.AUTDGainHoloAdd(g.P, C.double(pos[0]), C.double(pos[1]), C.double(pos[2]), C.double(amp))
}

func (g *Holo) SetConstraint(c Constraint) {
	C.AUTDSetConstraint(g.P, C.int(c.Ty()), c.Param())
}

type SDP struct {
	Holo
}

func NewSDP(backend IBackend, param ...float64) *SDP {
	g := new(SDP)
	g.P = unsafe.Pointer(nil)
	alpha := 1e-3
	lambda := 0.9
	var repeat uint64 = 100
	if len(param) != 0 {
		alpha = param[0]
	}
	if len(param) != 1 {
		lambda = param[1]
	}
	if len(param) != 2 {
		repeat = uint64(param[2])
	}

	C.AUTDGainHoloSDP(&g.P, backend.Ptr(), C.double(alpha), C.double(lambda), C.ulonglong(repeat))
	return g
}

type EVD struct {
	Holo
}

func NewEVD(backend IBackend, param ...float64) *EVD {
	g := new(EVD)
	g.P = unsafe.Pointer(nil)
	gamma := 1.0
	if len(param) != 0 {
		gamma = param[0]
	}

	C.AUTDGainHoloEVD(&g.P, backend.Ptr(), C.double(gamma))
	return g
}

type Naive struct {
	Holo
}

func NewNaive(backend IBackend) *Naive {
	g := new(Naive)
	g.P = unsafe.Pointer(nil)

	C.AUTDGainHoloNaive(&g.P, backend.Ptr())
	return g
}

type GS struct {
	Holo
}

func NewGS(backend IBackend, param ...uint64) *GS {
	g := new(GS)
	g.P = unsafe.Pointer(nil)

	var repeat uint64 = 100
	if len(param) != 0 {
		repeat = param[0]
	}

	C.AUTDGainHoloGS(&g.P, backend.Ptr(), C.ulonglong(repeat))
	return g
}

type GSPAT struct {
	Holo
}

func NewGSPAT(backend IBackend, param ...uint64) *GSPAT {
	g := new(GSPAT)
	g.P = unsafe.Pointer(nil)

	var repeat uint64 = 100
	if len(param) != 0 {
		repeat = param[0]
	}

	C.AUTDGainHoloGSPAT(&g.P, backend.Ptr(), C.ulonglong(repeat))
	return g
}

type LM struct {
	Holo
}

func NewLM(backend IBackend, param ...float64) *LM {
	g := new(LM)
	g.P = unsafe.Pointer(nil)

	eps1 := 1e-8
	eps2 := 1e-8
	tau := 1e-3
	var kMax uint64 = 5
	if len(param) != 0 {
		eps1 = param[0]
	}
	if len(param) != 1 {
		eps2 = param[1]
	}
	if len(param) != 2 {
		tau = param[2]
	}
	if len(param) != 3 {
		kMax = uint64(param[3])
	}

	C.AUTDGainHoloLM(&g.P, backend.Ptr(), C.double(eps1), C.double(eps2), C.double(tau), C.ulonglong(kMax), nil, C.int(0))
	return g
}

type Greedy struct {
	Holo
}

func NewGreedy(backend IBackend, param ...int) *Greedy {
	g := new(Greedy)
	g.P = unsafe.Pointer(nil)

	phaseDiv := 16
	if len(param) != 0 {
		phaseDiv = param[0]
	}

	C.AUTDGainHoloGreedy(&g.P, backend.Ptr(), C.int(phaseDiv))
	return g
}
