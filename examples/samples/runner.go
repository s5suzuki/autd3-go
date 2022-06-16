/*
 * File: runner.go
 * Project: examples
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
	"fmt"

	"github.com/shinolab/autd3-go/v2/autd3"
)

type F struct {
	f    func(*autd3.Controller)
	desc string
}

func Run(cnt *autd3.Controller) {
	tests := []F{F{f: Focus, desc: "Single focus"},
		F{f: Bessel, desc: "BesselBeam"},
		F{f: Holo, desc: "Holo gain"},
		F{f: PointSTM, desc: "PointSTM"},
		F{f: GainSTM, desc: "GainSTM"}}

	cnt.Clear()
	cnt.Synchronize()

	println("================================== Firmware information ==========================================")
	firmList := cnt.FirmwareInfoList()
	for _, info := range firmList {
		println(info)
	}
	println("==================================================================================================")

	for {
		for i, f := range tests {
			fmt.Printf("[%d]: %s\n", i, f.desc)
		}
		fmt.Print("[Other]: finish\n")

		var i int
		if _, err := fmt.Scanln(&i); err != nil {
			break
		}
		if i >= len(tests) {
			break
		}

		tests[i].f(cnt)

		println("press enter to finish...")
		var input string
		fmt.Scanln(&input)

		cnt.Stop()
	}

	cnt.Close()
}
