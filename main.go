/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *                                                                                     *
 * Copyright (c) 2015 guh                                                              *
 *                                                                                     *
 * Permission is hereby granted, free of charge, to any person obtaining a copy        *
 * of this software and associated documentation files (the "Software"), to deal       *
 * in the Software without restriction, including without limitation the rights        *
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell           *
 * copies of the Software, and to permit persons to whom the Software is               *
 * furnished to do so, subject to the following conditions:                            *
 *                                                                                     *
 * The above copyright notice and this permission notice shall be included in all      *
 * copies or substantial portions of the Software.                                     *
 *                                                                                     *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR          *
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,            *
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE         *
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER              *
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,       *
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE       *
 * SOFTWARE.                                                                           *
 *                                                                                     *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

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
	// Parse the flag
	flag.Parse()

	config, guhConfig := runConfiguration()

	m := martini.Classic()
	m.Use(martini.Static(config.StaticFolder))
	m.Use(render.Renderer())

	DefineBaseEndPoints(m, guhConfig)
	DefineDeviceEndPoints(m, guhConfig)
	DefineDeviceClassEndPoints(m, guhConfig)
	DefineVendorEndPoints(m, guhConfig)
	DefineRuleEndPoints(m, guhConfig)

	m.RunOnAddr(fmt.Sprintf("%v:%v", config.IP, config.Port))
}
