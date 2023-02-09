// Copyright  Splunk, Inc.
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

package gopsutilhelper

import (
	"fmt"
	"os"
)

// HostProc is the proc fs environment variable for gopsutil
const HostProc = "HOST_PROC"

// HostSys is the sys environment variable for gopsutil
const HostSys = "HOST_SYS"

// HostRun is the run environment variable for gopsutil
const HostRun = "HOST_RUN"

// HostEtc is the etc environment variable for gopsutil
const HostEtc = "HOST_ETC"

// HostVar is the var environment varialbe for gopsutil
const HostVar = "HOST_VAR"

// make array so we can loop over each
var envVars = []string{HostProc, HostSys, HostRun, HostEtc, HostVar}

// SetEnvVars sets environment variables from the config for gopustil
func SetEnvVars(paths map[string]string) error {
	for _, v := range envVars {
		if path, ok := paths[v]; ok && path != "" {
			if err := os.Setenv(v, path); err != nil {
				return fmt.Errorf("error setting %s env var %s", v, err.Error())
			}
		}
	}

	return nil
}
