// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of StorageAccountsQueueServices_Spec. Use v1beta20210401.StorageAccountsQueueServices_Spec instead
type StorageAccountsQueueServices_SpecARM struct {
	Location   *string                              `json:"location,omitempty"`
	Name       string                               `json:"name,omitempty"`
	Properties *QueueServicePropertiesPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsQueueServices_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (services StorageAccountsQueueServices_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (services StorageAccountsQueueServices_SpecARM) GetName() string {
	return services.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices"
func (services StorageAccountsQueueServices_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices"
}

// Deprecated version of QueueServicePropertiesProperties. Use v1beta20210401.QueueServicePropertiesProperties instead
type QueueServicePropertiesPropertiesARM struct {
	Cors *CorsRulesARM `json:"cors,omitempty"`
}
