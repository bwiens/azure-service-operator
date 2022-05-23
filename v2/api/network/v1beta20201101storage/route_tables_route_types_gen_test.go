// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

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

func Test_RouteTablesRoute_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablesRoute via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesRoute, RouteTablesRouteGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesRoute runs a test to see if a specific instance of RouteTablesRoute round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesRoute(subject RouteTablesRoute) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablesRoute
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

// Generator of RouteTablesRoute instances for property testing - lazily instantiated by RouteTablesRouteGenerator()
var routeTablesRouteGenerator gopter.Gen

// RouteTablesRouteGenerator returns a generator of RouteTablesRoute instances for property testing.
func RouteTablesRouteGenerator() gopter.Gen {
	if routeTablesRouteGenerator != nil {
		return routeTablesRouteGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRouteTablesRoute(generators)
	routeTablesRouteGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoute{}), generators)

	return routeTablesRouteGenerator
}

// AddRelatedPropertyGeneratorsForRouteTablesRoute is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTablesRoute(gens map[string]gopter.Gen) {
	gens["Spec"] = RouteTablesRoutesSpecGenerator()
	gens["Status"] = RouteStatusGenerator()
}

func Test_RouteTablesRoutes_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablesRoutes_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesRoutesSpec, RouteTablesRoutesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesRoutesSpec runs a test to see if a specific instance of RouteTablesRoutes_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesRoutesSpec(subject RouteTablesRoutes_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablesRoutes_Spec
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

// Generator of RouteTablesRoutes_Spec instances for property testing - lazily instantiated by
// RouteTablesRoutesSpecGenerator()
var routeTablesRoutesSpecGenerator gopter.Gen

// RouteTablesRoutesSpecGenerator returns a generator of RouteTablesRoutes_Spec instances for property testing.
func RouteTablesRoutesSpecGenerator() gopter.Gen {
	if routeTablesRoutesSpecGenerator != nil {
		return routeTablesRoutesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablesRoutesSpec(generators)
	routeTablesRoutesSpecGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoutes_Spec{}), generators)

	return routeTablesRoutesSpecGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablesRoutesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablesRoutesSpec(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["HasBgpOverride"] = gen.PtrOf(gen.Bool())
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_Route_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Route_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteStatus, RouteStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteStatus runs a test to see if a specific instance of Route_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteStatus(subject Route_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Route_Status
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

// Generator of Route_Status instances for property testing - lazily instantiated by RouteStatusGenerator()
var routeStatusGenerator gopter.Gen

// RouteStatusGenerator returns a generator of Route_Status instances for property testing.
func RouteStatusGenerator() gopter.Gen {
	if routeStatusGenerator != nil {
		return routeStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteStatus(generators)
	routeStatusGenerator = gen.Struct(reflect.TypeOf(Route_Status{}), generators)

	return routeStatusGenerator
}

// AddIndependentPropertyGeneratorsForRouteStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteStatus(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["HasBgpOverride"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
