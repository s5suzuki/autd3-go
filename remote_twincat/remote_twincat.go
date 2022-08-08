/*
 * File: remote_twincat.go
 * Project: remote_twincat
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 08/08/2022
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

	"github.com/shinolab/autd3-go/v2/autd3"
)

type RemoteTwinCAT struct {
	remoteIdAddr   string
	remoteAmsNetId string
	localAmsNetId  string
}

func NewRemoteTwinCAT(remoteIdAddr string, remoteAmsNetId string) *RemoteTwinCAT {
	l := new(RemoteTwinCAT)
	l.remoteIdAddr = remoteIdAddr
	l.remoteAmsNetId = remoteAmsNetId
	return l
}

func (l *RemoteTwinCAT) LocalAmsNetId(localAmsNetId string) *RemoteTwinCAT {
	l.localAmsNetId = localAmsNetId
	return l
}

func (l *RemoteTwinCAT) Build() *autd3.Link {
	link := new(autd3.Link)
	link.Ptr = unsafe.Pointer(nil)
	C.AUTDLinkRemoteTwinCAT(&link.Ptr, C.CString(l.remoteIdAddr), C.CString(l.remoteAmsNetId), C.CString(l.localAmsNetId))
	return link
}
