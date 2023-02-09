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

package metrics

import (
	"time"

	"github.com/signalfx/golib/v3/datapoint"
	k8sutil "github.com/signalfx/signalfx-agent/pkg/monitors/kubernetes/utils"
	atypes "github.com/signalfx/signalfx-agent/pkg/monitors/types"
	"github.com/signalfx/signalfx-agent/pkg/utils"
	v1 "k8s.io/api/core/v1"
)

func datapointsForReplicationController(rc *v1.ReplicationController) []*datapoint.Datapoint {
	dimensions := map[string]string{
		"metric_source":        "kubernetes",
		"kubernetes_namespace": rc.Namespace,
		"kubernetes_uid":       string(rc.UID),
		"kubernetes_name":      rc.Name,
	}

	if rc.Spec.Replicas == nil {
		return nil
	}
	return makeReplicaDPs("replication_controller", dimensions,
		*rc.Spec.Replicas, rc.Status.AvailableReplicas)
}

func dimensionForReplicationController(rc *v1.ReplicationController) *atypes.Dimension {
	props, tags := k8sutil.PropsAndTagsFromLabels(rc.Labels)
	props["kubernetes_workload_name"] = rc.Name
	props["kubernetes_workload"] = "ReplicationController"
	props["replication_controller_creation_timestamp"] = rc.GetCreationTimestamp().Format(time.RFC3339)

	for _, or := range rc.OwnerReferences {
		props[utils.LowercaseFirstChar(or.Kind)] = or.Name
		props[utils.LowercaseFirstChar(or.Kind)+"_uid"] = string(or.UID)
	}

	return &atypes.Dimension{
		Name:       "kubernetes_uid",
		Value:      string(rc.UID),
		Properties: props,
		Tags:       tags,
	}
}
