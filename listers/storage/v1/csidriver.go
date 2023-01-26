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
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	storagev1listers "k8s.io/client-go/listers/storage/v1"
	"k8s.io/client-go/tools/cache"
)

// CSIDriverClusterLister can list CSIDrivers across all workspaces, or scope down to a CSIDriverLister for one workspace.
// All objects returned here must be treated as read-only.
type CSIDriverClusterLister interface {
	// List lists all CSIDrivers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1.CSIDriver, err error)
	// Cluster returns a lister that can list and get CSIDrivers in one workspace.
	Cluster(clusterName logicalcluster.Name) storagev1listers.CSIDriverLister
	CSIDriverClusterListerExpansion
}

type cSIDriverClusterLister struct {
	indexer cache.Indexer
}

// NewCSIDriverClusterLister returns a new CSIDriverClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewCSIDriverClusterLister(indexer cache.Indexer) *cSIDriverClusterLister {
	return &cSIDriverClusterLister{indexer: indexer}
}

// List lists all CSIDrivers in the indexer across all workspaces.
func (s *cSIDriverClusterLister) List(selector labels.Selector) (ret []*storagev1.CSIDriver, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*storagev1.CSIDriver))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get CSIDrivers.
func (s *cSIDriverClusterLister) Cluster(clusterName logicalcluster.Name) storagev1listers.CSIDriverLister {
	return &cSIDriverLister{indexer: s.indexer, clusterName: clusterName}
}

// cSIDriverLister implements the storagev1listers.CSIDriverLister interface.
type cSIDriverLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all CSIDrivers in the indexer for a workspace.
func (s *cSIDriverLister) List(selector labels.Selector) (ret []*storagev1.CSIDriver, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*storagev1.CSIDriver))
	})
	return ret, err
}

// Get retrieves the CSIDriver from the indexer for a given workspace and name.
func (s *cSIDriverLister) Get(name string) (*storagev1.CSIDriver, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storagev1.Resource("csidrivers"), name)
	}
	return obj.(*storagev1.CSIDriver), nil
}
