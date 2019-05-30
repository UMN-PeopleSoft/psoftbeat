package main

import (
	"os"

	"github.com/elastic/beats/libbeat/cmd/instance"
	"github.com/elastic/beats/metricbeat/beater"
	"github.com/elastic/beats/libbeat/cfgfile"

	// Comment out the following line to exclude all official metricbeat modules and metricsets
	//_ "github.com/elastic/beats/metricbeat/include"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/UMN-PeopleSoft/psoftbeat/include"
)

var settings = instance.Settings{
	Name: "psoftbeat",
}

func main() {
	_ = cfgfile.ChangeDefaultCfgfileFlag(settings.Name)
	if err := instance.Run(settings, beater.DefaultCreator()); err != nil {
		os.Exit(1)
	}
}
