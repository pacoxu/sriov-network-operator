// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// SriovNetworks returns a SriovNetworkInformer.
	SriovNetworks() SriovNetworkInformer
	// SriovNetworkNodePolicies returns a SriovNetworkNodePolicyInformer.
	SriovNetworkNodePolicies() SriovNetworkNodePolicyInformer
	// SriovNetworkNodeStates returns a SriovNetworkNodeStateInformer.
	SriovNetworkNodeStates() SriovNetworkNodeStateInformer
	// SriovOperatorConfigs returns a SriovOperatorConfigInformer.
	SriovOperatorConfigs() SriovOperatorConfigInformer
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

// SriovNetworks returns a SriovNetworkInformer.
func (v *version) SriovNetworks() SriovNetworkInformer {
	return &sriovNetworkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SriovNetworkNodePolicies returns a SriovNetworkNodePolicyInformer.
func (v *version) SriovNetworkNodePolicies() SriovNetworkNodePolicyInformer {
	return &sriovNetworkNodePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SriovNetworkNodeStates returns a SriovNetworkNodeStateInformer.
func (v *version) SriovNetworkNodeStates() SriovNetworkNodeStateInformer {
	return &sriovNetworkNodeStateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SriovOperatorConfigs returns a SriovOperatorConfigInformer.
func (v *version) SriovOperatorConfigs() SriovOperatorConfigInformer {
	return &sriovOperatorConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
