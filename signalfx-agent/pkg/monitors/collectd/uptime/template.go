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

//go:build linux
// +build linux

package uptime

// AUTOGENERATED BY scripts/collectd-template-to-go.  DO NOT EDIT!!

import (
	"text/template"

	"github.com/signalfx/signalfx-agent/pkg/monitors/collectd"
)

// CollectdTemplate is a template for a uptime collectd config file
var CollectdTemplate = template.Must(collectd.InjectTemplateFuncs(template.New("uptime")).Parse(`
LoadPlugin uptime

<Chain "PostCache"> 
  <Rule "set_uptime_monitor_id"> 
    <Match "regex"> 
      Plugin "^uptime$" 
    </Match> 
    <Target "set"> 
      MetaData "monitorID" "{{.MonitorID}}" 
    </Target> 
  </Rule> 
</Chain>
`)).Option("missingkey=error")
