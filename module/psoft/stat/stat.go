package stat

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"

	"fmt"
	"github.com/UMN-PeopleSoft/psoftbeat/module/psoft"
	"github.com/UMN-PeopleSoft/psoftjmx"

)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("psoft", "stat", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	JmxClient *psoftjmx.PsoftJmxClient
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		JmxClient:     psoft.GetPsoftJMXClient(),
	}, nil
}

func (m *MetricSet) Fetch() ([]common.MapStr, error) {
	// FetchStats will start a queue/thread pool and load up all the current psoft targets
	//   and run JMX Queries to the Nailgun server.
	// global errors that fail all metric gathering will return in err return parameter.
	// single errors for a specific event/target will be included in metric data and processed by eventMapping
	metricData, err := psoft.FetchStats(m.JmxClient)
	if err != nil {
		return nil, fmt.Errorf("Failed fetching psoft Stats")
	}

	return eventsMapping(metricData), nil
}
