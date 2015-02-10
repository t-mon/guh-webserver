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
	"github.com/go-martini/martini"
	"github.com/guh/guh-libgo"
	"github.com/martini-contrib/render"
)

// DefineVendorEndPoints defines all routes related to vendors
func DefineVendorEndPoints(m *martini.ClassicMartini, config guh.Config) {

	// Lists all available vendors
	m.Get("/api/v1/vendors.json", func(r render.Render) {
		vendor := guh.NewVendor(config)
		vendors, err := vendor.All()

		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, vendors)
		}
	})

	// Gets one specific vendor identified by its ID
	m.Get("/api/v1/vendors/:id.json", func(r render.Render, params martini.Params) {
		device := guh.NewVendor(config)
		foundVendor, err := device.Find(params["id"])

		if err != nil {
			r.JSON(500, err)
		} else {
			if foundVendor == "" {
				r.JSON(404, make(map[string]interface{}))
			} else {
				r.JSON(200, foundVendor)
			}
		}
	})

	// Gets all available device classes of a specific vendor identified by his ID
	m.Get("/api/v1/vendors/:vendor_id/device_classes.json", func(r render.Render, params martini.Params) {
		deviceClassService := guh.NewDeviceClassService(config)
		deviceClasses, err := deviceClassService.AllByVendor(params["vendor_id"])

		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, deviceClasses)
		}
	})

}
