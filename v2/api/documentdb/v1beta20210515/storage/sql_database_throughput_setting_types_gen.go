// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515/storage"
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
// Storage version of v1beta20210515.SqlDatabaseThroughputSetting
// Deprecated version of SqlDatabaseThroughputSetting. Use v1api20210515.SqlDatabaseThroughputSetting instead
type SqlDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseThroughputSetting{}

// ConvertFrom populates our SqlDatabaseThroughputSetting from the provided hub SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20210515/storage/SqlDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_SqlDatabaseThroughputSetting(source)
}

// ConvertTo populates the provided hub SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20210515/storage/SqlDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_SqlDatabaseThroughputSetting(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseThroughputSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *SqlDatabaseThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (setting *SqlDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *SqlDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// AssignProperties_From_SqlDatabaseThroughputSetting populates our SqlDatabaseThroughputSetting from the provided source SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignProperties_From_SqlDatabaseThroughputSetting(source *v20210515s.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
	err = status.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseThroughputSetting populates the provided destination SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignProperties_To_SqlDatabaseThroughputSetting(destination *v20210515s.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
	err = setting.Status.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "SqlDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210515.SqlDatabaseThroughputSetting
// Deprecated version of SqlDatabaseThroughputSetting. Use v1api20210515.SqlDatabaseThroughputSetting instead
type SqlDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseThroughputSetting `json:"items"`
}

type augmentConversionForSqlDatabaseThroughputSetting interface {
	AssignPropertiesFrom(src *v20210515s.SqlDatabaseThroughputSetting) error
	AssignPropertiesTo(dst *v20210515s.SqlDatabaseThroughputSetting) error
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
type DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec from the provided source
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{}
	err := setting.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(dst)
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

// AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec populates our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec from the provided source DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(source *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	setting.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignProperties_From_ThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsResource() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		setting.PropertyBag = propertyBag
	} else {
		setting.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec populates the provided destination DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec from our DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec(destination *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(setting.PropertyBag)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if setting.Resource != nil {
		var resource v20210515s.ThroughputSettingsResource
		err := setting.Resource.AssignProperties_To_ThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
// Deprecated version of DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS. Use v1api20210515.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS instead
type DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS{}

// ConvertStatusFrom populates our DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS from the provided source
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS{}
	err := setting.AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(dst)
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

// AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS populates our DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS from the provided source DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) AssignProperties_From_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(source *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		setting.PropertyBag = propertyBag
	} else {
		setting.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS populates the provided destination DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS from our DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) AssignProperties_To_DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS(destination *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(setting.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// Resource
	if setting.Resource != nil {
		var resource v20210515s.ThroughputSettingsGetProperties_Resource_STATUS
		err := setting.Resource.AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_Spec interface {
	AssignPropertiesFrom(src *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) error
	AssignPropertiesTo(dst *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec) error
}

type augmentConversionForDatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS interface {
	AssignPropertiesFrom(src *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) error
	AssignPropertiesTo(dst *v20210515s.DatabaseAccounts_SqlDatabases_ThroughputSetting_STATUS) error
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseThroughputSetting{}, &SqlDatabaseThroughputSettingList{})
}