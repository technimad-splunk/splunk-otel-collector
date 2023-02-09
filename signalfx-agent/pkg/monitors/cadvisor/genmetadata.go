// Code generated by monitor-code-gen. DO NOT EDIT.

package cadvisor

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const (
	groupPodEphemeralStats = "podEphemeralStats"
)

var groupSet = map[string]bool{
	groupPodEphemeralStats: true,
}

const (
	containerCPUCfsPeriods                = "container_cpu_cfs_periods"
	containerCPUCfsThrottledPeriods       = "container_cpu_cfs_throttled_periods"
	containerCPUCfsThrottledTime          = "container_cpu_cfs_throttled_time"
	containerCPUPercent                   = "container_cpu_percent"
	containerCPUSystemSecondsTotal        = "container_cpu_system_seconds_total"
	containerCPUUsageSecondsTotal         = "container_cpu_usage_seconds_total"
	containerCPUUserSecondsTotal          = "container_cpu_user_seconds_total"
	containerCPUUtilization               = "container_cpu_utilization"
	containerCPUUtilizationPerCore        = "container_cpu_utilization_per_core"
	containerFsIoCurrent                  = "container_fs_io_current"
	containerFsIoTimeSecondsTotal         = "container_fs_io_time_seconds_total"
	containerFsIoTimeWeightedSecondsTotal = "container_fs_io_time_weighted_seconds_total"
	containerFsLimitBytes                 = "container_fs_limit_bytes"
	containerFsReadSecondsTotal           = "container_fs_read_seconds_total"
	containerFsReadsMergedTotal           = "container_fs_reads_merged_total"
	containerFsReadsTotal                 = "container_fs_reads_total"
	containerFsSectorReadsTotal           = "container_fs_sector_reads_total"
	containerFsSectorWritesTotal          = "container_fs_sector_writes_total"
	containerFsUsageBytes                 = "container_fs_usage_bytes"
	containerFsWriteSecondsTotal          = "container_fs_write_seconds_total"
	containerFsWritesMergedTotal          = "container_fs_writes_merged_total"
	containerFsWritesTotal                = "container_fs_writes_total"
	containerLastSeen                     = "container_last_seen"
	containerMemoryFailcnt                = "container_memory_failcnt"
	containerMemoryFailuresTotal          = "container_memory_failures_total"
	containerMemoryRss                    = "container_memory_rss"
	containerMemoryUsageBytes             = "container_memory_usage_bytes"
	containerMemoryWorkingSetBytes        = "container_memory_working_set_bytes"
	containerSpecCPUPeriod                = "container_spec_cpu_period"
	containerSpecCPUQuota                 = "container_spec_cpu_quota"
	containerSpecCPUShares                = "container_spec_cpu_shares"
	containerSpecMemoryLimitBytes         = "container_spec_memory_limit_bytes"
	containerSpecMemorySwapLimitBytes     = "container_spec_memory_swap_limit_bytes"
	containerStartTimeSeconds             = "container_start_time_seconds"
	containerTasksState                   = "container_tasks_state"
	machineCPUCores                       = "machine_cpu_cores"
	machineCPUFrequencyKhz                = "machine_cpu_frequency_khz"
	machineMemoryBytes                    = "machine_memory_bytes"
	podEphemeralStorageCapacityBytes      = "pod_ephemeral_storage_capacity_bytes"
	podEphemeralStorageUsedBytes          = "pod_ephemeral_storage_used_bytes"
	podNetworkReceiveBytesTotal           = "pod_network_receive_bytes_total"
	podNetworkReceiveErrorsTotal          = "pod_network_receive_errors_total"
	podNetworkReceivePacketsDroppedTotal  = "pod_network_receive_packets_dropped_total"
	podNetworkReceivePacketsTotal         = "pod_network_receive_packets_total"
	podNetworkTransmitBytesTotal          = "pod_network_transmit_bytes_total"
	podNetworkTransmitErrorsTotal         = "pod_network_transmit_errors_total"
	podNetworkTransmitPacketsDroppedTotal = "pod_network_transmit_packets_dropped_total"
	podNetworkTransmitPacketsTotal        = "pod_network_transmit_packets_total"
)

