// Code generated by monitor-code-gen. DO NOT EDIT.

package memory

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "memory"

var groupSet = map[string]bool{}

const (
	memoryAvailable   = "memory.available"
	memoryBuffered    = "memory.buffered"
	memoryCached      = "memory.cached"
	memoryFree        = "memory.free"
	memorySlabRecl    = "memory.slab_recl"
	memorySlabUnrecl  = "memory.slab_unrecl"
	memorySwapFree    = "memory.swap_free"
	memorySwapTotal   = "memory.swap_total"
	memorySwapUsed    = "memory.swap_used"
	memoryTotal       = "memory.total"
	memoryUsed        = "memory.used"
	memoryUtilization = "memory.utilization"
)

var metricSet = map[string]monitors.MetricInfo{
	memoryAvailable:   {Type: datapoint.Gauge},
	memoryBuffered:    {Type: datapoint.Gauge},
	memoryCached:      {Type: datapoint.Gauge},
	memoryFree:        {Type: datapoint.Gauge},
	memorySlabRecl:    {Type: datapoint.Gauge},
	memorySlabUnrecl:  {Type: datapoint.Gauge},
	memorySwapFree:    {Type: datapoint.Gauge},
	memorySwapTotal:   {Type: datapoint.Gauge},
	memorySwapUsed:    {Type: datapoint.Gauge},
	memoryTotal:       {Type: datapoint.Gauge},
	memoryUsed:        {Type: datapoint.Gauge},
	memoryUtilization: {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	memoryBuffered:    true,
	memoryCached:      true,
	memoryFree:        true,
	memorySlabRecl:    true,
	memorySlabUnrecl:  true,
	memoryTotal:       true,
	memoryUsed:        true,
	memoryUtilization: true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "memory",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
