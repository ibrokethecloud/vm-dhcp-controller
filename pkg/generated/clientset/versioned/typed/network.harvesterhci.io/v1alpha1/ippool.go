/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/harvester/vm-dhcp-controller/pkg/apis/network.harvesterhci.io/v1alpha1"
	scheme "github.com/harvester/vm-dhcp-controller/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IPPoolsGetter has a method to return a IPPoolInterface.
// A group's client should implement this interface.
type IPPoolsGetter interface {
	IPPools(namespace string) IPPoolInterface
}

// IPPoolInterface has methods to work with IPPool resources.
type IPPoolInterface interface {
	Create(ctx context.Context, iPPool *v1alpha1.IPPool, opts v1.CreateOptions) (*v1alpha1.IPPool, error)
	Update(ctx context.Context, iPPool *v1alpha1.IPPool, opts v1.UpdateOptions) (*v1alpha1.IPPool, error)
	UpdateStatus(ctx context.Context, iPPool *v1alpha1.IPPool, opts v1.UpdateOptions) (*v1alpha1.IPPool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.IPPool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.IPPoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IPPool, err error)
	IPPoolExpansion
}

// iPPools implements IPPoolInterface
type iPPools struct {
	client rest.Interface
	ns     string
}

// newIPPools returns a IPPools
func newIPPools(c *NetworkV1alpha1Client, namespace string) *iPPools {
	return &iPPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iPPool, and returns the corresponding iPPool object, and an error if there is any.
func (c *iPPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IPPool, err error) {
	result = &v1alpha1.IPPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ippools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IPPools that match those selectors.
func (c *iPPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IPPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IPPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iPPools.
func (c *iPPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iPPool and creates it.  Returns the server's representation of the iPPool, and an error, if there is any.
func (c *iPPools) Create(ctx context.Context, iPPool *v1alpha1.IPPool, opts v1.CreateOptions) (result *v1alpha1.IPPool, err error) {
	result = &v1alpha1.IPPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ippools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iPPool and updates it. Returns the server's representation of the iPPool, and an error, if there is any.
func (c *iPPools) Update(ctx context.Context, iPPool *v1alpha1.IPPool, opts v1.UpdateOptions) (result *v1alpha1.IPPool, err error) {
	result = &v1alpha1.IPPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ippools").
		Name(iPPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPPool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *iPPools) UpdateStatus(ctx context.Context, iPPool *v1alpha1.IPPool, opts v1.UpdateOptions) (result *v1alpha1.IPPool, err error) {
	result = &v1alpha1.IPPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ippools").
		Name(iPPool.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iPPool and deletes it. Returns an error if one occurs.
func (c *iPPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ippools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iPPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ippools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iPPool.
func (c *iPPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IPPool, err error) {
	result = &v1alpha1.IPPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ippools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
