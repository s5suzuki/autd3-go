/*
 * File: emulator.go
 * Project: examples
 * Created Date: 16/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 14/08/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package main

import (
	"os"

	"autd3-go-example/samples"

	"github.com/shinolab/autd3-go/v2/autd3"
	"github.com/shinolab/autd3-go/v2/emulator"
)

func main() {
	cnt := autd3.NewController()
	defer cnt.Delete()

	cnt.AddDevice([3]float64{0, 0, 0}, [3]float64{0, 0, 0})

	link := emulator.NewEmulator().Port(50632).Build()

	if !cnt.Open(link) {
		println(autd3.GetLastError())
		os.Exit(-1)
	}

	samples.Run(cnt)
}
