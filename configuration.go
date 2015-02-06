// Copyright (C) 2015 guh
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/guh/guh-libgo"
)

type Config struct {
	IP           string
	Port         int
	GuhIP        string
	GuhPort      int
	StaticFolder string
}

func runConfiguration() (Config, guh.Config) {
	// Load the config file, if it exists
	var config Config
	guhConfig := guh.Config{}
	if _, err := os.Stat(*confPathFlag); !os.IsNotExist(err) {
		if _, err := toml.DecodeFile(*confPathFlag, &config); err != nil {
			fmt.Println("Could not parse the config file!")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Set the Port
	if config.IP == "" && *IPFlag == "" {
		config.IP = "0.0.0.0"
	} else if *IPFlag != "" {
		config.IP = *IPFlag
	}

	// Set the Port
	if config.Port == 0 && *portFlag == 0 {
		config.Port = 3000
	} else if *portFlag > 0 {
		config.Port = *portFlag
	}

	// Set the guh IP
	if config.GuhIP == "" && *guhIPFlag == "" {
		config.GuhIP = "127.0.0.1"
	} else if *guhIPFlag != "" {
		config.GuhIP = *guhIPFlag
	}

	// Set the guh port
	if config.GuhPort == 0 && *guhPortFlag == 0 {
		config.GuhPort = 1234
	} else if *portFlag > 0 {
		config.GuhPort = *guhPortFlag
	}

	// Set the folder containing the static files
	if config.StaticFolder == "" && *staticFolderFlag == "" {
		config.StaticFolder = "./public"
	} else if *staticFolderFlag != "" {
		config.StaticFolder = *staticFolderFlag
	}

	fmt.Println("[guh-webserver]   ConfigFile:", *confPathFlag)
	fmt.Println("[guh-webserver]           IP:", config.IP)
	fmt.Println("[guh-webserver]         Port:", config.Port)
	fmt.Println("[guh-webserver] StaticFolder:", config.StaticFolder)

	guhConfig.IP = config.GuhIP
	guhConfig.Port = config.GuhPort

	return config, guhConfig
}
