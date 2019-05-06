package appserver

import (
	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"strconv"
)

var (
	schema = s.Schema{
		"domain_name":                  c.Str("domainName"),
		"domain_type":                  c.Str("domainType"),
		"purpose":                      c.Str("purpose"),
		"app":                          c.Str("app"),
		"env":                          c.Str("env"),
      "appenv":                       c.Str("appenv"),
		"tools_version":                c.Str("toolsVersion"),
		"server_name":                  c.Str("serverName"),
		"host":                         c.Str("host"),
		"health":                       c.Str("health", s.Optional),
		"load":                         c.Float("load"),
		"ps_load":                      c.Str("ps_load"),
		"server_count":                 c.Int("server_count"),
		"server_down_count":            c.Int("server_down_count"),
		"client_count":                 c.Int("client_count", s.Optional),
		"long_running_request_count":   c.Int("long_running_request_count", s.Optional),
		"requests.total":               c.Int("requests_total"),
		"requests.avg":                 c.Float("reqests_avg"),
		"requests.per_sec_total":       c.Float("request_per_sec_total"),
		"requests.per_sec_avg":         c.Float("request_per_sec_avg"),
		"requests.process_time_ms_avg": c.Float("request_process_time_ms_avg", s.Optional),
		"requests.service_time_ms_avg": c.Float("request_service_time_ms_avg", s.Optional),
		"client_trans_aborted_count":   c.Int("client_trans_aborted_count", s.Optional),
		"psappsrv_active_pct":          c.Float("psappsrv_active_pct"),

		"cache": s.Object{
			"memory_size_kb_total": c.Int("cache_memory_size_kb_total", s.Optional),
			"memory_size_kb_avg":   c.Int("cache_memory_size_kb_avg", s.Optional),
		},

		"queue": s.Object{
			"server_count":            c.Int("queue.server_count", s.Optional),
			"requests_per_sec":        c.Float("queue.requests_per_sec", s.Optional),
			"depth":                   c.Float("queue.depth", s.Optional),
			"ib.sub_requests_per_sec": c.Float("ib.sub_requests_per_sec", s.Optional),
			"ib.sub_max_queue_depth":  c.Float("ib.sub_max_queue_depth", s.Optional),
			"ib.pub_requests_per_sec": c.Float("ib.pub_requests_per_sec", s.Optional),
			"ib.pub_max_queue_depth":  c.Float("ib.pub_max_queue_depth", s.Optional),
		},

		"tcp_sockets": s.Object{
			"close_wait":  c.Int("tcp.close_wait_state"),
			"time_wait":   c.Int("tcp.time_wait_state"),
			"established": c.Int("tcp.established_state"),
			"fin_wait1":   c.Int("tcp.fin_wait1_state"),
			"fin_wait2":   c.Int("tcp.fin_wait2_state"),
		},
	}
)

func eventsMapping(appservStats []map[string]string) []common.MapStr {

	formattedEvents := []common.MapStr{}
	for i := range appservStats {
		formattedEvents = append(formattedEvents, eventMapping(appservStats[i]))
	}

	return formattedEvents
}

func eventMapping(appStat map[string]string) common.MapStr {
	// Convert the strings to interface
	var mappedData = make(map[string]interface{})
	var s interface{}
	var err error

	for fld, val := range appStat {
		if s, err = strconv.ParseInt(val, 10, 64); err == nil {
		} else if s, err = strconv.ParseFloat(val, 64); err == nil {
		} else {
			s = val
		}
		mappedData[fld] = s
	}

	event, _ := schema.Apply(mappedData)
	// if there was an error, attach special error key structure to event
	if appStat["errorMsg"] != "" {
		errMessage := common.MapStr{
			"error": common.MapStr{
				"message": appStat["errorMsg"],
			},
		}
		event.DeepUpdate(errMessage)
	}

	return event

}
