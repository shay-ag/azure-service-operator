//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20201101preview

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnectionProperties_STATUS) DeepCopyInto(out *PrivateEndpointConnectionProperties_STATUS) {
	*out = *in
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpointProperty_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionStateProperty_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(PrivateEndpointConnectionProperties_ProvisioningState_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnectionProperties_STATUS.
func (in *PrivateEndpointConnectionProperties_STATUS) DeepCopy() *PrivateEndpointConnectionProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnectionProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointProperty_STATUS) DeepCopyInto(out *PrivateEndpointProperty_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointProperty_STATUS.
func (in *PrivateEndpointProperty_STATUS) DeepCopy() *PrivateEndpointProperty_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointProperty_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionStateProperty_STATUS) DeepCopyInto(out *PrivateLinkServiceConnectionStateProperty_STATUS) {
	*out = *in
	if in.ActionsRequired != nil {
		in, out := &in.ActionsRequired, &out.ActionsRequired
		*out = new(PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PrivateLinkServiceConnectionStateProperty_Status_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionStateProperty_STATUS.
func (in *PrivateLinkServiceConnectionStateProperty_STATUS) DeepCopy() *PrivateLinkServiceConnectionStateProperty_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionStateProperty_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIdentity) DeepCopyInto(out *ResourceIdentity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ResourceIdentity_Type)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIdentity.
func (in *ResourceIdentity) DeepCopy() *ResourceIdentity {
	if in == nil {
		return nil
	}
	out := new(ResourceIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceIdentity_STATUS) DeepCopyInto(out *ResourceIdentity_STATUS) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ResourceIdentity_Type_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]UserIdentity_STATUS, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceIdentity_STATUS.
func (in *ResourceIdentity_STATUS) DeepCopy() *ResourceIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(ResourceIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerExternalAdministrator) DeepCopyInto(out *ServerExternalAdministrator) {
	*out = *in
	if in.AdministratorType != nil {
		in, out := &in.AdministratorType, &out.AdministratorType
		*out = new(ServerExternalAdministrator_AdministratorType)
		**out = **in
	}
	if in.AzureADOnlyAuthentication != nil {
		in, out := &in.AzureADOnlyAuthentication, &out.AzureADOnlyAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.Login != nil {
		in, out := &in.Login, &out.Login
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(ServerExternalAdministrator_PrincipalType)
		**out = **in
	}
	if in.Sid != nil {
		in, out := &in.Sid, &out.Sid
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerExternalAdministrator.
func (in *ServerExternalAdministrator) DeepCopy() *ServerExternalAdministrator {
	if in == nil {
		return nil
	}
	out := new(ServerExternalAdministrator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerExternalAdministrator_STATUS) DeepCopyInto(out *ServerExternalAdministrator_STATUS) {
	*out = *in
	if in.AdministratorType != nil {
		in, out := &in.AdministratorType, &out.AdministratorType
		*out = new(ServerExternalAdministrator_AdministratorType_STATUS)
		**out = **in
	}
	if in.AzureADOnlyAuthentication != nil {
		in, out := &in.AzureADOnlyAuthentication, &out.AzureADOnlyAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.Login != nil {
		in, out := &in.Login, &out.Login
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(ServerExternalAdministrator_PrincipalType_STATUS)
		**out = **in
	}
	if in.Sid != nil {
		in, out := &in.Sid, &out.Sid
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerExternalAdministrator_STATUS.
func (in *ServerExternalAdministrator_STATUS) DeepCopy() *ServerExternalAdministrator_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServerExternalAdministrator_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerObservation) DeepCopyInto(out *ServerObservation) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AdministratorLoginPassword != nil {
		in, out := &in.AdministratorLoginPassword, &out.AdministratorLoginPassword
		*out = new(string)
		**out = **in
	}
	if in.Administrators != nil {
		in, out := &in.Administrators, &out.Administrators
		*out = new(ServerExternalAdministrator_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ResourceIdentity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyId != nil {
		in, out := &in.KeyId, &out.KeyId
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimalTlsVersion != nil {
		in, out := &in.MinimalTlsVersion, &out.MinimalTlsVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrimaryUserAssignedIdentityId != nil {
		in, out := &in.PrimaryUserAssignedIdentityId, &out.PrimaryUserAssignedIdentityId
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]ServerPrivateEndpointConnection_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(ServerProperties_PublicNetworkAccess_STATUS)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.WorkspaceFeature != nil {
		in, out := &in.WorkspaceFeature, &out.WorkspaceFeature
		*out = new(ServerProperties_WorkspaceFeature_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerObservation.
func (in *ServerObservation) DeepCopy() *ServerObservation {
	if in == nil {
		return nil
	}
	out := new(ServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerParameters) DeepCopyInto(out *ServerParameters) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AdministratorLoginPassword != nil {
		in, out := &in.AdministratorLoginPassword, &out.AdministratorLoginPassword
		*out = new(string)
		**out = **in
	}
	if in.Administrators != nil {
		in, out := &in.Administrators, &out.Administrators
		*out = new(ServerExternalAdministrator)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(ResourceIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyId != nil {
		in, out := &in.KeyId, &out.KeyId
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimalTlsVersion != nil {
		in, out := &in.MinimalTlsVersion, &out.MinimalTlsVersion
		*out = new(string)
		**out = **in
	}
	if in.PrimaryUserAssignedIdentityId != nil {
		in, out := &in.PrimaryUserAssignedIdentityId, &out.PrimaryUserAssignedIdentityId
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(ServerProperties_PublicNetworkAccess)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerParameters.
func (in *ServerParameters) DeepCopy() *ServerParameters {
	if in == nil {
		return nil
	}
	out := new(ServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerPrivateEndpointConnection_STATUS) DeepCopyInto(out *ServerPrivateEndpointConnection_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PrivateEndpointConnectionProperties_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerPrivateEndpointConnection_STATUS.
func (in *ServerPrivateEndpointConnection_STATUS) DeepCopy() *ServerPrivateEndpointConnection_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServerPrivateEndpointConnection_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server_STATUS) DeepCopyInto(out *Server_STATUS) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server_STATUS.
func (in *Server_STATUS) DeepCopy() *Server_STATUS {
	if in == nil {
		return nil
	}
	out := new(Server_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server_Spec) DeepCopyInto(out *Server_Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server_Spec.
func (in *Server_Spec) DeepCopy() *Server_Spec {
	if in == nil {
		return nil
	}
	out := new(Server_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServersObservation) DeepCopyInto(out *ServersObservation) {
	*out = *in
	if in.AutoPauseDelay != nil {
		in, out := &in.AutoPauseDelay, &out.AutoPauseDelay
		*out = new(int)
		**out = **in
	}
	if in.CatalogCollation != nil {
		in, out := &in.CatalogCollation, &out.CatalogCollation
		*out = new(DatabaseProperties_CatalogCollation_STATUS)
		**out = **in
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
		*out = new(string)
		**out = **in
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(DatabaseProperties_CreateMode_STATUS)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.CurrentBackupStorageRedundancy != nil {
		in, out := &in.CurrentBackupStorageRedundancy, &out.CurrentBackupStorageRedundancy
		*out = new(DatabaseProperties_CurrentBackupStorageRedundancy_STATUS)
		**out = **in
	}
	if in.CurrentServiceObjectiveName != nil {
		in, out := &in.CurrentServiceObjectiveName, &out.CurrentServiceObjectiveName
		*out = new(string)
		**out = **in
	}
	if in.CurrentSku != nil {
		in, out := &in.CurrentSku, &out.CurrentSku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseId != nil {
		in, out := &in.DatabaseId, &out.DatabaseId
		*out = new(string)
		**out = **in
	}
	if in.DefaultSecondaryLocation != nil {
		in, out := &in.DefaultSecondaryLocation, &out.DefaultSecondaryLocation
		*out = new(string)
		**out = **in
	}
	if in.EarliestRestoreDate != nil {
		in, out := &in.EarliestRestoreDate, &out.EarliestRestoreDate
		*out = new(string)
		**out = **in
	}
	if in.ElasticPoolId != nil {
		in, out := &in.ElasticPoolId, &out.ElasticPoolId
		*out = new(string)
		**out = **in
	}
	if in.FailoverGroupId != nil {
		in, out := &in.FailoverGroupId, &out.FailoverGroupId
		*out = new(string)
		**out = **in
	}
	if in.HighAvailabilityReplicaCount != nil {
		in, out := &in.HighAvailabilityReplicaCount, &out.HighAvailabilityReplicaCount
		*out = new(int)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LicenseType != nil {
		in, out := &in.LicenseType, &out.LicenseType
		*out = new(DatabaseProperties_LicenseType_STATUS)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LongTermRetentionBackupResourceId != nil {
		in, out := &in.LongTermRetentionBackupResourceId, &out.LongTermRetentionBackupResourceId
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceConfigurationId != nil {
		in, out := &in.MaintenanceConfigurationId, &out.MaintenanceConfigurationId
		*out = new(string)
		**out = **in
	}
	if in.ManagedBy != nil {
		in, out := &in.ManagedBy, &out.ManagedBy
		*out = new(string)
		**out = **in
	}
	if in.MaxLogSizeBytes != nil {
		in, out := &in.MaxLogSizeBytes, &out.MaxLogSizeBytes
		*out = new(int)
		**out = **in
	}
	if in.MaxSizeBytes != nil {
		in, out := &in.MaxSizeBytes, &out.MaxSizeBytes
		*out = new(int)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PausedDate != nil {
		in, out := &in.PausedDate, &out.PausedDate
		*out = new(string)
		**out = **in
	}
	if in.ReadScale != nil {
		in, out := &in.ReadScale, &out.ReadScale
		*out = new(DatabaseProperties_ReadScale_STATUS)
		**out = **in
	}
	if in.RecoverableDatabaseId != nil {
		in, out := &in.RecoverableDatabaseId, &out.RecoverableDatabaseId
		*out = new(string)
		**out = **in
	}
	if in.RecoveryServicesRecoveryPointId != nil {
		in, out := &in.RecoveryServicesRecoveryPointId, &out.RecoveryServicesRecoveryPointId
		*out = new(string)
		**out = **in
	}
	if in.RequestedBackupStorageRedundancy != nil {
		in, out := &in.RequestedBackupStorageRedundancy, &out.RequestedBackupStorageRedundancy
		*out = new(DatabaseProperties_RequestedBackupStorageRedundancy_STATUS)
		**out = **in
	}
	if in.RequestedServiceObjectiveName != nil {
		in, out := &in.RequestedServiceObjectiveName, &out.RequestedServiceObjectiveName
		*out = new(string)
		**out = **in
	}
	if in.RestorableDroppedDatabaseId != nil {
		in, out := &in.RestorableDroppedDatabaseId, &out.RestorableDroppedDatabaseId
		*out = new(string)
		**out = **in
	}
	if in.RestorePointInTime != nil {
		in, out := &in.RestorePointInTime, &out.RestorePointInTime
		*out = new(string)
		**out = **in
	}
	if in.ResumedDate != nil {
		in, out := &in.ResumedDate, &out.ResumedDate
		*out = new(string)
		**out = **in
	}
	if in.SampleName != nil {
		in, out := &in.SampleName, &out.SampleName
		*out = new(DatabaseProperties_SampleName_STATUS)
		**out = **in
	}
	if in.SecondaryType != nil {
		in, out := &in.SecondaryType, &out.SecondaryType
		*out = new(DatabaseProperties_SecondaryType_STATUS)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceDatabaseDeletionDate != nil {
		in, out := &in.SourceDatabaseDeletionDate, &out.SourceDatabaseDeletionDate
		*out = new(string)
		**out = **in
	}
	if in.SourceDatabaseId != nil {
		in, out := &in.SourceDatabaseId, &out.SourceDatabaseId
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(DatabaseProperties_Status_STATUS)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ZoneRedundant != nil {
		in, out := &in.ZoneRedundant, &out.ZoneRedundant
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServersObservation.
func (in *ServersObservation) DeepCopy() *ServersObservation {
	if in == nil {
		return nil
	}
	out := new(ServersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servers_Database) DeepCopyInto(out *Servers_Database) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servers_Database.
func (in *Servers_Database) DeepCopy() *Servers_Database {
	if in == nil {
		return nil
	}
	out := new(Servers_Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Servers_Database) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servers_DatabaseList) DeepCopyInto(out *Servers_DatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Servers_Database, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servers_DatabaseList.
func (in *Servers_DatabaseList) DeepCopy() *Servers_DatabaseList {
	if in == nil {
		return nil
	}
	out := new(Servers_DatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Servers_DatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servers_DatabaseParameters) DeepCopyInto(out *Servers_DatabaseParameters) {
	*out = *in
	if in.AutoPauseDelay != nil {
		in, out := &in.AutoPauseDelay, &out.AutoPauseDelay
		*out = new(int)
		**out = **in
	}
	if in.CatalogCollation != nil {
		in, out := &in.CatalogCollation, &out.CatalogCollation
		*out = new(DatabaseProperties_CatalogCollation)
		**out = **in
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
		*out = new(string)
		**out = **in
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(DatabaseProperties_CreateMode)
		**out = **in
	}
	if in.ElasticPoolId != nil {
		in, out := &in.ElasticPoolId, &out.ElasticPoolId
		*out = new(string)
		**out = **in
	}
	if in.HighAvailabilityReplicaCount != nil {
		in, out := &in.HighAvailabilityReplicaCount, &out.HighAvailabilityReplicaCount
		*out = new(int)
		**out = **in
	}
	if in.LicenseType != nil {
		in, out := &in.LicenseType, &out.LicenseType
		*out = new(DatabaseProperties_LicenseType)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LongTermRetentionBackupResourceId != nil {
		in, out := &in.LongTermRetentionBackupResourceId, &out.LongTermRetentionBackupResourceId
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceConfigurationId != nil {
		in, out := &in.MaintenanceConfigurationId, &out.MaintenanceConfigurationId
		*out = new(string)
		**out = **in
	}
	if in.MaxSizeBytes != nil {
		in, out := &in.MaxSizeBytes, &out.MaxSizeBytes
		*out = new(int)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(float64)
		**out = **in
	}
	if in.ReadScale != nil {
		in, out := &in.ReadScale, &out.ReadScale
		*out = new(DatabaseProperties_ReadScale)
		**out = **in
	}
	if in.RecoverableDatabaseId != nil {
		in, out := &in.RecoverableDatabaseId, &out.RecoverableDatabaseId
		*out = new(string)
		**out = **in
	}
	if in.RecoveryServicesRecoveryPointId != nil {
		in, out := &in.RecoveryServicesRecoveryPointId, &out.RecoveryServicesRecoveryPointId
		*out = new(string)
		**out = **in
	}
	if in.RequestedBackupStorageRedundancy != nil {
		in, out := &in.RequestedBackupStorageRedundancy, &out.RequestedBackupStorageRedundancy
		*out = new(DatabaseProperties_RequestedBackupStorageRedundancy)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RestorableDroppedDatabaseId != nil {
		in, out := &in.RestorableDroppedDatabaseId, &out.RestorableDroppedDatabaseId
		*out = new(string)
		**out = **in
	}
	if in.RestorePointInTime != nil {
		in, out := &in.RestorePointInTime, &out.RestorePointInTime
		*out = new(string)
		**out = **in
	}
	if in.SampleName != nil {
		in, out := &in.SampleName, &out.SampleName
		*out = new(DatabaseProperties_SampleName)
		**out = **in
	}
	if in.SecondaryType != nil {
		in, out := &in.SecondaryType, &out.SecondaryType
		*out = new(DatabaseProperties_SecondaryType)
		**out = **in
	}
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerNameSelector != nil {
		in, out := &in.ServerNameSelector, &out.ServerNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceDatabaseDeletionDate != nil {
		in, out := &in.SourceDatabaseDeletionDate, &out.SourceDatabaseDeletionDate
		*out = new(string)
		**out = **in
	}
	if in.SourceDatabaseId != nil {
		in, out := &in.SourceDatabaseId, &out.SourceDatabaseId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ZoneRedundant != nil {
		in, out := &in.ZoneRedundant, &out.ZoneRedundant
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servers_DatabaseParameters.
func (in *Servers_DatabaseParameters) DeepCopy() *Servers_DatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(Servers_DatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servers_Database_STATUS) DeepCopyInto(out *Servers_Database_STATUS) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servers_Database_STATUS.
func (in *Servers_Database_STATUS) DeepCopy() *Servers_Database_STATUS {
	if in == nil {
		return nil
	}
	out := new(Servers_Database_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Servers_Database_Spec) DeepCopyInto(out *Servers_Database_Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Servers_Database_Spec.
func (in *Servers_Database_Spec) DeepCopy() *Servers_Database_Spec {
	if in == nil {
		return nil
	}
	out := new(Servers_Database_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku.
func (in *Sku) DeepCopy() *Sku {
	if in == nil {
		return nil
	}
	out := new(Sku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS) DeepCopyInto(out *Sku_STATUS) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS.
func (in *Sku_STATUS) DeepCopy() *Sku_STATUS {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserIdentity_STATUS) DeepCopyInto(out *UserIdentity_STATUS) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserIdentity_STATUS.
func (in *UserIdentity_STATUS) DeepCopy() *UserIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(UserIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}
