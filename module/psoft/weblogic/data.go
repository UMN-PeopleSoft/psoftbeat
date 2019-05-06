package weblogic

import (
	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
	"strconv"
)

var (
	schema = s.Schema{
		"domain_name":      c.Str("domainName"),
		"domain_type":      c.Str("domainType"),
		"purpose":          c.Str("purpose"),
		"app":              c.Str("app"),
		"env":              c.Str("env"),
      "appenv":           c.Str("appenv"),
		"tools_version":    c.Str("toolsVersion"),
		"weblogic_version": c.Str("weblogicVersion"),
		"server_name":      c.Str("serverName"),
		"host":             c.Str("host"),
		"health":           c.Str("health", s.Optional),
		"load":             c.Str("load"),
		"sessions.current": c.Int("sessions_current"),
		"sessions.high":    c.Int("sessions_high"),

		"jolt_service": s.Object{
			"req_serialize_avg_ms":    c.Float("jolt.req_serialize_avg"),
			"resp_deserialize_avg_ms": c.Float("jolt.resp_deserize_avg", s.Optional),
			"response_avg_ms":         c.Float("jolt.response_avg", s.Optional),
			"total_service_avg_ms":    c.Float("jolt.total_service_avg", s.Optional),
			"service_error_count":     c.Int("jolt.service_error_count", s.Optional),
			"requests_per_sec":        c.Float("jolt.requests_per_sec", s.Optional),
		},

		"jolt_session_pool": s.Object{
			"avail_count":  c.Int("jolt_pool.session_avail_count", s.Optional),
			"in_use_count": c.Int("jolt_pool.session_in_use_count", s.Optional),
		},

		"tcp_sockets": s.Object{
			"close_wait":  c.Int("tcp.close_wait"),
			"time_wait":   c.Int("tcp.time_wait"),
			"established": c.Int("tcp.established"),
			"fin_wait1":   c.Int("tcp.fin_wait1"),
			"fin_wait2":   c.Int("tcp.fin_wait2"),
		},

		"thread_pool": s.Object{
			"queue_length":      c.Int("queue_length"),
			"pending_requests":  c.Int("pending_requests"),
			"standby_count":     c.Int("standby_count"),
			"execute.total":     c.Int("execute.total"),
			"execute.idle":      c.Int("execute.idle"),
			"hogging_count":     c.Int("hogging_count"),
			"stuck_count":       c.Int("stuck_count"),
			"overload_rejected": c.Int("overload_rejected_count"),
			"requests_per_sec":  c.Float("requests_per_sec"),
		},

		"jvm": s.Object{
			"process.cpu_load":     c.Float("jvm.process_cpu_load", s.Optional),
			"system.cpu_load":      c.Float("jvm.system_cpu_load", s.Optional),
			"system.load_avg_1min": c.Float("jvm.system_load_avg_1min", s.Optional),
			"heap_free.pct":        c.Int("jvm.heap_free_pct"),
			"heap_free.current":    c.Int("jvm.heap_free_current"),
		},

		"security": s.Object{
			"invalid_login_attempts": c.Int("security.invalid_login_attempts", s.Optional),
			"current_locked_users":   c.Int("security.current_locked_users", s.Optional),
		},

		"servlets": s.Object{
			"all.exec_time_avg":        c.Float("all.exec_time_avg", s.Optional),
			"all.exec_time_max":        c.Float("all.exec_time_max", s.Optional),
			"all.request_count":        c.Int("all.request_count", s.Optional),
			"psp.exec_time_avg":        c.Float("psp.exec_time_avg", s.Optional),
			"psp.exec_time_max":        c.Float("psp.exec_time_max", s.Optional),
			"psp.request_count":        c.Int("psp.request_count", s.Optional),
			"psc.exec_time_avg":        c.Float("psc.exec_time_avg", s.Optional),
			"psc.exec_time_max":        c.Float("psc.exec_time_max", s.Optional),
			"psc.request_count":        c.Int("psc.request_count", s.Optional),
			"ib_rest.exec_time_avg":    c.Float("ib.rest.exec_time_avg", s.Optional),
			"ib_rest.exec_time_max":    c.Float("ib.rest.exec_time_max", s.Optional),
			"ib_rest.request_count":    c.Int("ib.rest.request_count", s.Optional),
			"ib_service.exec_time_avg": c.Float("ib.service.exec_time_avg", s.Optional),
			"ib_service.exec_time_max": c.Float("ib.service.exec_time_max", s.Optional),
			"ib_service.request_count": c.Int("ib.service.request_count", s.Optional),
			"ib_plc.exec_time_avg":     c.Float("ib.plc.exec_time_avg", s.Optional),
			"ib_plc.exec_time_max":     c.Float("ib.plc.exec_time_max", s.Optional),
			"ib_plc.request_count":     c.Int("ib.plc.request_count", s.Optional),
		},
	}
)

func eventsMapping(weblogicStats []map[string]string) []common.MapStr {

	formattedEvents := []common.MapStr{}
	for i := range weblogicStats {
		formattedEvents = append(formattedEvents, eventMapping(weblogicStats[i]))
	}

	return formattedEvents
}

func eventMapping(webStat map[string]string) common.MapStr {
	// Convert the strings to interface
	var mappedData = make(map[string]interface{})
	var s interface{}
	var err error

	for fld, val := range webStat {
		if s, err = strconv.ParseInt(val, 10, 64); err == nil {
		} else if s, err = strconv.ParseFloat(val, 64); err == nil {
		} else {
			s = val
		}
		mappedData[fld] = s
	}

	event, _ := schema.Apply(mappedData)
	// if there was an error, attach special error key structure to event
	if webStat["errorMsg"] != "" {
		errMessage := common.MapStr{
			"error": common.MapStr{
				"message": webStat["errorMsg"],
			},
		}
		event.DeepUpdate(errMessage)
	}

	return event

}
