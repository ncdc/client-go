//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package v2beta2

import (
	"github.com/kcp-dev/logicalcluster/v3"

	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	"k8s.io/client-go/rest"

	kcpautoscalingv2beta2 "github.com/kcp-dev/client-go/kubernetes/typed/autoscaling/v2beta2"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpautoscalingv2beta2.AutoscalingV2beta2ClusterInterface = (*AutoscalingV2beta2ClusterClient)(nil)

type AutoscalingV2beta2ClusterClient struct {
	*kcptesting.Fake
}

func (c *AutoscalingV2beta2ClusterClient) Cluster(clusterPath logicalcluster.Path) autoscalingv2beta2.AutoscalingV2beta2Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AutoscalingV2beta2Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *AutoscalingV2beta2ClusterClient) HorizontalPodAutoscalers() kcpautoscalingv2beta2.HorizontalPodAutoscalerClusterInterface {
	return &horizontalPodAutoscalersClusterClient{Fake: c.Fake}
}

var _ autoscalingv2beta2.AutoscalingV2beta2Interface = (*AutoscalingV2beta2Client)(nil)

type AutoscalingV2beta2Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *AutoscalingV2beta2Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AutoscalingV2beta2Client) HorizontalPodAutoscalers(namespace string) autoscalingv2beta2.HorizontalPodAutoscalerInterface {
	return &horizontalPodAutoscalersClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
