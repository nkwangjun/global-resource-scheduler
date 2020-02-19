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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	scheme "k8s.io/metrics/pkg/client/clientset/versioned/scheme"
)

// PodMetricsesGetter has a method to return a PodMetricsInterface.
// A group's client should implement this interface.
type PodMetricsesGetter interface {
	PodMetricses(namespace string) PodMetricsInterface
	PodMetricsesWithMultiTenancy(namespace string, tenant string) PodMetricsInterface
}

// PodMetricsInterface has methods to work with PodMetrics resources.
type PodMetricsInterface interface {
	Get(name string, options v1.GetOptions) (*v1beta1.PodMetrics, error)
	List(opts v1.ListOptions) (*v1beta1.PodMetricsList, error)
	Watch(opts v1.ListOptions) watch.AggregatedWatchInterface
	PodMetricsExpansion
}

// podMetricses implements PodMetricsInterface
type podMetricses struct {
	client  rest.Interface
	clients []rest.Interface
	ns      string
	te      string
}

// newPodMetricses returns a PodMetricses
func newPodMetricses(c *MetricsV1beta1Client, namespace string) *podMetricses {
	return newPodMetricsesWithMultiTenancy(c, namespace, "default")
}

func newPodMetricsesWithMultiTenancy(c *MetricsV1beta1Client, namespace string, tenant string) *podMetricses {
	return &podMetricses{
		client:  c.RESTClient(),
		clients: c.RESTClients(),
		ns:      namespace,
		te:      tenant,
	}
}

// Get takes name of the podMetrics, and returns the corresponding podMetrics object, and an error if there is any.
func (c *podMetricses) Get(name string, options v1.GetOptions) (result *v1beta1.PodMetrics, err error) {
	result = &v1beta1.PodMetrics{}
	err = c.client.Get().
		Tenant(c.te).
		Namespace(c.ns).
		Resource("pods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)

	return
}

// List takes label and field selectors, and returns the list of PodMetricses that match those selectors.
func (c *podMetricses) List(opts v1.ListOptions) (result *v1beta1.PodMetricsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.PodMetricsList{}
	err = c.client.Get().
		Tenant(c.te).Namespace(c.ns).
		Resource("pods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)

	return
}

// Watch returns a watch.Interface that watches the requested podMetricses.
func (c *podMetricses) Watch(opts v1.ListOptions) watch.AggregatedWatchInterface {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	aggWatch := watch.NewAggregatedWatcher()
	for _, client := range c.clients {
		watcher, err := client.Get().
			Tenant(c.te).
			Namespace(c.ns).
			Resource("pods").
			VersionedParams(&opts, scheme.ParameterCodec).
			Timeout(timeout).
			Watch()
		aggWatch.AddWatchInterface(watcher, err)
	}
	return aggWatch
}
