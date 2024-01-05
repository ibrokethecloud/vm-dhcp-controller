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
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type VirtualMachineNetworkConfigHandler func(string, *v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error)

type VirtualMachineNetworkConfigController interface {
	generic.ControllerMeta
	VirtualMachineNetworkConfigClient

	OnChange(ctx context.Context, name string, sync VirtualMachineNetworkConfigHandler)
	OnRemove(ctx context.Context, name string, sync VirtualMachineNetworkConfigHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() VirtualMachineNetworkConfigCache
}

type VirtualMachineNetworkConfigClient interface {
	Create(*v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error)
	Update(*v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error)
	UpdateStatus(*v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.VirtualMachineNetworkConfig, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha1.VirtualMachineNetworkConfigList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineNetworkConfig, err error)
}

type VirtualMachineNetworkConfigCache interface {
	Get(namespace, name string) (*v1alpha1.VirtualMachineNetworkConfig, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha1.VirtualMachineNetworkConfig, error)

	AddIndexer(indexName string, indexer VirtualMachineNetworkConfigIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha1.VirtualMachineNetworkConfig, error)
}

type VirtualMachineNetworkConfigIndexer func(obj *v1alpha1.VirtualMachineNetworkConfig) ([]string, error)

type virtualMachineNetworkConfigController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewVirtualMachineNetworkConfigController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VirtualMachineNetworkConfigController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &virtualMachineNetworkConfigController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromVirtualMachineNetworkConfigHandlerToHandler(sync VirtualMachineNetworkConfigHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha1.VirtualMachineNetworkConfig
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha1.VirtualMachineNetworkConfig))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *virtualMachineNetworkConfigController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha1.VirtualMachineNetworkConfig))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVirtualMachineNetworkConfigDeepCopyOnChange(client VirtualMachineNetworkConfigClient, obj *v1alpha1.VirtualMachineNetworkConfig, handler func(obj *v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error)) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *virtualMachineNetworkConfigController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *virtualMachineNetworkConfigController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *virtualMachineNetworkConfigController) OnChange(ctx context.Context, name string, sync VirtualMachineNetworkConfigHandler) {
	c.AddGenericHandler(ctx, name, FromVirtualMachineNetworkConfigHandlerToHandler(sync))
}

func (c *virtualMachineNetworkConfigController) OnRemove(ctx context.Context, name string, sync VirtualMachineNetworkConfigHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVirtualMachineNetworkConfigHandlerToHandler(sync)))
}

func (c *virtualMachineNetworkConfigController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *virtualMachineNetworkConfigController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *virtualMachineNetworkConfigController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *virtualMachineNetworkConfigController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *virtualMachineNetworkConfigController) Cache() VirtualMachineNetworkConfigCache {
	return &virtualMachineNetworkConfigCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *virtualMachineNetworkConfigController) Create(obj *v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	result := &v1alpha1.VirtualMachineNetworkConfig{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *virtualMachineNetworkConfigController) Update(obj *v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	result := &v1alpha1.VirtualMachineNetworkConfig{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *virtualMachineNetworkConfigController) UpdateStatus(obj *v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	result := &v1alpha1.VirtualMachineNetworkConfig{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *virtualMachineNetworkConfigController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *virtualMachineNetworkConfigController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	result := &v1alpha1.VirtualMachineNetworkConfig{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *virtualMachineNetworkConfigController) List(namespace string, opts metav1.ListOptions) (*v1alpha1.VirtualMachineNetworkConfigList, error) {
	result := &v1alpha1.VirtualMachineNetworkConfigList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *virtualMachineNetworkConfigController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *virtualMachineNetworkConfigController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	result := &v1alpha1.VirtualMachineNetworkConfig{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type virtualMachineNetworkConfigCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *virtualMachineNetworkConfigCache) Get(namespace, name string) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1alpha1.VirtualMachineNetworkConfig), nil
}

func (c *virtualMachineNetworkConfigCache) List(namespace string, selector labels.Selector) (ret []*v1alpha1.VirtualMachineNetworkConfig, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineNetworkConfig))
	})

	return ret, err
}

func (c *virtualMachineNetworkConfigCache) AddIndexer(indexName string, indexer VirtualMachineNetworkConfigIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha1.VirtualMachineNetworkConfig))
		},
	}))
}

func (c *virtualMachineNetworkConfigCache) GetByIndex(indexName, key string) (result []*v1alpha1.VirtualMachineNetworkConfig, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1alpha1.VirtualMachineNetworkConfig, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha1.VirtualMachineNetworkConfig))
	}
	return result, nil
}

type VirtualMachineNetworkConfigStatusHandler func(obj *v1alpha1.VirtualMachineNetworkConfig, status v1alpha1.VirtualMachineNetworkConfigStatus) (v1alpha1.VirtualMachineNetworkConfigStatus, error)

type VirtualMachineNetworkConfigGeneratingHandler func(obj *v1alpha1.VirtualMachineNetworkConfig, status v1alpha1.VirtualMachineNetworkConfigStatus) ([]runtime.Object, v1alpha1.VirtualMachineNetworkConfigStatus, error)

func RegisterVirtualMachineNetworkConfigStatusHandler(ctx context.Context, controller VirtualMachineNetworkConfigController, condition condition.Cond, name string, handler VirtualMachineNetworkConfigStatusHandler) {
	statusHandler := &virtualMachineNetworkConfigStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromVirtualMachineNetworkConfigHandlerToHandler(statusHandler.sync))
}

func RegisterVirtualMachineNetworkConfigGeneratingHandler(ctx context.Context, controller VirtualMachineNetworkConfigController, apply apply.Apply,
	condition condition.Cond, name string, handler VirtualMachineNetworkConfigGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &virtualMachineNetworkConfigGeneratingHandler{
		VirtualMachineNetworkConfigGeneratingHandler: handler,
		apply: apply,
		name:  name,
		gvk:   controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterVirtualMachineNetworkConfigStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type virtualMachineNetworkConfigStatusHandler struct {
	client    VirtualMachineNetworkConfigClient
	condition condition.Cond
	handler   VirtualMachineNetworkConfigStatusHandler
}

func (a *virtualMachineNetworkConfigStatusHandler) sync(key string, obj *v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type virtualMachineNetworkConfigGeneratingHandler struct {
	VirtualMachineNetworkConfigGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *virtualMachineNetworkConfigGeneratingHandler) Remove(key string, obj *v1alpha1.VirtualMachineNetworkConfig) (*v1alpha1.VirtualMachineNetworkConfig, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1alpha1.VirtualMachineNetworkConfig{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *virtualMachineNetworkConfigGeneratingHandler) Handle(obj *v1alpha1.VirtualMachineNetworkConfig, status v1alpha1.VirtualMachineNetworkConfigStatus) (v1alpha1.VirtualMachineNetworkConfigStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.VirtualMachineNetworkConfigGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
