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
    StaticFolder = "/my/custom/webserver.conf"

### License & Copyright

Copyright (c) 2015 guh

This software may be modified and distributed under the terms of the MIT license. See the LICENSE file for details.

### TODO

 - [ ] Properly handle deviceErrors etc. in responses
 - [ ] Refactor code to use something like "RegisterEndPoint()" to be able to autogenerate meta information (required params, relationship to guh core, etc.) about the API
 - [x] Convert ENV params to real command line params
   - [x] Params for guh-ip, guh-port, config-path, port
 - [x] Serve static files from configurable directory (http://stackoverflow.com/a/14187941/641032)

