/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// PackageImport represents an import of a name from a package
type PackageImport struct {
	packageReference PackageReference
	name             string
}

var _ fmt.Stringer = &PackageImport{}

// NewPackageImport creates a new package import from a reference
func NewPackageImport(packageReference PackageReference) PackageImport {
	return PackageImport{
		packageReference: packageReference,
	}
}

// WithName creates a new package reference with a friendly name
func (pi PackageImport) WithName(name string) PackageImport {
	pi.name = name
	return pi
}

func (pi PackageImport) AsImportSpec() *dst.ImportSpec {
	var name *dst.Ident
	if pi.name != "" {
		name = dst.NewIdent(pi.name)
	}

	return &dst.ImportSpec{
		Name: name,
		Path: astbuilder.StringLiteral(pi.packageReference.PackagePath()),
	}
}

// PackageName is the package name of the package reference
func (pi PackageImport) PackageName() string {
	if pi.HasExplicitName() {
		return pi.name
	}

	return pi.packageReference.PackageName()
}

// HasExplicitName returns true if this package import has an explicitly defined name
func (pi PackageImport) HasExplicitName() bool {
	return pi.name != ""
}

// Equals returns true if the passed package reference references the same package, false otherwise
func (pi PackageImport) Equals(ref PackageImport) bool {
	packagesEqual := pi.packageReference.Equals(ref.packageReference)
	namesEqual := pi.name == ref.name

	return packagesEqual && namesEqual
}

func (pi PackageImport) String() string {
	if len(pi.name) > 0 {
		return fmt.Sprintf("%s %s", pi.name, pi.packageReference)
	}

	return pi.packageReference.String()
}

// WithImportAlias creates a copy of this import with a name following the specified rules
func (pi PackageImport) WithImportAlias(style PackageImportStyle) PackageImport {
	var alias string
	switch ref := pi.packageReference.(type) {
	case LocalPackageReference:
		alias = pi.createImportAliasForLocalPackageReference(ref, style)
	case StoragePackageReference:
		alias = pi.createImportAliasForStoragePackageReference(ref, style)
	default:
		msg := fmt.Sprintf("cannot create import alias for external package reference %s", pi.packageReference)
		panic(msg)
	}

	return pi.WithName(alias)
}

// createImportAliasForLocalPackageReference creates a custom alias for importing this reference
// ref is the local package reference for which we want an alias
// style is the kind of alias to generate
func (pi PackageImport) createImportAliasForLocalPackageReference(
	ref LocalPackageReference,
	style PackageImportStyle) string {
	switch style {
	case VersionOnly:
		return fmt.Sprintf(
			"%s%s",
			pi.simplifiedGeneratorVersion(ref.GeneratorVersion()),
			pi.simplifiedApiVersion(ref.ApiVersion()))
	case GroupOnly:
		return ref.Group()
	case GroupAndVersion:
		return fmt.Sprintf(
			"%s_%s%s",
			ref.Group(),
			pi.simplifiedGeneratorVersion(ref.GeneratorVersion()),
			pi.simplifiedApiVersion(ref.ApiVersion()))
	default:
		panic(fmt.Sprintf("didn't expect PackageImportStyle %q", style))
	}
}

// createImportAliasForStoragePackageReference creates a custom alias for importing this reference
func (pi PackageImport) createImportAliasForStoragePackageReference(
	ref StoragePackageReference,
	style PackageImportStyle) string {
	localImport := pi.createImportAliasForLocalPackageReference(ref.Local(), style)
	switch style {
	case VersionOnly:
		return localImport + "s"
	case GroupOnly:
		return localImport
	case GroupAndVersion:
		return localImport + "s"
	}

	panic(fmt.Sprintf("didn't expect PackageImportStyle %q", style))
}

func (pi PackageImport) simplifiedApiVersion(version string) string {
	return strings.ToLower(pi.simplify(version, apiVersionSimplifications))
}

var apiVersionSimplifications = map[string]string{
	"alpha":   "a",
	"beta":    "b",
	"preview": "p",
	"-":       "",
}

func (pi PackageImport) simplifiedGeneratorVersion(version string) string {
	return pi.simplify(version, generatorVersionSimplifications)
}

var generatorVersionSimplifications = map[string]string{
	"v1alpha1api": "alpha",
	"v1beta":      "v",
}

func (pi PackageImport) simplify(result string, simplifications map[string]string) string {
	for l, s := range simplifications {
		result = strings.Replace(result, l, s, -1)
	}
	return result
}
