// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230815

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_DatabaseStatistics_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseStatistics_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseStatistics_STATUS_ARM, DatabaseStatistics_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseStatistics_STATUS_ARM runs a test to see if a specific instance of DatabaseStatistics_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseStatistics_STATUS_ARM(subject DatabaseStatistics_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseStatistics_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DatabaseStatistics_STATUS_ARM instances for property testing - lazily instantiated by
// DatabaseStatistics_STATUS_ARMGenerator()
var databaseStatistics_STATUS_ARMGenerator gopter.Gen

// DatabaseStatistics_STATUS_ARMGenerator returns a generator of DatabaseStatistics_STATUS_ARM instances for property testing.
func DatabaseStatistics_STATUS_ARMGenerator() gopter.Gen {
	if databaseStatistics_STATUS_ARMGenerator != nil {
		return databaseStatistics_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS_ARM(generators)
	databaseStatistics_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseStatistics_STATUS_ARM{}), generators)

	return databaseStatistics_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Size"] = gen.PtrOf(gen.Float64())
}

func Test_Database_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase_STATUS_ARM, Database_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase_STATUS_ARM runs a test to see if a specific instance of Database_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase_STATUS_ARM(subject Database_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Database_STATUS_ARM instances for property testing - lazily instantiated by
// Database_STATUS_ARMGenerator()
var database_STATUS_ARMGenerator gopter.Gen

// Database_STATUS_ARMGenerator returns a generator of Database_STATUS_ARM instances for property testing.
func Database_STATUS_ARMGenerator() gopter.Gen {
	if database_STATUS_ARMGenerator != nil {
		return database_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDatabase_STATUS_ARM(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		gens = append(gens, gen.Struct(reflect.TypeOf(Database_STATUS_ARM{}), map[string]gopter.Gen{propName: propGen}))
	}
	database_STATUS_ARMGenerator = gen.OneGenOf(gens...)

	return database_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ReadOnlyFollowing"] = ReadOnlyFollowingDatabase_STATUS_ARMGenerator().Map(func(it ReadOnlyFollowingDatabase_STATUS_ARM) *ReadOnlyFollowingDatabase_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
	gens["ReadWrite"] = ReadWriteDatabase_STATUS_ARMGenerator().Map(func(it ReadWriteDatabase_STATUS_ARM) *ReadWriteDatabase_STATUS_ARM {
		return &it
	}) // generate one case for OneOf type
}

func Test_ReadOnlyFollowingDatabaseProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadOnlyFollowingDatabaseProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties_STATUS_ARM, ReadOnlyFollowingDatabaseProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties_STATUS_ARM runs a test to see if a specific instance of ReadOnlyFollowingDatabaseProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForReadOnlyFollowingDatabaseProperties_STATUS_ARM(subject ReadOnlyFollowingDatabaseProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadOnlyFollowingDatabaseProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ReadOnlyFollowingDatabaseProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ReadOnlyFollowingDatabaseProperties_STATUS_ARMGenerator()
var readOnlyFollowingDatabaseProperties_STATUS_ARMGenerator gopter.Gen

// ReadOnlyFollowingDatabaseProperties_STATUS_ARMGenerator returns a generator of ReadOnlyFollowingDatabaseProperties_STATUS_ARM instances for property testing.
// We first initialize readOnlyFollowingDatabaseProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadOnlyFollowingDatabaseProperties_STATUS_ARMGenerator() gopter.Gen {
	if readOnlyFollowingDatabaseProperties_STATUS_ARMGenerator != nil {
		return readOnlyFollowingDatabaseProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS_ARM(generators)
	readOnlyFollowingDatabaseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabaseProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS_ARM(generators)
	readOnlyFollowingDatabaseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabaseProperties_STATUS_ARM{}), generators)

	return readOnlyFollowingDatabaseProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AttachedDatabaseConfigurationName"] = gen.PtrOf(gen.AlphaString())
	gens["DatabaseShareOrigin"] = gen.PtrOf(gen.OneConstOf(DatabaseShareOrigin_STATUS_ARM_DataShare, DatabaseShareOrigin_STATUS_ARM_Direct, DatabaseShareOrigin_STATUS_ARM_Other))
	gens["HotCachePeriod"] = gen.PtrOf(gen.AlphaString())
	gens["LeaderClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalDatabaseName"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalsModificationKind"] = gen.PtrOf(gen.OneConstOf(ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_None, ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_Replace, ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_Union))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_ARM_Canceled,
		ProvisioningState_STATUS_ARM_Creating,
		ProvisioningState_STATUS_ARM_Deleting,
		ProvisioningState_STATUS_ARM_Failed,
		ProvisioningState_STATUS_ARM_Moving,
		ProvisioningState_STATUS_ARM_Running,
		ProvisioningState_STATUS_ARM_Succeeded))
	gens["SoftDeletePeriod"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabaseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Statistics"] = gen.PtrOf(DatabaseStatistics_STATUS_ARMGenerator())
	gens["SuspensionDetails"] = gen.PtrOf(SuspensionDetails_STATUS_ARMGenerator())
	gens["TableLevelSharingProperties"] = gen.PtrOf(TableLevelSharingProperties_STATUS_ARMGenerator())
}

func Test_ReadOnlyFollowingDatabase_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadOnlyFollowingDatabase_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadOnlyFollowingDatabase_STATUS_ARM, ReadOnlyFollowingDatabase_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadOnlyFollowingDatabase_STATUS_ARM runs a test to see if a specific instance of ReadOnlyFollowingDatabase_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForReadOnlyFollowingDatabase_STATUS_ARM(subject ReadOnlyFollowingDatabase_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadOnlyFollowingDatabase_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ReadOnlyFollowingDatabase_STATUS_ARM instances for property testing - lazily instantiated by
// ReadOnlyFollowingDatabase_STATUS_ARMGenerator()
var readOnlyFollowingDatabase_STATUS_ARMGenerator gopter.Gen

// ReadOnlyFollowingDatabase_STATUS_ARMGenerator returns a generator of ReadOnlyFollowingDatabase_STATUS_ARM instances for property testing.
// We first initialize readOnlyFollowingDatabase_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadOnlyFollowingDatabase_STATUS_ARMGenerator() gopter.Gen {
	if readOnlyFollowingDatabase_STATUS_ARMGenerator != nil {
		return readOnlyFollowingDatabase_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS_ARM(generators)
	readOnlyFollowingDatabase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabase_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS_ARM(generators)
	readOnlyFollowingDatabase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadOnlyFollowingDatabase_STATUS_ARM{}), generators)

	return readOnlyFollowingDatabase_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.OneConstOf(ReadOnlyFollowingDatabase_Kind_STATUS_ARM_ReadOnlyFollowing)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadOnlyFollowingDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReadOnlyFollowingDatabaseProperties_STATUS_ARMGenerator())
}

