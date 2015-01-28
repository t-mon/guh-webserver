// Copyright (C) 2015 guh
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package main

import (
	"flag"
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var IPFlag = flag.String("ip", "", "The ip address of the webserver (default: 0.0.0.0)")
var portFlag = flag.Int("port", 0, "The port of the webserver (default: 3000)")
var guhIPFlag = flag.String("guh_ip", "", "The IP address of the guh daemon (default: 127.0.0.1)")
var guhPortFlag = flag.Int("guh_port", 0, "The port of the guh daemon (default: 1234)")
var staticFolderFlag = flag.String("static_folder", "", "The location of the folder containing the static files (default: ./public)")
var confPathFlag = flag.String("conf_path", "/etc/guh/guh-webserver.conf", "The location of the config file (default: /etc/guh/guh-webserver.conf)")

func main() {
	// Prase the flag
	flag.Parse()

	config, guhConfig := runConfiguration()

	m := martini.Classic()
	m.Use(martini.Static("/Users/christoph/Desktop/guh-webinterface/build"))
	m.Use(render.Renderer())

	DefineBaseEndPoints(m, guhConfig)
	DefineDeviceEndPoints(m, guhConfig)
	DefineDeviceClassEndPoints(m, guhConfig)
	DefineVendorEndPoints(m, guhConfig)

	m.RunOnAddr(fmt.Sprintf("%v:%v", config.IP, config.Port))
}
