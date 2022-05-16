// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersFirewallRules_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the server firewall rule.
	Name string `json:"name,omitempty"`

	// Properties: The properties of a server firewall rule.
	Properties *FirewallRulePropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersFirewallRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (rules FlexibleServersFirewallRules_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (rules FlexibleServersFirewallRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/firewallRules"
func (rules FlexibleServersFirewallRules_SpecARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/firewallRules"
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/FirewallRuleProperties
type FirewallRulePropertiesARM struct {
	// EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}