func Test_ReadWriteDatabaseProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabaseProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS_ARM, ReadWriteDatabaseProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS_ARM runs a test to see if a specific instance of ReadWriteDatabaseProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS_ARM(subject ReadWriteDatabaseProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabaseProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ReadWriteDatabaseProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ReadWriteDatabaseProperties_STATUS_ARMGenerator()
var readWriteDatabaseProperties_STATUS_ARMGenerator gopter.Gen

// ReadWriteDatabaseProperties_STATUS_ARMGenerator returns a generator of ReadWriteDatabaseProperties_STATUS_ARM instances for property testing.
// We first initialize readWriteDatabaseProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabaseProperties_STATUS_ARMGenerator() gopter.Gen {
	if readWriteDatabaseProperties_STATUS_ARMGenerator != nil {
		return readWriteDatabaseProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS_ARM(generators)
	readWriteDatabaseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS_ARM(generators)
	readWriteDatabaseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties_STATUS_ARM{}), generators)

	return readWriteDatabaseProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["HotCachePeriod"] = gen.PtrOf(gen.AlphaString())
	gens["IsFollowed"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_ARM_Canceled,
		ProvisioningState_STATUS_ARM_Creating,
		ProvisioningState_STATUS_ARM_Deleting,
		ProvisioningState_STATUS_ARM_Failed,
		ProvisioningState_STATUS_ARM_Moving,
		ProvisioningState_STATUS_ARM_Running,
		ProvisioningState_STATUS_ARM_Succeeded))
	gens["SoftDeletePeriod"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_STATUS_ARMGenerator())
	gens["Statistics"] = gen.PtrOf(DatabaseStatistics_STATUS_ARMGenerator())
	gens["SuspensionDetails"] = gen.PtrOf(SuspensionDetails_STATUS_ARMGenerator())
}

