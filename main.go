package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/guh/guh.go"
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
