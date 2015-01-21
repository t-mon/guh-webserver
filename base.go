package main

import (
	"github.com/go-martini/martini"
	"github.com/guh/guh.go"
	"github.com/martini-contrib/render"
)

func DefineBaseEndPoints(m *martini.ClassicMartini, config guh.Config) {
	m.Get("/api/v1/introspect.json", func(r render.Render) {
		base := guh.NewDevice(config)
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
