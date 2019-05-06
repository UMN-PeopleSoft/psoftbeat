package scheduler

import (
	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"strconv"
)

var (
	schema = s.Schema{
		"domain_name":         c.Str("domainName"),
		"domain_type":         c.Str("domainType"),
		"purpose":             c.Str("purpose"),
		"app":                 c.Str("app"),
		"env":                 c.Str("env"),
      "appenv":              c.Str("appenv"),
		"tools_version":       c.Str("toolsVersion"),
		"server_name":         c.Str("serverName"),
		"host":                c.Str("host"),
		"health":              c.Str("health", s.Optional),
		"load":                c.Float("load"),
		"server_count":        c.Int("server_count"),
		"server_down_count":   c.Int("server_down_count"),
		"requests.total":      c.Int("requests_total"),
		"psaeserv_active_pct": c.Float("psaesrv_active_pct"),

		"cache": s.Object{
			"memory_size_kb_total": c.Int("memory_size_kb_total", s.Optional),
		},

		"tcp_sockets": s.Object{
			"close_wait":  c.Int("tcp.close_wait"),
			"time_wait":   c.Int("tcp.time_wait"),
			"established": c.Int("tcp.established"),
			"fin_wait1":   c.Int("tcp.fin_wait1"),
			"fin_wait2":   c.Int("tcp.fin_wait2"),
		},
	}
)

func eventsMapping(schedulerStats []map[string]string) []common.MapStr {

	formattedEvents := []common.MapStr{}
	for i := range schedulerStats {
		formattedEvents = append(formattedEvents, eventMapping(schedulerStats[i]))
	}

	return formattedEvents
}

func eventMapping(schedulerStat map[string]string) common.MapStr {
	// Convert the strings to interface
	var mappedData = make(map[string]interface{})
	var s interface{}
	var err error

	for fld, val := range schedulerStat {
		if s, err = strconv.ParseInt(val, 10, 64); err == nil {
		} else if s, err = strconv.ParseFloat(val, 64); err == nil {
		} else {
			s = val
		}
		mappedData[fld] = s
	}

	event, _ := schema.Apply(mappedData)
	// if there was an error, attach special error key structure to event
	if schedulerStat["errorMsg"] != "" {
		errMessage := common.MapStr{
			"error": common.MapStr{
				"message": schedulerStat["errorMsg"],
			},
		}
		event.DeepUpdate(errMessage)
	}

	return event

}
