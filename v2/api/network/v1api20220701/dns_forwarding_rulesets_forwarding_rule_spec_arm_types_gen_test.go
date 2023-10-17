// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_DnsForwardingRulesets_ForwardingRule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRulesets_ForwardingRule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRulesets_ForwardingRule_Spec_ARM, DnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRulesets_ForwardingRule_Spec_ARM runs a test to see if a specific instance of DnsForwardingRulesets_ForwardingRule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRulesets_ForwardingRule_Spec_ARM(subject DnsForwardingRulesets_ForwardingRule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRulesets_ForwardingRule_Spec_ARM
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

// Generator of DnsForwardingRulesets_ForwardingRule_Spec_ARM instances for property testing - lazily instantiated by
// DnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator()
var dnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator gopter.Gen

// DnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator returns a generator of DnsForwardingRulesets_ForwardingRule_Spec_ARM instances for property testing.
// We first initialize dnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator() gopter.Gen {
	if dnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator != nil {
		return dnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesets_ForwardingRule_Spec_ARM(generators)
	dnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesets_ForwardingRule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRulesets_ForwardingRule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRulesets_ForwardingRule_Spec_ARM(generators)
	dnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRulesets_ForwardingRule_Spec_ARM{}), generators)

	return dnsForwardingRulesets_ForwardingRule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRulesets_ForwardingRule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRulesets_ForwardingRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDnsForwardingRulesets_ForwardingRule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRulesets_ForwardingRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ForwardingRuleProperties_ARMGenerator())
}

func Test_ForwardingRuleProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ForwardingRuleProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForForwardingRuleProperties_ARM, ForwardingRuleProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForForwardingRuleProperties_ARM runs a test to see if a specific instance of ForwardingRuleProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForForwardingRuleProperties_ARM(subject ForwardingRuleProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ForwardingRuleProperties_ARM
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

// Generator of ForwardingRuleProperties_ARM instances for property testing - lazily instantiated by
// ForwardingRuleProperties_ARMGenerator()
var forwardingRuleProperties_ARMGenerator gopter.Gen

// ForwardingRuleProperties_ARMGenerator returns a generator of ForwardingRuleProperties_ARM instances for property testing.
// We first initialize forwardingRuleProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ForwardingRuleProperties_ARMGenerator() gopter.Gen {
	if forwardingRuleProperties_ARMGenerator != nil {
		return forwardingRuleProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForForwardingRuleProperties_ARM(generators)
	forwardingRuleProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ForwardingRuleProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForForwardingRuleProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForForwardingRuleProperties_ARM(generators)
	forwardingRuleProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ForwardingRuleProperties_ARM{}), generators)

	return forwardingRuleProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForForwardingRuleProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForForwardingRuleProperties_ARM(gens map[string]gopter.Gen) {
	gens["DomainName"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardingRuleState"] = gen.PtrOf(gen.OneConstOf(ForwardingRuleProperties_ForwardingRuleState_Disabled, ForwardingRuleProperties_ForwardingRuleState_Enabled))
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForForwardingRuleProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForForwardingRuleProperties_ARM(gens map[string]gopter.Gen) {
	gens["TargetDnsServers"] = gen.SliceOf(TargetDnsServer_ARMGenerator())
}

func Test_TargetDnsServer_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TargetDnsServer_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTargetDnsServer_ARM, TargetDnsServer_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTargetDnsServer_ARM runs a test to see if a specific instance of TargetDnsServer_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTargetDnsServer_ARM(subject TargetDnsServer_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TargetDnsServer_ARM
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

// Generator of TargetDnsServer_ARM instances for property testing - lazily instantiated by
// TargetDnsServer_ARMGenerator()
var targetDnsServer_ARMGenerator gopter.Gen

// TargetDnsServer_ARMGenerator returns a generator of TargetDnsServer_ARM instances for property testing.
func TargetDnsServer_ARMGenerator() gopter.Gen {
	if targetDnsServer_ARMGenerator != nil {
		return targetDnsServer_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTargetDnsServer_ARM(generators)
	targetDnsServer_ARMGenerator = gen.Struct(reflect.TypeOf(TargetDnsServer_ARM{}), generators)

	return targetDnsServer_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTargetDnsServer_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTargetDnsServer_ARM(gens map[string]gopter.Gen) {
	gens["IpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Port"] = gen.PtrOf(gen.Int())
}