func Test_ReadWriteDatabase_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabase_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabase_STATUS_ARM, ReadWriteDatabase_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabase_STATUS_ARM runs a test to see if a specific instance of ReadWriteDatabase_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabase_STATUS_ARM(subject ReadWriteDatabase_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabase_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ReadWriteDatabase_STATUS_ARM instances for property testing - lazily instantiated by
// ReadWriteDatabase_STATUS_ARMGenerator()
var readWriteDatabase_STATUS_ARMGenerator gopter.Gen

// ReadWriteDatabase_STATUS_ARMGenerator returns a generator of ReadWriteDatabase_STATUS_ARM instances for property testing.
// We first initialize readWriteDatabase_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabase_STATUS_ARMGenerator() gopter.Gen {
	if readWriteDatabase_STATUS_ARMGenerator != nil {
		return readWriteDatabase_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS_ARM(generators)
	readWriteDatabase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS_ARM(generators)
	readWriteDatabase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase_STATUS_ARM{}), generators)

	return readWriteDatabase_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.OneConstOf(ReadWriteDatabase_Kind_STATUS_ARM_ReadWrite)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReadWriteDatabaseProperties_STATUS_ARMGenerator())
}

func Test_SuspensionDetails_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SuspensionDetails_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSuspensionDetails_STATUS_ARM, SuspensionDetails_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSuspensionDetails_STATUS_ARM runs a test to see if a specific instance of SuspensionDetails_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSuspensionDetails_STATUS_ARM(subject SuspensionDetails_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SuspensionDetails_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SuspensionDetails_STATUS_ARM instances for property testing - lazily instantiated by
// SuspensionDetails_STATUS_ARMGenerator()
var suspensionDetails_STATUS_ARMGenerator gopter.Gen

// SuspensionDetails_STATUS_ARMGenerator returns a generator of SuspensionDetails_STATUS_ARM instances for property testing.
func SuspensionDetails_STATUS_ARMGenerator() gopter.Gen {
	if suspensionDetails_STATUS_ARMGenerator != nil {
		return suspensionDetails_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS_ARM(generators)
	suspensionDetails_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SuspensionDetails_STATUS_ARM{}), generators)

	return suspensionDetails_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SuspensionStartDate"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableLevelSharingProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableLevelSharingProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableLevelSharingProperties_STATUS_ARM, TableLevelSharingProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableLevelSharingProperties_STATUS_ARM runs a test to see if a specific instance of TableLevelSharingProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTableLevelSharingProperties_STATUS_ARM(subject TableLevelSharingProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableLevelSharingProperties_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of TableLevelSharingProperties_STATUS_ARM instances for property testing - lazily instantiated by
// TableLevelSharingProperties_STATUS_ARMGenerator()
var tableLevelSharingProperties_STATUS_ARMGenerator gopter.Gen

// TableLevelSharingProperties_STATUS_ARMGenerator returns a generator of TableLevelSharingProperties_STATUS_ARM instances for property testing.
func TableLevelSharingProperties_STATUS_ARMGenerator() gopter.Gen {
	if tableLevelSharingProperties_STATUS_ARMGenerator != nil {
		return tableLevelSharingProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableLevelSharingProperties_STATUS_ARM(generators)
	tableLevelSharingProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TableLevelSharingProperties_STATUS_ARM{}), generators)

	return tableLevelSharingProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTableLevelSharingProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableLevelSharingProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ExternalTablesToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["ExternalTablesToInclude"] = gen.SliceOf(gen.AlphaString())
	gens["FunctionsToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["FunctionsToInclude"] = gen.SliceOf(gen.AlphaString())
	gens["MaterializedViewsToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["MaterializedViewsToInclude"] = gen.SliceOf(gen.AlphaString())
	gens["TablesToExclude"] = gen.SliceOf(gen.AlphaString())
	gens["TablesToInclude"] = gen.SliceOf(gen.AlphaString())
}
