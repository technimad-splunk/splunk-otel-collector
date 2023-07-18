// Copyright Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build integration

package tests

import (
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"

	"github.com/signalfx/splunk-otel-collector/tests/testutils"
)

func TestJMXReceiverProvidesAllJVMMetrics(t *testing.T) {

	containers := []testutils.Container{
		testutils.NewContainer().WithContext(
			path.Join(".", "testdata", "server"),
		).WithExposedPorts("7199:7199").WithName("jmx").WillWaitForPorts("7199"),
	}

	// JMX requires a local directory that can be written to, so we must mount a local dir
	// that the collector has write access to.
	tmp_dir := "/etc/otel/collector/tmp"
	working_dir, err := os.Getwd()
	require.NoError(t, err)

	testutils.AssertAllMetricsReceived(t, "all.yaml",
		"all_metrics_config.yaml", containers,
		[]testutils.CollectorBuilder{
			func(collector testutils.Collector) testutils.Collector {
				return collector.WithEnv(map[string]string{
					"TMPDIR": tmp_dir,
				}).WithMount(path.Join(working_dir, "testdata"), tmp_dir)
			},
		})
}