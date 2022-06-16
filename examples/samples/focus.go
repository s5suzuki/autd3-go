/*
 * File: focus.go
 * Project: samples
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package samples

import "github.com/shinolab/autd3-go/v2/autd3"

func Focus(cnt *autd3.Controller) {
	config := autd3.NewSilencerConfig()
	defer config.Delete()
	cnt.Send(config)

	g := autd3.NewFocus([3]float64{90, 80, 150})
	defer g.Delete()
	m := autd3.NewSine(150)
	defer m.Delete()

	cnt.Send(m, g)
}
