// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901storage

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

func Test_StorageAccountsBlobServicesContainer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServicesContainer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesContainer, StorageAccountsBlobServicesContainerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesContainer runs a test to see if a specific instance of StorageAccountsBlobServicesContainer round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesContainer(subject StorageAccountsBlobServicesContainer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServicesContainer
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

// Generator of StorageAccountsBlobServicesContainer instances for property testing - lazily instantiated by
// StorageAccountsBlobServicesContainerGenerator()
var storageAccountsBlobServicesContainerGenerator gopter.Gen

// StorageAccountsBlobServicesContainerGenerator returns a generator of StorageAccountsBlobServicesContainer instances for property testing.
func StorageAccountsBlobServicesContainerGenerator() gopter.Gen {
	if storageAccountsBlobServicesContainerGenerator != nil {
		return storageAccountsBlobServicesContainerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer(generators)
	storageAccountsBlobServicesContainerGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainer{}), generators)

	return storageAccountsBlobServicesContainerGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccounts_BlobServices_Container_SpecGenerator()
	gens["Status"] = StorageAccounts_BlobServices_Container_STATUSGenerator()
}

func Test_StorageAccounts_BlobServices_Container_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_BlobServices_Container_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_BlobServices_Container_Spec, StorageAccounts_BlobServices_Container_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_BlobServices_Container_Spec runs a test to see if a specific instance of StorageAccounts_BlobServices_Container_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_BlobServices_Container_Spec(subject StorageAccounts_BlobServices_Container_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_BlobServices_Container_Spec
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

// Generator of StorageAccounts_BlobServices_Container_Spec instances for property testing - lazily instantiated by
// StorageAccounts_BlobServices_Container_SpecGenerator()
var storageAccounts_BlobServices_Container_SpecGenerator gopter.Gen

// StorageAccounts_BlobServices_Container_SpecGenerator returns a generator of StorageAccounts_BlobServices_Container_Spec instances for property testing.
// We first initialize storageAccounts_BlobServices_Container_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_BlobServices_Container_SpecGenerator() gopter.Gen {
	if storageAccounts_BlobServices_Container_SpecGenerator != nil {
		return storageAccounts_BlobServices_Container_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec(generators)
	storageAccounts_BlobServices_Container_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_BlobServices_Container_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec(generators)
	storageAccounts_BlobServices_Container_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_BlobServices_Container_Spec{}), generators)

	return storageAccounts_BlobServices_Container_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3AllSquash"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3RootSquash"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PublicAccess"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_Spec(gens map[string]gopter.Gen) {
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioningGenerator())
}

func Test_StorageAccounts_BlobServices_Container_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_BlobServices_Container_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_BlobServices_Container_STATUS, StorageAccounts_BlobServices_Container_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_BlobServices_Container_STATUS runs a test to see if a specific instance of StorageAccounts_BlobServices_Container_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_BlobServices_Container_STATUS(subject StorageAccounts_BlobServices_Container_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_BlobServices_Container_STATUS
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

// Generator of StorageAccounts_BlobServices_Container_STATUS instances for property testing - lazily instantiated by
// StorageAccounts_BlobServices_Container_STATUSGenerator()
var storageAccounts_BlobServices_Container_STATUSGenerator gopter.Gen

// StorageAccounts_BlobServices_Container_STATUSGenerator returns a generator of StorageAccounts_BlobServices_Container_STATUS instances for property testing.
// We first initialize storageAccounts_BlobServices_Container_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_BlobServices_Container_STATUSGenerator() gopter.Gen {
	if storageAccounts_BlobServices_Container_STATUSGenerator != nil {
		return storageAccounts_BlobServices_Container_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_STATUS(generators)
	storageAccounts_BlobServices_Container_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_BlobServices_Container_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_STATUS(generators)
	storageAccounts_BlobServices_Container_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_BlobServices_Container_STATUS{}), generators)

	return storageAccounts_BlobServices_Container_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_BlobServices_Container_STATUS(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3AllSquash"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3RootSquash"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["HasImmutabilityPolicy"] = gen.PtrOf(gen.Bool())
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseState"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.AlphaString())
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_BlobServices_Container_STATUS(gens map[string]gopter.Gen) {
	gens["ImmutabilityPolicy"] = gen.PtrOf(ImmutabilityPolicyProperties_STATUSGenerator())
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioning_STATUSGenerator())
	gens["LegalHold"] = gen.PtrOf(LegalHoldProperties_STATUSGenerator())
}

