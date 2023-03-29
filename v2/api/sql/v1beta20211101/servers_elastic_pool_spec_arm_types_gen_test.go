// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_Servers_ElasticPool_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_ElasticPool_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_ElasticPool_Spec_ARM, Servers_ElasticPool_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_ElasticPool_Spec_ARM runs a test to see if a specific instance of Servers_ElasticPool_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_ElasticPool_Spec_ARM(subject Servers_ElasticPool_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_ElasticPool_Spec_ARM
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

// Generator of Servers_ElasticPool_Spec_ARM instances for property testing - lazily instantiated by
// Servers_ElasticPool_Spec_ARMGenerator()
var servers_ElasticPool_Spec_ARMGenerator gopter.Gen

// Servers_ElasticPool_Spec_ARMGenerator returns a generator of Servers_ElasticPool_Spec_ARM instances for property testing.
// We first initialize servers_ElasticPool_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_ElasticPool_Spec_ARMGenerator() gopter.Gen {
	if servers_ElasticPool_Spec_ARMGenerator != nil {
		return servers_ElasticPool_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec_ARM(generators)
	servers_ElasticPool_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec_ARM(generators)
	servers_ElasticPool_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Servers_ElasticPool_Spec_ARM{}), generators)

	return servers_ElasticPool_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_ElasticPool_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_ElasticPool_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ElasticPoolProperties_ARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
}

func Test_ElasticPoolProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolProperties_ARM, ElasticPoolProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolProperties_ARM runs a test to see if a specific instance of ElasticPoolProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolProperties_ARM(subject ElasticPoolProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolProperties_ARM
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

// Generator of ElasticPoolProperties_ARM instances for property testing - lazily instantiated by
// ElasticPoolProperties_ARMGenerator()
var elasticPoolProperties_ARMGenerator gopter.Gen

// ElasticPoolProperties_ARMGenerator returns a generator of ElasticPoolProperties_ARM instances for property testing.
// We first initialize elasticPoolProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ElasticPoolProperties_ARMGenerator() gopter.Gen {
	if elasticPoolProperties_ARMGenerator != nil {
		return elasticPoolProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolProperties_ARM(generators)
	elasticPoolProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ElasticPoolProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForElasticPoolProperties_ARM(generators)
	elasticPoolProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ElasticPoolProperties_ARM{}), generators)

	return elasticPoolProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolProperties_ARM(gens map[string]gopter.Gen) {
	gens["HighAvailabilityReplicaCount"] = gen.PtrOf(gen.Int())
	gens["LicenseType"] = gen.PtrOf(gen.OneConstOf(ElasticPoolProperties_LicenseType_BasePrice, ElasticPoolProperties_LicenseType_LicenseIncluded))
	gens["MaintenanceConfigurationId"] = gen.PtrOf(gen.AlphaString())
	gens["MaxSizeBytes"] = gen.PtrOf(gen.Int())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
	gens["ZoneRedundant"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForElasticPoolProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForElasticPoolProperties_ARM(gens map[string]gopter.Gen) {
	gens["PerDatabaseSettings"] = gen.PtrOf(ElasticPoolPerDatabaseSettings_ARMGenerator())
}

func Test_ElasticPoolPerDatabaseSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ElasticPoolPerDatabaseSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForElasticPoolPerDatabaseSettings_ARM, ElasticPoolPerDatabaseSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForElasticPoolPerDatabaseSettings_ARM runs a test to see if a specific instance of ElasticPoolPerDatabaseSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForElasticPoolPerDatabaseSettings_ARM(subject ElasticPoolPerDatabaseSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ElasticPoolPerDatabaseSettings_ARM
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

// Generator of ElasticPoolPerDatabaseSettings_ARM instances for property testing - lazily instantiated by
// ElasticPoolPerDatabaseSettings_ARMGenerator()
var elasticPoolPerDatabaseSettings_ARMGenerator gopter.Gen

// ElasticPoolPerDatabaseSettings_ARMGenerator returns a generator of ElasticPoolPerDatabaseSettings_ARM instances for property testing.
func ElasticPoolPerDatabaseSettings_ARMGenerator() gopter.Gen {
	if elasticPoolPerDatabaseSettings_ARMGenerator != nil {
		return elasticPoolPerDatabaseSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_ARM(generators)
	elasticPoolPerDatabaseSettings_ARMGenerator = gen.Struct(reflect.TypeOf(ElasticPoolPerDatabaseSettings_ARM{}), generators)

	return elasticPoolPerDatabaseSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForElasticPoolPerDatabaseSettings_ARM(gens map[string]gopter.Gen) {
	gens["MaxCapacity"] = gen.PtrOf(gen.Float64())
	gens["MinCapacity"] = gen.PtrOf(gen.Float64())
}