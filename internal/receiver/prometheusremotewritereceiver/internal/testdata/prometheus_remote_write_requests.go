// Copyright Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testdata

import (
	"time"

	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

var (
	Jan20 = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
)

func SampleCounterTs() []prompb.TimeSeries {
	return []prompb.TimeSeries{
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "http_requests_total"},
				{Name: "method", Value: "GET"},
				{Name: "status", Value: "200"},
			},
			Samples: []prompb.Sample{
				{Value: 1024, Timestamp: Jan20.UnixMilli()},
			},
		},
	}
}
func SampleCounterWq() *prompb.WriteRequest {
	return &prompb.WriteRequest{Timeseries: SampleCounterTs()}
}

func SampleGaugeTs() []prompb.TimeSeries {
	return []prompb.TimeSeries{
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "go_goroutines"},
			},
			Samples: []prompb.Sample{
				{Value: 42, Timestamp: Jan20.UnixMilli()},
			},
		},
	}
}

func SampleGaugeWq() *prompb.WriteRequest { return &prompb.WriteRequest{Timeseries: SampleGaugeTs()} }

func SampleHistogramTs() []prompb.TimeSeries {
	return []prompb.TimeSeries{
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "api_request_duration_seconds_bucket"},
				{Name: "le", Value: "0.1"},
			},
			Samples: []prompb.Sample{
				{Value: 500, Timestamp: Jan20.UnixMilli()},
			},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "api_request_duration_seconds_bucket"},
				{Name: "le", Value: "0.2"},
			},
			Samples: []prompb.Sample{
				{Value: 1500, Timestamp: Jan20.UnixMilli()},
			},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "api_request_duration_seconds_count"},
			},
			Samples: []prompb.Sample{
				{Value: 2500, Timestamp: Jan20.UnixMilli()},
			},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "api_request_duration_seconds_sum"},
			},
			Samples: []prompb.Sample{
				{Value: 350, Timestamp: Jan20.UnixMilli()},
			},
		},
	}
}

func SampleHistogramWq() *prompb.WriteRequest {
	return &prompb.WriteRequest{
		Timeseries: SampleHistogramTs(),
	}
}

func SampleSummaryTs() []prompb.TimeSeries {
	return []prompb.TimeSeries{
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "rpc_duration_seconds"},
				{Name: "quantile", Value: "0.5"},
			},
			Samples: []prompb.Sample{
				{Value: 0.25, Timestamp: Jan20.UnixMilli()},
			},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "rpc_duration_seconds"},
				{Name: "quantile", Value: "0.9"},
			},
			Samples: []prompb.Sample{
				{Value: 0.35, Timestamp: Jan20.Add(1 * time.Second).UnixMilli()},
			},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "rpc_duration_seconds_sum"},
			},
			Samples: []prompb.Sample{
				{Value: 123.5, Timestamp: Jan20.UnixMilli()},
			},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "rpc_duration_seconds_count"},
			},
			Samples: []prompb.Sample{
				{Value: 1500, Timestamp: Jan20.UnixMilli()},
			},
		},
	}
}

func SampleSummaryWq() *prompb.WriteRequest {
	return &prompb.WriteRequest{
		Timeseries: SampleSummaryTs(),
	}
}

func ExpectedCounter(sfxCompat bool) pmetric.Metrics {
	result := pmetric.NewMetrics()
	resourceMetrics := result.ResourceMetrics().AppendEmpty()
	scopeMetrics := resourceMetrics.ScopeMetrics().AppendEmpty()
	scopeMetrics.Scope().SetName("prometheusremotewrite")
	scopeMetrics.Scope().SetVersion("0.1")
	metric := scopeMetrics.Metrics().AppendEmpty()
	metric.SetName("http_requests_total")
	counter := metric.SetEmptySum()
	counter.SetIsMonotonic(true)
	counter.SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	dp := counter.DataPoints().AppendEmpty()
	dp.SetTimestamp(pcommon.NewTimestampFromTime(Jan20))
	dp.SetStartTimestamp(pcommon.NewTimestampFromTime(Jan20))
	dp.SetIntValue(1024)
	dp.Attributes().PutStr("method", "GET")
	dp.Attributes().PutStr("status", "200")

	if sfxCompat {
		metric := scopeMetrics.Metrics().AppendEmpty()
		metric.SetName("prometheus.total_NaN_datapoints")
		counter := metric.SetEmptySum()
		counter.SetIsMonotonic(true)
		counter.SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	}

	return result
}

func GetWriteRequestsOfAllTypesWithoutMetadata() []*prompb.WriteRequest {
	var sampleWriteRequestsNoMetadata = []*prompb.WriteRequest{
		// Counter
		SampleCounterWq(),
		// Gauge
		SampleGaugeWq(),
		// Histogram
		SampleHistogramWq(),
		// Summary
		SampleSummaryWq(),
	}
	return sampleWriteRequestsNoMetadata
}

func FlattenWriteRequests(request []*prompb.WriteRequest) *prompb.WriteRequest {
	var ts []prompb.TimeSeries
	for _, req := range request {
		for _, t := range req.Timeseries {
			ts = append(ts, t)
		}
	}
	var md []prompb.MetricMetadata
	for _, req := range request {
		for _, t := range req.Metadata {
			md = append(md, t)
		}
	}
	return &prompb.WriteRequest{
		Timeseries: ts,
		Metadata:   md,
	}
}