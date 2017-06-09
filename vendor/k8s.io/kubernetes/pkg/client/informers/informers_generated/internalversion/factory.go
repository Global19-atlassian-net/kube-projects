/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	admissionregistration "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/admissionregistration"
	apps "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/apps"
	autoscaling "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/autoscaling"
	batch "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/batch"
	certificates "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/certificates"
	core "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/core"
	extensions "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/extensions"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	networking "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/networking"
	policy "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/policy"
	rbac "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/rbac"
	settings "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/settings"
	storage "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/storage"
	reflect "reflect"
	sync "sync"
	time "time"
)

type sharedInformerFactory struct {
	client        internalclientset.Interface
	lock          sync.Mutex
	defaultResync time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory
func NewSharedInformerFactory(client internalclientset.Interface, defaultResync time.Duration) SharedInformerFactory {
	return &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
	}
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}
	informer = newFunc(f.client, f.defaultResync)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Admissionregistration() admissionregistration.Interface
	Apps() apps.Interface
	Autoscaling() autoscaling.Interface
	Batch() batch.Interface
	Certificates() certificates.Interface
	Core() core.Interface
	Extensions() extensions.Interface
	Networking() networking.Interface
	Policy() policy.Interface
	Rbac() rbac.Interface
	Settings() settings.Interface
	Storage() storage.Interface
}

func (f *sharedInformerFactory) Admissionregistration() admissionregistration.Interface {
	return admissionregistration.New(f)
}

func (f *sharedInformerFactory) Apps() apps.Interface {
	return apps.New(f)
}

func (f *sharedInformerFactory) Autoscaling() autoscaling.Interface {
	return autoscaling.New(f)
}

func (f *sharedInformerFactory) Batch() batch.Interface {
	return batch.New(f)
}

func (f *sharedInformerFactory) Certificates() certificates.Interface {
	return certificates.New(f)
}

func (f *sharedInformerFactory) Core() core.Interface {
	return core.New(f)
}

func (f *sharedInformerFactory) Extensions() extensions.Interface {
	return extensions.New(f)
}

func (f *sharedInformerFactory) Networking() networking.Interface {
	return networking.New(f)
}

func (f *sharedInformerFactory) Policy() policy.Interface {
	return policy.New(f)
}

func (f *sharedInformerFactory) Rbac() rbac.Interface {
	return rbac.New(f)
}

func (f *sharedInformerFactory) Settings() settings.Interface {
	return settings.New(f)
}

func (f *sharedInformerFactory) Storage() storage.Interface {
	return storage.New(f)
}
