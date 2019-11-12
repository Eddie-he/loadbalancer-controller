/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/alerting/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlertingSubRuleLister helps list AlertingSubRules.
type AlertingSubRuleLister interface {
	// List lists all AlertingSubRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AlertingSubRule, err error)
	// Get retrieves the AlertingSubRule from the index for a given name.
	Get(name string) (*v1alpha1.AlertingSubRule, error)
	AlertingSubRuleListerExpansion
}

// alertingSubRuleLister implements the AlertingSubRuleLister interface.
type alertingSubRuleLister struct {
	indexer cache.Indexer
}

// NewAlertingSubRuleLister returns a new AlertingSubRuleLister.
func NewAlertingSubRuleLister(indexer cache.Indexer) AlertingSubRuleLister {
	return &alertingSubRuleLister{indexer: indexer}
}

// List lists all AlertingSubRules in the indexer.
func (s *alertingSubRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AlertingSubRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlertingSubRule))
	})
	return ret, err
}

// Get retrieves the AlertingSubRule from the index for a given name.
func (s *alertingSubRuleLister) Get(name string) (*v1alpha1.AlertingSubRule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alertingsubrule"), name)
	}
	return obj.(*v1alpha1.AlertingSubRule), nil
}