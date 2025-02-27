//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20220901

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAadOrApiKeyAuthOption) DeepCopyInto(out *DataPlaneAadOrApiKeyAuthOption) {
	*out = *in
	if in.AadAuthFailureMode != nil {
		in, out := &in.AadAuthFailureMode, &out.AadAuthFailureMode
		*out = new(DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAadOrApiKeyAuthOption.
func (in *DataPlaneAadOrApiKeyAuthOption) DeepCopy() *DataPlaneAadOrApiKeyAuthOption {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAadOrApiKeyAuthOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAadOrApiKeyAuthOption_ARM) DeepCopyInto(out *DataPlaneAadOrApiKeyAuthOption_ARM) {
	*out = *in
	if in.AadAuthFailureMode != nil {
		in, out := &in.AadAuthFailureMode, &out.AadAuthFailureMode
		*out = new(DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAadOrApiKeyAuthOption_ARM.
func (in *DataPlaneAadOrApiKeyAuthOption_ARM) DeepCopy() *DataPlaneAadOrApiKeyAuthOption_ARM {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAadOrApiKeyAuthOption_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAadOrApiKeyAuthOption_STATUS) DeepCopyInto(out *DataPlaneAadOrApiKeyAuthOption_STATUS) {
	*out = *in
	if in.AadAuthFailureMode != nil {
		in, out := &in.AadAuthFailureMode, &out.AadAuthFailureMode
		*out = new(DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAadOrApiKeyAuthOption_STATUS.
func (in *DataPlaneAadOrApiKeyAuthOption_STATUS) DeepCopy() *DataPlaneAadOrApiKeyAuthOption_STATUS {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAadOrApiKeyAuthOption_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAadOrApiKeyAuthOption_STATUS_ARM) DeepCopyInto(out *DataPlaneAadOrApiKeyAuthOption_STATUS_ARM) {
	*out = *in
	if in.AadAuthFailureMode != nil {
		in, out := &in.AadAuthFailureMode, &out.AadAuthFailureMode
		*out = new(DataPlaneAadOrApiKeyAuthOption_AadAuthFailureMode_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAadOrApiKeyAuthOption_STATUS_ARM.
func (in *DataPlaneAadOrApiKeyAuthOption_STATUS_ARM) DeepCopy() *DataPlaneAadOrApiKeyAuthOption_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAadOrApiKeyAuthOption_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAuthOptions) DeepCopyInto(out *DataPlaneAuthOptions) {
	*out = *in
	if in.AadOrApiKey != nil {
		in, out := &in.AadOrApiKey, &out.AadOrApiKey
		*out = new(DataPlaneAadOrApiKeyAuthOption)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAuthOptions.
func (in *DataPlaneAuthOptions) DeepCopy() *DataPlaneAuthOptions {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAuthOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAuthOptions_ARM) DeepCopyInto(out *DataPlaneAuthOptions_ARM) {
	*out = *in
	if in.AadOrApiKey != nil {
		in, out := &in.AadOrApiKey, &out.AadOrApiKey
		*out = new(DataPlaneAadOrApiKeyAuthOption_ARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAuthOptions_ARM.
func (in *DataPlaneAuthOptions_ARM) DeepCopy() *DataPlaneAuthOptions_ARM {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAuthOptions_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAuthOptions_STATUS) DeepCopyInto(out *DataPlaneAuthOptions_STATUS) {
	*out = *in
	if in.AadOrApiKey != nil {
		in, out := &in.AadOrApiKey, &out.AadOrApiKey
		*out = new(DataPlaneAadOrApiKeyAuthOption_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiKeyOnly != nil {
		in, out := &in.ApiKeyOnly, &out.ApiKeyOnly
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAuthOptions_STATUS.
func (in *DataPlaneAuthOptions_STATUS) DeepCopy() *DataPlaneAuthOptions_STATUS {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAuthOptions_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneAuthOptions_STATUS_ARM) DeepCopyInto(out *DataPlaneAuthOptions_STATUS_ARM) {
	*out = *in
	if in.AadOrApiKey != nil {
		in, out := &in.AadOrApiKey, &out.AadOrApiKey
		*out = new(DataPlaneAadOrApiKeyAuthOption_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiKeyOnly != nil {
		in, out := &in.ApiKeyOnly, &out.ApiKeyOnly
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneAuthOptions_STATUS_ARM.
func (in *DataPlaneAuthOptions_STATUS_ARM) DeepCopy() *DataPlaneAuthOptions_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(DataPlaneAuthOptions_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionWithCmk) DeepCopyInto(out *EncryptionWithCmk) {
	*out = *in
	if in.Enforcement != nil {
		in, out := &in.Enforcement, &out.Enforcement
		*out = new(EncryptionWithCmk_Enforcement)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionWithCmk.
func (in *EncryptionWithCmk) DeepCopy() *EncryptionWithCmk {
	if in == nil {
		return nil
	}
	out := new(EncryptionWithCmk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionWithCmk_ARM) DeepCopyInto(out *EncryptionWithCmk_ARM) {
	*out = *in
	if in.Enforcement != nil {
		in, out := &in.Enforcement, &out.Enforcement
		*out = new(EncryptionWithCmk_Enforcement)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionWithCmk_ARM.
func (in *EncryptionWithCmk_ARM) DeepCopy() *EncryptionWithCmk_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionWithCmk_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionWithCmk_STATUS) DeepCopyInto(out *EncryptionWithCmk_STATUS) {
	*out = *in
	if in.EncryptionComplianceStatus != nil {
		in, out := &in.EncryptionComplianceStatus, &out.EncryptionComplianceStatus
		*out = new(EncryptionWithCmk_EncryptionComplianceStatus_STATUS)
		**out = **in
	}
	if in.Enforcement != nil {
		in, out := &in.Enforcement, &out.Enforcement
		*out = new(EncryptionWithCmk_Enforcement_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionWithCmk_STATUS.
func (in *EncryptionWithCmk_STATUS) DeepCopy() *EncryptionWithCmk_STATUS {
	if in == nil {
		return nil
	}
	out := new(EncryptionWithCmk_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionWithCmk_STATUS_ARM) DeepCopyInto(out *EncryptionWithCmk_STATUS_ARM) {
	*out = *in
	if in.EncryptionComplianceStatus != nil {
		in, out := &in.EncryptionComplianceStatus, &out.EncryptionComplianceStatus
		*out = new(EncryptionWithCmk_EncryptionComplianceStatus_STATUS)
		**out = **in
	}
	if in.Enforcement != nil {
		in, out := &in.Enforcement, &out.Enforcement
		*out = new(EncryptionWithCmk_Enforcement_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionWithCmk_STATUS_ARM.
func (in *EncryptionWithCmk_STATUS_ARM) DeepCopy() *EncryptionWithCmk_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionWithCmk_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(Identity_Type)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_ARM) DeepCopyInto(out *Identity_ARM) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(Identity_Type)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_ARM.
func (in *Identity_ARM) DeepCopy() *Identity_ARM {
	if in == nil {
		return nil
	}
	out := new(Identity_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_STATUS) DeepCopyInto(out *Identity_STATUS) {
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
		*out = new(Identity_Type_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_STATUS.
func (in *Identity_STATUS) DeepCopy() *Identity_STATUS {
	if in == nil {
		return nil
	}
	out := new(Identity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_STATUS_ARM) DeepCopyInto(out *Identity_STATUS_ARM) {
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
		*out = new(Identity_Type_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_STATUS_ARM.
func (in *Identity_STATUS_ARM) DeepCopy() *Identity_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(Identity_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRule) DeepCopyInto(out *IpRule) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRule.
func (in *IpRule) DeepCopy() *IpRule {
	if in == nil {
		return nil
	}
	out := new(IpRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRule_ARM) DeepCopyInto(out *IpRule_ARM) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRule_ARM.
func (in *IpRule_ARM) DeepCopy() *IpRule_ARM {
	if in == nil {
		return nil
	}
	out := new(IpRule_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRule_STATUS) DeepCopyInto(out *IpRule_STATUS) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRule_STATUS.
func (in *IpRule_STATUS) DeepCopy() *IpRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(IpRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRule_STATUS_ARM) DeepCopyInto(out *IpRule_STATUS_ARM) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRule_STATUS_ARM.
func (in *IpRule_STATUS_ARM) DeepCopy() *IpRule_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(IpRule_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRuleSet) DeepCopyInto(out *NetworkRuleSet) {
	*out = *in
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IpRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRuleSet.
func (in *NetworkRuleSet) DeepCopy() *NetworkRuleSet {
	if in == nil {
		return nil
	}
	out := new(NetworkRuleSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRuleSet_ARM) DeepCopyInto(out *NetworkRuleSet_ARM) {
	*out = *in
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IpRule_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRuleSet_ARM.
func (in *NetworkRuleSet_ARM) DeepCopy() *NetworkRuleSet_ARM {
	if in == nil {
		return nil
	}
	out := new(NetworkRuleSet_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRuleSet_STATUS) DeepCopyInto(out *NetworkRuleSet_STATUS) {
	*out = *in
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IpRule_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRuleSet_STATUS.
func (in *NetworkRuleSet_STATUS) DeepCopy() *NetworkRuleSet_STATUS {
	if in == nil {
		return nil
	}
	out := new(NetworkRuleSet_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRuleSet_STATUS_ARM) DeepCopyInto(out *NetworkRuleSet_STATUS_ARM) {
	*out = *in
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IpRule_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRuleSet_STATUS_ARM.
func (in *NetworkRuleSet_STATUS_ARM) DeepCopy() *NetworkRuleSet_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(NetworkRuleSet_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS) DeepCopyInto(out *PrivateEndpointConnection_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS.
func (in *PrivateEndpointConnection_STATUS) DeepCopy() *PrivateEndpointConnection_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS_ARM) DeepCopyInto(out *PrivateEndpointConnection_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS_ARM.
func (in *PrivateEndpointConnection_STATUS_ARM) DeepCopy() *PrivateEndpointConnection_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchService) DeepCopyInto(out *SearchService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchService.
func (in *SearchService) DeepCopy() *SearchService {
	if in == nil {
		return nil
	}
	out := new(SearchService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchServiceList) DeepCopyInto(out *SearchServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SearchService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchServiceList.
func (in *SearchServiceList) DeepCopy() *SearchServiceList {
	if in == nil {
		return nil
	}
	out := new(SearchServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchServiceOperatorSecrets) DeepCopyInto(out *SearchServiceOperatorSecrets) {
	*out = *in
	if in.AdminPrimaryKey != nil {
		in, out := &in.AdminPrimaryKey, &out.AdminPrimaryKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.AdminSecondaryKey != nil {
		in, out := &in.AdminSecondaryKey, &out.AdminSecondaryKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
	if in.QueryKey != nil {
		in, out := &in.QueryKey, &out.QueryKey
		*out = new(genruntime.SecretDestination)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchServiceOperatorSecrets.
func (in *SearchServiceOperatorSecrets) DeepCopy() *SearchServiceOperatorSecrets {
	if in == nil {
		return nil
	}
	out := new(SearchServiceOperatorSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchServiceOperatorSpec) DeepCopyInto(out *SearchServiceOperatorSpec) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(SearchServiceOperatorSecrets)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchServiceOperatorSpec.
func (in *SearchServiceOperatorSpec) DeepCopy() *SearchServiceOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(SearchServiceOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchServiceProperties_ARM) DeepCopyInto(out *SearchServiceProperties_ARM) {
	*out = *in
	if in.AuthOptions != nil {
		in, out := &in.AuthOptions, &out.AuthOptions
		*out = new(DataPlaneAuthOptions_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionWithCmk != nil {
		in, out := &in.EncryptionWithCmk, &out.EncryptionWithCmk
		*out = new(EncryptionWithCmk_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.HostingMode != nil {
		in, out := &in.HostingMode, &out.HostingMode
		*out = new(SearchServiceProperties_HostingMode)
		**out = **in
	}
	if in.NetworkRuleSet != nil {
		in, out := &in.NetworkRuleSet, &out.NetworkRuleSet
		*out = new(NetworkRuleSet_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(int)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(SearchServiceProperties_PublicNetworkAccess)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchServiceProperties_ARM.
func (in *SearchServiceProperties_ARM) DeepCopy() *SearchServiceProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(SearchServiceProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchServiceProperties_STATUS_ARM) DeepCopyInto(out *SearchServiceProperties_STATUS_ARM) {
	*out = *in
	if in.AuthOptions != nil {
		in, out := &in.AuthOptions, &out.AuthOptions
		*out = new(DataPlaneAuthOptions_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionWithCmk != nil {
		in, out := &in.EncryptionWithCmk, &out.EncryptionWithCmk
		*out = new(EncryptionWithCmk_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.HostingMode != nil {
		in, out := &in.HostingMode, &out.HostingMode
		*out = new(SearchServiceProperties_HostingMode_STATUS)
		**out = **in
	}
	if in.NetworkRuleSet != nil {
		in, out := &in.NetworkRuleSet, &out.NetworkRuleSet
		*out = new(NetworkRuleSet_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(int)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(SearchServiceProperties_ProvisioningState_STATUS)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(SearchServiceProperties_PublicNetworkAccess_STATUS)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int)
		**out = **in
	}
	if in.SharedPrivateLinkResources != nil {
		in, out := &in.SharedPrivateLinkResources, &out.SharedPrivateLinkResources
		*out = make([]SharedPrivateLinkResource_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(SearchServiceProperties_Status_STATUS)
		**out = **in
	}
	if in.StatusDetails != nil {
		in, out := &in.StatusDetails, &out.StatusDetails
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchServiceProperties_STATUS_ARM.
func (in *SearchServiceProperties_STATUS_ARM) DeepCopy() *SearchServiceProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SearchServiceProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchService_STATUS) DeepCopyInto(out *SearchService_STATUS) {
	*out = *in
	if in.AuthOptions != nil {
		in, out := &in.AuthOptions, &out.AuthOptions
		*out = new(DataPlaneAuthOptions_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionWithCmk != nil {
		in, out := &in.EncryptionWithCmk, &out.EncryptionWithCmk
		*out = new(EncryptionWithCmk_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.HostingMode != nil {
		in, out := &in.HostingMode, &out.HostingMode
		*out = new(SearchServiceProperties_HostingMode_STATUS)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkRuleSet != nil {
		in, out := &in.NetworkRuleSet, &out.NetworkRuleSet
		*out = new(NetworkRuleSet_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(int)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(SearchServiceProperties_ProvisioningState_STATUS)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(SearchServiceProperties_PublicNetworkAccess_STATUS)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int)
		**out = **in
	}
	if in.SharedPrivateLinkResources != nil {
		in, out := &in.SharedPrivateLinkResources, &out.SharedPrivateLinkResources
		*out = make([]SharedPrivateLinkResource_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(SearchServiceProperties_Status_STATUS)
		**out = **in
	}
	if in.StatusDetails != nil {
		in, out := &in.StatusDetails, &out.StatusDetails
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchService_STATUS.
func (in *SearchService_STATUS) DeepCopy() *SearchService_STATUS {
	if in == nil {
		return nil
	}
	out := new(SearchService_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchService_STATUS_ARM) DeepCopyInto(out *SearchService_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(SearchServiceProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS_ARM)
		(*in).DeepCopyInto(*out)
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchService_STATUS_ARM.
func (in *SearchService_STATUS_ARM) DeepCopy() *SearchService_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SearchService_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchService_Spec) DeepCopyInto(out *SearchService_Spec) {
	*out = *in
	if in.AuthOptions != nil {
		in, out := &in.AuthOptions, &out.AuthOptions
		*out = new(DataPlaneAuthOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableLocalAuth != nil {
		in, out := &in.DisableLocalAuth, &out.DisableLocalAuth
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionWithCmk != nil {
		in, out := &in.EncryptionWithCmk, &out.EncryptionWithCmk
		*out = new(EncryptionWithCmk)
		(*in).DeepCopyInto(*out)
	}
	if in.HostingMode != nil {
		in, out := &in.HostingMode, &out.HostingMode
		*out = new(SearchServiceProperties_HostingMode)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.NetworkRuleSet != nil {
		in, out := &in.NetworkRuleSet, &out.NetworkRuleSet
		*out = new(NetworkRuleSet)
		(*in).DeepCopyInto(*out)
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(SearchServiceOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(int)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(SearchServiceProperties_PublicNetworkAccess)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchService_Spec.
func (in *SearchService_Spec) DeepCopy() *SearchService_Spec {
	if in == nil {
		return nil
	}
	out := new(SearchService_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchService_Spec_ARM) DeepCopyInto(out *SearchService_Spec_ARM) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(SearchServiceProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchService_Spec_ARM.
func (in *SearchService_Spec_ARM) DeepCopy() *SearchService_Spec_ARM {
	if in == nil {
		return nil
	}
	out := new(SearchService_Spec_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedPrivateLinkResource_STATUS) DeepCopyInto(out *SharedPrivateLinkResource_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedPrivateLinkResource_STATUS.
func (in *SharedPrivateLinkResource_STATUS) DeepCopy() *SharedPrivateLinkResource_STATUS {
	if in == nil {
		return nil
	}
	out := new(SharedPrivateLinkResource_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedPrivateLinkResource_STATUS_ARM) DeepCopyInto(out *SharedPrivateLinkResource_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedPrivateLinkResource_STATUS_ARM.
func (in *SharedPrivateLinkResource_STATUS_ARM) DeepCopy() *SharedPrivateLinkResource_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SharedPrivateLinkResource_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name)
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
func (in *Sku_ARM) DeepCopyInto(out *Sku_ARM) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_ARM.
func (in *Sku_ARM) DeepCopy() *Sku_ARM {
	if in == nil {
		return nil
	}
	out := new(Sku_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS) DeepCopyInto(out *Sku_STATUS) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name_STATUS)
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
func (in *Sku_STATUS_ARM) DeepCopyInto(out *Sku_STATUS_ARM) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(Sku_Name_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS_ARM.
func (in *Sku_STATUS_ARM) DeepCopy() *Sku_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}
