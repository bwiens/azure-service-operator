// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240401

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type Workspace_Spec_ARM struct {
	// Identity: The identity of the resource.
	Identity *ManagedServiceIdentity_ARM `json:"identity,omitempty"`
	Kind     *string                     `json:"kind,omitempty"`

	// Location: Specifies the location of the resource.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The properties of the machine learning workspace.
	Properties *WorkspaceProperties_ARM `json:"properties,omitempty"`

	// Sku: The sku of the workspace.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Contains resource tags defined as key/value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspace_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-04-01"
func (workspace Workspace_Spec_ARM) GetAPIVersion() string {
	return "2024-04-01"
}

// GetName returns the Name of the resource
func (workspace *Workspace_Spec_ARM) GetName() string {
	return workspace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces"
func (workspace *Workspace_Spec_ARM) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces"
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity_ARM struct {
	// Type: Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type                   *ManagedServiceIdentityType_ARM            `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// The resource model definition representing SKU
type Sku_ARM struct {
	// Capacity: If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible
	// for the resource this may be omitted.
	Capacity *int `json:"capacity,omitempty"`

	// Family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// Name: The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`

	// Size: The SKU size. When the name field is the combination of tier and some other value, this would be the standalone
	// code.
	Size *string `json:"size,omitempty"`

	// Tier: This field is required to be implemented by the Resource Provider if the service has more than one tier, but is
	// not  required on a PUT.
	Tier *SkuTier_ARM `json:"tier,omitempty"`
}

// The properties of a machine learning workspace.
type WorkspaceProperties_ARM struct {
	// AllowPublicAccessWhenBehindVnet: The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet *bool    `json:"allowPublicAccessWhenBehindVnet,omitempty"`
	ApplicationInsights             *string  `json:"applicationInsights,omitempty"`
	AssociatedWorkspaces            []string `json:"associatedWorkspaces,omitempty"`
	ContainerRegistry               *string  `json:"containerRegistry,omitempty"`

	// Description: The description of this workspace.
	Description *string `json:"description,omitempty"`

	// DiscoveryUrl: Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl        *string `json:"discoveryUrl,omitempty"`
	EnableDataIsolation *bool   `json:"enableDataIsolation,omitempty"`

	// Encryption: The encryption settings of Azure ML workspace.
	Encryption *EncryptionProperty_ARM `json:"encryption,omitempty"`

	// FeatureStoreSettings: Settings for feature store type workspace.
	FeatureStoreSettings *FeatureStoreSettings_ARM `json:"featureStoreSettings,omitempty"`

	// FriendlyName: The friendly name for this workspace. This name in mutable
	FriendlyName *string `json:"friendlyName,omitempty"`

	// HbiWorkspace: The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace  *bool   `json:"hbiWorkspace,omitempty"`
	HubResourceId *string `json:"hubResourceId,omitempty"`

	// ImageBuildCompute: The compute name for image build
	ImageBuildCompute *string `json:"imageBuildCompute,omitempty"`
	KeyVault          *string `json:"keyVault,omitempty"`

	// ManagedNetwork: Managed Network settings for a machine learning workspace.
	ManagedNetwork              *ManagedNetworkSettings_ARM `json:"managedNetwork,omitempty"`
	PrimaryUserAssignedIdentity *string                     `json:"primaryUserAssignedIdentity,omitempty"`

	// PublicNetworkAccess: Whether requests from Public Network are allowed.
	PublicNetworkAccess *WorkspaceProperties_PublicNetworkAccess_ARM `json:"publicNetworkAccess,omitempty"`

	// ServerlessComputeSettings: Settings for serverless compute created in the workspace
	ServerlessComputeSettings *ServerlessComputeSettings_ARM `json:"serverlessComputeSettings,omitempty"`

	// ServiceManagedResourcesSettings: The service managed resource settings.
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings_ARM `json:"serviceManagedResourcesSettings,omitempty"`

	// SharedPrivateLinkResources: The list of shared private link resources in this workspace.
	SharedPrivateLinkResources []SharedPrivateLinkResource_ARM `json:"sharedPrivateLinkResources,omitempty"`
	StorageAccount             *string                         `json:"storageAccount,omitempty"`

	// V1LegacyMode: Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
	V1LegacyMode *bool `json:"v1LegacyMode,omitempty"`

	// WorkspaceHubConfig: WorkspaceHub's configuration object.
	WorkspaceHubConfig *WorkspaceHubConfig_ARM `json:"workspaceHubConfig,omitempty"`
}

type EncryptionProperty_ARM struct {
	// Identity: The identity that will be used to access the key vault for encryption at rest.
	Identity *IdentityForCmk_ARM `json:"identity,omitempty"`

	// KeyVaultProperties: Customer Key vault properties.
	KeyVaultProperties *EncryptionKeyVaultProperties_ARM `json:"keyVaultProperties,omitempty"`
}

// Settings for feature store type workspace.
type FeatureStoreSettings_ARM struct {
	// ComputeRuntime: Compute runtime config for feature store type workspace.
	ComputeRuntime             *ComputeRuntimeDto_ARM `json:"computeRuntime,omitempty"`
	OfflineStoreConnectionName *string                `json:"offlineStoreConnectionName,omitempty"`
	OnlineStoreConnectionName  *string                `json:"onlineStoreConnectionName,omitempty"`
}

// Managed Network settings for a machine learning workspace.
type ManagedNetworkSettings_ARM struct {
	// IsolationMode: Isolation mode for the managed network of a machine learning workspace.
	IsolationMode *IsolationMode_ARM          `json:"isolationMode,omitempty"`
	OutboundRules map[string]OutboundRule_ARM `json:"outboundRules,omitempty"`

	// Status: Status of the Provisioning for the managed network of a machine learning workspace.
	Status *ManagedNetworkProvisionStatus_ARM `json:"status,omitempty"`
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ManagedServiceIdentityType_ARM string

const (
	ManagedServiceIdentityType_ARM_None                       = ManagedServiceIdentityType_ARM("None")
	ManagedServiceIdentityType_ARM_SystemAssigned             = ManagedServiceIdentityType_ARM("SystemAssigned")
	ManagedServiceIdentityType_ARM_SystemAssignedUserAssigned = ManagedServiceIdentityType_ARM("SystemAssigned,UserAssigned")
	ManagedServiceIdentityType_ARM_UserAssigned               = ManagedServiceIdentityType_ARM("UserAssigned")
)

// Mapping from string to ManagedServiceIdentityType_ARM
var managedServiceIdentityType_ARM_Values = map[string]ManagedServiceIdentityType_ARM{
	"none":                        ManagedServiceIdentityType_ARM_None,
	"systemassigned":              ManagedServiceIdentityType_ARM_SystemAssigned,
	"systemassigned,userassigned": ManagedServiceIdentityType_ARM_SystemAssignedUserAssigned,
	"userassigned":                ManagedServiceIdentityType_ARM_UserAssigned,
}

type ServerlessComputeSettings_ARM struct {
	ServerlessComputeCustomSubnet *string `json:"serverlessComputeCustomSubnet,omitempty"`

	// ServerlessComputeNoPublicIP: The flag to signal if serverless compute nodes deployed in custom vNet would have no public
	// IP addresses for a workspace with private endpoint
	ServerlessComputeNoPublicIP *bool `json:"serverlessComputeNoPublicIP,omitempty"`
}

type ServiceManagedResourcesSettings_ARM struct {
	// CosmosDb: The settings for the service managed cosmosdb account.
	CosmosDb *CosmosDbSettings_ARM `json:"cosmosDb,omitempty"`
}

type SharedPrivateLinkResource_ARM struct {
	// Name: Unique name of the private link.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *SharedPrivateLinkResourceProperty_ARM `json:"properties,omitempty"`
}

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not
// required on a PUT.
// +kubebuilder:validation:Enum={"Basic","Free","Premium","Standard"}
type SkuTier_ARM string

const (
	SkuTier_ARM_Basic    = SkuTier_ARM("Basic")
	SkuTier_ARM_Free     = SkuTier_ARM("Free")
	SkuTier_ARM_Premium  = SkuTier_ARM("Premium")
	SkuTier_ARM_Standard = SkuTier_ARM("Standard")
)

// Mapping from string to SkuTier_ARM
var skuTier_ARM_Values = map[string]SkuTier_ARM{
	"basic":    SkuTier_ARM_Basic,
	"free":     SkuTier_ARM_Free,
	"premium":  SkuTier_ARM_Premium,
	"standard": SkuTier_ARM_Standard,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// WorkspaceHub's configuration object.
type WorkspaceHubConfig_ARM struct {
	AdditionalWorkspaceStorageAccounts []string `json:"additionalWorkspaceStorageAccounts,omitempty"`
	DefaultWorkspaceResourceGroup      *string  `json:"defaultWorkspaceResourceGroup,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type WorkspaceProperties_PublicNetworkAccess_ARM string

const (
	WorkspaceProperties_PublicNetworkAccess_ARM_Disabled = WorkspaceProperties_PublicNetworkAccess_ARM("Disabled")
	WorkspaceProperties_PublicNetworkAccess_ARM_Enabled  = WorkspaceProperties_PublicNetworkAccess_ARM("Enabled")
)

// Mapping from string to WorkspaceProperties_PublicNetworkAccess_ARM
var workspaceProperties_PublicNetworkAccess_ARM_Values = map[string]WorkspaceProperties_PublicNetworkAccess_ARM{
	"disabled": WorkspaceProperties_PublicNetworkAccess_ARM_Disabled,
	"enabled":  WorkspaceProperties_PublicNetworkAccess_ARM_Enabled,
}

// Compute runtime config for feature store type workspace.
type ComputeRuntimeDto_ARM struct {
	SparkRuntimeVersion *string `json:"sparkRuntimeVersion,omitempty"`
}

type CosmosDbSettings_ARM struct {
	// CollectionsThroughput: The throughput of the collections in cosmosdb database
	CollectionsThroughput *int `json:"collectionsThroughput,omitempty"`
}

type EncryptionKeyVaultProperties_ARM struct {
	// IdentityClientId: For future use - The client id of the identity which will be used to access key vault.
	IdentityClientId *string `json:"identityClientId,omitempty" optionalConfigMapPair:"IdentityClientId"`

	// KeyIdentifier: Key vault uri to access the encryption key.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
	KeyVaultArmId *string `json:"keyVaultArmId,omitempty"`
}

// Identity that will be used to access key vault for encryption at rest
type IdentityForCmk_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Isolation mode for the managed network of a machine learning workspace.
// +kubebuilder:validation:Enum={"AllowInternetOutbound","AllowOnlyApprovedOutbound","Disabled"}
type IsolationMode_ARM string

const (
	IsolationMode_ARM_AllowInternetOutbound     = IsolationMode_ARM("AllowInternetOutbound")
	IsolationMode_ARM_AllowOnlyApprovedOutbound = IsolationMode_ARM("AllowOnlyApprovedOutbound")
	IsolationMode_ARM_Disabled                  = IsolationMode_ARM("Disabled")
)

// Mapping from string to IsolationMode_ARM
var isolationMode_ARM_Values = map[string]IsolationMode_ARM{
	"allowinternetoutbound":     IsolationMode_ARM_AllowInternetOutbound,
	"allowonlyapprovedoutbound": IsolationMode_ARM_AllowOnlyApprovedOutbound,
	"disabled":                  IsolationMode_ARM_Disabled,
}

// Status of the Provisioning for the managed network of a machine learning workspace.
type ManagedNetworkProvisionStatus_ARM struct {
	SparkReady *bool `json:"sparkReady,omitempty"`

	// Status: Status for the managed network of a machine learning workspace.
	Status *ManagedNetworkStatus_ARM `json:"status,omitempty"`
}

type OutboundRule_ARM struct {
	// FQDN: Mutually exclusive with all other properties
	FQDN *FqdnOutboundRule_ARM `json:"fqdn,omitempty"`

	// PrivateEndpoint: Mutually exclusive with all other properties
	PrivateEndpoint *PrivateEndpointOutboundRule_ARM `json:"privateEndpoint,omitempty"`

	// ServiceTag: Mutually exclusive with all other properties
	ServiceTag *ServiceTagOutboundRule_ARM `json:"serviceTag,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because OutboundRule_ARM represents a discriminated union (JSON OneOf)
func (rule OutboundRule_ARM) MarshalJSON() ([]byte, error) {
	if rule.FQDN != nil {
		return json.Marshal(rule.FQDN)
	}
	if rule.PrivateEndpoint != nil {
		return json.Marshal(rule.PrivateEndpoint)
	}
	if rule.ServiceTag != nil {
		return json.Marshal(rule.ServiceTag)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the OutboundRule_ARM
func (rule *OutboundRule_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "FQDN" {
		rule.FQDN = &FqdnOutboundRule_ARM{}
		return json.Unmarshal(data, rule.FQDN)
	}
	if discriminator == "PrivateEndpoint" {
		rule.PrivateEndpoint = &PrivateEndpointOutboundRule_ARM{}
		return json.Unmarshal(data, rule.PrivateEndpoint)
	}
	if discriminator == "ServiceTag" {
		rule.ServiceTag = &ServiceTagOutboundRule_ARM{}
		return json.Unmarshal(data, rule.ServiceTag)
	}

	// No error
	return nil
}

// Properties of a shared private link resource.
type SharedPrivateLinkResourceProperty_ARM struct {
	// GroupId: The private link resource group id.
	GroupId               *string `json:"groupId,omitempty"`
	PrivateLinkResourceId *string `json:"privateLinkResourceId,omitempty"`

	// RequestMessage: Request message.
	RequestMessage *string `json:"requestMessage,omitempty"`

	// Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus_ARM `json:"status,omitempty"`
}

type FqdnOutboundRule_ARM struct {
	// Category: Category of a managed network Outbound Rule of a machine learning workspace.
	Category    *RuleCategory_ARM `json:"category,omitempty"`
	Destination *string           `json:"destination,omitempty"`

	// Status: Type of a managed network Outbound Rule of a machine learning workspace.
	Status *RuleStatus_ARM           `json:"status,omitempty"`
	Type   FqdnOutboundRule_Type_ARM `json:"type,omitempty"`
}

// Status for the managed network of a machine learning workspace.
// +kubebuilder:validation:Enum={"Active","Inactive"}
type ManagedNetworkStatus_ARM string

const (
	ManagedNetworkStatus_ARM_Active   = ManagedNetworkStatus_ARM("Active")
	ManagedNetworkStatus_ARM_Inactive = ManagedNetworkStatus_ARM("Inactive")
)

// Mapping from string to ManagedNetworkStatus_ARM
var managedNetworkStatus_ARM_Values = map[string]ManagedNetworkStatus_ARM{
	"active":   ManagedNetworkStatus_ARM_Active,
	"inactive": ManagedNetworkStatus_ARM_Inactive,
}

type PrivateEndpointOutboundRule_ARM struct {
	// Category: Category of a managed network Outbound Rule of a machine learning workspace.
	Category *RuleCategory_ARM `json:"category,omitempty"`

	// Destination: Private Endpoint destination for a Private Endpoint Outbound Rule for the managed network of a machine
	// learning  workspace.
	Destination *PrivateEndpointDestination_ARM `json:"destination,omitempty"`

	// Status: Type of a managed network Outbound Rule of a machine learning workspace.
	Status *RuleStatus_ARM                      `json:"status,omitempty"`
	Type   PrivateEndpointOutboundRule_Type_ARM `json:"type,omitempty"`
}

// The private endpoint connection status.
// +kubebuilder:validation:Enum={"Approved","Disconnected","Pending","Rejected","Timeout"}
type PrivateEndpointServiceConnectionStatus_ARM string

const (
	PrivateEndpointServiceConnectionStatus_ARM_Approved     = PrivateEndpointServiceConnectionStatus_ARM("Approved")
	PrivateEndpointServiceConnectionStatus_ARM_Disconnected = PrivateEndpointServiceConnectionStatus_ARM("Disconnected")
	PrivateEndpointServiceConnectionStatus_ARM_Pending      = PrivateEndpointServiceConnectionStatus_ARM("Pending")
	PrivateEndpointServiceConnectionStatus_ARM_Rejected     = PrivateEndpointServiceConnectionStatus_ARM("Rejected")
	PrivateEndpointServiceConnectionStatus_ARM_Timeout      = PrivateEndpointServiceConnectionStatus_ARM("Timeout")
)

// Mapping from string to PrivateEndpointServiceConnectionStatus_ARM
var privateEndpointServiceConnectionStatus_ARM_Values = map[string]PrivateEndpointServiceConnectionStatus_ARM{
	"approved":     PrivateEndpointServiceConnectionStatus_ARM_Approved,
	"disconnected": PrivateEndpointServiceConnectionStatus_ARM_Disconnected,
	"pending":      PrivateEndpointServiceConnectionStatus_ARM_Pending,
	"rejected":     PrivateEndpointServiceConnectionStatus_ARM_Rejected,
	"timeout":      PrivateEndpointServiceConnectionStatus_ARM_Timeout,
}

type ServiceTagOutboundRule_ARM struct {
	// Category: Category of a managed network Outbound Rule of a machine learning workspace.
	Category *RuleCategory_ARM `json:"category,omitempty"`

	// Destination: Service Tag destination for a Service Tag Outbound Rule for the managed network of a machine learning
	// workspace.
	Destination *ServiceTagDestination_ARM `json:"destination,omitempty"`

	// Status: Type of a managed network Outbound Rule of a machine learning workspace.
	Status *RuleStatus_ARM                 `json:"status,omitempty"`
	Type   ServiceTagOutboundRule_Type_ARM `json:"type,omitempty"`
}

// +kubebuilder:validation:Enum={"FQDN"}
type FqdnOutboundRule_Type_ARM string

const FqdnOutboundRule_Type_ARM_FQDN = FqdnOutboundRule_Type_ARM("FQDN")

// Mapping from string to FqdnOutboundRule_Type_ARM
var fqdnOutboundRule_Type_ARM_Values = map[string]FqdnOutboundRule_Type_ARM{
	"fqdn": FqdnOutboundRule_Type_ARM_FQDN,
}

// Private Endpoint destination for a Private Endpoint Outbound Rule for the managed network of a machine learning
// workspace.
type PrivateEndpointDestination_ARM struct {
	ServiceResourceId *string `json:"serviceResourceId,omitempty"`
	SparkEnabled      *bool   `json:"sparkEnabled,omitempty"`

	// SparkStatus: Type of a managed network Outbound Rule of a machine learning workspace.
	SparkStatus       *RuleStatus_ARM `json:"sparkStatus,omitempty"`
	SubresourceTarget *string         `json:"subresourceTarget,omitempty"`
}

// +kubebuilder:validation:Enum={"PrivateEndpoint"}
type PrivateEndpointOutboundRule_Type_ARM string

const PrivateEndpointOutboundRule_Type_ARM_PrivateEndpoint = PrivateEndpointOutboundRule_Type_ARM("PrivateEndpoint")

// Mapping from string to PrivateEndpointOutboundRule_Type_ARM
var privateEndpointOutboundRule_Type_ARM_Values = map[string]PrivateEndpointOutboundRule_Type_ARM{
	"privateendpoint": PrivateEndpointOutboundRule_Type_ARM_PrivateEndpoint,
}

// Category of a managed network Outbound Rule of a machine learning workspace.
// +kubebuilder:validation:Enum={"Dependency","Recommended","Required","UserDefined"}
type RuleCategory_ARM string

const (
	RuleCategory_ARM_Dependency  = RuleCategory_ARM("Dependency")
	RuleCategory_ARM_Recommended = RuleCategory_ARM("Recommended")
	RuleCategory_ARM_Required    = RuleCategory_ARM("Required")
	RuleCategory_ARM_UserDefined = RuleCategory_ARM("UserDefined")
)

// Mapping from string to RuleCategory_ARM
var ruleCategory_ARM_Values = map[string]RuleCategory_ARM{
	"dependency":  RuleCategory_ARM_Dependency,
	"recommended": RuleCategory_ARM_Recommended,
	"required":    RuleCategory_ARM_Required,
	"userdefined": RuleCategory_ARM_UserDefined,
}

// Type of a managed network Outbound Rule of a machine learning workspace.
// +kubebuilder:validation:Enum={"Active","Inactive"}
type RuleStatus_ARM string

const (
	RuleStatus_ARM_Active   = RuleStatus_ARM("Active")
	RuleStatus_ARM_Inactive = RuleStatus_ARM("Inactive")
)

// Mapping from string to RuleStatus_ARM
var ruleStatus_ARM_Values = map[string]RuleStatus_ARM{
	"active":   RuleStatus_ARM_Active,
	"inactive": RuleStatus_ARM_Inactive,
}

// Service Tag destination for a Service Tag Outbound Rule for the managed network of a machine learning workspace.
type ServiceTagDestination_ARM struct {
	// Action: The action enum for networking rule.
	Action     *RuleAction_ARM `json:"action,omitempty"`
	PortRanges *string         `json:"portRanges,omitempty"`
	Protocol   *string         `json:"protocol,omitempty"`
	ServiceTag *string         `json:"serviceTag,omitempty"`
}

// +kubebuilder:validation:Enum={"ServiceTag"}
type ServiceTagOutboundRule_Type_ARM string

const ServiceTagOutboundRule_Type_ARM_ServiceTag = ServiceTagOutboundRule_Type_ARM("ServiceTag")

// Mapping from string to ServiceTagOutboundRule_Type_ARM
var serviceTagOutboundRule_Type_ARM_Values = map[string]ServiceTagOutboundRule_Type_ARM{
	"servicetag": ServiceTagOutboundRule_Type_ARM_ServiceTag,
}

// The action enum for networking rule.
// +kubebuilder:validation:Enum={"Allow","Deny"}
type RuleAction_ARM string

const (
	RuleAction_ARM_Allow = RuleAction_ARM("Allow")
	RuleAction_ARM_Deny  = RuleAction_ARM("Deny")
)

// Mapping from string to RuleAction_ARM
var ruleAction_ARM_Values = map[string]RuleAction_ARM{
	"allow": RuleAction_ARM_Allow,
	"deny":  RuleAction_ARM_Deny,
}
