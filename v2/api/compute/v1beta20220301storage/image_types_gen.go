// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=compute.azure.com,resources=images,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=compute.azure.com,resources={images/status,images/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20220301.Image
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/image.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Image_Spec   `json:"spec,omitempty"`
	Status            Image_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Image{}

// GetConditions returns the conditions of the resource
func (image *Image) GetConditions() conditions.Conditions {
	return image.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (image *Image) SetConditions(conditions conditions.Conditions) {
	image.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Image{}

// AzureName returns the Azure name of the resource
func (image *Image) AzureName() string {
	return image.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (image Image) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (image *Image) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (image *Image) GetSpec() genruntime.ConvertibleSpec {
	return &image.Spec
}

// GetStatus returns the status of this resource
func (image *Image) GetStatus() genruntime.ConvertibleStatus {
	return &image.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/images"
func (image *Image) GetType() string {
	return "Microsoft.Compute/images"
}

// NewEmptyStatus returns a new empty (blank) status
func (image *Image) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Image_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (image *Image) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(image.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  image.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (image *Image) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Image_STATUS); ok {
		image.Status = *st
		return nil
	}

	// Convert status to required version
	var st Image_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	image.Status = st
	return nil
}

// Hub marks that this Image is the hub type for conversion
func (image *Image) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (image *Image) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: image.Spec.OriginalVersion,
		Kind:    "Image",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20220301.Image
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/image.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Storage version of v1beta20220301.APIVersion
// +kubebuilder:validation:Enum={"2022-03-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-03-01")

// Storage version of v1beta20220301.Image_Spec
type Image_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string            `json:"azureName,omitempty"`
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`
	HyperVGeneration *string           `json:"hyperVGeneration,omitempty"`
	Location         *string           `json:"location,omitempty"`
	OriginalVersion  string            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag          genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SourceVirtualMachine *SubResource                       `json:"sourceVirtualMachine,omitempty"`
	StorageProfile       *ImageStorageProfile               `json:"storageProfile,omitempty"`
	Tags                 map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Image_Spec{}

// ConvertSpecFrom populates our Image_Spec from the provided source
func (image *Image_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == image {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(image)
}

// ConvertSpecTo populates the provided destination from our Image_Spec
func (image *Image_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == image {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(image)
}

// Storage version of v1beta20220301.Image_STATUS
// The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual
// machine. If SourceImage is provided, the destination virtual hard drive must not exist.
type Image_STATUS struct {
	Conditions           []conditions.Condition      `json:"conditions,omitempty"`
	ExtendedLocation     *ExtendedLocation_STATUS    `json:"extendedLocation,omitempty"`
	HyperVGeneration     *string                     `json:"hyperVGeneration,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Location             *string                     `json:"location,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                     `json:"provisioningState,omitempty"`
	SourceVirtualMachine *SubResource_STATUS         `json:"sourceVirtualMachine,omitempty"`
	StorageProfile       *ImageStorageProfile_STATUS `json:"storageProfile,omitempty"`
	Tags                 map[string]string           `json:"tags,omitempty"`
	Type                 *string                     `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Image_STATUS{}

// ConvertStatusFrom populates our Image_STATUS from the provided source
func (image *Image_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == image {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(image)
}

// ConvertStatusTo populates the provided destination from our Image_STATUS
func (image *Image_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == image {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(image)
}

// Storage version of v1beta20220301.ExtendedLocation
// The complex type of the extended location.
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20220301.ExtendedLocation_STATUS
// The complex type of the extended location.
type ExtendedLocation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20220301.ImageStorageProfile
// Describes a storage profile.
type ImageStorageProfile struct {
	DataDisks     []ImageDataDisk        `json:"dataDisks,omitempty"`
	OsDisk        *ImageOSDisk           `json:"osDisk,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ZoneResilient *bool                  `json:"zoneResilient,omitempty"`
}

// Storage version of v1beta20220301.ImageStorageProfile_STATUS
// Describes a storage profile.
type ImageStorageProfile_STATUS struct {
	DataDisks     []ImageDataDisk_STATUS `json:"dataDisks,omitempty"`
	OsDisk        *ImageOSDisk_STATUS    `json:"osDisk,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ZoneResilient *bool                  `json:"zoneResilient,omitempty"`
}

// Storage version of v1beta20220301.SubResource
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource Id
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20220301.SubResource_STATUS
type SubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20220301.ImageDataDisk
// Describes a data disk.
type ImageDataDisk struct {
	BlobUri            *string                `json:"blobUri,omitempty"`
	Caching            *string                `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource           `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                   `json:"diskSizeGB,omitempty"`
	Lun                *int                   `json:"lun,omitempty"`
	ManagedDisk        *SubResource           `json:"managedDisk,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource           `json:"snapshot,omitempty"`
	StorageAccountType *string                `json:"storageAccountType,omitempty"`
}

// Storage version of v1beta20220301.ImageDataDisk_STATUS
// Describes a data disk.
type ImageDataDisk_STATUS struct {
	BlobUri            *string                `json:"blobUri,omitempty"`
	Caching            *string                `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource_STATUS    `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                   `json:"diskSizeGB,omitempty"`
	Lun                *int                   `json:"lun,omitempty"`
	ManagedDisk        *SubResource_STATUS    `json:"managedDisk,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource_STATUS    `json:"snapshot,omitempty"`
	StorageAccountType *string                `json:"storageAccountType,omitempty"`
}

// Storage version of v1beta20220301.ImageOSDisk
// Describes an Operating System disk.
type ImageOSDisk struct {
	BlobUri            *string                `json:"blobUri,omitempty"`
	Caching            *string                `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource           `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                   `json:"diskSizeGB,omitempty"`
	ManagedDisk        *SubResource           `json:"managedDisk,omitempty"`
	OsState            *string                `json:"osState,omitempty"`
	OsType             *string                `json:"osType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource           `json:"snapshot,omitempty"`
	StorageAccountType *string                `json:"storageAccountType,omitempty"`
}

// Storage version of v1beta20220301.ImageOSDisk_STATUS
// Describes an Operating System disk.
type ImageOSDisk_STATUS struct {
	BlobUri            *string                `json:"blobUri,omitempty"`
	Caching            *string                `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource_STATUS    `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                   `json:"diskSizeGB,omitempty"`
	ManagedDisk        *SubResource_STATUS    `json:"managedDisk,omitempty"`
	OsState            *string                `json:"osState,omitempty"`
	OsType             *string                `json:"osType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource_STATUS    `json:"snapshot,omitempty"`
	StorageAccountType *string                `json:"storageAccountType,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}