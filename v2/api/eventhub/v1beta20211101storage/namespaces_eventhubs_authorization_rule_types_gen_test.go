// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"encoding/json"
	v1api20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101storage"
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

func Test_NamespacesEventhubsAuthorizationRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsAuthorizationRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhubsAuthorizationRule, NamespacesEventhubsAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhubsAuthorizationRule tests if a specific instance of NamespacesEventhubsAuthorizationRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhubsAuthorizationRule(subject NamespacesEventhubsAuthorizationRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20211101s.NamespacesEventhubsAuthorizationRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhubsAuthorizationRule
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

func Test_NamespacesEventhubsAuthorizationRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsAuthorizationRule to NamespacesEventhubsAuthorizationRule via AssignProperties_To_NamespacesEventhubsAuthorizationRule & AssignProperties_From_NamespacesEventhubsAuthorizationRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRule, NamespacesEventhubsAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRule tests if a specific instance of NamespacesEventhubsAuthorizationRule can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsAuthorizationRule(subject NamespacesEventhubsAuthorizationRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211101s.NamespacesEventhubsAuthorizationRule
	err := copied.AssignProperties_To_NamespacesEventhubsAuthorizationRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsAuthorizationRule
	err = actual.AssignProperties_From_NamespacesEventhubsAuthorizationRule(&other)
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

func Test_NamespacesEventhubsAuthorizationRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsAuthorizationRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule, NamespacesEventhubsAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule runs a test to see if a specific instance of NamespacesEventhubsAuthorizationRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule(subject NamespacesEventhubsAuthorizationRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsAuthorizationRule
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

// Generator of NamespacesEventhubsAuthorizationRule instances for property testing - lazily instantiated by
// NamespacesEventhubsAuthorizationRuleGenerator()
var namespacesEventhubsAuthorizationRuleGenerator gopter.Gen

// NamespacesEventhubsAuthorizationRuleGenerator returns a generator of NamespacesEventhubsAuthorizationRule instances for property testing.
func NamespacesEventhubsAuthorizationRuleGenerator() gopter.Gen {
	if namespacesEventhubsAuthorizationRuleGenerator != nil {
		return namespacesEventhubsAuthorizationRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule(generators)
	namespacesEventhubsAuthorizationRuleGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsAuthorizationRule{}), generators)

	return namespacesEventhubsAuthorizationRuleGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule(gens map[string]gopter.Gen) {
	gens["Spec"] = Namespaces_Eventhubs_AuthorizationRule_SpecGenerator()
	gens["Status"] = Namespaces_Eventhubs_AuthorizationRule_STATUSGenerator()
}

func Test_Namespaces_Eventhubs_AuthorizationRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Eventhubs_AuthorizationRule_Spec to Namespaces_Eventhubs_AuthorizationRule_Spec via AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec & AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Eventhubs_AuthorizationRule_Spec, Namespaces_Eventhubs_AuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Eventhubs_AuthorizationRule_Spec tests if a specific instance of Namespaces_Eventhubs_AuthorizationRule_Spec can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Eventhubs_AuthorizationRule_Spec(subject Namespaces_Eventhubs_AuthorizationRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec
	err := copied.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Eventhubs_AuthorizationRule_Spec
	err = actual.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec(&other)
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

func Test_Namespaces_Eventhubs_AuthorizationRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_AuthorizationRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Spec, Namespaces_Eventhubs_AuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Spec runs a test to see if a specific instance of Namespaces_Eventhubs_AuthorizationRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Spec(subject Namespaces_Eventhubs_AuthorizationRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_AuthorizationRule_Spec
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

// Generator of Namespaces_Eventhubs_AuthorizationRule_Spec instances for property testing - lazily instantiated by
// Namespaces_Eventhubs_AuthorizationRule_SpecGenerator()
var namespaces_Eventhubs_AuthorizationRule_SpecGenerator gopter.Gen

// Namespaces_Eventhubs_AuthorizationRule_SpecGenerator returns a generator of Namespaces_Eventhubs_AuthorizationRule_Spec instances for property testing.
func Namespaces_Eventhubs_AuthorizationRule_SpecGenerator() gopter.Gen {
	if namespaces_Eventhubs_AuthorizationRule_SpecGenerator != nil {
		return namespaces_Eventhubs_AuthorizationRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Spec(generators)
	namespaces_Eventhubs_AuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_AuthorizationRule_Spec{}), generators)

	return namespaces_Eventhubs_AuthorizationRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
}

func Test_Namespaces_Eventhubs_AuthorizationRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Namespaces_Eventhubs_AuthorizationRule_STATUS to Namespaces_Eventhubs_AuthorizationRule_STATUS via AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS & AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespaces_Eventhubs_AuthorizationRule_STATUS, Namespaces_Eventhubs_AuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespaces_Eventhubs_AuthorizationRule_STATUS tests if a specific instance of Namespaces_Eventhubs_AuthorizationRule_STATUS can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespaces_Eventhubs_AuthorizationRule_STATUS(subject Namespaces_Eventhubs_AuthorizationRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS
	err := copied.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Namespaces_Eventhubs_AuthorizationRule_STATUS
	err = actual.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS(&other)
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

func Test_Namespaces_Eventhubs_AuthorizationRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_AuthorizationRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_STATUS, Namespaces_Eventhubs_AuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_STATUS runs a test to see if a specific instance of Namespaces_Eventhubs_AuthorizationRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_STATUS(subject Namespaces_Eventhubs_AuthorizationRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_AuthorizationRule_STATUS
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

// Generator of Namespaces_Eventhubs_AuthorizationRule_STATUS instances for property testing - lazily instantiated by
// Namespaces_Eventhubs_AuthorizationRule_STATUSGenerator()
var namespaces_Eventhubs_AuthorizationRule_STATUSGenerator gopter.Gen

// Namespaces_Eventhubs_AuthorizationRule_STATUSGenerator returns a generator of Namespaces_Eventhubs_AuthorizationRule_STATUS instances for property testing.
// We first initialize namespaces_Eventhubs_AuthorizationRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhubs_AuthorizationRule_STATUSGenerator() gopter.Gen {
	if namespaces_Eventhubs_AuthorizationRule_STATUSGenerator != nil {
		return namespaces_Eventhubs_AuthorizationRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_STATUS(generators)
	namespaces_Eventhubs_AuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_AuthorizationRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_STATUS(generators)
	namespaces_Eventhubs_AuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_AuthorizationRule_STATUS{}), generators)

	return namespaces_Eventhubs_AuthorizationRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Rights"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}
