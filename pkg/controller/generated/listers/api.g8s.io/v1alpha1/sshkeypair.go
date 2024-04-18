/*
Copyright 2024 James Riley O'Donnell.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/jrodonnell/g8s/pkg/controller/apis/api.g8s.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SSHKeyPairLister helps list SSHKeyPairs.
// All objects returned here must be treated as read-only.
type SSHKeyPairLister interface {
	// List lists all SSHKeyPairs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SSHKeyPair, err error)
	// SSHKeyPairs returns an object that can list and get SSHKeyPairs.
	SSHKeyPairs(namespace string) SSHKeyPairNamespaceLister
	SSHKeyPairListerExpansion
}

// sSHKeyPairLister implements the SSHKeyPairLister interface.
type sSHKeyPairLister struct {
	indexer cache.Indexer
}

// NewSSHKeyPairLister returns a new SSHKeyPairLister.
func NewSSHKeyPairLister(indexer cache.Indexer) SSHKeyPairLister {
	return &sSHKeyPairLister{indexer: indexer}
}

// List lists all SSHKeyPairs in the indexer.
func (s *sSHKeyPairLister) List(selector labels.Selector) (ret []*v1alpha1.SSHKeyPair, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SSHKeyPair))
	})
	return ret, err
}

// SSHKeyPairs returns an object that can list and get SSHKeyPairs.
func (s *sSHKeyPairLister) SSHKeyPairs(namespace string) SSHKeyPairNamespaceLister {
	return sSHKeyPairNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SSHKeyPairNamespaceLister helps list and get SSHKeyPairs.
// All objects returned here must be treated as read-only.
type SSHKeyPairNamespaceLister interface {
	// List lists all SSHKeyPairs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SSHKeyPair, err error)
	// Get retrieves the SSHKeyPair from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SSHKeyPair, error)
	SSHKeyPairNamespaceListerExpansion
}

// sSHKeyPairNamespaceLister implements the SSHKeyPairNamespaceLister
// interface.
type sSHKeyPairNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SSHKeyPairs in the indexer for a given namespace.
func (s sSHKeyPairNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SSHKeyPair, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SSHKeyPair))
	})
	return ret, err
}

// Get retrieves the SSHKeyPair from the indexer for a given namespace and name.
func (s sSHKeyPairNamespaceLister) Get(name string) (*v1alpha1.SSHKeyPair, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sshkeypair"), name)
	}
	return obj.(*v1alpha1.SSHKeyPair), nil
}
