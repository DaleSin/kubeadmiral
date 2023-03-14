// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PropagationPolicyLister helps list PropagationPolicies.
// All objects returned here must be treated as read-only.
type PropagationPolicyLister interface {
	// List lists all PropagationPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PropagationPolicy, err error)
	// PropagationPolicies returns an object that can list and get PropagationPolicies.
	PropagationPolicies(namespace string) PropagationPolicyNamespaceLister
	PropagationPolicyListerExpansion
}

// propagationPolicyLister implements the PropagationPolicyLister interface.
type propagationPolicyLister struct {
	indexer cache.Indexer
}

// NewPropagationPolicyLister returns a new PropagationPolicyLister.
func NewPropagationPolicyLister(indexer cache.Indexer) PropagationPolicyLister {
	return &propagationPolicyLister{indexer: indexer}
}

// List lists all PropagationPolicies in the indexer.
func (s *propagationPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.PropagationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PropagationPolicy))
	})
	return ret, err
}

// PropagationPolicies returns an object that can list and get PropagationPolicies.
func (s *propagationPolicyLister) PropagationPolicies(namespace string) PropagationPolicyNamespaceLister {
	return propagationPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PropagationPolicyNamespaceLister helps list and get PropagationPolicies.
// All objects returned here must be treated as read-only.
type PropagationPolicyNamespaceLister interface {
	// List lists all PropagationPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PropagationPolicy, err error)
	// Get retrieves the PropagationPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PropagationPolicy, error)
	PropagationPolicyNamespaceListerExpansion
}

// propagationPolicyNamespaceLister implements the PropagationPolicyNamespaceLister
// interface.
type propagationPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PropagationPolicies in the indexer for a given namespace.
func (s propagationPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PropagationPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PropagationPolicy))
	})
	return ret, err
}

// Get retrieves the PropagationPolicy from the indexer for a given namespace and name.
func (s propagationPolicyNamespaceLister) Get(name string) (*v1alpha1.PropagationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("propagationpolicy"), name)
	}
	return obj.(*v1alpha1.PropagationPolicy), nil
}
