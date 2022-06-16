/*
 * File: bessel.go
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
	"math"

	"github.com/shinolab/autd3-go/v2/autd3"
)

func Bessel(cnt *autd3.Controller) {
	config := autd3.NewSilencerConfig()
	defer config.Delete()
	cnt.Send(config)

	g := autd3.NewBesselBeam([3]float64{90, 80, 0}, [3]float64{0, 0, 1}, 13.0/180.0*math.Pi)
	defer g.Delete()
	m := autd3.NewSine(150)
	defer m.Delete()

	cnt.Send(m, g)
}
