/*
 * File: header.go
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

import "unsafe"

type Sendable interface {
	Ptr() unsafe.Pointer
	IsHeader() bool
}
