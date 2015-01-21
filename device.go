package main

import (
	"github.com/go-martini/martini"
	"github.com/guh/guh.go"
	"github.com/martini-contrib/render"
)

// DefineDeviceEndPoints defines all routes related to devices
func DefineDeviceEndPoints(m *martini.ClassicMartini, config guh.Config) {

	// Lists all available Devices
	m.Get("/api/v1/devices.json", func(r render.Render) {
		device := guh.NewDevice(config)
		devices, err := device.All()

		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, devices)
		}
	})

	// Shows one specific device identified by its ID
	m.Get("/api/v1/devices/:id.json", func(r render.Render, params martini.Params) {
		device := guh.NewDevice(config)
		foundDevice, err := device.Find(params["id"])

		if err != nil {
			r.JSON(500, err)
		} else {
			if foundDevice == "" {
				r.JSON(404, "{}")
			} else {
				r.JSON(200, foundDevice)
			}
		}
	})

	// Creates a new device
	m.Post("/api/v1/devices.json", func(r render.Render, params martini.Params) {
		// deviceClassID := params["device"]["deviceClassId"]
		// delete(params["device"], "deviceClassId")
		//
		// device := guh.NewDevice(config)
		// createdDevice, err := device.create(deviceClassID, params["deviceDescriptorId"], params["device"]["deviceParams"])
	})

	// Removes a configured device identified by its ID
	m.Delete("/api/v1/devices/:id.json", func(r render.Render, params martini.Params) {
		device := guh.NewDevice(config)
		deletedDevice, err := device.Remove(params["id"])

		if err != nil {
			r.JSON(500, err)
		} else {
			if deletedDevice == "" {
				r.JSON(404, "{}")
			} else {
				r.JSON(200, deletedDevice)
			}
		}
	})

	// Returns a list of all discovered devices
	m.Get("/api/v1/devices/discover.json", func(r render.Render, params martini.Params) {

	})

	m.Get("/api/v1/devices/pair.json", func(r render.Render, params martini.Params) {

	})

	m.Post("/api/v1/devices/confirm_paring.json", func(r render.Render, params martini.Params) {

	})

	m.Get("/api/v1/devices/:device_id/actions.json", func(r render.Render, params martini.Params) {

	})

	m.Get("/api/v1/devices/:device_id/actions/:id/execute.json", func(r render.Render, params martini.Params) {

	})

	m.Get("/api/v1/devices/:device_id/states.json", func(r render.Render, params martini.Params) {

	})

}
