// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601storage

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

func Test_Database_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase runs a test to see if a specific instance of Database round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase(subject Database) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database
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

// Generator of Database instances for property testing - lazily instantiated by DatabaseGenerator()
var databaseGenerator gopter.Gen

// DatabaseGenerator returns a generator of Database instances for property testing.
func DatabaseGenerator() gopter.Gen {
	if databaseGenerator != nil {
		return databaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDatabase(generators)
	databaseGenerator = gen.Struct(reflect.TypeOf(Database{}), generators)

	return databaseGenerator
}

// AddRelatedPropertyGeneratorsForDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Database_SpecGenerator()
	gens["Status"] = Servers_Database_STATUSGenerator()
}

func Test_Servers_Database_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Database_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Database_Spec, Servers_Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Database_Spec runs a test to see if a specific instance of Servers_Database_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Database_Spec(subject Servers_Database_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Database_Spec
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

// Generator of Servers_Database_Spec instances for property testing - lazily instantiated by
// Servers_Database_SpecGenerator()
var servers_Database_SpecGenerator gopter.Gen

// Servers_Database_SpecGenerator returns a generator of Servers_Database_Spec instances for property testing.
func Servers_Database_SpecGenerator() gopter.Gen {
	if servers_Database_SpecGenerator != nil {
		return servers_Database_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Database_Spec(generators)
	servers_Database_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Database_Spec{}), generators)

	return servers_Database_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Database_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Database_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_Servers_Database_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Database_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Database_STATUS, Servers_Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Database_STATUS runs a test to see if a specific instance of Servers_Database_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Database_STATUS(subject Servers_Database_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Database_STATUS
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

// Generator of Servers_Database_STATUS instances for property testing - lazily instantiated by
// Servers_Database_STATUSGenerator()
var servers_Database_STATUSGenerator gopter.Gen

// Servers_Database_STATUSGenerator returns a generator of Servers_Database_STATUS instances for property testing.
func Servers_Database_STATUSGenerator() gopter.Gen {
	if servers_Database_STATUSGenerator != nil {
		return servers_Database_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Database_STATUS(generators)
	servers_Database_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Database_STATUS{}), generators)

	return servers_Database_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Database_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Database_STATUS(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
