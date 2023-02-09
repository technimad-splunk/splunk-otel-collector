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

package hostid

import (
	systemdutil "github.com/coreos/go-systemd/util"
)

// MachineID returns the Linux machine-id, which is present on most newer Linux
// distros.  It is more useful than hostname as a unique identifier.  See
// http://man7.org/linux/man-pages/man5/machine-id.5.html
func MachineID() string {
	mid, err := systemdutil.GetMachineID()
	if err != nil {
		return ""
	}
	return mid
}
