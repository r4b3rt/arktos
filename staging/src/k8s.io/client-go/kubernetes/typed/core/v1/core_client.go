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

package v1

import (
	"time"

	v1 "k8s.io/api/core/v1"
	rand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type CoreV1Interface interface {
	RESTClient() rest.Interface
	RESTClients() []rest.Interface
	ActionsGetter
	ComponentStatusesGetter
	ConfigMapsGetter
	ControllerInstancesGetter
	EndpointsGetter
	EventsGetter
	LimitRangesGetter
	NamespacesGetter
	NodesGetter
	PersistentVolumesGetter
	PersistentVolumeClaimsGetter
	PodsGetter
	PodTemplatesGetter
	ReplicationControllersGetter
	ResourceQuotasGetter
	SecretsGetter
	ServicesGetter
	ServiceAccountsGetter
	TenantsGetter
}

// CoreV1Client is used to interact with features provided by the  group.
type CoreV1Client struct {
	restClients []rest.Interface
}

func (c *CoreV1Client) Actions(namespace string) ActionInterface {
	return newActionsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) ActionsWithMultiTenancy(namespace string, tenant string) ActionInterface {
	return newActionsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) ComponentStatuses() ComponentStatusInterface {
	return newComponentStatuses(c)
}

func (c *CoreV1Client) ConfigMaps(namespace string) ConfigMapInterface {
	return newConfigMapsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) ConfigMapsWithMultiTenancy(namespace string, tenant string) ConfigMapInterface {
	return newConfigMapsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) ControllerInstances() ControllerInstanceInterface {
	return newControllerInstances(c)
}

func (c *CoreV1Client) Endpoints(namespace string) EndpointsInterface {
	return newEndpointsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) EndpointsWithMultiTenancy(namespace string, tenant string) EndpointsInterface {
	return newEndpointsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) Events(namespace string) EventInterface {
	return newEventsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) EventsWithMultiTenancy(namespace string, tenant string) EventInterface {
	return newEventsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) LimitRanges(namespace string) LimitRangeInterface {
	return newLimitRangesWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) LimitRangesWithMultiTenancy(namespace string, tenant string) LimitRangeInterface {
	return newLimitRangesWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) Namespaces() NamespaceInterface {
	return newNamespacesWithMultiTenancy(c, "default")
}

func (c *CoreV1Client) NamespacesWithMultiTenancy(tenant string) NamespaceInterface {
	return newNamespacesWithMultiTenancy(c, tenant)
}

func (c *CoreV1Client) Nodes() NodeInterface {
	return newNodes(c)
}

func (c *CoreV1Client) PersistentVolumes() PersistentVolumeInterface {
	return newPersistentVolumesWithMultiTenancy(c, "default")
}

func (c *CoreV1Client) PersistentVolumesWithMultiTenancy(tenant string) PersistentVolumeInterface {
	return newPersistentVolumesWithMultiTenancy(c, tenant)
}

func (c *CoreV1Client) PersistentVolumeClaims(namespace string) PersistentVolumeClaimInterface {
	return newPersistentVolumeClaimsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) PersistentVolumeClaimsWithMultiTenancy(namespace string, tenant string) PersistentVolumeClaimInterface {
	return newPersistentVolumeClaimsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) Pods(namespace string) PodInterface {
	return newPodsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) PodsWithMultiTenancy(namespace string, tenant string) PodInterface {
	return newPodsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) PodTemplates(namespace string) PodTemplateInterface {
	return newPodTemplatesWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) PodTemplatesWithMultiTenancy(namespace string, tenant string) PodTemplateInterface {
	return newPodTemplatesWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) ReplicationControllers(namespace string) ReplicationControllerInterface {
	return newReplicationControllersWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) ReplicationControllersWithMultiTenancy(namespace string, tenant string) ReplicationControllerInterface {
	return newReplicationControllersWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) ResourceQuotas(namespace string) ResourceQuotaInterface {
	return newResourceQuotasWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) ResourceQuotasWithMultiTenancy(namespace string, tenant string) ResourceQuotaInterface {
	return newResourceQuotasWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) Secrets(namespace string) SecretInterface {
	return newSecretsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) SecretsWithMultiTenancy(namespace string, tenant string) SecretInterface {
	return newSecretsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) Services(namespace string) ServiceInterface {
	return newServicesWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) ServicesWithMultiTenancy(namespace string, tenant string) ServiceInterface {
	return newServicesWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) ServiceAccounts(namespace string) ServiceAccountInterface {
	return newServiceAccountsWithMultiTenancy(c, namespace, "default")
}

func (c *CoreV1Client) ServiceAccountsWithMultiTenancy(namespace string, tenant string) ServiceAccountInterface {
	return newServiceAccountsWithMultiTenancy(c, namespace, tenant)
}

func (c *CoreV1Client) Tenants() TenantInterface {
	return newTenants(c)
}

// NewForConfig creates a new CoreV1Client for the given config.
func NewForConfig(c *rest.Config) (*CoreV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	clients := []rest.Interface{client}
	return &CoreV1Client{clients}, nil
}

// NewForConfigOrDie creates a new CoreV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoreV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoreV1Client for the given RESTClient.
func New(c rest.Interface) *CoreV1Client {
	clients := []rest.Interface{c}
	return &CoreV1Client{clients}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/api"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CoreV1Client) RESTClient() rest.Interface {
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
func (c *CoreV1Client) RESTClients() []rest.Interface {
	if c == nil {
		return nil
	}

	return c.restClients
}
