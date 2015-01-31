// Copyright (C) 2015 guh
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package main

// GenerateErrorMessage is a helper method that turns an error into a map
func GenerateErrorMessage(err error) map[string]string {
	message := make(map[string]string)
	message["errorMessage"] = err.Error()

	return message
}
