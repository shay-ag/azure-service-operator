// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"encoding/json"
	v1api20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601storage"
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

func Test_PrivateDnsZonesCNAMERecord_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesCNAMERecord to hub returns original",
		prop.ForAll(RunResourceConversionTestForPrivateDnsZonesCNAMERecord, PrivateDnsZonesCNAMERecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForPrivateDnsZonesCNAMERecord tests if a specific instance of PrivateDnsZonesCNAMERecord round trips to the hub storage version and back losslessly
func RunResourceConversionTestForPrivateDnsZonesCNAMERecord(subject PrivateDnsZonesCNAMERecord) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20200601s.PrivateDnsZonesCNAMERecord
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual PrivateDnsZonesCNAMERecord
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

func Test_PrivateDnsZonesCNAMERecord_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZonesCNAMERecord to PrivateDnsZonesCNAMERecord via AssignProperties_To_PrivateDnsZonesCNAMERecord & AssignProperties_From_PrivateDnsZonesCNAMERecord returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZonesCNAMERecord, PrivateDnsZonesCNAMERecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZonesCNAMERecord tests if a specific instance of PrivateDnsZonesCNAMERecord can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZonesCNAMERecord(subject PrivateDnsZonesCNAMERecord) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20200601s.PrivateDnsZonesCNAMERecord
	err := copied.AssignProperties_To_PrivateDnsZonesCNAMERecord(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZonesCNAMERecord
	err = actual.AssignProperties_From_PrivateDnsZonesCNAMERecord(&other)
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

func Test_PrivateDnsZonesCNAMERecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesCNAMERecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesCNAMERecord, PrivateDnsZonesCNAMERecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesCNAMERecord runs a test to see if a specific instance of PrivateDnsZonesCNAMERecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesCNAMERecord(subject PrivateDnsZonesCNAMERecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesCNAMERecord
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

// Generator of PrivateDnsZonesCNAMERecord instances for property testing - lazily instantiated by
// PrivateDnsZonesCNAMERecordGenerator()
var privateDnsZonesCNAMERecordGenerator gopter.Gen

// PrivateDnsZonesCNAMERecordGenerator returns a generator of PrivateDnsZonesCNAMERecord instances for property testing.
func PrivateDnsZonesCNAMERecordGenerator() gopter.Gen {
	if privateDnsZonesCNAMERecordGenerator != nil {
		return privateDnsZonesCNAMERecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord(generators)
	privateDnsZonesCNAMERecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord{}), generators)

	return privateDnsZonesCNAMERecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZones_CNAME_SpecGenerator()
	gens["Status"] = PrivateDnsZones_CNAME_STATUSGenerator()
}

func Test_PrivateDnsZones_CNAME_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_CNAME_Spec to PrivateDnsZones_CNAME_Spec via AssignProperties_To_PrivateDnsZones_CNAME_Spec & AssignProperties_From_PrivateDnsZones_CNAME_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_CNAME_Spec, PrivateDnsZones_CNAME_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_CNAME_Spec tests if a specific instance of PrivateDnsZones_CNAME_Spec can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_CNAME_Spec(subject PrivateDnsZones_CNAME_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20200601s.PrivateDnsZones_CNAME_Spec
	err := copied.AssignProperties_To_PrivateDnsZones_CNAME_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_CNAME_Spec
	err = actual.AssignProperties_From_PrivateDnsZones_CNAME_Spec(&other)
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

func Test_PrivateDnsZones_CNAME_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_CNAME_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_CNAME_Spec, PrivateDnsZones_CNAME_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_CNAME_Spec runs a test to see if a specific instance of PrivateDnsZones_CNAME_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_CNAME_Spec(subject PrivateDnsZones_CNAME_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_CNAME_Spec
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

// Generator of PrivateDnsZones_CNAME_Spec instances for property testing - lazily instantiated by
// PrivateDnsZones_CNAME_SpecGenerator()
var privateDnsZones_CNAME_SpecGenerator gopter.Gen

// PrivateDnsZones_CNAME_SpecGenerator returns a generator of PrivateDnsZones_CNAME_Spec instances for property testing.
// We first initialize privateDnsZones_CNAME_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_CNAME_SpecGenerator() gopter.Gen {
	if privateDnsZones_CNAME_SpecGenerator != nil {
		return privateDnsZones_CNAME_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_Spec(generators)
	privateDnsZones_CNAME_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_CNAME_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_CNAME_Spec(generators)
	privateDnsZones_CNAME_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_CNAME_Spec{}), generators)

	return privateDnsZones_CNAME_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_CNAME_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_CNAME_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}

func Test_PrivateDnsZones_CNAME_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateDnsZones_CNAME_STATUS to PrivateDnsZones_CNAME_STATUS via AssignProperties_To_PrivateDnsZones_CNAME_STATUS & AssignProperties_From_PrivateDnsZones_CNAME_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateDnsZones_CNAME_STATUS, PrivateDnsZones_CNAME_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateDnsZones_CNAME_STATUS tests if a specific instance of PrivateDnsZones_CNAME_STATUS can be assigned to v1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateDnsZones_CNAME_STATUS(subject PrivateDnsZones_CNAME_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20200601s.PrivateDnsZones_CNAME_STATUS
	err := copied.AssignProperties_To_PrivateDnsZones_CNAME_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateDnsZones_CNAME_STATUS
	err = actual.AssignProperties_From_PrivateDnsZones_CNAME_STATUS(&other)
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

func Test_PrivateDnsZones_CNAME_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZones_CNAME_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZones_CNAME_STATUS, PrivateDnsZones_CNAME_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZones_CNAME_STATUS runs a test to see if a specific instance of PrivateDnsZones_CNAME_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZones_CNAME_STATUS(subject PrivateDnsZones_CNAME_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZones_CNAME_STATUS
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

// Generator of PrivateDnsZones_CNAME_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZones_CNAME_STATUSGenerator()
var privateDnsZones_CNAME_STATUSGenerator gopter.Gen

// PrivateDnsZones_CNAME_STATUSGenerator returns a generator of PrivateDnsZones_CNAME_STATUS instances for property testing.
// We first initialize privateDnsZones_CNAME_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZones_CNAME_STATUSGenerator() gopter.Gen {
	if privateDnsZones_CNAME_STATUSGenerator != nil {
		return privateDnsZones_CNAME_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_STATUS(generators)
	privateDnsZones_CNAME_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_CNAME_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZones_CNAME_STATUS(generators)
	privateDnsZones_CNAME_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZones_CNAME_STATUS{}), generators)

	return privateDnsZones_CNAME_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZones_CNAME_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAutoRegistered"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZones_CNAME_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZones_CNAME_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}
