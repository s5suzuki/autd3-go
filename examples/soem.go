/*
 * File: main.go
 * Project: gomod-test
 * Created Date: 15/06/2022
 * Author: Shun Suzuki
 * -----
 * Last Modified: 16/06/2022
 * Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
 * -----
 * Copyright (c) 2022 Shun Suzuki. All rights reserved.
 *
 */

package main

import (
	"fmt"
	"os"

	"autd3-go-example/samples"

	"github.com/shinolab/autd3-go/v2/autd3"
	"github.com/shinolab/autd3-go/v2/soem"
)

func onLost(msg string) {
	println(msg)
	os.Exit(-1)
}

func getAdapter() string {
	adapters := soem.EnumerateAdapters()
	for i, adapter := range adapters {
		fmt.Printf("[%d]: %s, %s\n", i, adapter.Desc, adapter.Name)
	}

	fmt.Print("choose: ")

	var i int
	if _, err := fmt.Scanln(&i); err != nil {
		fmt.Printf("failed to read integer: %s\n", err)
		os.Exit(-1)
	}

	if i >= len(adapters) {
		fmt.Print("index out of range\n")
		os.Exit(-1)
	}

	return adapters[i].Name
}

func main() {
	cnt := autd3.NewController()
	defer cnt.Delete()

	cnt.AddDevice([3]float64{0, 0, 0}, [3]float64{0, 0, 0})

	ifname := getAdapter()
	soem.RegisterOnLostCallback(onLost)
	link := soem.NewSOEM(ifname, cnt.NumDevices(), 1, true)

	if !cnt.Open(link) {
		println(autd3.GetLastError())
		os.Exit(-1)
	}

	samples.Run(cnt)
}
