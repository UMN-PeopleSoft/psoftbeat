package psoft

import (
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"

	"github.com/UMN-PeopleSoft/psoftjmx"
)

// Select psoftJMX API version
const psoftjmxAPIVersion = "1.1"

var client *psoftjmx.PsoftJmxClient

func init() {
	// Register the ModuleFactory function for the "psoft" module.
	if err := mb.Registry.AddModule("psoft", NewModule); err != nil {
		panic(err)
	}
}

type Module struct {
	mb.BaseModule
}

func NewModule(base mb.BaseModule) (mb.Module, error) {

	config := DefaultConfig()
	if err := base.UnpackConfig(&config); err != nil {
		return nil, err
	}
	var err error

	if client == nil {
		// map configs
		jmxClientConfig := &psoftjmx.JMXConfig{
			PathInventoryFile: config.PathInventoryFile,
			PathBlackoutFile:  config.PathBlackoutFile,
			PathExclusionFile: config.PathExclusionFile,
			AttribWebMetrics:  config.AttribWebMetrics,
			AttribAppMetrics:  config.AttribAppMetrics,
			AttribPrcMetrics:  config.AttribPrcMetrics,
			LogLevel:          config.LogLevel,
			ConcurrentWorkers: config.ConcurrentWorkers,
			NailgunServerConn: config.NailgunServerConn,
			JavaPath:          config.JavaPath,
			ConcatenateDomainWithHost: config.ConcatenateDomainWithHost,
			UseLastXCharactersOfHost: config.UseLastXCharactersOfHost, 
			LocalInventoryOnly:config.LocalInventoryOnly,
		
		}
		// send custom configs to the common JMXClient and capture only this type of metricset targets
		// NewClient will setup the Nailgun server, load JMX queries and attribute lists
		//   An err will indicate a configuration problem or Nailgun server not running/communicating
		client, err = psoftjmx.NewClient(jmxClientConfig)
		if err != nil {
			return nil, err
		}

		logp.Info("Started Nailgun Server")
	}
	return &Module{BaseModule: base}, nil
}

// NewPsoftJMXClient initializes and returns a new PsoftJMX client
func GetPsoftJMXClient() *psoftjmx.PsoftJmxClient {
	//return the existing client at module level
	return client
}

// FetchStats all psoft metric data for specific domains for the metricset Type
func FetchStats(client *psoftjmx.PsoftJmxClient) ([]map[string]interface{}, error) {

	// now get all the metrics for the filtered targets
	psoftMetrics, err := client.GetMetrics()
	return psoftMetrics, err
}