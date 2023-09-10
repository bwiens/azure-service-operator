// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

import (
	"encoding/json"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
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

func Test_LoadBalancersInboundNatRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from LoadBalancersInboundNatRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForLoadBalancersInboundNatRule, LoadBalancersInboundNatRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForLoadBalancersInboundNatRule tests if a specific instance of LoadBalancersInboundNatRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForLoadBalancersInboundNatRule(subject LoadBalancersInboundNatRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.LoadBalancersInboundNatRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual LoadBalancersInboundNatRule
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_LoadBalancersInboundNatRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from LoadBalancersInboundNatRule to LoadBalancersInboundNatRule via AssignProperties_To_LoadBalancersInboundNatRule & AssignProperties_From_LoadBalancersInboundNatRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForLoadBalancersInboundNatRule, LoadBalancersInboundNatRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForLoadBalancersInboundNatRule tests if a specific instance of LoadBalancersInboundNatRule can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForLoadBalancersInboundNatRule(subject LoadBalancersInboundNatRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.LoadBalancersInboundNatRule
	err := copied.AssignProperties_To_LoadBalancersInboundNatRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual LoadBalancersInboundNatRule
	err = actual.AssignProperties_From_LoadBalancersInboundNatRule(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_LoadBalancersInboundNatRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancersInboundNatRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancersInboundNatRule, LoadBalancersInboundNatRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancersInboundNatRule runs a test to see if a specific instance of LoadBalancersInboundNatRule round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancersInboundNatRule(subject LoadBalancersInboundNatRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancersInboundNatRule
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

// Generator of LoadBalancersInboundNatRule instances for property testing - lazily instantiated by
// LoadBalancersInboundNatRuleGenerator()
var loadBalancersInboundNatRuleGenerator gopter.Gen

// LoadBalancersInboundNatRuleGenerator returns a generator of LoadBalancersInboundNatRule instances for property testing.
func LoadBalancersInboundNatRuleGenerator() gopter.Gen {
	if loadBalancersInboundNatRuleGenerator != nil {
		return loadBalancersInboundNatRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule(generators)
	loadBalancersInboundNatRuleGenerator = gen.Struct(reflect.TypeOf(LoadBalancersInboundNatRule{}), generators)

	return loadBalancersInboundNatRuleGenerator
}

// AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule(gens map[string]gopter.Gen) {
	gens["Spec"] = LoadBalancers_InboundNatRule_SpecGenerator()
	gens["Status"] = LoadBalancers_InboundNatRule_STATUSGenerator()
}

func Test_LoadBalancers_InboundNatRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from LoadBalancers_InboundNatRule_Spec to LoadBalancers_InboundNatRule_Spec via AssignProperties_To_LoadBalancers_InboundNatRule_Spec & AssignProperties_From_LoadBalancers_InboundNatRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForLoadBalancers_InboundNatRule_Spec, LoadBalancers_InboundNatRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForLoadBalancers_InboundNatRule_Spec tests if a specific instance of LoadBalancers_InboundNatRule_Spec can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForLoadBalancers_InboundNatRule_Spec(subject LoadBalancers_InboundNatRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.LoadBalancers_InboundNatRule_Spec
	err := copied.AssignProperties_To_LoadBalancers_InboundNatRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual LoadBalancers_InboundNatRule_Spec
	err = actual.AssignProperties_From_LoadBalancers_InboundNatRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_LoadBalancers_InboundNatRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancers_InboundNatRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancers_InboundNatRule_Spec, LoadBalancers_InboundNatRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancers_InboundNatRule_Spec runs a test to see if a specific instance of LoadBalancers_InboundNatRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancers_InboundNatRule_Spec(subject LoadBalancers_InboundNatRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancers_InboundNatRule_Spec
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

// Generator of LoadBalancers_InboundNatRule_Spec instances for property testing - lazily instantiated by
// LoadBalancers_InboundNatRule_SpecGenerator()
var loadBalancers_InboundNatRule_SpecGenerator gopter.Gen

// LoadBalancers_InboundNatRule_SpecGenerator returns a generator of LoadBalancers_InboundNatRule_Spec instances for property testing.
// We first initialize loadBalancers_InboundNatRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LoadBalancers_InboundNatRule_SpecGenerator() gopter.Gen {
	if loadBalancers_InboundNatRule_SpecGenerator != nil {
		return loadBalancers_InboundNatRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_Spec(generators)
	loadBalancers_InboundNatRule_SpecGenerator = gen.Struct(reflect.TypeOf(LoadBalancers_InboundNatRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_Spec(generators)
	AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_Spec(generators)
	loadBalancers_InboundNatRule_SpecGenerator = gen.Struct(reflect.TypeOf(LoadBalancers_InboundNatRule_Spec{}), generators)

	return loadBalancers_InboundNatRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["BackendPort"] = gen.PtrOf(gen.Int())
	gens["EnableFloatingIP"] = gen.PtrOf(gen.Bool())
	gens["EnableTcpReset"] = gen.PtrOf(gen.Bool())
	gens["FrontendPort"] = gen.PtrOf(gen.Int())
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(TransportProtocol_All, TransportProtocol_Tcp, TransportProtocol_Udp))
}

// AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_Spec(gens map[string]gopter.Gen) {
	gens["FrontendIPConfiguration"] = gen.PtrOf(SubResourceGenerator())
}

func Test_LoadBalancers_InboundNatRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from LoadBalancers_InboundNatRule_STATUS to LoadBalancers_InboundNatRule_STATUS via AssignProperties_To_LoadBalancers_InboundNatRule_STATUS & AssignProperties_From_LoadBalancers_InboundNatRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForLoadBalancers_InboundNatRule_STATUS, LoadBalancers_InboundNatRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForLoadBalancers_InboundNatRule_STATUS tests if a specific instance of LoadBalancers_InboundNatRule_STATUS can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForLoadBalancers_InboundNatRule_STATUS(subject LoadBalancers_InboundNatRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.LoadBalancers_InboundNatRule_STATUS
	err := copied.AssignProperties_To_LoadBalancers_InboundNatRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual LoadBalancers_InboundNatRule_STATUS
	err = actual.AssignProperties_From_LoadBalancers_InboundNatRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_LoadBalancers_InboundNatRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancers_InboundNatRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancers_InboundNatRule_STATUS, LoadBalancers_InboundNatRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancers_InboundNatRule_STATUS runs a test to see if a specific instance of LoadBalancers_InboundNatRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancers_InboundNatRule_STATUS(subject LoadBalancers_InboundNatRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancers_InboundNatRule_STATUS
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

// Generator of LoadBalancers_InboundNatRule_STATUS instances for property testing - lazily instantiated by
// LoadBalancers_InboundNatRule_STATUSGenerator()
var loadBalancers_InboundNatRule_STATUSGenerator gopter.Gen

// LoadBalancers_InboundNatRule_STATUSGenerator returns a generator of LoadBalancers_InboundNatRule_STATUS instances for property testing.
// We first initialize loadBalancers_InboundNatRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LoadBalancers_InboundNatRule_STATUSGenerator() gopter.Gen {
	if loadBalancers_InboundNatRule_STATUSGenerator != nil {
		return loadBalancers_InboundNatRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS(generators)
	loadBalancers_InboundNatRule_STATUSGenerator = gen.Struct(reflect.TypeOf(LoadBalancers_InboundNatRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS(generators)
	loadBalancers_InboundNatRule_STATUSGenerator = gen.Struct(reflect.TypeOf(LoadBalancers_InboundNatRule_STATUS{}), generators)

	return loadBalancers_InboundNatRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS(gens map[string]gopter.Gen) {
	gens["BackendPort"] = gen.PtrOf(gen.Int())
	gens["EnableFloatingIP"] = gen.PtrOf(gen.Bool())
	gens["EnableTcpReset"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["FrontendPort"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(TransportProtocol_STATUS_All, TransportProtocol_STATUS_Tcp, TransportProtocol_STATUS_Udp))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLoadBalancers_InboundNatRule_STATUS(gens map[string]gopter.Gen) {
	gens["BackendIPConfiguration"] = gen.PtrOf(NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator())
	gens["FrontendIPConfiguration"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded to NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded via AssignProperties_To_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded & AssignProperties_From_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded, NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded tests if a specific instance of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(subject NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded
	err := copied.AssignProperties_To_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded
	err = actual.AssignProperties_From_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded, NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded runs a test to see if a specific instance of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(subject NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded
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

// Generator of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator()
var networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator gopter.Gen

// NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator returns a generator of NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded instances for property testing.
func NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator != nil {
		return networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(generators)
	networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded{}), generators)

	return networkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubResource to SubResource via AssignProperties_To_SubResource & AssignProperties_From_SubResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubResource tests if a specific instance of SubResource can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForSubResource(subject SubResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.SubResource
	err := copied.AssignProperties_To_SubResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubResource
	err = actual.AssignProperties_From_SubResource(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource runs a test to see if a specific instance of SubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource(subject SubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource
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

// Generator of SubResource instances for property testing - lazily instantiated by SubResourceGenerator()
var subResourceGenerator gopter.Gen

// SubResourceGenerator returns a generator of SubResource instances for property testing.
func SubResourceGenerator() gopter.Gen {
	if subResourceGenerator != nil {
		return subResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	subResourceGenerator = gen.Struct(reflect.TypeOf(SubResource{}), generators)

	return subResourceGenerator
}

func Test_SubResource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubResource_STATUS to SubResource_STATUS via AssignProperties_To_SubResource_STATUS & AssignProperties_From_SubResource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubResource_STATUS tests if a specific instance of SubResource_STATUS can be assigned to v1api20201101storage and back losslessly
func RunPropertyAssignmentTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.SubResource_STATUS
	err := copied.AssignProperties_To_SubResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubResource_STATUS
	err = actual.AssignProperties_From_SubResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_STATUS runs a test to see if a specific instance of SubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUS
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

// Generator of SubResource_STATUS instances for property testing - lazily instantiated by SubResource_STATUSGenerator()
var subResource_STATUSGenerator gopter.Gen

// SubResource_STATUSGenerator returns a generator of SubResource_STATUS instances for property testing.
func SubResource_STATUSGenerator() gopter.Gen {
	if subResource_STATUSGenerator != nil {
		return subResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_STATUS(generators)
	subResource_STATUSGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUS{}), generators)

	return subResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}