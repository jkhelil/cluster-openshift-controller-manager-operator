// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/cluster-openshift-controller-manager-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KubeApiserverOperatorConfigs returns a KubeApiserverOperatorConfigInformer.
	KubeApiserverOperatorConfigs() KubeApiserverOperatorConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// KubeApiserverOperatorConfigs returns a KubeApiserverOperatorConfigInformer.
func (v *version) KubeApiserverOperatorConfigs() KubeApiserverOperatorConfigInformer {
	return &kubeApiserverOperatorConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
