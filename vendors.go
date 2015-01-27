// Copyright (C) 2015 guh
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

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
		deviceClass := guh.NewDeviceClass(config)
		deviceClasses, err := deviceClass.AllByVendor(params["vendor_id"])

		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, deviceClasses)
		}
	})

}
