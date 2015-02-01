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

// DefineRuleEndPoints defines all routes related to rules
func DefineRuleEndPoints(m *martini.ClassicMartini, config guh.Config) {

	// Lists all available devices classes
	m.Get("/api/v1/rules.json", func(r render.Render) {
		ruleService := guh.NewRuleService(config)
		rules, err := ruleService.All()

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, rules)
		}
	})

	// Finds a specific device class identified by its ID
	m.Get("/api/v1/rules/:id.json", func(r render.Render, params martini.Params) {
		ruleService := guh.NewRuleService(config)

		rule, err := ruleService.Find(params["id"])

		if err != nil {
			if err.Error() == guh.RecordNotFoundError {
				r.JSON(404, make(map[string]string))
			} else {
				r.JSON(500, GenerateErrorMessage(err))
			}
		} else {
			r.JSON(200, rule)
		}
	})

	// Enable an existing rule
	m.Patch("/api/v1/rules/:id/enable.json", func(r render.Render, params martini.Params) {
		ruleService := guh.NewRuleService(config)

		err := ruleService.Enable(params["id"])

		if err != nil {
			if err.Error() == guh.RecordNotFoundError {
				r.JSON(404, make(map[string]string))
			} else {
				r.JSON(500, GenerateErrorMessage(err))
			}
		} else {
			r.Status(204)
		}
	})

	// Disables an existing rule
	m.Patch("/api/v1/rules/:id/disable.json", func(r render.Render, params martini.Params) {
		ruleService := guh.NewRuleService(config)

		err := ruleService.Disable(params["id"])

		if err != nil {
			if err.Error() == guh.RecordNotFoundError {
				r.JSON(404, make(map[string]string))
			} else {
				r.JSON(500, GenerateErrorMessage(err))
			}
		} else {
			r.Status(204)
		}
	})

	// Creates a new rule
	m.Post("/api/v1/rules.json", func(r render.Render, params martini.Params, request *http.Request) {

		decoder := json.NewDecoder(request.Body)
		var requestBody map[string]interface{}
		err := decoder.Decode(&requestBody)

		newRuleID := ""
		var newRule guh.Rule
		ruleService := guh.NewRuleService(config)

		if err == nil {
			rule := requestBody["rule"].(map[string]interface{})
			newRuleID, err = ruleService.Add(rule)

			if err == nil {
				newRule, err = ruleService.Find(newRuleID)
			}
		}

		if err != nil {
			r.JSON(500, GenerateErrorMessage(err))
		} else {
			r.JSON(200, newRule)
		}
	})

	// Removes an existing rule permanently
	m.Delete("/api/v1/rules/:id.json", func(r render.Render, params martini.Params) {
		ruleService := guh.NewRuleService(config)

		err := ruleService.Remove(params["id"])

		if err != nil {
			if err.Error() == guh.RecordNotFoundError {
				r.JSON(404, make(map[string]string))
			} else {
				r.JSON(500, GenerateErrorMessage(err))
			}
		} else {
			r.Status(204)
		}
	})

}
