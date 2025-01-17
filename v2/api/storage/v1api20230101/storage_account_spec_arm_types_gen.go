// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccount_Spec_ARM struct {
	// ExtendedLocation: Optional. Set the extended location of the resource. If not set, the storage account will be created
	// in Azure main region. Otherwise it will be created in the specified extended location
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Identity: The identity of the resource.
	Identity *Identity_ARM `json:"identity,omitempty"`

	// Kind: Required. Indicates the type of storage account.
	Kind *StorageAccount_Kind_Spec_ARM `json:"kind,omitempty"`

	// Location: Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure
	// Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is
	// created, but if an identical geo region is specified on update, the request will succeed.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The parameters used to create the storage account.
	Properties *StorageAccountPropertiesCreateParameters_ARM `json:"properties,omitempty"`

	// Sku: Required. Gets or sets the SKU name.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping
	// this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key
	// with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags map[string]string `json:"tags"`
}

var _ genruntime.ARMResourceSpec = &StorageAccount_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (account StorageAccount_Spec_ARM) GetAPIVersion() string {
	return "2023-01-01"
}

// GetName returns the Name of the resource
func (account *StorageAccount_Spec_ARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts"
func (account *StorageAccount_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts"
}

// The complex type of the extended location.
type ExtendedLocation_ARM struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType_ARM `json:"type,omitempty"`
}

