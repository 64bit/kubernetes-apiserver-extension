// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BlueDataClusterLister helps list BlueDataClusters.
type BlueDataClusterLister interface {
	// List lists all BlueDataClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BlueDataCluster, err error)
	// BlueDataClusters returns an object that can list and get BlueDataClusters.
	BlueDataClusters(namespace string) BlueDataClusterNamespaceLister
	BlueDataClusterListerExpansion
}

// blueDataClusterLister implements the BlueDataClusterLister interface.
type blueDataClusterLister struct {
	indexer cache.Indexer
}

// NewBlueDataClusterLister returns a new BlueDataClusterLister.
func NewBlueDataClusterLister(indexer cache.Indexer) BlueDataClusterLister {
	return &blueDataClusterLister{indexer: indexer}
}

// List lists all BlueDataClusters in the indexer.
func (s *blueDataClusterLister) List(selector labels.Selector) (ret []*v1alpha1.BlueDataCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BlueDataCluster))
	})
	return ret, err
}

// BlueDataClusters returns an object that can list and get BlueDataClusters.
func (s *blueDataClusterLister) BlueDataClusters(namespace string) BlueDataClusterNamespaceLister {
	return blueDataClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BlueDataClusterNamespaceLister helps list and get BlueDataClusters.
type BlueDataClusterNamespaceLister interface {
	// List lists all BlueDataClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BlueDataCluster, err error)
	// Get retrieves the BlueDataCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BlueDataCluster, error)
	BlueDataClusterNamespaceListerExpansion
}

// blueDataClusterNamespaceLister implements the BlueDataClusterNamespaceLister
// interface.
type blueDataClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BlueDataClusters in the indexer for a given namespace.
func (s blueDataClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BlueDataCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BlueDataCluster))
	})
	return ret, err
}

// Get retrieves the BlueDataCluster from the indexer for a given namespace and name.
func (s blueDataClusterNamespaceLister) Get(name string) (*v1alpha1.BlueDataCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bluedatacluster"), name)
	}
	return obj.(*v1alpha1.BlueDataCluster), nil
}
