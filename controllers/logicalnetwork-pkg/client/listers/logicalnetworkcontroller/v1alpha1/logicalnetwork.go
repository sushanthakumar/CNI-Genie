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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "github.com/Huawei-PaaS/CNI-Genie/utils"
	r "github.com/Huawei-PaaS/CNI-Genie/controllers/logicalnetwork-pkg/apis/logicalnetworkcontroller/v1alpha1"
)

// LogicalNetworkLister helps list LogicalNetworks.
type LogicalNetworkLister interface {
	// List lists all LogicalNetworks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LogicalNetwork, err error)
	// LogicalNetworks returns an object that can list and get LogicalNetworks.
	LogicalNetworks(namespace string) LogicalNetworkNamespaceLister
	LogicalNetworkListerExpansion
}

// logicalNetworkLister implements the LogicalNetworkLister interface.
type logicalNetworkLister struct {
	indexer cache.Indexer
}

// NewLogicalNetworkLister returns a new LogicalNetworkLister.
func NewLogicalNetworkLister(indexer cache.Indexer) LogicalNetworkLister {
	return &logicalNetworkLister{indexer: indexer}
}

// List lists all LogicalNetworks in the indexer.
func (s *logicalNetworkLister) List(selector labels.Selector) (ret []*v1alpha1.LogicalNetwork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicalNetwork))
	})
	return ret, err
}

// LogicalNetworks returns an object that can list and get LogicalNetworks.
func (s *logicalNetworkLister) LogicalNetworks(namespace string) LogicalNetworkNamespaceLister {
	return logicalNetworkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LogicalNetworkNamespaceLister helps list and get LogicalNetworks.
type LogicalNetworkNamespaceLister interface {
	// List lists all LogicalNetworks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LogicalNetwork, err error)
	// Get retrieves the LogicalNetwork from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LogicalNetwork, error)
	LogicalNetworkNamespaceListerExpansion
}

// logicalNetworkNamespaceLister implements the LogicalNetworkNamespaceLister
// interface.
type logicalNetworkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LogicalNetworks in the indexer for a given namespace.
func (s logicalNetworkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LogicalNetwork, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicalNetwork))
	})
	return ret, err
}

// Get retrieves the LogicalNetwork from the indexer for a given namespace and name.
func (s logicalNetworkNamespaceLister) Get(name string) (*v1alpha1.LogicalNetwork, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(r.Resource("logicalnetwork"), name)
	}
	return obj.(*v1alpha1.LogicalNetwork), nil
}
