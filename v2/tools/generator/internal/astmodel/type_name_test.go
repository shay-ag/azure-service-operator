/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSingular_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		expected string
	}{
		{"Account", "Account"},
		{"Accounts", "Account"},
		{"Address", "Address"},
		{"Addresses", "Address"},
		{"Batch", "Batch"},
		{"Batches", "Batch"},
		{"ImportServices", "ImportService"},
		{"Exportservices", "Exportservice"},
		{"AzureRedis", "AzureRedis"},
		{"Aliases", "Alias"},
		{"AdoptedFoxes", "AdoptedFox"},
	}

	ref := makeTestLocalPackageReference("Demo", "v2010")

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			name := MakeTypeName(ref, c.name)
			result := name.Singular()
			g.Expect(result.name).To(Equal(c.expected))
		})
	}
}

func TestPlural_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		expected string
	}{
		{"Account", "Accounts"},
		{"Accounts", "Accounts"},
		{"Batch", "Batches"},
		{"Batches", "Batches"},
		{"ImportService", "ImportServices"},
		{"Exportservice", "Exportservices"},
		{"AzureRedis", "AzureRedis"},
	}

	ref := makeTestLocalPackageReference("Demo", "v2010")

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			name := MakeTypeName(ref, c.name)
			result := name.Plural()
			g.Expect(result.name).To(Equal(c.expected))
		})
	}
}

func TestTypeName_IsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ref := makeTestLocalPackageReference("Demo", "v2010")
	blank := TypeName{}
	name := MakeTypeName(ref, "Person")

	g.Expect(blank.IsEmpty()).To(BeTrue())
	g.Expect(name.IsEmpty()).To(BeFalse())
}

func TestSortTypeName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pkgv2 := makeTestLocalPackageReference("crm", "v2")
	pkgv3 := makeTestLocalPackageReference("crm", "v3")

	// Test cases
	testCases := []struct {
		name     string
		left     TypeName
		right    TypeName
		expected bool
	}{
		{
			name:     "Package v2 sorts before v3",
			left:     MakeTypeName(pkgv2, "TypeA"),
			right:    MakeTypeName(pkgv3, "TypeB"),
			expected: true,
		},
		{
			name:     "Package v3 can't be sorted before v2",
			left:     MakeTypeName(pkgv3, "TypeA"),
			right:    MakeTypeName(pkgv2, "TypeB"),
			expected: false,
		},
		{
			name:     "Package v2 of TypeA sorts before Package v2 of TypeB",
			left:     MakeTypeName(pkgv2, "TypeA"),
			right:    MakeTypeName(pkgv2, "TypeB"),
			expected: true,
		},
		{
			name:     "Package v2 can't be sorted before Package v2 of same Type",
			left:     MakeTypeName(pkgv2, "TypeB"),
			right:    MakeTypeName(pkgv2, "TypeB"),
			expected: false,
		},
		{
			name:     "Package v2 can't be sorted before Package v2 of same Type",
			left:     MakeTypeName(pkgv2, "TypeA"),
			right:    MakeTypeName(pkgv2, "TypeA"),
			expected: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := SortTypeName(tc.left, tc.right)
			g.Expect(result).To(Equal(tc.expected))
		})
	}
}
