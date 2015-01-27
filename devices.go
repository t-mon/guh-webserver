// Copyright (C) 2015 guh
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/guh/guh-libgo"
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
			if foundDevice == nil {
				r.JSON(404, make(map[string]interface{}))
			} else {
				r.JSON(200, foundDevice)
			}
		}
	})

	// Creates a new device
	// TODO maybe split this up in several endpoints to prevent errors where
	// devices support multiple conflicting createMethods
	m.Post("/api/v1/devices.json", func(r render.Render, params martini.Params, request *http.Request) {

		decoder := json.NewDecoder(request.Body)
		var requestBody map[string]interface{}
		err := decoder.Decode(&requestBody)

		device := requestBody["device"].(map[string]interface{})

		deviceClassID := device["deviceClassId"].(string)
		delete(device, "deviceClassId")

		deviceDescriptorID := device["deviceDescriptorId"].(string)
		delete(device, "deviceDescriptorID")

		deviceService := guh.NewDevice(config)
		deviceService.Add(deviceClassID, deviceDescriptorID, device["deviceParams"].([]interface{}))

		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, make(map[string]interface{}))
		}

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
				r.JSON(404, make(map[string]interface{}))
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