// Identity for the resource.
type Identity_ARM struct {
	// Type: The identity type.
	Type                   *Identity_Type_ARM                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// The SKU of the storage account.
type Sku_ARM struct {
	// Name: The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called
	//  accountType.
	Name *SkuName_ARM `json:"name,omitempty"`

	// Tier: The SKU tier. This is based on the SKU name.
	Tier *Tier_ARM `json:"tier,omitempty"`
}

// +kubebuilder:validation:Enum={"BlobStorage","BlockBlobStorage","FileStorage","Storage","StorageV2"}
type StorageAccount_Kind_Spec_ARM string

const (
	StorageAccount_Kind_Spec_ARM_BlobStorage      = StorageAccount_Kind_Spec_ARM("BlobStorage")
	StorageAccount_Kind_Spec_ARM_BlockBlobStorage = StorageAccount_Kind_Spec_ARM("BlockBlobStorage")
	StorageAccount_Kind_Spec_ARM_FileStorage      = StorageAccount_Kind_Spec_ARM("FileStorage")
	StorageAccount_Kind_Spec_ARM_Storage          = StorageAccount_Kind_Spec_ARM("Storage")
	StorageAccount_Kind_Spec_ARM_StorageV2        = StorageAccount_Kind_Spec_ARM("StorageV2")
)

// Mapping from string to StorageAccount_Kind_Spec_ARM
var storageAccount_Kind_Spec_ARM_Values = map[string]StorageAccount_Kind_Spec_ARM{
	"blobstorage":      StorageAccount_Kind_Spec_ARM_BlobStorage,
	"blockblobstorage": StorageAccount_Kind_Spec_ARM_BlockBlobStorage,
	"filestorage":      StorageAccount_Kind_Spec_ARM_FileStorage,
	"storage":          StorageAccount_Kind_Spec_ARM_Storage,
	"storagev2":        StorageAccount_Kind_Spec_ARM_StorageV2,
}

// The parameters used to create the storage account.
type StorageAccountPropertiesCreateParameters_ARM struct {
	// AccessTier: Required for storage accounts where kind = BlobStorage. The access tier is used for billing. The 'Premium'
	// access tier is the default value for premium block blobs storage account type and it cannot be changed for the premium
	// block blobs storage account type.
	AccessTier *StorageAccountPropertiesCreateParameters_AccessTier_ARM `json:"accessTier,omitempty"`

	// AllowBlobPublicAccess: Allow or disallow public access to all blobs or containers in the storage account. The default
	// interpretation is false for this property.
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty"`

	// AllowCrossTenantReplication: Allow or disallow cross AAD tenant object replication. Set this property to true for new or
	// existing accounts only if object replication policies will involve storage accounts in different AAD tenants. The
	// default interpretation is false for new accounts to follow best security practices by default.
	AllowCrossTenantReplication *bool `json:"allowCrossTenantReplication,omitempty"`

	// AllowSharedKeyAccess: Indicates whether the storage account permits requests to be authorized with the account access
	// key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure
	// Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `json:"allowSharedKeyAccess,omitempty"`

	// AllowedCopyScope: Restrict copy to and from Storage Accounts within an AAD tenant or with Private Links to the same VNet.
	AllowedCopyScope *StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM `json:"allowedCopyScope,omitempty"`

	// AzureFilesIdentityBasedAuthentication: Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication_ARM `json:"azureFilesIdentityBasedAuthentication,omitempty"`

	// CustomDomain: User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported
	// per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name
	// property.
	CustomDomain *CustomDomain_ARM `json:"customDomain,omitempty"`

	// DefaultToOAuthAuthentication: A boolean flag which indicates whether the default authentication is OAuth or not. The
	// default interpretation is false for this property.
	DefaultToOAuthAuthentication *bool `json:"defaultToOAuthAuthentication,omitempty"`

	// DnsEndpointType: Allows you to specify the type of endpoint. Set this to AzureDNSZone to create a large number of
	// accounts in a single subscription, which creates accounts in an Azure DNS Zone and the endpoint URL will have an
	// alphanumeric DNS Zone identifier.
	DnsEndpointType *StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM `json:"dnsEndpointType,omitempty"`

	// Encryption: Encryption settings to be used for server-side encryption for the storage account.
	Encryption *Encryption_ARM `json:"encryption,omitempty"`

	// ImmutableStorageWithVersioning: The property is immutable and can only be set to true at the account creation time. When
	// set to true, it enables object level immutability for all the new containers in the account by default.
	ImmutableStorageWithVersioning *ImmutableStorageAccount_ARM `json:"immutableStorageWithVersioning,omitempty"`

	// IsHnsEnabled: Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty"`

	// IsLocalUserEnabled: Enables local users feature, if set to true
	IsLocalUserEnabled *bool `json:"isLocalUserEnabled,omitempty"`

	// IsNfsV3Enabled: NFS 3.0 protocol support enabled if set to true.
	IsNfsV3Enabled *bool `json:"isNfsV3Enabled,omitempty"`

	// IsSftpEnabled: Enables Secure File Transfer Protocol, if set to true
	IsSftpEnabled *bool `json:"isSftpEnabled,omitempty"`

	// KeyPolicy: KeyPolicy assigned to the storage account.
	KeyPolicy *KeyPolicy_ARM `json:"keyPolicy,omitempty"`

	// LargeFileSharesState: Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState *StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM `json:"largeFileSharesState,omitempty"`

	// MinimumTlsVersion: Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS
	// 1.0 for this property.
	MinimumTlsVersion *StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM `json:"minimumTlsVersion,omitempty"`

	// NetworkAcls: Network rule set
	NetworkAcls *NetworkRuleSet_ARM `json:"networkAcls,omitempty"`

	// PublicNetworkAccess: Allow or disallow public network access to Storage Account. Value is optional but if passed in,
	// must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM `json:"publicNetworkAccess,omitempty"`

	// RoutingPreference: Maintains information about the network routing choice opted by the user for data transfer
	RoutingPreference *RoutingPreference_ARM `json:"routingPreference,omitempty"`

	// SasPolicy: SasPolicy assigned to the storage account.
	SasPolicy *SasPolicy_ARM `json:"sasPolicy,omitempty"`

	// SupportsHttpsTrafficOnly: Allows https traffic only to storage service if sets to true. The default value is true since
	// API version 2019-04-01.
	SupportsHttpsTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

// Settings for Azure Files identity based authentication.
type AzureFilesIdentityBasedAuthentication_ARM struct {
	// ActiveDirectoryProperties: Required if directoryServiceOptions are AD, optional if they are AADKERB.
	ActiveDirectoryProperties *ActiveDirectoryProperties_ARM `json:"activeDirectoryProperties,omitempty"`

	// DefaultSharePermission: Default share permission for users using Kerberos authentication if RBAC role is not assigned.
	DefaultSharePermission *AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM `json:"defaultSharePermission,omitempty"`

	// DirectoryServiceOptions: Indicates the directory service used. Note that this enum may be extended in the future.
	DirectoryServiceOptions *AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM `json:"directoryServiceOptions,omitempty"`
}

// The custom domain assigned to this storage account. This can be set via Update.
type CustomDomain_ARM struct {
	// Name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name *string `json:"name,omitempty"`

	// UseSubDomainName: Indicates whether indirect CName validation is enabled. Default value is false. This should only be
	// set on updates.
	UseSubDomainName *bool `json:"useSubDomainName,omitempty"`
}

// The encryption settings on the storage account.
type Encryption_ARM struct {
	// Identity: The identity to be used with service-side encryption at rest.
	Identity *EncryptionIdentity_ARM `json:"identity,omitempty"`

	// KeySource: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage,
	// Microsoft.Keyvault
	KeySource *Encryption_KeySource_ARM `json:"keySource,omitempty"`

	// Keyvaultproperties: Properties provided by key vault.
	Keyvaultproperties *KeyVaultProperties_ARM `json:"keyvaultproperties,omitempty"`

	// RequireInfrastructureEncryption: A boolean indicating whether or not the service applies a secondary layer of encryption
	// with platform managed keys for data at rest.
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`

	// Services: List of services which support encryption.
	Services *EncryptionServices_ARM `json:"services,omitempty"`
}

// The type of extendedLocation.
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType_ARM string

const ExtendedLocationType_ARM_EdgeZone = ExtendedLocationType_ARM("EdgeZone")

// Mapping from string to ExtendedLocationType_ARM
var extendedLocationType_ARM_Values = map[string]ExtendedLocationType_ARM{
	"edgezone": ExtendedLocationType_ARM_EdgeZone,
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type Identity_Type_ARM string

const (
	Identity_Type_ARM_None                       = Identity_Type_ARM("None")
	Identity_Type_ARM_SystemAssigned             = Identity_Type_ARM("SystemAssigned")
	Identity_Type_ARM_SystemAssignedUserAssigned = Identity_Type_ARM("SystemAssigned,UserAssigned")
	Identity_Type_ARM_UserAssigned               = Identity_Type_ARM("UserAssigned")
)

// Mapping from string to Identity_Type_ARM
var identity_Type_ARM_Values = map[string]Identity_Type_ARM{
	"none":                        Identity_Type_ARM_None,
	"systemassigned":              Identity_Type_ARM_SystemAssigned,
	"systemassigned,userassigned": Identity_Type_ARM_SystemAssignedUserAssigned,
	"userassigned":                Identity_Type_ARM_UserAssigned,
}

// This property enables and defines account-level immutability. Enabling the feature auto-enables Blob Versioning.
type ImmutableStorageAccount_ARM struct {
	// Enabled: A boolean flag which enables account-level immutability. All the containers under such an account have
	// object-level immutability enabled by default.
	Enabled *bool `json:"enabled,omitempty"`

	// ImmutabilityPolicy: Specifies the default account-level immutability policy which is inherited and applied to objects
	// that do not possess an explicit immutability policy at the object level. The object-level immutability policy has higher
	// precedence than the container-level immutability policy, which has a higher precedence than the account-level
	// immutability policy.
	ImmutabilityPolicy *AccountImmutabilityPolicyProperties_ARM `json:"immutabilityPolicy,omitempty"`
}

// KeyPolicy assigned to the storage account.
type KeyPolicy_ARM struct {
	// KeyExpirationPeriodInDays: The key expiration period in days.
	KeyExpirationPeriodInDays *int `json:"keyExpirationPeriodInDays,omitempty"`
}

// Network rule set
type NetworkRuleSet_ARM struct {
	// Bypass: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of
	// Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
	Bypass *NetworkRuleSet_Bypass_ARM `json:"bypass,omitempty"`

	// DefaultAction: Specifies the default action of allow or deny when no other rules match.
	DefaultAction *NetworkRuleSet_DefaultAction_ARM `json:"defaultAction,omitempty"`

	// IpRules: Sets the IP ACL rules
	IpRules []IPRule_ARM `json:"ipRules"`

	// ResourceAccessRules: Sets the resource access rules
	ResourceAccessRules []ResourceAccessRule_ARM `json:"resourceAccessRules"`

	// VirtualNetworkRules: Sets the virtual network rules
	VirtualNetworkRules []VirtualNetworkRule_ARM `json:"virtualNetworkRules"`
}

// Routing preference defines the type of network, either microsoft or internet routing to be used to deliver the user
// data, the default option is microsoft routing
type RoutingPreference_ARM struct {
	// PublishInternetEndpoints: A boolean flag which indicates whether internet routing storage endpoints are to be published
	PublishInternetEndpoints *bool `json:"publishInternetEndpoints,omitempty"`

	// PublishMicrosoftEndpoints: A boolean flag which indicates whether microsoft routing storage endpoints are to be published
	PublishMicrosoftEndpoints *bool `json:"publishMicrosoftEndpoints,omitempty"`

	// RoutingChoice: Routing Choice defines the kind of network routing opted by the user.
	RoutingChoice *RoutingPreference_RoutingChoice_ARM `json:"routingChoice,omitempty"`
}

// SasPolicy assigned to the storage account.
type SasPolicy_ARM struct {
	// ExpirationAction: The SAS expiration action. Can only be Log.
	ExpirationAction *SasPolicy_ExpirationAction_ARM `json:"expirationAction,omitempty"`

	// SasExpirationPeriod: The SAS expiration period, DD.HH:MM:SS.
	SasExpirationPeriod *string `json:"sasExpirationPeriod,omitempty"`
}

// The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called
// accountType.
// +kubebuilder:validation:Enum={"Premium_LRS","Premium_ZRS","Standard_GRS","Standard_GZRS","Standard_LRS","Standard_RAGRS","Standard_RAGZRS","Standard_ZRS"}
type SkuName_ARM string

const (
	SkuName_ARM_Premium_LRS     = SkuName_ARM("Premium_LRS")
	SkuName_ARM_Premium_ZRS     = SkuName_ARM("Premium_ZRS")
	SkuName_ARM_Standard_GRS    = SkuName_ARM("Standard_GRS")
	SkuName_ARM_Standard_GZRS   = SkuName_ARM("Standard_GZRS")
	SkuName_ARM_Standard_LRS    = SkuName_ARM("Standard_LRS")
	SkuName_ARM_Standard_RAGRS  = SkuName_ARM("Standard_RAGRS")
	SkuName_ARM_Standard_RAGZRS = SkuName_ARM("Standard_RAGZRS")
	SkuName_ARM_Standard_ZRS    = SkuName_ARM("Standard_ZRS")
)

// Mapping from string to SkuName_ARM
var skuName_ARM_Values = map[string]SkuName_ARM{
	"premium_lrs":     SkuName_ARM_Premium_LRS,
	"premium_zrs":     SkuName_ARM_Premium_ZRS,
	"standard_grs":    SkuName_ARM_Standard_GRS,
	"standard_gzrs":   SkuName_ARM_Standard_GZRS,
	"standard_lrs":    SkuName_ARM_Standard_LRS,
	"standard_ragrs":  SkuName_ARM_Standard_RAGRS,
	"standard_ragzrs": SkuName_ARM_Standard_RAGZRS,
	"standard_zrs":    SkuName_ARM_Standard_ZRS,
}

// +kubebuilder:validation:Enum={"Cool","Hot","Premium"}
type StorageAccountPropertiesCreateParameters_AccessTier_ARM string

const (
	StorageAccountPropertiesCreateParameters_AccessTier_ARM_Cool    = StorageAccountPropertiesCreateParameters_AccessTier_ARM("Cool")
	StorageAccountPropertiesCreateParameters_AccessTier_ARM_Hot     = StorageAccountPropertiesCreateParameters_AccessTier_ARM("Hot")
	StorageAccountPropertiesCreateParameters_AccessTier_ARM_Premium = StorageAccountPropertiesCreateParameters_AccessTier_ARM("Premium")
)

// Mapping from string to StorageAccountPropertiesCreateParameters_AccessTier_ARM
var storageAccountPropertiesCreateParameters_AccessTier_ARM_Values = map[string]StorageAccountPropertiesCreateParameters_AccessTier_ARM{
	"cool":    StorageAccountPropertiesCreateParameters_AccessTier_ARM_Cool,
	"hot":     StorageAccountPropertiesCreateParameters_AccessTier_ARM_Hot,
	"premium": StorageAccountPropertiesCreateParameters_AccessTier_ARM_Premium,
}

// +kubebuilder:validation:Enum={"AAD","PrivateLink"}
type StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM string

const (
	StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM_AAD         = StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM("AAD")
	StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM_PrivateLink = StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM("PrivateLink")
)

// Mapping from string to StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM
var storageAccountPropertiesCreateParameters_AllowedCopyScope_ARM_Values = map[string]StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM{
	"aad":         StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM_AAD,
	"privatelink": StorageAccountPropertiesCreateParameters_AllowedCopyScope_ARM_PrivateLink,
}

// +kubebuilder:validation:Enum={"AzureDnsZone","Standard"}
type StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM string

const (
	StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM_AzureDnsZone = StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM("AzureDnsZone")
	StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM_Standard     = StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM("Standard")
)

// Mapping from string to StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM
var storageAccountPropertiesCreateParameters_DnsEndpointType_ARM_Values = map[string]StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM{
	"azurednszone": StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM_AzureDnsZone,
	"standard":     StorageAccountPropertiesCreateParameters_DnsEndpointType_ARM_Standard,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM string

const (
	StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM_Disabled = StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM("Disabled")
	StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM_Enabled  = StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM("Enabled")
)

// Mapping from string to StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM
var storageAccountPropertiesCreateParameters_LargeFileSharesState_ARM_Values = map[string]StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM{
	"disabled": StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM_Disabled,
	"enabled":  StorageAccountPropertiesCreateParameters_LargeFileSharesState_ARM_Enabled,
}

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2"}
type StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM string

const (
	StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM_TLS1_0 = StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM("TLS1_0")
	StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM_TLS1_1 = StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM("TLS1_1")
	StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM_TLS1_2 = StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM("TLS1_2")
)

// Mapping from string to StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM
var storageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM_Values = map[string]StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM{
	"tls1_0": StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM_TLS1_0,
	"tls1_1": StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM_TLS1_1,
	"tls1_2": StorageAccountPropertiesCreateParameters_MinimumTlsVersion_ARM_TLS1_2,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM string

const (
	StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM_Disabled = StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM("Disabled")
	StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM_Enabled  = StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM("Enabled")
)

// Mapping from string to StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM
var storageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM_Values = map[string]StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM{
	"disabled": StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM_Disabled,
	"enabled":  StorageAccountPropertiesCreateParameters_PublicNetworkAccess_ARM_Enabled,
}

// The SKU tier. This is based on the SKU name.
// +kubebuilder:validation:Enum={"Premium","Standard"}
type Tier_ARM string

const (
	Tier_ARM_Premium  = Tier_ARM("Premium")
	Tier_ARM_Standard = Tier_ARM("Standard")
)

// Mapping from string to Tier_ARM
var tier_ARM_Values = map[string]Tier_ARM{
	"premium":  Tier_ARM_Premium,
	"standard": Tier_ARM_Standard,
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// This defines account-level immutability policy properties.
type AccountImmutabilityPolicyProperties_ARM struct {
	// AllowProtectedAppendWrites: This property can only be changed for disabled and unlocked time-based retention policies.
	// When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only
	// new blocks can be added and any existing blocks cannot be modified or deleted.
	AllowProtectedAppendWrites *bool `json:"allowProtectedAppendWrites,omitempty"`

	// ImmutabilityPeriodSinceCreationInDays: The immutability period for the blobs in the container since the policy creation,
	// in days.
	ImmutabilityPeriodSinceCreationInDays *int `json:"immutabilityPeriodSinceCreationInDays,omitempty"`

	// State: The ImmutabilityPolicy state defines the mode of the policy. Disabled state disables the policy, Unlocked state
	// allows increase and decrease of immutability retention time and also allows toggling allowProtectedAppendWrites
	// property, Locked state only allows the increase of the immutability retention time. A policy can only be created in a
	// Disabled or Unlocked state and can be toggled between the two states. Only a policy in an Unlocked state can transition
	// to a Locked state which cannot be reverted.
	State *AccountImmutabilityPolicyProperties_State_ARM `json:"state,omitempty"`
}

// Settings properties for Active Directory (AD).
type ActiveDirectoryProperties_ARM struct {
	// AccountType: Specifies the Active Directory account type for Azure Storage.
	AccountType *ActiveDirectoryProperties_AccountType_ARM `json:"accountType,omitempty"`

	// AzureStorageSid: Specifies the security identifier (SID) for Azure Storage.
	AzureStorageSid *string `json:"azureStorageSid,omitempty"`

	// DomainGuid: Specifies the domain GUID.
	DomainGuid *string `json:"domainGuid,omitempty"`

	// DomainName: Specifies the primary domain that the AD DNS server is authoritative for.
	DomainName *string `json:"domainName,omitempty"`

	// DomainSid: Specifies the security identifier (SID).
	DomainSid *string `json:"domainSid,omitempty"`

	// ForestName: Specifies the Active Directory forest to get.
	ForestName *string `json:"forestName,omitempty"`

	// NetBiosDomainName: Specifies the NetBIOS domain name.
	NetBiosDomainName *string `json:"netBiosDomainName,omitempty"`

	// SamAccountName: Specifies the Active Directory SAMAccountName for Azure Storage.
	SamAccountName *string `json:"samAccountName,omitempty"`
}

// +kubebuilder:validation:Enum={"None","StorageFileDataSmbShareContributor","StorageFileDataSmbShareElevatedContributor","StorageFileDataSmbShareReader"}
type AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM string

const (
	AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_None                                       = AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM("None")
	AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_StorageFileDataSmbShareContributor         = AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM("StorageFileDataSmbShareContributor")
	AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_StorageFileDataSmbShareElevatedContributor = AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM("StorageFileDataSmbShareElevatedContributor")
	AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_StorageFileDataSmbShareReader              = AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM("StorageFileDataSmbShareReader")
)

// Mapping from string to AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM
var azureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_Values = map[string]AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM{
	"none":                               AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_None,
	"storagefiledatasmbsharecontributor": AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_StorageFileDataSmbShareContributor,
	"storagefiledatasmbshareelevatedcontributor": AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_StorageFileDataSmbShareElevatedContributor,
	"storagefiledatasmbsharereader":              AzureFilesIdentityBasedAuthentication_DefaultSharePermission_ARM_StorageFileDataSmbShareReader,
}

// +kubebuilder:validation:Enum={"AADDS","AADKERB","AD","None"}
type AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM string

const (
	AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_AADDS   = AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM("AADDS")
	AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_AADKERB = AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM("AADKERB")
	AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_AD      = AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM("AD")
	AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_None    = AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM("None")
)

// Mapping from string to AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM
var azureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_Values = map[string]AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM{
	"aadds":   AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_AADDS,
	"aadkerb": AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_AADKERB,
	"ad":      AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_AD,
	"none":    AzureFilesIdentityBasedAuthentication_DirectoryServiceOptions_ARM_None,
}

// +kubebuilder:validation:Enum={"Microsoft.Keyvault","Microsoft.Storage"}
type Encryption_KeySource_ARM string

const (
	Encryption_KeySource_ARM_MicrosoftKeyvault = Encryption_KeySource_ARM("Microsoft.Keyvault")
	Encryption_KeySource_ARM_MicrosoftStorage  = Encryption_KeySource_ARM("Microsoft.Storage")
)

// Mapping from string to Encryption_KeySource_ARM
var encryption_KeySource_ARM_Values = map[string]Encryption_KeySource_ARM{
	"microsoft.keyvault": Encryption_KeySource_ARM_MicrosoftKeyvault,
	"microsoft.storage":  Encryption_KeySource_ARM_MicrosoftStorage,
}

// Encryption identity for the storage account.
type EncryptionIdentity_ARM struct {
	// FederatedIdentityClientId: ClientId of the multi-tenant application to be used in conjunction with the user-assigned
	// identity for cross-tenant customer-managed-keys server-side encryption on the storage account.
	FederatedIdentityClientId *string `json:"federatedIdentityClientId,omitempty"`
	UserAssignedIdentity      *string `json:"userAssignedIdentity,omitempty"`
}

// A list of services that support encryption.
type EncryptionServices_ARM struct {
	// Blob: The encryption function of the blob storage service.
	Blob *EncryptionService_ARM `json:"blob,omitempty"`

	// File: The encryption function of the file storage service.
	File *EncryptionService_ARM `json:"file,omitempty"`

	// Queue: The encryption function of the queue storage service.
	Queue *EncryptionService_ARM `json:"queue,omitempty"`

	// Table: The encryption function of the table storage service.
	Table *EncryptionService_ARM `json:"table,omitempty"`
}

// IP rule with specific IP or IP range in CIDR format.
type IPRule_ARM struct {
	// Action: The action of IP ACL rule.
	Action *IPRule_Action_ARM `json:"action,omitempty"`

	// Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value *string `json:"value,omitempty"`
}

// Properties of key vault.
type KeyVaultProperties_ARM struct {
	// Keyname: The name of KeyVault key.
	Keyname *string `json:"keyname,omitempty"`

	// Keyvaulturi: The Uri of KeyVault.
	Keyvaulturi *string `json:"keyvaulturi,omitempty"`

	// Keyversion: The version of KeyVault key.
	Keyversion *string `json:"keyversion,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureServices","Logging","Metrics","None"}
type NetworkRuleSet_Bypass_ARM string

const (
	NetworkRuleSet_Bypass_ARM_AzureServices = NetworkRuleSet_Bypass_ARM("AzureServices")
	NetworkRuleSet_Bypass_ARM_Logging       = NetworkRuleSet_Bypass_ARM("Logging")
	NetworkRuleSet_Bypass_ARM_Metrics       = NetworkRuleSet_Bypass_ARM("Metrics")
	NetworkRuleSet_Bypass_ARM_None          = NetworkRuleSet_Bypass_ARM("None")
)

// Mapping from string to NetworkRuleSet_Bypass_ARM
var networkRuleSet_Bypass_ARM_Values = map[string]NetworkRuleSet_Bypass_ARM{
	"azureservices": NetworkRuleSet_Bypass_ARM_AzureServices,
	"logging":       NetworkRuleSet_Bypass_ARM_Logging,
	"metrics":       NetworkRuleSet_Bypass_ARM_Metrics,
	"none":          NetworkRuleSet_Bypass_ARM_None,
}

// +kubebuilder:validation:Enum={"Allow","Deny"}
type NetworkRuleSet_DefaultAction_ARM string

const (
	NetworkRuleSet_DefaultAction_ARM_Allow = NetworkRuleSet_DefaultAction_ARM("Allow")
	NetworkRuleSet_DefaultAction_ARM_Deny  = NetworkRuleSet_DefaultAction_ARM("Deny")
)

// Mapping from string to NetworkRuleSet_DefaultAction_ARM
var networkRuleSet_DefaultAction_ARM_Values = map[string]NetworkRuleSet_DefaultAction_ARM{
	"allow": NetworkRuleSet_DefaultAction_ARM_Allow,
	"deny":  NetworkRuleSet_DefaultAction_ARM_Deny,
}

// Resource Access Rule.
type ResourceAccessRule_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`

	// TenantId: Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

// +kubebuilder:validation:Enum={"InternetRouting","MicrosoftRouting"}
type RoutingPreference_RoutingChoice_ARM string

const (
	RoutingPreference_RoutingChoice_ARM_InternetRouting  = RoutingPreference_RoutingChoice_ARM("InternetRouting")
	RoutingPreference_RoutingChoice_ARM_MicrosoftRouting = RoutingPreference_RoutingChoice_ARM("MicrosoftRouting")
)

// Mapping from string to RoutingPreference_RoutingChoice_ARM
var routingPreference_RoutingChoice_ARM_Values = map[string]RoutingPreference_RoutingChoice_ARM{
	"internetrouting":  RoutingPreference_RoutingChoice_ARM_InternetRouting,
	"microsoftrouting": RoutingPreference_RoutingChoice_ARM_MicrosoftRouting,
}

// +kubebuilder:validation:Enum={"Log"}
type SasPolicy_ExpirationAction_ARM string

const SasPolicy_ExpirationAction_ARM_Log = SasPolicy_ExpirationAction_ARM("Log")

// Mapping from string to SasPolicy_ExpirationAction_ARM
var sasPolicy_ExpirationAction_ARM_Values = map[string]SasPolicy_ExpirationAction_ARM{
	"log": SasPolicy_ExpirationAction_ARM_Log,
}

// Virtual Network rule.
type VirtualNetworkRule_ARM struct {
	// Action: The action of virtual network rule.
	Action *VirtualNetworkRule_Action_ARM `json:"action,omitempty"`
	Id     *string                        `json:"id,omitempty"`

	// State: Gets the state of virtual network rule.
	State *VirtualNetworkRule_State_ARM `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Locked","Unlocked"}
type AccountImmutabilityPolicyProperties_State_ARM string

const (
	AccountImmutabilityPolicyProperties_State_ARM_Disabled = AccountImmutabilityPolicyProperties_State_ARM("Disabled")
	AccountImmutabilityPolicyProperties_State_ARM_Locked   = AccountImmutabilityPolicyProperties_State_ARM("Locked")
	AccountImmutabilityPolicyProperties_State_ARM_Unlocked = AccountImmutabilityPolicyProperties_State_ARM("Unlocked")
)

// Mapping from string to AccountImmutabilityPolicyProperties_State_ARM
var accountImmutabilityPolicyProperties_State_ARM_Values = map[string]AccountImmutabilityPolicyProperties_State_ARM{
	"disabled": AccountImmutabilityPolicyProperties_State_ARM_Disabled,
	"locked":   AccountImmutabilityPolicyProperties_State_ARM_Locked,
	"unlocked": AccountImmutabilityPolicyProperties_State_ARM_Unlocked,
}

// +kubebuilder:validation:Enum={"Computer","User"}
type ActiveDirectoryProperties_AccountType_ARM string

const (
	ActiveDirectoryProperties_AccountType_ARM_Computer = ActiveDirectoryProperties_AccountType_ARM("Computer")
	ActiveDirectoryProperties_AccountType_ARM_User     = ActiveDirectoryProperties_AccountType_ARM("User")
)

// Mapping from string to ActiveDirectoryProperties_AccountType_ARM
var activeDirectoryProperties_AccountType_ARM_Values = map[string]ActiveDirectoryProperties_AccountType_ARM{
	"computer": ActiveDirectoryProperties_AccountType_ARM_Computer,
	"user":     ActiveDirectoryProperties_AccountType_ARM_User,
}

// A service that allows server-side encryption to be used.
type EncryptionService_ARM struct {
	// Enabled: A boolean indicating whether or not the service encrypts the data as it is stored. Encryption at rest is
	// enabled by default today and cannot be disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// KeyType: Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped
	// encryption key will be used. 'Service' key type implies that a default service key is used.
	KeyType *EncryptionService_KeyType_ARM `json:"keyType,omitempty"`
}

// +kubebuilder:validation:Enum={"Allow"}
type IPRule_Action_ARM string

const IPRule_Action_ARM_Allow = IPRule_Action_ARM("Allow")

// Mapping from string to IPRule_Action_ARM
var iPRule_Action_ARM_Values = map[string]IPRule_Action_ARM{
	"allow": IPRule_Action_ARM_Allow,
}

// +kubebuilder:validation:Enum={"Allow"}
type VirtualNetworkRule_Action_ARM string

const VirtualNetworkRule_Action_ARM_Allow = VirtualNetworkRule_Action_ARM("Allow")

// Mapping from string to VirtualNetworkRule_Action_ARM
var virtualNetworkRule_Action_ARM_Values = map[string]VirtualNetworkRule_Action_ARM{
	"allow": VirtualNetworkRule_Action_ARM_Allow,
}

// +kubebuilder:validation:Enum={"Deprovisioning","Failed","NetworkSourceDeleted","Provisioning","Succeeded"}
type VirtualNetworkRule_State_ARM string

const (
	VirtualNetworkRule_State_ARM_Deprovisioning       = VirtualNetworkRule_State_ARM("Deprovisioning")
	VirtualNetworkRule_State_ARM_Failed               = VirtualNetworkRule_State_ARM("Failed")
	VirtualNetworkRule_State_ARM_NetworkSourceDeleted = VirtualNetworkRule_State_ARM("NetworkSourceDeleted")
	VirtualNetworkRule_State_ARM_Provisioning         = VirtualNetworkRule_State_ARM("Provisioning")
	VirtualNetworkRule_State_ARM_Succeeded            = VirtualNetworkRule_State_ARM("Succeeded")
)

// Mapping from string to VirtualNetworkRule_State_ARM
var virtualNetworkRule_State_ARM_Values = map[string]VirtualNetworkRule_State_ARM{
	"deprovisioning":       VirtualNetworkRule_State_ARM_Deprovisioning,
	"failed":               VirtualNetworkRule_State_ARM_Failed,
	"networksourcedeleted": VirtualNetworkRule_State_ARM_NetworkSourceDeleted,
	"provisioning":         VirtualNetworkRule_State_ARM_Provisioning,
	"succeeded":            VirtualNetworkRule_State_ARM_Succeeded,
}

// +kubebuilder:validation:Enum={"Account","Service"}
type EncryptionService_KeyType_ARM string

const (
	EncryptionService_KeyType_ARM_Account = EncryptionService_KeyType_ARM("Account")
	EncryptionService_KeyType_ARM_Service = EncryptionService_KeyType_ARM("Service")
)

// Mapping from string to EncryptionService_KeyType_ARM
var encryptionService_KeyType_ARM_Values = map[string]EncryptionService_KeyType_ARM{
	"account": EncryptionService_KeyType_ARM_Account,
	"service": EncryptionService_KeyType_ARM_Service,
}
