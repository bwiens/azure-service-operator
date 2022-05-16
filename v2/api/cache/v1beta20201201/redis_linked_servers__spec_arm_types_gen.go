// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisLinkedServers_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the linked server that is being added to the Redis cache.
	Name string `json:"name,omitempty"`

	// Properties: Create properties for a linked server
	Properties *RedisLinkedServerCreatePropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisLinkedServers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (servers RedisLinkedServers_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (servers RedisLinkedServers_SpecARM) GetName() string {
	return servers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (servers RedisLinkedServers_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisLinkedServerCreateProperties
type RedisLinkedServerCreatePropertiesARM struct {
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	// LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerCreatePropertiesServerRole `json:"serverRole,omitempty"`
}
