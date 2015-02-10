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
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/guh/guh-libgo"
	"github.com/martini-contrib/render"
)

// DefineDeviceEndPoints defines all routes related to devices
func DefineDeviceEndPoints(m *martini.ClassicMartini, config guh.Config) {

	// Lists all available Devices
	m.Get("/api/v1/devices.json", func(r render.Render) {
		deviceService := guh.NewDeviceService(config)
		devices, err := deviceService.All()

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, devices)
		}
	})

	// Shows one specific device identified by its ID
	m.Get("/api/v1/devices/:id.json", func(r render.Render, params martini.Params) {
		deviceService := guh.NewDeviceService(config)
		foundDevice, err := deviceService.Find(params["id"])

		if err != nil {
			if err.Error() == guh.RecordNotFoundError {
				r.JSON(404, make(map[string]string))
			} else {
				r.JSON(500, GenerateErrorMessage(err))
			}
		} else {
			r.JSON(200, foundDevice)
		}
	})

	// Creates a new device
	// TODO maybe split this up in several endpoints to prevent errors where
	// devices support multiple conflicting createMethods
	m.Post("/api/v1/devices.json", func(r render.Render, params martini.Params, request *http.Request) {

		newDevice := guh.Device{}

		decoder := json.NewDecoder(request.Body)
		var requestBody map[string]interface{}
		err := decoder.Decode(&requestBody)

		if err == nil {
			device := requestBody["device"].(map[string]interface{})

			deviceClassID := device["deviceClassId"].(string)
			delete(device, "deviceClassId")

			// Check if there is a deviceDescriptorID in the POST body
			var deviceDescriptorID string
			var ok bool
			if deviceDescriptorID, ok = device["deviceDescriptorId"].(string); ok {
				delete(device, "deviceDescriptorID")
			}

			deviceService := guh.NewDeviceService(config)
			newDeviceID := ""
			newDeviceID, err = deviceService.Add(deviceClassID, deviceDescriptorID, device["deviceParams"].([]interface{}))

			if err == nil {
				newDevice, err = deviceService.Find(newDeviceID)
			}
		} else {
			err = errors.New("Error parsing request body")
		}

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, newDevice)
		}
	})

	// Removes a configured device identified by its ID
	m.Delete("/api/v1/devices/:id.json", func(r render.Render, params martini.Params) {
		deviceService := guh.NewDeviceService(config)
		_, err := deviceService.Remove(params["id"])

		if err != nil {
			if err.Error() == guh.RecordNotFoundError {
				r.Status(404)
			} else {
				r.JSON(500, GenerateErrorMessage(err))
			}
		} else {
			r.Status(204)
		}
	})

	m.Get("/api/v1/devices/:id/states.json", func(r render.Render, params martini.Params) {
		stateService := guh.NewStateService(config)
		states, err := stateService.All(params["id"])

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, states)
		}
	})

	m.Post("/api/v1/devices/confirm_paring.json", func(r render.Render, params martini.Params) {

	})

	m.Post("/api/v1/devices/:device_id/actions/:id/execute.json", func(r render.Render, params martini.Params, request *http.Request) {

		// Parse the request body
		decoder := json.NewDecoder(request.Body)
		var requestBody map[string]interface{}
		err := decoder.Decode(&requestBody)
		var actionParams []interface{}
		if requestBody["params"] != nil {
			actionParams = requestBody["params"].([]interface{})
		}

		if err == nil {
			actionTypeService := guh.NewActionTypeService(config)
			response, err := actionTypeService.Execute(params["device_id"], params["id"], actionParams)
			if err == nil {
				if response != "DeviceErrorNoError" {
					err = errors.New(response)
				}
			}
		} else {
			err = errors.New("Error parsing request body")
		}

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.Status(204)
		}
	})

}
