package main

import (
	"os"

	"github.com/UMN-PeopleSoft/psoftbeat/cmd"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/UMN-PeopleSoft/psoftbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
