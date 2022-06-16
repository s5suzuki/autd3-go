/*
 * File: holo.go
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

import (
	"github.com/shinolab/autd3-go/v2/autd3"
	"github.com/shinolab/autd3-go/v2/holo"
)

func Holo(cnt *autd3.Controller) {
	config := autd3.NewSilencerConfig()
	defer config.Delete()
	cnt.Send(config)

	b := holo.NewBackendEigen()
	defer b.Delete()

	g := holo.NewGSPAT(b)
	defer g.Delete()
	g.Add([3]float64{60, 80, 150}, 1.0)
	g.Add([3]float64{120, 80, 150}, 1.0)
	m := autd3.NewSine(150)
	defer m.Delete()

	cnt.Send(m, g)
}
