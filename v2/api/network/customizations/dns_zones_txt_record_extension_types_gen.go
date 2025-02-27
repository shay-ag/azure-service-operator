// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20180501 "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501"
	v1api20180501s "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DnsZonesTXTRecordExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *DnsZonesTXTRecordExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20180501.DnsZonesTXTRecord{},
		&v1api20180501s.DnsZonesTXTRecord{}}
}
