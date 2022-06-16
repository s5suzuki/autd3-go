/*
 * File: cuda.go
 * Project: cuda
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package cuda

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi-backend-cuda
//
// #include "cuda_backend.h"
import "C"
import (
	"unsafe"

	"github.com/shinolab/autd3-go/v2/holo"
)

type BackendCUDA struct {
	holo.Backend
}

func NewBackendCUDA() *BackendCUDA {
	b := new(BackendCUDA)
	b.P = unsafe.Pointer(nil)

	C.AUTDCUDABackend(&b.P)
	return b
}

func (l *BackendCUDA) Ptr() unsafe.Pointer {
	return l.P
}
