// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

// Describes a DNS forwarding ruleset.
type DnsForwardingRuleset_STATUS_ARM struct {
	// Etag: ETag of the DNS forwarding ruleset.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the DNS forwarding ruleset.
	Properties *DnsForwardingRulesetProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Represents the properties of a DNS forwarding ruleset.
type DnsForwardingRulesetProperties_STATUS_ARM struct {
	// DnsResolverOutboundEndpoints: The reference to the DNS resolver outbound endpoints that are used to route DNS queries
	// matching the forwarding rules in the ruleset to the target DNS servers.
	DnsResolverOutboundEndpoints []DnsresolverSubResource_STATUS_ARM `json:"dnsResolverOutboundEndpoints,omitempty"`

	// ProvisioningState: The current provisioning state of the DNS forwarding ruleset. This is a read-only property and any
	// attempt to set this value will be ignored.
	ProvisioningState *DnsresolverProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resourceGuid for the DNS forwarding ruleset.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Reference to another ARM resource.
type DnsresolverSubResource_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)
