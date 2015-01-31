## guh REST

## Installation

## Getting started

To start the web server simply type

    ./guh-webserver

You may configure the process with the following flags:

`ip`: The IP of the web server (default `0.0.0.0`)

`port`: The port of the web server (default `3000`)

`guh_ip`: The IP of the guhd server (default `127.0.0.1`)

`guh_port`: The port of the guhd server (default `1234`)

`static_folder`: The folder containing the static files of guh-webinterface (default: `./public`)

`conf_path`: A path pointing at a config file (default: `/etc/guh/guh-webserver.conf`)

For example:

    ./guh-webserver --guh_port=192.168.0.2 --guh_port=80

Here is an example for `guh-webserver.conf`:

    IP = "127.0.0.1"
    Port = 3000
    GuhIP = "192.168.0.3"
    GuhPort = 1234
    StaticFolder = "/my/cusom/public"

### License & Copyright

Copyright (c) 2015 guh

This software may be modified and distributed under the terms of the MIT license. See the LICENSE file for details.

### TODO

 - [ ] Refactor code and rename structs properly (Entity and EntityService)
 - [x] Properly handle deviceErrors etc. in responses
 - [ ] Refactor code to use something like "RegisterEndPoint()" to be able to autogenerate meta information (required params, relationship to guh core, etc.) about the API
 - [x] Convert ENV params to real command line params
   - [x] Params for guh-ip, guh-port, config-path, port
 - [x] Serve static files from configurable directory (http://stackoverflow.com/a/14187941/641032)


#### Endpoints

 - [x] get /core/introspect.json
   - [x] JSONRPC.Introspect
 - [x] get /core/version.json
   - [x] JSONRPC.Version
 - [x] get /devices.json
   - [x] Devices.GetConfiguredDevices
 - [x] get /devices/:id.json
   - [x] Devices.GetConfiguredDevices (filtered)
 - [x] delete /devices/:id.json
   - [x] Devices.RemoveConfiguredDevice
 - [x] post /devices.json
   - [x] Devices.AddConfiguredDevice
   - [x] Devices.PairDevice
 - [ ] post /devices/confirm_pairing.json
   - [ ] Devices.ConfirmPairing
 - [ ] get /devices/:device_id/actions.json
 - [ ] get /devices/:device_id/actions/:id.json
 - [ ] post /devices/:device_id/execute/:id.json
 - [x] get /device_classes.json
   - [x] Devices.GetSupportedDevices
 - [x] get /device_classes/:id.json
   - [x] Devices.GetSupportedDevices (filtered)
 - [x] get /device_classes/:device_class_id/action_types.json
   - [x] Devices.GetActionTypes
 - [x] get /device_classes/:device_class_id/state_types.json
   - [x] Devices.GetStateTypes
 - [x] get /device_classes/:id/discover.json
   - [x] Devices.GetDiscoveredDevices
 - [x] get /rules.json
   - [x] Rules.GetRules
 - [x] get /rules/:id.json
   - [x] Rules.GetRuleDetails
 - [x] post /rules.json
   - [x] Rules.AddRule
 - [ ] patch /rules/:id/disable.json
   - [ ] Rules.DisableRule
 - [ ] patch /rules/:id/enable.json
   - [ ] Rules.EnableRule
 - [ ] delete /rules/:id.json
   - [ ] Rules.RemoveRule
 - [x] get /vendors.json do
   - [x] Devices.GetSupportedVendors
 - [x] get /vendors/:id.json
   - [x] Devices.GetSupportedVendors (filtered)
 - [x] get /vendor/:id/device_classes.json
   - [x] Devices.GetSupportedDevices (filtered)
 - [ ] get /ws
   - [ ] JSONRPC.SetNotificationStatus

#### RPC

These RPCs are available but don't have a corresponding REST endpoint yet.

 - [ ] Actions.ExecuteAction
 - [ ] Actions.GetActionType
 - [ ] Devices.GetEventTypes
 - [ ] Devices.GetPluginConfiguration
 - [ ] Devices.GetPlugins
 - [ ] Devices.GetStateValue
 - [ ] Devices.GetStateValues
 - [ ] Devices.SetPluginConfiguration
 - [ ] Events.GetEventType
 - [ ] Logging.GetLogEntries
 - [ ] Rules.FindRules
 - [ ] States.GetStateType
