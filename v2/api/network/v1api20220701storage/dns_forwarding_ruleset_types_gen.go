// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=dnsforwardingrulesets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnsforwardingrulesets/status,dnsforwardingrulesets/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.DnsForwardingRuleset
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}
type DnsForwardingRuleset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsForwardingRuleset_Spec   `json:"spec,omitempty"`
	Status            DnsForwardingRuleset_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsForwardingRuleset{}

// GetConditions returns the conditions of the resource
func (ruleset *DnsForwardingRuleset) GetConditions() conditions.Conditions {
	return ruleset.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (ruleset *DnsForwardingRuleset) SetConditions(conditions conditions.Conditions) {
	ruleset.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DnsForwardingRuleset{}

// AzureName returns the Azure name of the resource
func (ruleset *DnsForwardingRuleset) AzureName() string {
	return ruleset.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (ruleset DnsForwardingRuleset) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (ruleset *DnsForwardingRuleset) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (ruleset *DnsForwardingRuleset) GetSpec() genruntime.ConvertibleSpec {
	return &ruleset.Spec
}

// GetStatus returns the status of this resource
func (ruleset *DnsForwardingRuleset) GetStatus() genruntime.ConvertibleStatus {
	return &ruleset.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsForwardingRulesets"
func (ruleset *DnsForwardingRuleset) GetType() string {
	return "Microsoft.Network/dnsForwardingRulesets"
}

// NewEmptyStatus returns a new empty (blank) status
func (ruleset *DnsForwardingRuleset) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsForwardingRuleset_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (ruleset *DnsForwardingRuleset) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(ruleset.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  ruleset.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (ruleset *DnsForwardingRuleset) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsForwardingRuleset_STATUS); ok {
		ruleset.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsForwardingRuleset_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	ruleset.Status = st
	return nil
}

// Hub marks that this DnsForwardingRuleset is the hub type for conversion
func (ruleset *DnsForwardingRuleset) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (ruleset *DnsForwardingRuleset) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: ruleset.Spec.OriginalVersion,
		Kind:    "DnsForwardingRuleset",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.DnsForwardingRuleset
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}
type DnsForwardingRulesetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsForwardingRuleset `json:"items"`
}

// Storage version of v1api20220701.DnsForwardingRuleset_Spec
type DnsForwardingRuleset_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                    string                   `json:"azureName,omitempty"`
	DnsResolverOutboundEndpoints []DnsresolverSubResource `json:"dnsResolverOutboundEndpoints,omitempty"`
	Location                     *string                  `json:"location,omitempty"`
	OriginalVersion              string                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsForwardingRuleset_Spec{}

// ConvertSpecFrom populates our DnsForwardingRuleset_Spec from the provided source
func (ruleset *DnsForwardingRuleset_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == ruleset {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(ruleset)
}

// ConvertSpecTo populates the provided destination from our DnsForwardingRuleset_Spec
func (ruleset *DnsForwardingRuleset_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == ruleset {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(ruleset)
}

// Storage version of v1api20220701.DnsForwardingRuleset_STATUS
// Describes a DNS forwarding ruleset.
type DnsForwardingRuleset_STATUS struct {
	Conditions                   []conditions.Condition          `json:"conditions,omitempty"`
	DnsResolverOutboundEndpoints []DnsresolverSubResource_STATUS `json:"dnsResolverOutboundEndpoints,omitempty"`
	Etag                         *string                         `json:"etag,omitempty"`
	Id                           *string                         `json:"id,omitempty"`
	Location                     *string                         `json:"location,omitempty"`
	Name                         *string                         `json:"name,omitempty"`
	PropertyBag                  genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                         `json:"provisioningState,omitempty"`
	ResourceGuid                 *string                         `json:"resourceGuid,omitempty"`
	SystemData                   *SystemData_STATUS              `json:"systemData,omitempty"`
	Tags                         map[string]string               `json:"tags,omitempty"`
	Type                         *string                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsForwardingRuleset_STATUS{}

// ConvertStatusFrom populates our DnsForwardingRuleset_STATUS from the provided source
func (ruleset *DnsForwardingRuleset_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == ruleset {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(ruleset)
}

// ConvertStatusTo populates the provided destination from our DnsForwardingRuleset_STATUS
func (ruleset *DnsForwardingRuleset_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == ruleset {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(ruleset)
}

// Storage version of v1api20220701.DnsresolverSubResource
// Reference to another ARM resource.
type DnsresolverSubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20220701.DnsresolverSubResource_STATUS
// Reference to another ARM resource.
type DnsresolverSubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DnsForwardingRuleset{}, &DnsForwardingRulesetList{})
}
