// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

type SqlDatabaseGetResults_StatusARM struct {
	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	//Properties: The properties of an Azure Cosmos DB SQL database
	Properties *SqlDatabaseGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                   `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type SqlDatabaseGetProperties_StatusARM struct {
	Options  *OptionsResource_StatusARM                   `json:"options,omitempty"`
	Resource *SqlDatabaseGetProperties_Status_ResourceARM `json:"resource,omitempty"`
}

type SqlDatabaseGetProperties_Status_ResourceARM struct {
	//Colls: A system generated property that specified the addressable path of the collections resource.
	Colls *string `json:"_colls,omitempty"`

	//Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	//Id: Name of the Cosmos DB SQL database
	Id string `json:"id"`

	//Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	//Ts: A system generated property that denotes the last updated timestamp of the resource.
	Ts *float64 `json:"_ts,omitempty"`

	//Users: A system generated property that specifies the addressable path of the users resource.
	Users *string `json:"_users,omitempty"`
}