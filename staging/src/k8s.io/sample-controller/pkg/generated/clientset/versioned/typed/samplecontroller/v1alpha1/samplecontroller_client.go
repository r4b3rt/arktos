/*
Copyright The Kubernetes Authors.
Copyright 2020 Authors of Alkaid - file modified.

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

package v1alpha1

import (
	"time"

	rand "k8s.io/apimachinery/pkg/util/rand"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
	"k8s.io/sample-controller/pkg/generated/clientset/versioned/scheme"
)

type SamplecontrollerV1alpha1Interface interface {
	RESTClient() rest.Interface
	RESTClients() []rest.Interface
	FoosGetter
}

// SamplecontrollerV1alpha1Client is used to interact with features provided by the samplecontroller.k8s.io group.
type SamplecontrollerV1alpha1Client struct {
	restClients []rest.Interface
}

func (c *SamplecontrollerV1alpha1Client) Foos(namespace string) FooInterface {
	return newFoosWithMultiTenancy(c, namespace, "default")
}

func (c *SamplecontrollerV1alpha1Client) FoosWithMultiTenancy(namespace string, tenant string) FooInterface {
	return newFoosWithMultiTenancy(c, namespace, tenant)
}

// NewForConfig creates a new SamplecontrollerV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*SamplecontrollerV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	clients := []rest.Interface{client}
	return &SamplecontrollerV1alpha1Client{clients}, nil
}

// NewForConfigOrDie creates a new SamplecontrollerV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SamplecontrollerV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SamplecontrollerV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SamplecontrollerV1alpha1Client {
	clients := []rest.Interface{c}
	return &SamplecontrollerV1alpha1Client{clients}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SamplecontrollerV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}

	max := len(c.restClients)
	if max == 0 {
		return nil
	}
	if max == 1 {
		return c.restClients[0]
	}

	rand.Seed(time.Now().UnixNano())
	ran := rand.IntnRange(0, max-1)
	return c.restClients[ran]
}

// RESTClients returns all RESTClient that are used to communicate
// with all API servers by this client implementation.
func (c *SamplecontrollerV1alpha1Client) RESTClients() []rest.Interface {
	if c == nil {
		return nil
	}

	return c.restClients
}
