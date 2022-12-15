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

package v1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/testing"

	kcpcorev1 "github.com/kcp-dev/client-go/kubernetes/typed/core/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var limitRangesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "limitranges"}
var limitRangesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "LimitRange"}

type limitRangesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *limitRangesClusterClient) Cluster(clusterPath logicalcluster.Path) kcpcorev1.LimitRangesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &limitRangesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of LimitRanges that match those selectors across all clusters.
func (c *limitRangesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.LimitRangeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(limitRangesResource, limitRangesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &corev1.LimitRangeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.LimitRangeList{ListMeta: obj.(*corev1.LimitRangeList).ListMeta}
	for _, item := range obj.(*corev1.LimitRangeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested LimitRanges across all clusters.
func (c *limitRangesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(limitRangesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type limitRangesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *limitRangesNamespacer) Namespace(namespace string) corev1client.LimitRangeInterface {
	return &limitRangesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type limitRangesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *limitRangesClient) Create(ctx context.Context, limitRange *corev1.LimitRange, opts metav1.CreateOptions) (*corev1.LimitRange, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(limitRangesResource, c.ClusterPath, c.Namespace, limitRange), &corev1.LimitRange{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

func (c *limitRangesClient) Update(ctx context.Context, limitRange *corev1.LimitRange, opts metav1.UpdateOptions) (*corev1.LimitRange, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(limitRangesResource, c.ClusterPath, c.Namespace, limitRange), &corev1.LimitRange{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

func (c *limitRangesClient) UpdateStatus(ctx context.Context, limitRange *corev1.LimitRange, opts metav1.UpdateOptions) (*corev1.LimitRange, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(limitRangesResource, c.ClusterPath, "status", c.Namespace, limitRange), &corev1.LimitRange{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

func (c *limitRangesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(limitRangesResource, c.ClusterPath, c.Namespace, name, opts), &corev1.LimitRange{})
	return err
}

func (c *limitRangesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(limitRangesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.LimitRangeList{})
	return err
}

func (c *limitRangesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.LimitRange, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(limitRangesResource, c.ClusterPath, c.Namespace, name), &corev1.LimitRange{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

// List takes label and field selectors, and returns the list of LimitRanges that match those selectors.
func (c *limitRangesClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.LimitRangeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(limitRangesResource, limitRangesKind, c.ClusterPath, c.Namespace, opts), &corev1.LimitRangeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.LimitRangeList{ListMeta: obj.(*corev1.LimitRangeList).ListMeta}
	for _, item := range obj.(*corev1.LimitRangeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *limitRangesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(limitRangesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *limitRangesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.LimitRange, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(limitRangesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &corev1.LimitRange{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

func (c *limitRangesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.LimitRangeApplyConfiguration, opts metav1.ApplyOptions) (*corev1.LimitRange, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(limitRangesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &corev1.LimitRange{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

func (c *limitRangesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.LimitRangeApplyConfiguration, opts metav1.ApplyOptions) (*corev1.LimitRange, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(limitRangesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &corev1.LimitRange{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}
