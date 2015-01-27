## guh REST

## Installation

## Getting started

To start the web server simply type

    ./guh_rest

You may configure the process with the following `ENV` variables:

`IP`: The IP of the web server (default `0.0.0.0`)

`PORT`: The port of the web server (default `3000`)

`GUH_IP`: The IP of the guhd server (default `127.0.0.1`)

`GUH_PORT`: The port of the guhd server (default `1234`)

For example:

    GUH_IP=10.0.0.11 GUH_PORT=12345 IP=127.0.0.1 PORT=80 ./guh_rest

### License & Copyright

Copyright (c) 2015 guh

This software may be modified and distributed under the terms of the MIT license. See the LICENSE file for details.

### TODO

 - [ ] Properly handle deviceErrors etc. in responses
 - [ ] Convert ENV params to real command line params
   - [ ] Params for guh-ip, guh-port, config-path, port
 - [ ] Create a config file in the user's home directory (~/.config/guh-webserver)

