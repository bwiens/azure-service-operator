// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230315preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Fleet_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The resource-specific properties for this resource.
	Properties *FleetProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags"`
}

var _ genruntime.ARMResourceSpec = &Fleet_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-03-15-preview"
func (fleet Fleet_Spec_ARM) GetAPIVersion() string {
	return "2023-03-15-preview"
}

// GetName returns the Name of the resource
func (fleet *Fleet_Spec_ARM) GetName() string {
	return fleet.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/fleets"
func (fleet *Fleet_Spec_ARM) GetType() string {
	return "Microsoft.ContainerService/fleets"
}

// Fleet properties.
type FleetProperties_ARM struct {
	// HubProfile: The FleetHubProfile configures the Fleet's hub.
	HubProfile *FleetHubProfile_ARM `json:"hubProfile,omitempty"`
}

// The FleetHubProfile configures the fleet hub.
type FleetHubProfile_ARM struct {
	// DnsPrefix: DNS prefix used to create the FQDN for the Fleet hub.
	DnsPrefix *string `json:"dnsPrefix,omitempty"`
}
