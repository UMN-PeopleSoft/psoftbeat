package stat

import (
	"github.com/elastic/beats/libbeat/common"
)


func eventsMapping(stats []map[string]interface{}) []common.MapStr {

	formattedEvents := []common.MapStr{}
	for i := range stats {
		formattedEvents = append(formattedEvents, eventMapping(stats[i]))
	}

	return formattedEvents
}

func eventMapping(stat map[string]interface{}) common.MapStr {
	// Convert map to the MapStr
	event := common.MapStr{}
   for fld, val := range stat {
		event.Put(fld, val)
	}

	// if there was an error, attach special error key structure to event
	if stat["errorMsg"] != "" {
		errMessage := common.MapStr{
			"error": common.MapStr{
				"message": stat["errorMsg"],
			},
		}
		event.DeepUpdate(errMessage)
	}

	return event
}