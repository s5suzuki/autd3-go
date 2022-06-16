/*
 * File: gain_stm.go
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

func GainSTM(cnt *autd3.Controller) {
	config := autd3.NewSilencerConfig(0xFFFF, 4096)
	defer config.Delete()
	cnt.Send(config)

	stm := autd3.NewGainSTM(cnt)
	defer stm.Delete()

	radius := 30.0
	size := 200
	x := 90.0
	y := 80.0
	z := 150.0
	for i := 0; i < size; i++ {
		theta := 2.0 * math.Pi * float64(i) / float64(size)
		p := [3]float64{x + radius*math.Cos(theta), y + radius*math.Sin(theta), z}
		g := autd3.NewFocus(p)
		stm.Add(g)
		g.Delete()
	}
	stm.SetFrequency(1.0)

	m := autd3.NewStatic()
	defer m.Delete()

	cnt.Send(m, stm)
}
