// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

// Deprecated version of Image_STATUS. Use v1api20210701.Image_STATUS instead
type Image_STATUS_ARM struct {
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`
	Id               *string                      `json:"id,omitempty"`
	Location         *string                      `json:"location,omitempty"`
	Name             *string                      `json:"name,omitempty"`
	Properties       *ImageProperties_STATUS_ARM  `json:"properties,omitempty"`
	Tags             map[string]string            `json:"tags,omitempty"`
	Type             *string                      `json:"type,omitempty"`
}

// Deprecated version of ExtendedLocation_STATUS. Use v1api20210701.ExtendedLocation_STATUS instead
type ExtendedLocation_STATUS_ARM struct {
	Name *string                      `json:"name,omitempty"`
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Deprecated version of ImageProperties_STATUS. Use v1api20210701.ImageProperties_STATUS instead
type ImageProperties_STATUS_ARM struct {
	HyperVGeneration     *HyperVGenerationType_STATUS    `json:"hyperVGeneration,omitempty"`
	ProvisioningState    *string                         `json:"provisioningState,omitempty"`
	SourceVirtualMachine *SubResource_STATUS_ARM         `json:"sourceVirtualMachine,omitempty"`
	StorageProfile       *ImageStorageProfile_STATUS_ARM `json:"storageProfile,omitempty"`
}

// Deprecated version of ExtendedLocationType_STATUS. Use v1api20210701.ExtendedLocationType_STATUS instead
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Deprecated version of ImageStorageProfile_STATUS. Use v1api20210701.ImageStorageProfile_STATUS instead
type ImageStorageProfile_STATUS_ARM struct {
	DataDisks     []ImageDataDisk_STATUS_ARM `json:"dataDisks,omitempty"`
	OsDisk        *ImageOSDisk_STATUS_ARM    `json:"osDisk,omitempty"`
	ZoneResilient *bool                      `json:"zoneResilient,omitempty"`
}

// Deprecated version of SubResource_STATUS. Use v1api20210701.SubResource_STATUS instead
type SubResource_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of ImageDataDisk_STATUS. Use v1api20210701.ImageDataDisk_STATUS instead
type ImageDataDisk_STATUS_ARM struct {
	BlobUri            *string                       `json:"blobUri,omitempty"`
	Caching            *ImageDataDisk_Caching_STATUS `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource_STATUS_ARM       `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                          `json:"diskSizeGB,omitempty"`
	Lun                *int                          `json:"lun,omitempty"`
	ManagedDisk        *SubResource_STATUS_ARM       `json:"managedDisk,omitempty"`
	Snapshot           *SubResource_STATUS_ARM       `json:"snapshot,omitempty"`
	StorageAccountType *StorageAccountType_STATUS    `json:"storageAccountType,omitempty"`
}

// Deprecated version of ImageOSDisk_STATUS. Use v1api20210701.ImageOSDisk_STATUS instead
type ImageOSDisk_STATUS_ARM struct {
	BlobUri            *string                     `json:"blobUri,omitempty"`
	Caching            *ImageOSDisk_Caching_STATUS `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource_STATUS_ARM     `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                        `json:"diskSizeGB,omitempty"`
	ManagedDisk        *SubResource_STATUS_ARM     `json:"managedDisk,omitempty"`
	OsState            *ImageOSDisk_OsState_STATUS `json:"osState,omitempty"`
	OsType             *ImageOSDisk_OsType_STATUS  `json:"osType,omitempty"`
	Snapshot           *SubResource_STATUS_ARM     `json:"snapshot,omitempty"`
	StorageAccountType *StorageAccountType_STATUS  `json:"storageAccountType,omitempty"`
}