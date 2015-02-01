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

func DefineBaseEndPoints(m *martini.ClassicMartini, config guh.Config) {
	m.Get("/api/v1/introspect.json", func(r render.Render) {
		base := guh.NewBase(config)
		introspect, err := base.Introspect()

		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, introspect)
		}
	})

	m.Get("/api/v1/version.json", func(r render.Render) {
		base := guh.NewBase(config)
		version, err := base.Version()
		if err != nil {
			r.JSON(500, err)
		} else {
			r.JSON(200, version)
		}
	})
}
