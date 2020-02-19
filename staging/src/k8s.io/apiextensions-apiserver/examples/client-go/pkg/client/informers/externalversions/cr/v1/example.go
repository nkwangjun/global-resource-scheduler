/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	crv1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1"
	versioned "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/clientset/versioned"
	internalinterfaces "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/informers/externalversions/internalinterfaces"
	v1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/listers/cr/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ExampleInformer provides access to a shared informer and lister for
// Examples.
type ExampleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ExampleLister
}

type exampleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
	tenant           string
}

// NewExampleInformer constructs a new informer for Example type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExampleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExampleInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewExampleInformerWithMultiTenancy(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tenant string) cache.SharedIndexInformer {
	return NewFilteredExampleInformerWithMultiTenancy(client, namespace, resyncPeriod, indexers, nil, tenant)
}

// NewFilteredExampleInformer constructs a new informer for Example type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExampleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return NewFilteredExampleInformerWithMultiTenancy(client, namespace, resyncPeriod, indexers, tweakListOptions, "default")
}

func NewFilteredExampleInformerWithMultiTenancy(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc, tenant string) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1().ExamplesWithMultiTenancy(namespace, tenant).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) watch.AggregatedWatchInterface {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1().ExamplesWithMultiTenancy(namespace, tenant).Watch(options)
			},
		},
		&crv1.Example{},
		resyncPeriod,
		indexers,
	)
}

func (f *exampleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExampleInformerWithMultiTenancy(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions, f.tenant)
}

func (f *exampleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crv1.Example{}, f.defaultInformer)
}

func (f *exampleInformer) Lister() v1.ExampleLister {
	return v1.NewExampleLister(f.Informer().GetIndexer())
}
