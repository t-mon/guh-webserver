// Copyright (C) 2015 guh
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-martini/martini"
	"github.com/guh/guh-libgo"
	"github.com/martini-contrib/render"
)

func main() {

	var config guh.Config
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		config = guh.Config{IP: os.Getenv("GUH_IP"), Port: os.Getenv("GUH_PORT")}
	} else {
		config = guh.LoadConfig("config.json")
	}

	flag.Parse()

	fmt.Println("  guh_ip:", config.IP)
	fmt.Println("guh_port:", config.Port)

	m := martini.Classic()
	m.Use(render.Renderer())

	DefineBaseEndPoints(m, config)
	DefineDeviceEndPoints(m, config)
	DefineDeviceClassEndPoints(m, config)
	DefineVendorEndPoints(m, config)

	m.Run()

}
