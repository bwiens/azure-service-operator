// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701storage

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

func Test_DnsForwardingRuleset_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleset via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleset, DnsForwardingRulesetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleset runs a test to see if a specific instance of DnsForwardingRuleset round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleset(subject DnsForwardingRuleset) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleset
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

// Generator of DnsForwardingRuleset instances for property testing - lazily instantiated by
// DnsForwardingRulesetGenerator()
var dnsForwardingRulesetGenerator gopter.Gen

// DnsForwardingRulesetGenerator returns a generator of DnsForwardingRuleset instances for property testing.
func DnsForwardingRulesetGenerator() gopter.Gen {
	if dnsForwardingRulesetGenerator != nil {
		return dnsForwardingRulesetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleset(generators)
	dnsForwardingRulesetGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset{}), generators)

	return dnsForwardingRulesetGenerator
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleset is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleset(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsForwardingRuleset_SpecGenerator()
	gens["Status"] = DnsForwardingRuleset_STATUSGenerator()
}

func Test_DnsForwardingRuleset_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleset_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleset_Spec, DnsForwardingRuleset_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleset_Spec runs a test to see if a specific instance of DnsForwardingRuleset_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleset_Spec(subject DnsForwardingRuleset_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleset_Spec
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

// Generator of DnsForwardingRuleset_Spec instances for property testing - lazily instantiated by
// DnsForwardingRuleset_SpecGenerator()
var dnsForwardingRuleset_SpecGenerator gopter.Gen

// DnsForwardingRuleset_SpecGenerator returns a generator of DnsForwardingRuleset_Spec instances for property testing.
// We first initialize dnsForwardingRuleset_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRuleset_SpecGenerator() gopter.Gen {
	if dnsForwardingRuleset_SpecGenerator != nil {
		return dnsForwardingRuleset_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_Spec(generators)
	dnsForwardingRuleset_SpecGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleset_Spec(generators)
	dnsForwardingRuleset_SpecGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_Spec{}), generators)

	return dnsForwardingRuleset_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRuleset_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRuleset_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleset_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleset_Spec(gens map[string]gopter.Gen) {
	gens["DnsResolverOutboundEndpoints"] = gen.SliceOf(DnsresolverSubResourceGenerator())
}

func Test_DnsForwardingRuleset_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleset_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleset_STATUS, DnsForwardingRuleset_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleset_STATUS runs a test to see if a specific instance of DnsForwardingRuleset_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleset_STATUS(subject DnsForwardingRuleset_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleset_STATUS
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

// Generator of DnsForwardingRuleset_STATUS instances for property testing - lazily instantiated by
// DnsForwardingRuleset_STATUSGenerator()
var dnsForwardingRuleset_STATUSGenerator gopter.Gen

// DnsForwardingRuleset_STATUSGenerator returns a generator of DnsForwardingRuleset_STATUS instances for property testing.
// We first initialize dnsForwardingRuleset_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRuleset_STATUSGenerator() gopter.Gen {
	if dnsForwardingRuleset_STATUSGenerator != nil {
		return dnsForwardingRuleset_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS(generators)
	dnsForwardingRuleset_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS(generators)
	dnsForwardingRuleset_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleset_STATUS{}), generators)

	return dnsForwardingRuleset_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRuleset_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleset_STATUS(gens map[string]gopter.Gen) {
	gens["DnsResolverOutboundEndpoints"] = gen.SliceOf(DnsresolverSubResource_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_DnsresolverSubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsresolverSubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsresolverSubResource, DnsresolverSubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsresolverSubResource runs a test to see if a specific instance of DnsresolverSubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsresolverSubResource(subject DnsresolverSubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsresolverSubResource
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

// Generator of DnsresolverSubResource instances for property testing - lazily instantiated by
// DnsresolverSubResourceGenerator()
var dnsresolverSubResourceGenerator gopter.Gen

// DnsresolverSubResourceGenerator returns a generator of DnsresolverSubResource instances for property testing.
func DnsresolverSubResourceGenerator() gopter.Gen {
	if dnsresolverSubResourceGenerator != nil {
		return dnsresolverSubResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	dnsresolverSubResourceGenerator = gen.Struct(reflect.TypeOf(DnsresolverSubResource{}), generators)

	return dnsresolverSubResourceGenerator
}

func Test_DnsresolverSubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsresolverSubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsresolverSubResource_STATUS, DnsresolverSubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsresolverSubResource_STATUS runs a test to see if a specific instance of DnsresolverSubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsresolverSubResource_STATUS(subject DnsresolverSubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsresolverSubResource_STATUS
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

// Generator of DnsresolverSubResource_STATUS instances for property testing - lazily instantiated by
// DnsresolverSubResource_STATUSGenerator()
var dnsresolverSubResource_STATUSGenerator gopter.Gen

// DnsresolverSubResource_STATUSGenerator returns a generator of DnsresolverSubResource_STATUS instances for property testing.
func DnsresolverSubResource_STATUSGenerator() gopter.Gen {
	if dnsresolverSubResource_STATUSGenerator != nil {
		return dnsresolverSubResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS(generators)
	dnsresolverSubResource_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsresolverSubResource_STATUS{}), generators)

	return dnsresolverSubResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsresolverSubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
