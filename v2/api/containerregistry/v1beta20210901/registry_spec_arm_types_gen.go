// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210901

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Registry_Spec. Use v1api20210901.Registry_Spec instead
type Registry_Spec_ARM struct {
	Identity   *IdentityProperties_ARM `json:"identity,omitempty"`
	Location   *string                 `json:"location,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Properties *RegistryProperties_ARM `json:"properties,omitempty"`
	Sku        *Sku_ARM                `json:"sku,omitempty"`
	Tags       map[string]string       `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Registry_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-09-01"
func (registry Registry_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (registry *Registry_Spec_ARM) GetName() string {
	return registry.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerRegistry/registries"
func (registry *Registry_Spec_ARM) GetType() string {
	return "Microsoft.ContainerRegistry/registries"
}

// Deprecated version of IdentityProperties. Use v1api20210901.IdentityProperties instead
type IdentityProperties_ARM struct {
	PrincipalId            *string                                    `json:"principalId,omitempty"`
	TenantId               *string                                    `json:"tenantId,omitempty"`
	Type                   *IdentityProperties_Type                   `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of RegistryProperties. Use v1api20210901.RegistryProperties instead
type RegistryProperties_ARM struct {
	AdminUserEnabled         *bool                                        `json:"adminUserEnabled,omitempty"`
	DataEndpointEnabled      *bool                                        `json:"dataEndpointEnabled,omitempty"`
	Encryption               *EncryptionProperty_ARM                      `json:"encryption,omitempty"`
	NetworkRuleBypassOptions *RegistryProperties_NetworkRuleBypassOptions `json:"networkRuleBypassOptions,omitempty"`
	NetworkRuleSet           *NetworkRuleSet_ARM                          `json:"networkRuleSet,omitempty"`
	Policies                 *Policies_ARM                                `json:"policies,omitempty"`
	PublicNetworkAccess      *RegistryProperties_PublicNetworkAccess      `json:"publicNetworkAccess,omitempty"`
	ZoneRedundancy           *RegistryProperties_ZoneRedundancy           `json:"zoneRedundancy,omitempty"`
}

// Deprecated version of Sku. Use v1api20210901.Sku instead
type Sku_ARM struct {
	Name *Sku_Name `json:"name,omitempty"`
}

// Deprecated version of EncryptionProperty. Use v1api20210901.EncryptionProperty instead
type EncryptionProperty_ARM struct {
	KeyVaultProperties *KeyVaultProperties_ARM    `json:"keyVaultProperties,omitempty"`
	Status             *EncryptionProperty_Status `json:"status,omitempty"`
}

// Deprecated version of IdentityProperties_Type. Use v1api20210901.IdentityProperties_Type instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type IdentityProperties_Type string

const (
	IdentityProperties_Type_None                       = IdentityProperties_Type("None")
	IdentityProperties_Type_SystemAssigned             = IdentityProperties_Type("SystemAssigned")
	IdentityProperties_Type_SystemAssignedUserAssigned = IdentityProperties_Type("SystemAssigned, UserAssigned")
	IdentityProperties_Type_UserAssigned               = IdentityProperties_Type("UserAssigned")
)

// Deprecated version of NetworkRuleSet. Use v1api20210901.NetworkRuleSet instead
type NetworkRuleSet_ARM struct {
	DefaultAction *NetworkRuleSet_DefaultAction `json:"defaultAction,omitempty"`
	IpRules       []IPRule_ARM                  `json:"ipRules,omitempty"`
}

// Deprecated version of Policies. Use v1api20210901.Policies instead
type Policies_ARM struct {
	ExportPolicy     *ExportPolicy_ARM     `json:"exportPolicy,omitempty"`
	QuarantinePolicy *QuarantinePolicy_ARM `json:"quarantinePolicy,omitempty"`
	RetentionPolicy  *RetentionPolicy_ARM  `json:"retentionPolicy,omitempty"`
	TrustPolicy      *TrustPolicy_ARM      `json:"trustPolicy,omitempty"`
}

// Deprecated version of Sku_Name. Use v1api20210901.Sku_Name instead
// +kubebuilder:validation:Enum={"Basic","Classic","Premium","Standard"}
type Sku_Name string

const (
	Sku_Name_Basic    = Sku_Name("Basic")
	Sku_Name_Classic  = Sku_Name("Classic")
	Sku_Name_Premium  = Sku_Name("Premium")
	Sku_Name_Standard = Sku_Name("Standard")
)

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Deprecated version of ExportPolicy. Use v1api20210901.ExportPolicy instead
type ExportPolicy_ARM struct {
	Status *ExportPolicy_Status `json:"status,omitempty"`
}

// Deprecated version of IPRule. Use v1api20210901.IPRule instead
type IPRule_ARM struct {
	Action *IPRule_Action `json:"action,omitempty"`
	Value  *string        `json:"value,omitempty"`
}

// Deprecated version of KeyVaultProperties. Use v1api20210901.KeyVaultProperties instead
type KeyVaultProperties_ARM struct {
	Identity      *string `json:"identity,omitempty"`
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}

// Deprecated version of QuarantinePolicy. Use v1api20210901.QuarantinePolicy instead
type QuarantinePolicy_ARM struct {
	Status *QuarantinePolicy_Status `json:"status,omitempty"`
}

// Deprecated version of RetentionPolicy. Use v1api20210901.RetentionPolicy instead
type RetentionPolicy_ARM struct {
	Days   *int                    `json:"days,omitempty"`
	Status *RetentionPolicy_Status `json:"status,omitempty"`
}

// Deprecated version of TrustPolicy. Use v1api20210901.TrustPolicy instead
type TrustPolicy_ARM struct {
	Status *TrustPolicy_Status `json:"status,omitempty"`
	Type   *TrustPolicy_Type   `json:"type,omitempty"`
}