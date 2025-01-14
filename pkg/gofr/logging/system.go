package logging

import (
	"os"
	"runtime"
)

var hostname string

func fetchSystemStats() map[string]interface{} {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	if hostname == "" {
		hostname, _ = os.Hostname()
	}

	stats := make(map[string]interface{})
	stats["alloc"] = m.Alloc
	stats["totalAlloc"] = m.TotalAlloc
	stats["sys"] = m.Sys
	stats["numGC"] = m.NumGC
	stats["goRoutines"] = runtime.NumGoroutine()
	stats["host"] = hostname

	return stats
}
