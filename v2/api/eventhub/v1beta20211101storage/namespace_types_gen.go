// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventhub.azure.com,resources=namespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventhub.azure.com,resources={namespaces/status,namespaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20211101.Namespace
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces
type Namespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Spec    `json:"spec,omitempty"`
	Status            EHNamespace_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Namespace{}

// GetConditions returns the conditions of the resource
func (namespace *Namespace) GetConditions() conditions.Conditions {
	return namespace.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (namespace *Namespace) SetConditions(conditions conditions.Conditions) {
	namespace.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Namespace{}

// AzureName returns the Azure name of the resource
func (namespace *Namespace) AzureName() string {
	return namespace.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (namespace Namespace) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (namespace *Namespace) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (namespace *Namespace) GetSpec() genruntime.ConvertibleSpec {
	return &namespace.Spec
}

// GetStatus returns the status of this resource
func (namespace *Namespace) GetStatus() genruntime.ConvertibleStatus {
	return &namespace.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces"
func (namespace *Namespace) GetType() string {
	return "Microsoft.EventHub/namespaces"
}

// NewEmptyStatus returns a new empty (blank) status
func (namespace *Namespace) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &EHNamespace_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (namespace *Namespace) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(namespace.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  namespace.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (namespace *Namespace) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*EHNamespace_Status); ok {
		namespace.Status = *st
		return nil
	}

	// Convert status to required version
	var st EHNamespace_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	namespace.Status = st
	return nil
}

// Hub marks that this Namespace is the hub type for conversion
func (namespace *Namespace) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (namespace *Namespace) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: namespace.Spec.OriginalVersion,
		Kind:    "Namespace",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20211101.Namespace
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces
type NamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Namespace `json:"items"`
}

// Storage version of v1beta20211101.EHNamespace_Status
type EHNamespace_Status struct {
	AlternateName              *string                                                `json:"alternateName,omitempty"`
	ClusterArmId               *string                                                `json:"clusterArmId,omitempty"`
	Conditions                 []conditions.Condition                                 `json:"conditions,omitempty"`
	CreatedAt                  *string                                                `json:"createdAt,omitempty"`
	DisableLocalAuth           *bool                                                  `json:"disableLocalAuth,omitempty"`
	Encryption                 *Encryption_Status                                     `json:"encryption,omitempty"`
	Id                         *string                                                `json:"id,omitempty"`
	Identity                   *Identity_Status                                       `json:"identity,omitempty"`
	IsAutoInflateEnabled       *bool                                                  `json:"isAutoInflateEnabled,omitempty"`
	KafkaEnabled               *bool                                                  `json:"kafkaEnabled,omitempty"`
	Location                   *string                                                `json:"location,omitempty"`
	MaximumThroughputUnits     *int                                                   `json:"maximumThroughputUnits,omitempty"`
	MetricId                   *string                                                `json:"metricId,omitempty"`
	Name                       *string                                                `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                `json:"provisioningState,omitempty"`
	ServiceBusEndpoint         *string                                                `json:"serviceBusEndpoint,omitempty"`
	Sku                        *Sku_Status                                            `json:"sku,omitempty"`
	Status                     *string                                                `json:"status,omitempty"`
	SystemData                 *SystemData_Status                                     `json:"systemData,omitempty"`
	Tags                       map[string]string                                      `json:"tags,omitempty"`
	Type                       *string                                                `json:"type,omitempty"`
	UpdatedAt                  *string                                                `json:"updatedAt,omitempty"`
	ZoneRedundant              *bool                                                  `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleStatus = &EHNamespace_Status{}

// ConvertStatusFrom populates our EHNamespace_Status from the provided source
func (namespace *EHNamespace_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == namespace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(namespace)
}

// ConvertStatusTo populates the provided destination from our EHNamespace_Status
func (namespace *EHNamespace_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == namespace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(namespace)
}

// Storage version of v1beta20211101.Namespaces_Spec
type Namespaces_Spec struct {
	AlternateName *string `json:"alternateName,omitempty"`

	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=6
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// ClusterArmReference: Cluster ARM ID of the Namespace.
	ClusterArmReference    *genruntime.ResourceReference `armReference:"ClusterArmId" json:"clusterArmReference,omitempty"`
	DisableLocalAuth       *bool                         `json:"disableLocalAuth,omitempty"`
	Encryption             *Encryption                   `json:"encryption,omitempty"`
	Identity               *Identity                     `json:"identity,omitempty"`
	IsAutoInflateEnabled   *bool                         `json:"isAutoInflateEnabled,omitempty"`
	KafkaEnabled           *bool                         `json:"kafkaEnabled,omitempty"`
	Location               *string                       `json:"location,omitempty"`
	MaximumThroughputUnits *int                          `json:"maximumThroughputUnits,omitempty"`
	OriginalVersion        string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                      *genruntime.KnownResourceReference                      `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PrivateEndpointConnections []Namespaces_Spec_Properties_PrivateEndpointConnections `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                  `json:"$propertyBag,omitempty"`
	Sku                        *Sku                                                    `json:"sku,omitempty"`
	Tags                       map[string]string                                       `json:"tags,omitempty"`
	ZoneRedundant              *bool                                                   `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_Spec{}

// ConvertSpecFrom populates our Namespaces_Spec from the provided source
func (namespaces *Namespaces_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == namespaces {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(namespaces)
}

// ConvertSpecTo populates the provided destination from our Namespaces_Spec
func (namespaces *Namespaces_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == namespaces {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(namespaces)
}

// Storage version of v1beta20211101.Encryption
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Encryption
type Encryption struct {
	KeySource                       *string                `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties   `json:"keyVaultProperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                  `json:"requireInfrastructureEncryption,omitempty"`
}

// Storage version of v1beta20211101.Encryption_Status
type Encryption_Status struct {
	KeySource                       *string                     `json:"keySource,omitempty"`
	KeyVaultProperties              []KeyVaultProperties_Status `json:"keyVaultProperties,omitempty"`
	PropertyBag                     genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequireInfrastructureEncryption *bool                       `json:"requireInfrastructureEncryption,omitempty"`
}

// Storage version of v1beta20211101.Identity
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Identity
type Identity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20211101.Identity_Status
type Identity_Status struct {
	PrincipalId            *string                                `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	TenantId               *string                                `json:"tenantId,omitempty"`
	Type                   *string                                `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentity_Status `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20211101.Namespaces_Spec_Properties_PrivateEndpointConnections
type Namespaces_Spec_Properties_PrivateEndpointConnections struct {
	PrivateEndpoint *PrivateEndpoint       `json:"privateEndpoint,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.PrivateEndpointConnection_Status_SubResourceEmbedded
type PrivateEndpointConnection_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_Status     `json:"systemData,omitempty"`
}

// Storage version of v1beta20211101.Sku
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Sku
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20211101.Sku_Status
type Sku_Status struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20211101.SystemData_Status
type SystemData_Status struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.KeyVaultProperties
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/KeyVaultProperties
type KeyVaultProperties struct {
	Identity    *UserAssignedIdentityProperties `json:"identity,omitempty"`
	KeyName     *string                         `json:"keyName,omitempty"`
	KeyVaultUri *string                         `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                         `json:"keyVersion,omitempty"`
	PropertyBag genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.KeyVaultProperties_Status
type KeyVaultProperties_Status struct {
	Identity    *UserAssignedIdentityProperties_Status `json:"identity,omitempty"`
	KeyName     *string                                `json:"keyName,omitempty"`
	KeyVaultUri *string                                `json:"keyVaultUri,omitempty"`
	KeyVersion  *string                                `json:"keyVersion,omitempty"`
	PropertyBag genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.PrivateEndpoint
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpoint
type PrivateEndpoint struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: The ARM identifier for Private Endpoint.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20211101.UserAssignedIdentity_Status
type UserAssignedIdentity_Status struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.UserAssignedIdentityProperties
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/UserAssignedIdentityProperties
type UserAssignedIdentityProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// UserAssignedIdentityReference: ARM ID of user Identity selected for encryption
	UserAssignedIdentityReference *genruntime.ResourceReference `armReference:"UserAssignedIdentity" json:"userAssignedIdentityReference,omitempty"`
}

// Storage version of v1beta20211101.UserAssignedIdentityProperties_Status
type UserAssignedIdentityProperties_Status struct {
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UserAssignedIdentity *string                `json:"userAssignedIdentity,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Namespace{}, &NamespaceList{})
}
