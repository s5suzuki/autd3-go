/*
 * File: link.go
 * Project: link
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 08/08/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package autd3

import "unsafe"

type Link struct {
	Ptr unsafe.Pointer
}
