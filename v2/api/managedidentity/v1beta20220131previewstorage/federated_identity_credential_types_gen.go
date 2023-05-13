// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220131previewstorage

import (
	"fmt"
	v1api20220131ps "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20220131previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20220131preview.FederatedIdentityCredential
// Deprecated version of FederatedIdentityCredential. Use v1api20220131preview.FederatedIdentityCredential instead
type FederatedIdentityCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserAssignedIdentities_FederatedIdentityCredential_Spec   `json:"spec,omitempty"`
	Status            UserAssignedIdentities_FederatedIdentityCredential_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FederatedIdentityCredential{}

// GetConditions returns the conditions of the resource
func (credential *FederatedIdentityCredential) GetConditions() conditions.Conditions {
	return credential.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (credential *FederatedIdentityCredential) SetConditions(conditions conditions.Conditions) {
	credential.Status.Conditions = conditions
}

var _ conversion.Convertible = &FederatedIdentityCredential{}

// ConvertFrom populates our FederatedIdentityCredential from the provided hub FederatedIdentityCredential
func (credential *FederatedIdentityCredential) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20220131ps.FederatedIdentityCredential)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1api20220131previewstorage/FederatedIdentityCredential but received %T instead", hub)
	}

	return credential.AssignProperties_From_FederatedIdentityCredential(source)
}

// ConvertTo populates the provided hub FederatedIdentityCredential from our FederatedIdentityCredential
func (credential *FederatedIdentityCredential) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20220131ps.FederatedIdentityCredential)
	if !ok {
		return fmt.Errorf("expected managedidentity/v1api20220131previewstorage/FederatedIdentityCredential but received %T instead", hub)
	}

	return credential.AssignProperties_To_FederatedIdentityCredential(destination)
}

var _ genruntime.KubernetesResource = &FederatedIdentityCredential{}

// AzureName returns the Azure name of the resource
func (credential *FederatedIdentityCredential) AzureName() string {
	return credential.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-31-preview"
func (credential FederatedIdentityCredential) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (credential *FederatedIdentityCredential) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (credential *FederatedIdentityCredential) GetSpec() genruntime.ConvertibleSpec {
	return &credential.Spec
}

// GetStatus returns the status of this resource
func (credential *FederatedIdentityCredential) GetStatus() genruntime.ConvertibleStatus {
	return &credential.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
func (credential *FederatedIdentityCredential) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
}

// NewEmptyStatus returns a new empty (blank) status
func (credential *FederatedIdentityCredential) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &UserAssignedIdentities_FederatedIdentityCredential_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (credential *FederatedIdentityCredential) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(credential.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  credential.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (credential *FederatedIdentityCredential) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*UserAssignedIdentities_FederatedIdentityCredential_STATUS); ok {
		credential.Status = *st
		return nil
	}

	// Convert status to required version
	var st UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	credential.Status = st
	return nil
}

// AssignProperties_From_FederatedIdentityCredential populates our FederatedIdentityCredential from the provided source FederatedIdentityCredential
func (credential *FederatedIdentityCredential) AssignProperties_From_FederatedIdentityCredential(source *v1api20220131ps.FederatedIdentityCredential) error {

	// ObjectMeta
	credential.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec UserAssignedIdentities_FederatedIdentityCredential_Spec
	err := spec.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec() to populate field Spec")
	}
	credential.Spec = spec

	// Status
	var status UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err = status.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS() to populate field Status")
	}
	credential.Status = status

	// Invoke the augmentConversionForFederatedIdentityCredential interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential); ok {
		err := augmentedCredential.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FederatedIdentityCredential populates the provided destination FederatedIdentityCredential from our FederatedIdentityCredential
func (credential *FederatedIdentityCredential) AssignProperties_To_FederatedIdentityCredential(destination *v1api20220131ps.FederatedIdentityCredential) error {

	// ObjectMeta
	destination.ObjectMeta = *credential.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec
	err := credential.Spec.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err = credential.Status.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFederatedIdentityCredential interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForFederatedIdentityCredential); ok {
		err := augmentedCredential.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (credential *FederatedIdentityCredential) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: credential.Spec.OriginalVersion,
		Kind:    "FederatedIdentityCredential",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20220131preview.FederatedIdentityCredential
// Deprecated version of FederatedIdentityCredential. Use v1api20220131preview.FederatedIdentityCredential instead
type FederatedIdentityCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedIdentityCredential `json:"items"`
}

// Storage version of v1beta20220131preview.APIVersion
// Deprecated version of APIVersion. Use v1api20220131preview.APIVersion instead
// +kubebuilder:validation:Enum={"2022-01-31-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-01-31-preview")

type augmentConversionForFederatedIdentityCredential interface {
	AssignPropertiesFrom(src *v1api20220131ps.FederatedIdentityCredential) error
	AssignPropertiesTo(dst *v1api20220131ps.FederatedIdentityCredential) error
}

