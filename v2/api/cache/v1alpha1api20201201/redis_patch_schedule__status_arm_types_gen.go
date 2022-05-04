// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

//Deprecated version of RedisPatchSchedule_Status. Use v1beta20201201.RedisPatchSchedule_Status instead
type RedisPatchSchedule_StatusARM struct {
	Id         *string                    `json:"id,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Properties *ScheduleEntries_StatusARM `json:"properties,omitempty"`
	Type       *string                    `json:"type,omitempty"`
}

//Deprecated version of ScheduleEntries_Status. Use v1beta20201201.ScheduleEntries_Status instead
type ScheduleEntries_StatusARM struct {
	ScheduleEntries []ScheduleEntry_StatusARM `json:"scheduleEntries,omitempty"`
}

//Deprecated version of ScheduleEntry_Status. Use v1beta20201201.ScheduleEntry_Status instead
type ScheduleEntry_StatusARM struct {
	DayOfWeek         *ScheduleEntryStatusDayOfWeek `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                       `json:"maintenanceWindow,omitempty"`
	StartHourUtc      *int                          `json:"startHourUtc,omitempty"`
}