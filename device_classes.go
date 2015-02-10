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
	"net/http"

	"github.com/go-martini/martini"
	"github.com/guh/guh-libgo"
	"github.com/martini-contrib/render"
)

// DefineDeviceClassEndPoints defines all routes related to device classes
func DefineDeviceClassEndPoints(m *martini.ClassicMartini, config guh.Config) {

	// Lists all available devices classes
	m.Get("/api/v1/device_classes.json", func(r render.Render) {
		deviceClassService := guh.NewDeviceClassService(config)
		deviceClasses, err := deviceClassService.All()

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, deviceClasses)
		}
	})

	// Finds a specific device class identified by its ID
	m.Get("/api/v1/device_classes/:id.json", func(r render.Render, params martini.Params) {
		deviceClassService := guh.NewDeviceClassService(config)

		foundDeviceClass, err := deviceClassService.Find(params["id"])

		if err != nil {
			if err.Error() == guh.RecordNotFoundError {
				r.JSON(404, make(map[string]string))
			} else {
				r.JSON(500, GenerateErrorMessage(err))
			}
		} else {
			r.JSON(200, foundDeviceClass)
		}
	})

	// Returns a list of all discovered devices
	m.Get("/api/v1/device_classes/:device_class_id/discover.json", func(r render.Render, params martini.Params, request *http.Request) {

		deviceClassService := guh.NewDeviceClassService(config)

		var discoveryParams []interface{}
		err := json.Unmarshal([]byte(request.FormValue("discovery_params")), &discoveryParams)

		discoveredDevices, err := deviceClassService.Discover(params["device_class_id"], discoveryParams)

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, discoveredDevices["deviceDescriptors"])
		}
	})

	// Lists all available state types of a device class
	m.Get("/api/v1/device_classes/:device_class_id/action_types.json", func(r render.Render, params martini.Params) {
		actionType := guh.NewActionTypeService(config)

		actionTypes, err := actionType.All(params["device_class_id"])

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, actionTypes)
		}
	})

	// Lists all available event types of a device class
	m.Get("/api/v1/device_classes/:device_class_id/event_types.json", func(r render.Render, params martini.Params) {
		eventTypeService := guh.NewEventTypeService(config)

		eventTypes, err := eventTypeService.All(params["device_class_id"])

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, eventTypes)
		}
	})

	// Lists all available state types of a device class
	m.Get("/api/v1/device_classes/:device_class_id/state_types.json", func(r render.Render, params martini.Params) {
		stateTypeService := guh.NewStateTypeService(config)

		stateTypes, err := stateTypeService.All(params["device_class_id"])

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, stateTypes)
		}
	})
}