var metricSet = map[string]monitors.MetricInfo{
	containerCPUCfsPeriods:                {Type: datapoint.Counter},
	containerCPUCfsThrottledPeriods:       {Type: datapoint.Counter},
	containerCPUCfsThrottledTime:          {Type: datapoint.Counter},
	containerCPUPercent:                   {Type: datapoint.Counter},
	containerCPUSystemSecondsTotal:        {Type: datapoint.Counter},
	containerCPUUsageSecondsTotal:         {Type: datapoint.Counter},
	containerCPUUserSecondsTotal:          {Type: datapoint.Counter},
	containerCPUUtilization:               {Type: datapoint.Counter},
	containerCPUUtilizationPerCore:        {Type: datapoint.Counter},
	containerFsIoCurrent:                  {Type: datapoint.Gauge},
	containerFsIoTimeSecondsTotal:         {Type: datapoint.Counter},
	containerFsIoTimeWeightedSecondsTotal: {Type: datapoint.Counter},
	containerFsLimitBytes:                 {Type: datapoint.Gauge},
	containerFsReadSecondsTotal:           {Type: datapoint.Counter},
	containerFsReadsMergedTotal:           {Type: datapoint.Counter},
	containerFsReadsTotal:                 {Type: datapoint.Counter},
	containerFsSectorReadsTotal:           {Type: datapoint.Counter},
	containerFsSectorWritesTotal:          {Type: datapoint.Counter},
	containerFsUsageBytes:                 {Type: datapoint.Gauge},
	containerFsWriteSecondsTotal:          {Type: datapoint.Counter},
	containerFsWritesMergedTotal:          {Type: datapoint.Counter},
	containerFsWritesTotal:                {Type: datapoint.Counter},
	containerLastSeen:                     {Type: datapoint.Gauge},
	containerMemoryFailcnt:                {Type: datapoint.Counter},
	containerMemoryFailuresTotal:          {Type: datapoint.Counter},
	containerMemoryRss:                    {Type: datapoint.Gauge},
	containerMemoryUsageBytes:             {Type: datapoint.Gauge},
	containerMemoryWorkingSetBytes:        {Type: datapoint.Gauge},
	containerSpecCPUPeriod:                {Type: datapoint.Gauge},
	containerSpecCPUQuota:                 {Type: datapoint.Gauge},
	containerSpecCPUShares:                {Type: datapoint.Gauge},
	containerSpecMemoryLimitBytes:         {Type: datapoint.Gauge},
	containerSpecMemorySwapLimitBytes:     {Type: datapoint.Gauge},
	containerStartTimeSeconds:             {Type: datapoint.Gauge},
	containerTasksState:                   {Type: datapoint.Gauge},
	machineCPUCores:                       {Type: datapoint.Gauge},
	machineCPUFrequencyKhz:                {Type: datapoint.Gauge},
	machineMemoryBytes:                    {Type: datapoint.Gauge},
	podEphemeralStorageCapacityBytes:      {Type: datapoint.Gauge, Group: groupPodEphemeralStats},
	podEphemeralStorageUsedBytes:          {Type: datapoint.Gauge, Group: groupPodEphemeralStats},
	podNetworkReceiveBytesTotal:           {Type: datapoint.Counter},
	podNetworkReceiveErrorsTotal:          {Type: datapoint.Counter},
	podNetworkReceivePacketsDroppedTotal:  {Type: datapoint.Counter},
	podNetworkReceivePacketsTotal:         {Type: datapoint.Counter},
	podNetworkTransmitBytesTotal:          {Type: datapoint.Counter},
	podNetworkTransmitErrorsTotal:         {Type: datapoint.Counter},
	podNetworkTransmitPacketsDroppedTotal: {Type: datapoint.Counter},
	podNetworkTransmitPacketsTotal:        {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	containerCPUPercent:           true,
	containerCPUUtilization:       true,
	containerFsReadSecondsTotal:   true,
	containerFsReadsTotal:         true,
	containerFsUsageBytes:         true,
	containerFsWriteSecondsTotal:  true,
	containerFsWritesTotal:        true,
	containerMemoryFailuresTotal:  true,
	containerMemoryUsageBytes:     true,
	containerSpecCPUPeriod:        true,
	containerSpecCPUQuota:         true,
	containerSpecMemoryLimitBytes: true,
	machineCPUCores:               true,
	machineMemoryBytes:            true,
	podNetworkReceiveBytesTotal:   true,
	podNetworkReceiveErrorsTotal:  true,
	podNetworkTransmitBytesTotal:  true,
	podNetworkTransmitErrorsTotal: true,
}

var groupMetricsMap = map[string][]string{
	groupPodEphemeralStats: {
		podEphemeralStorageCapacityBytes,
		podEphemeralStorageUsedBytes,
	},
}

var cadvisorMonitorMetadata = monitors.Metadata{
	MonitorType:     "cadvisor",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}

var kubeletStatsMonitorMetadata = monitors.Metadata{
	MonitorType:     "kubelet-stats",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