func Test_ImmutabilityPolicyProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyProperties_STATUS, ImmutabilityPolicyProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyProperties_STATUS runs a test to see if a specific instance of ImmutabilityPolicyProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyProperties_STATUS(subject ImmutabilityPolicyProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperties_STATUS
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

// Generator of ImmutabilityPolicyProperties_STATUS instances for property testing - lazily instantiated by
// ImmutabilityPolicyProperties_STATUSGenerator()
var immutabilityPolicyProperties_STATUSGenerator gopter.Gen

// ImmutabilityPolicyProperties_STATUSGenerator returns a generator of ImmutabilityPolicyProperties_STATUS instances for property testing.
// We first initialize immutabilityPolicyProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImmutabilityPolicyProperties_STATUSGenerator() gopter.Gen {
	if immutabilityPolicyProperties_STATUSGenerator != nil {
		return immutabilityPolicyProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(generators)
	immutabilityPolicyProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(generators)
	immutabilityPolicyProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUS{}), generators)

	return immutabilityPolicyProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["AllowProtectedAppendWritesAll"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(gens map[string]gopter.Gen) {
	gens["UpdateHistory"] = gen.SliceOf(UpdateHistoryProperty_STATUSGenerator())
}

func Test_ImmutableStorageWithVersioning_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning, ImmutableStorageWithVersioningGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning runs a test to see if a specific instance of ImmutableStorageWithVersioning round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning(subject ImmutableStorageWithVersioning) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning
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

// Generator of ImmutableStorageWithVersioning instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioningGenerator()
var immutableStorageWithVersioningGenerator gopter.Gen

// ImmutableStorageWithVersioningGenerator returns a generator of ImmutableStorageWithVersioning instances for property testing.
func ImmutableStorageWithVersioningGenerator() gopter.Gen {
	if immutableStorageWithVersioningGenerator != nil {
		return immutableStorageWithVersioningGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning(generators)
	immutableStorageWithVersioningGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning{}), generators)

	return immutableStorageWithVersioningGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_ImmutableStorageWithVersioning_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning_STATUS, ImmutableStorageWithVersioning_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning_STATUS runs a test to see if a specific instance of ImmutableStorageWithVersioning_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning_STATUS(subject ImmutableStorageWithVersioning_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_STATUS
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

// Generator of ImmutableStorageWithVersioning_STATUS instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioning_STATUSGenerator()
var immutableStorageWithVersioning_STATUSGenerator gopter.Gen

// ImmutableStorageWithVersioning_STATUSGenerator returns a generator of ImmutableStorageWithVersioning_STATUS instances for property testing.
func ImmutableStorageWithVersioning_STATUSGenerator() gopter.Gen {
	if immutableStorageWithVersioning_STATUSGenerator != nil {
		return immutableStorageWithVersioning_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUS(generators)
	immutableStorageWithVersioning_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_STATUS{}), generators)

	return immutableStorageWithVersioning_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["MigrationState"] = gen.PtrOf(gen.AlphaString())
	gens["TimeStamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_LegalHoldProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LegalHoldProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLegalHoldProperties_STATUS, LegalHoldProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLegalHoldProperties_STATUS runs a test to see if a specific instance of LegalHoldProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLegalHoldProperties_STATUS(subject LegalHoldProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LegalHoldProperties_STATUS
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

// Generator of LegalHoldProperties_STATUS instances for property testing - lazily instantiated by
// LegalHoldProperties_STATUSGenerator()
var legalHoldProperties_STATUSGenerator gopter.Gen

// LegalHoldProperties_STATUSGenerator returns a generator of LegalHoldProperties_STATUS instances for property testing.
// We first initialize legalHoldProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LegalHoldProperties_STATUSGenerator() gopter.Gen {
	if legalHoldProperties_STATUSGenerator != nil {
		return legalHoldProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS(generators)
	legalHoldProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUS(generators)
	legalHoldProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUS{}), generators)

	return legalHoldProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS(gens map[string]gopter.Gen) {
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProtectedAppendWritesHistory"] = gen.PtrOf(ProtectedAppendWritesHistory_STATUSGenerator())
	gens["Tags"] = gen.SliceOf(TagProperty_STATUSGenerator())
}

func Test_ProtectedAppendWritesHistory_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProtectedAppendWritesHistory_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProtectedAppendWritesHistory_STATUS, ProtectedAppendWritesHistory_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProtectedAppendWritesHistory_STATUS runs a test to see if a specific instance of ProtectedAppendWritesHistory_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProtectedAppendWritesHistory_STATUS(subject ProtectedAppendWritesHistory_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProtectedAppendWritesHistory_STATUS
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

// Generator of ProtectedAppendWritesHistory_STATUS instances for property testing - lazily instantiated by
// ProtectedAppendWritesHistory_STATUSGenerator()
var protectedAppendWritesHistory_STATUSGenerator gopter.Gen

// ProtectedAppendWritesHistory_STATUSGenerator returns a generator of ProtectedAppendWritesHistory_STATUS instances for property testing.
func ProtectedAppendWritesHistory_STATUSGenerator() gopter.Gen {
	if protectedAppendWritesHistory_STATUSGenerator != nil {
		return protectedAppendWritesHistory_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProtectedAppendWritesHistory_STATUS(generators)
	protectedAppendWritesHistory_STATUSGenerator = gen.Struct(reflect.TypeOf(ProtectedAppendWritesHistory_STATUS{}), generators)

	return protectedAppendWritesHistory_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProtectedAppendWritesHistory_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProtectedAppendWritesHistory_STATUS(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWritesAll"] = gen.PtrOf(gen.Bool())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_TagProperty_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagProperty_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagProperty_STATUS, TagProperty_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagProperty_STATUS runs a test to see if a specific instance of TagProperty_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTagProperty_STATUS(subject TagProperty_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagProperty_STATUS
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

// Generator of TagProperty_STATUS instances for property testing - lazily instantiated by TagProperty_STATUSGenerator()
var tagProperty_STATUSGenerator gopter.Gen

// TagProperty_STATUSGenerator returns a generator of TagProperty_STATUS instances for property testing.
func TagProperty_STATUSGenerator() gopter.Gen {
	if tagProperty_STATUSGenerator != nil {
		return tagProperty_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagProperty_STATUS(generators)
	tagProperty_STATUSGenerator = gen.Struct(reflect.TypeOf(TagProperty_STATUS{}), generators)

	return tagProperty_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTagProperty_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagProperty_STATUS(gens map[string]gopter.Gen) {
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

func Test_UpdateHistoryProperty_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UpdateHistoryProperty_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUpdateHistoryProperty_STATUS, UpdateHistoryProperty_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUpdateHistoryProperty_STATUS runs a test to see if a specific instance of UpdateHistoryProperty_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUpdateHistoryProperty_STATUS(subject UpdateHistoryProperty_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UpdateHistoryProperty_STATUS
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

// Generator of UpdateHistoryProperty_STATUS instances for property testing - lazily instantiated by
// UpdateHistoryProperty_STATUSGenerator()
var updateHistoryProperty_STATUSGenerator gopter.Gen

// UpdateHistoryProperty_STATUSGenerator returns a generator of UpdateHistoryProperty_STATUS instances for property testing.
func UpdateHistoryProperty_STATUSGenerator() gopter.Gen {
	if updateHistoryProperty_STATUSGenerator != nil {
		return updateHistoryProperty_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUS(generators)
	updateHistoryProperty_STATUSGenerator = gen.Struct(reflect.TypeOf(UpdateHistoryProperty_STATUS{}), generators)

	return updateHistoryProperty_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUS(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["AllowProtectedAppendWritesAll"] = gen.PtrOf(gen.Bool())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Update"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}
