/*
 * File: audio_file.go
 * Project: audio_file
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package audio_file

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi-modulation-audio-file
//
// #include "audio_file_modulation.h"
import "C"
import (
	"unsafe"

	"github.com/shinolab/autd3-go/v2/autd3"
)

type RawPCM struct {
	autd3.Modulation
}

func NewRawPCM(filename string, samplingFreq float64, samplingFreqDiv uint32) *RawPCM {
	m := new(RawPCM)
	m.P = unsafe.Pointer(nil)

	C.AUTDModulationRawPCM(&m.P, C.CString(filename), C.double(samplingFreq), C.uint(samplingFreqDiv))
	return m
}

type Wav struct {
	autd3.Modulation
}

func NewWav(filename string, samplingFreqDiv uint32) *Wav {
	m := new(Wav)
	m.P = unsafe.Pointer(nil)

	C.AUTDModulationWav(&m.P, C.CString(filename), C.uint(samplingFreqDiv))
	return m
}
