// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	v1api20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
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

func Test_ServersOutboundFirewallRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersOutboundFirewallRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersOutboundFirewallRule, ServersOutboundFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersOutboundFirewallRule tests if a specific instance of ServersOutboundFirewallRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersOutboundFirewallRule(subject ServersOutboundFirewallRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20211101s.ServersOutboundFirewallRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersOutboundFirewallRule
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

func Test_ServersOutboundFirewallRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersOutboundFirewallRule to ServersOutboundFirewallRule via AssignProperties_To_ServersOutboundFirewallRule & AssignProperties_From_ServersOutboundFirewallRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersOutboundFirewallRule, ServersOutboundFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersOutboundFirewallRule tests if a specific instance of ServersOutboundFirewallRule can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServersOutboundFirewallRule(subject ServersOutboundFirewallRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211101s.ServersOutboundFirewallRule
	err := copied.AssignProperties_To_ServersOutboundFirewallRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersOutboundFirewallRule
	err = actual.AssignProperties_From_ServersOutboundFirewallRule(&other)
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

func Test_ServersOutboundFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersOutboundFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersOutboundFirewallRule, ServersOutboundFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersOutboundFirewallRule runs a test to see if a specific instance of ServersOutboundFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForServersOutboundFirewallRule(subject ServersOutboundFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersOutboundFirewallRule
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

// Generator of ServersOutboundFirewallRule instances for property testing - lazily instantiated by
// ServersOutboundFirewallRuleGenerator()
var serversOutboundFirewallRuleGenerator gopter.Gen

// ServersOutboundFirewallRuleGenerator returns a generator of ServersOutboundFirewallRule instances for property testing.
func ServersOutboundFirewallRuleGenerator() gopter.Gen {
	if serversOutboundFirewallRuleGenerator != nil {
		return serversOutboundFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersOutboundFirewallRule(generators)
	serversOutboundFirewallRuleGenerator = gen.Struct(reflect.TypeOf(ServersOutboundFirewallRule{}), generators)

	return serversOutboundFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForServersOutboundFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersOutboundFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_OutboundFirewallRule_SpecGenerator()
	gens["Status"] = Servers_OutboundFirewallRule_STATUSGenerator()
}

func Test_Servers_OutboundFirewallRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_OutboundFirewallRule_Spec to Servers_OutboundFirewallRule_Spec via AssignProperties_To_Servers_OutboundFirewallRule_Spec & AssignProperties_From_Servers_OutboundFirewallRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_OutboundFirewallRule_Spec, Servers_OutboundFirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_OutboundFirewallRule_Spec tests if a specific instance of Servers_OutboundFirewallRule_Spec can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_OutboundFirewallRule_Spec(subject Servers_OutboundFirewallRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211101s.Servers_OutboundFirewallRule_Spec
	err := copied.AssignProperties_To_Servers_OutboundFirewallRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_OutboundFirewallRule_Spec
	err = actual.AssignProperties_From_Servers_OutboundFirewallRule_Spec(&other)
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

func Test_Servers_OutboundFirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_OutboundFirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_OutboundFirewallRule_Spec, Servers_OutboundFirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_OutboundFirewallRule_Spec runs a test to see if a specific instance of Servers_OutboundFirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_OutboundFirewallRule_Spec(subject Servers_OutboundFirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_OutboundFirewallRule_Spec
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

// Generator of Servers_OutboundFirewallRule_Spec instances for property testing - lazily instantiated by
// Servers_OutboundFirewallRule_SpecGenerator()
var servers_OutboundFirewallRule_SpecGenerator gopter.Gen

// Servers_OutboundFirewallRule_SpecGenerator returns a generator of Servers_OutboundFirewallRule_Spec instances for property testing.
func Servers_OutboundFirewallRule_SpecGenerator() gopter.Gen {
	if servers_OutboundFirewallRule_SpecGenerator != nil {
		return servers_OutboundFirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_OutboundFirewallRule_Spec(generators)
	servers_OutboundFirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_OutboundFirewallRule_Spec{}), generators)

	return servers_OutboundFirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_OutboundFirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_OutboundFirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

func Test_Servers_OutboundFirewallRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_OutboundFirewallRule_STATUS to Servers_OutboundFirewallRule_STATUS via AssignProperties_To_Servers_OutboundFirewallRule_STATUS & AssignProperties_From_Servers_OutboundFirewallRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_OutboundFirewallRule_STATUS, Servers_OutboundFirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_OutboundFirewallRule_STATUS tests if a specific instance of Servers_OutboundFirewallRule_STATUS can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_OutboundFirewallRule_STATUS(subject Servers_OutboundFirewallRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211101s.Servers_OutboundFirewallRule_STATUS
	err := copied.AssignProperties_To_Servers_OutboundFirewallRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_OutboundFirewallRule_STATUS
	err = actual.AssignProperties_From_Servers_OutboundFirewallRule_STATUS(&other)
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

func Test_Servers_OutboundFirewallRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_OutboundFirewallRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_OutboundFirewallRule_STATUS, Servers_OutboundFirewallRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_OutboundFirewallRule_STATUS runs a test to see if a specific instance of Servers_OutboundFirewallRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_OutboundFirewallRule_STATUS(subject Servers_OutboundFirewallRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_OutboundFirewallRule_STATUS
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

// Generator of Servers_OutboundFirewallRule_STATUS instances for property testing - lazily instantiated by
// Servers_OutboundFirewallRule_STATUSGenerator()
var servers_OutboundFirewallRule_STATUSGenerator gopter.Gen

// Servers_OutboundFirewallRule_STATUSGenerator returns a generator of Servers_OutboundFirewallRule_STATUS instances for property testing.
func Servers_OutboundFirewallRule_STATUSGenerator() gopter.Gen {
	if servers_OutboundFirewallRule_STATUSGenerator != nil {
		return servers_OutboundFirewallRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_OutboundFirewallRule_STATUS(generators)
	servers_OutboundFirewallRule_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_OutboundFirewallRule_STATUS{}), generators)

	return servers_OutboundFirewallRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_OutboundFirewallRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_OutboundFirewallRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
