/*
 * File: remote_twincat.go
 * Project: remote_twincat
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package remote_twincat

// #cgo CFLAGS: -I${SRCDIR}/../include/
// #cgo LDFLAGS: -L${SRCDIR}/../ -lautd3capi-link-remote-twincat
//
// #include "remote_twincat_link.h"
import "C"
import (
	"unsafe"
)

type RemoteTwinCAT struct {
	ptr unsafe.Pointer
}

func NewRemoteTwinCAT(remoteIdAddr string, remoteAmsNetId string, localAmsNetId string) *RemoteTwinCAT {
	l := new(RemoteTwinCAT)
	l.ptr = unsafe.Pointer(nil)

	C.AUTDLinkRemoteTwinCAT(&l.ptr, C.CString(remoteIdAddr), C.CString(remoteAmsNetId), C.CString(localAmsNetId))
	return l
}

func (l *RemoteTwinCAT) Ptr() unsafe.Pointer {
	return l.ptr
}