// Storage version of v1beta20220131preview.UserAssignedIdentities_FederatedIdentityCredential_Spec
type UserAssignedIdentities_FederatedIdentityCredential_Spec struct {
	Audiences []string `json:"audiences,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Issuer          *string `json:"issuer,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a managedidentity.azure.com/UserAssignedIdentity resource
	Owner       *genruntime.KnownResourceReference `group:"managedidentity.azure.com" json:"owner,omitempty" kind:"UserAssignedIdentity"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Subject     *string                            `json:"subject,omitempty"`
}

var _ genruntime.ConvertibleSpec = &UserAssignedIdentities_FederatedIdentityCredential_Spec{}

// ConvertSpecFrom populates our UserAssignedIdentities_FederatedIdentityCredential_Spec from the provided source
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec)
	if ok {
		// Populate our instance from source
		return credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentities_FederatedIdentityCredential_Spec
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec)
	if ok {
		// Populate destination from our instance
		return credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec{}
	err := credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec populates our UserAssignedIdentities_FederatedIdentityCredential_Spec from the provided source UserAssignedIdentities_FederatedIdentityCredential_Spec
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(source *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Audiences
	credential.Audiences = genruntime.CloneSliceOfString(source.Audiences)

	// AzureName
	credential.AzureName = source.AzureName

	// Issuer
	credential.Issuer = genruntime.ClonePointerToString(source.Issuer)

	// OriginalVersion
	credential.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		credential.Owner = &owner
	} else {
		credential.Owner = nil
	}

	// Subject
	credential.Subject = genruntime.ClonePointerToString(source.Subject)

	// Update the property bag
	if len(propertyBag) > 0 {
		credential.PropertyBag = propertyBag
	} else {
		credential.PropertyBag = nil
	}

	// Invoke the augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_Spec interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_Spec); ok {
		err := augmentedCredential.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec populates the provided destination UserAssignedIdentities_FederatedIdentityCredential_Spec from our UserAssignedIdentities_FederatedIdentityCredential_Spec
func (credential *UserAssignedIdentities_FederatedIdentityCredential_Spec) AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(destination *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(credential.PropertyBag)

	// Audiences
	destination.Audiences = genruntime.CloneSliceOfString(credential.Audiences)

	// AzureName
	destination.AzureName = credential.AzureName

	// Issuer
	destination.Issuer = genruntime.ClonePointerToString(credential.Issuer)

	// OriginalVersion
	destination.OriginalVersion = credential.OriginalVersion

	// Owner
	if credential.Owner != nil {
		owner := credential.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Subject
	destination.Subject = genruntime.ClonePointerToString(credential.Subject)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_Spec interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_Spec); ok {
		err := augmentedCredential.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20220131preview.UserAssignedIdentities_FederatedIdentityCredential_STATUS
// Deprecated version of UserAssignedIdentities_FederatedIdentityCredential_STATUS. Use v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_STATUS instead
type UserAssignedIdentities_FederatedIdentityCredential_STATUS struct {
	Audiences   []string               `json:"audiences,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Issuer      *string                `json:"issuer,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subject     *string                `json:"subject,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &UserAssignedIdentities_FederatedIdentityCredential_STATUS{}

// ConvertStatusFrom populates our UserAssignedIdentities_FederatedIdentityCredential_STATUS from the provided source
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS)
	if ok {
		// Populate our instance from source
		return credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = credential.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS)
	if ok {
		// Populate destination from our instance
		return credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS{}
	err := credential.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS populates our UserAssignedIdentities_FederatedIdentityCredential_STATUS from the provided source UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(source *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Audiences
	credential.Audiences = genruntime.CloneSliceOfString(source.Audiences)

	// Conditions
	credential.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	credential.Id = genruntime.ClonePointerToString(source.Id)

	// Issuer
	credential.Issuer = genruntime.ClonePointerToString(source.Issuer)

	// Name
	credential.Name = genruntime.ClonePointerToString(source.Name)

	// Subject
	credential.Subject = genruntime.ClonePointerToString(source.Subject)

	// Type
	credential.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		credential.PropertyBag = propertyBag
	} else {
		credential.PropertyBag = nil
	}

	// Invoke the augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_STATUS interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_STATUS); ok {
		err := augmentedCredential.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS populates the provided destination UserAssignedIdentities_FederatedIdentityCredential_STATUS from our UserAssignedIdentities_FederatedIdentityCredential_STATUS
func (credential *UserAssignedIdentities_FederatedIdentityCredential_STATUS) AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(destination *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(credential.PropertyBag)

	// Audiences
	destination.Audiences = genruntime.CloneSliceOfString(credential.Audiences)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(credential.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(credential.Id)

	// Issuer
	destination.Issuer = genruntime.ClonePointerToString(credential.Issuer)

	// Name
	destination.Name = genruntime.ClonePointerToString(credential.Name)

	// Subject
	destination.Subject = genruntime.ClonePointerToString(credential.Subject)

	// Type
	destination.Type = genruntime.ClonePointerToString(credential.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_STATUS interface (if implemented) to customize the conversion
	var credentialAsAny any = credential
	if augmentedCredential, ok := credentialAsAny.(augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_STATUS); ok {
		err := augmentedCredential.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_Spec interface {
	AssignPropertiesFrom(src *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec) error
	AssignPropertiesTo(dst *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec) error
}

type augmentConversionForUserAssignedIdentities_FederatedIdentityCredential_STATUS interface {
	AssignPropertiesFrom(src *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS) error
	AssignPropertiesTo(dst *v1api20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS) error
}

func init() {
	SchemeBuilder.Register(&FederatedIdentityCredential{}, &FederatedIdentityCredentialList{})
}