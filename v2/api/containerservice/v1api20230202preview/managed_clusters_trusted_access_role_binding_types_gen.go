// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230202preview

import (
	"fmt"
	v1api20230202ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202previewstorage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-02-02-preview/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}
type ManagedClustersTrustedAccessRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedClusters_TrustedAccessRoleBinding_Spec   `json:"spec,omitempty"`
	Status            ManagedClusters_TrustedAccessRoleBinding_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ManagedClustersTrustedAccessRoleBinding{}

// GetConditions returns the conditions of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) GetConditions() conditions.Conditions {
	return binding.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (binding *ManagedClustersTrustedAccessRoleBinding) SetConditions(conditions conditions.Conditions) {
	binding.Status.Conditions = conditions
}

var _ conversion.Convertible = &ManagedClustersTrustedAccessRoleBinding{}

// ConvertFrom populates our ManagedClustersTrustedAccessRoleBinding from the provided hub ManagedClustersTrustedAccessRoleBinding
func (binding *ManagedClustersTrustedAccessRoleBinding) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20230202ps.ManagedClustersTrustedAccessRoleBinding)
	if !ok {
		return fmt.Errorf("expected containerservice/v1api20230202previewstorage/ManagedClustersTrustedAccessRoleBinding but received %T instead", hub)
	}

	return binding.AssignProperties_From_ManagedClustersTrustedAccessRoleBinding(source)
}

// ConvertTo populates the provided hub ManagedClustersTrustedAccessRoleBinding from our ManagedClustersTrustedAccessRoleBinding
func (binding *ManagedClustersTrustedAccessRoleBinding) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20230202ps.ManagedClustersTrustedAccessRoleBinding)
	if !ok {
		return fmt.Errorf("expected containerservice/v1api20230202previewstorage/ManagedClustersTrustedAccessRoleBinding but received %T instead", hub)
	}

	return binding.AssignProperties_To_ManagedClustersTrustedAccessRoleBinding(destination)
}

// +kubebuilder:webhook:path=/mutate-containerservice-azure-com-v1api20230202preview-managedclusterstrustedaccessrolebinding,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=managedclusterstrustedaccessrolebindings,verbs=create;update,versions=v1api20230202preview,name=default.v1api20230202preview.managedclusterstrustedaccessrolebindings.containerservice.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ManagedClustersTrustedAccessRoleBinding{}

// Default applies defaults to the ManagedClustersTrustedAccessRoleBinding resource
func (binding *ManagedClustersTrustedAccessRoleBinding) Default() {
	binding.defaultImpl()
	var temp any = binding
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (binding *ManagedClustersTrustedAccessRoleBinding) defaultAzureName() {
	if binding.Spec.AzureName == "" {
		binding.Spec.AzureName = binding.Name
	}
}

// defaultImpl applies the code generated defaults to the ManagedClustersTrustedAccessRoleBinding resource
func (binding *ManagedClustersTrustedAccessRoleBinding) defaultImpl() { binding.defaultAzureName() }

var _ genruntime.ImportableResource = &ManagedClustersTrustedAccessRoleBinding{}

// InitializeSpec initializes the spec for this resource from the given status
func (binding *ManagedClustersTrustedAccessRoleBinding) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*ManagedClusters_TrustedAccessRoleBinding_STATUS); ok {
		return binding.Spec.Initialize_From_ManagedClusters_TrustedAccessRoleBinding_STATUS(s)
	}

	return fmt.Errorf("expected Status of type ManagedClusters_TrustedAccessRoleBinding_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ManagedClustersTrustedAccessRoleBinding{}

// AzureName returns the Azure name of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) AzureName() string {
	return binding.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-02-02-preview"
func (binding ManagedClustersTrustedAccessRoleBinding) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (binding *ManagedClustersTrustedAccessRoleBinding) GetSpec() genruntime.ConvertibleSpec {
	return &binding.Spec
}

// GetStatus returns the status of this resource
func (binding *ManagedClustersTrustedAccessRoleBinding) GetStatus() genruntime.ConvertibleStatus {
	return &binding.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
func (binding *ManagedClustersTrustedAccessRoleBinding) GetType() string {
	return "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
}

// NewEmptyStatus returns a new empty (blank) status
func (binding *ManagedClustersTrustedAccessRoleBinding) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ManagedClusters_TrustedAccessRoleBinding_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (binding *ManagedClustersTrustedAccessRoleBinding) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(binding.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  binding.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (binding *ManagedClustersTrustedAccessRoleBinding) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ManagedClusters_TrustedAccessRoleBinding_STATUS); ok {
		binding.Status = *st
		return nil
	}

	// Convert status to required version
	var st ManagedClusters_TrustedAccessRoleBinding_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	binding.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-containerservice-azure-com-v1api20230202preview-managedclusterstrustedaccessrolebinding,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=managedclusterstrustedaccessrolebindings,verbs=create;update,versions=v1api20230202preview,name=validate.v1api20230202preview.managedclusterstrustedaccessrolebindings.containerservice.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ManagedClustersTrustedAccessRoleBinding{}

// ValidateCreate validates the creation of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) ValidateCreate() error {
	validations := binding.createValidations()
	var temp any = binding
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) ValidateDelete() error {
	validations := binding.deleteValidations()
	var temp any = binding
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) ValidateUpdate(old runtime.Object) error {
	validations := binding.updateValidations()
	var temp any = binding
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) createValidations() []func() error {
	return []func() error{binding.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return binding.validateResourceReferences()
		},
		binding.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (binding *ManagedClustersTrustedAccessRoleBinding) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&binding.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (binding *ManagedClustersTrustedAccessRoleBinding) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*ManagedClustersTrustedAccessRoleBinding)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, binding)
}

// AssignProperties_From_ManagedClustersTrustedAccessRoleBinding populates our ManagedClustersTrustedAccessRoleBinding from the provided source ManagedClustersTrustedAccessRoleBinding
func (binding *ManagedClustersTrustedAccessRoleBinding) AssignProperties_From_ManagedClustersTrustedAccessRoleBinding(source *v1api20230202ps.ManagedClustersTrustedAccessRoleBinding) error {

	// ObjectMeta
	binding.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec ManagedClusters_TrustedAccessRoleBinding_Spec
	err := spec.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec() to populate field Spec")
	}
	binding.Spec = spec

	// Status
	var status ManagedClusters_TrustedAccessRoleBinding_STATUS
	err = status.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS() to populate field Status")
	}
	binding.Status = status

	// No error
	return nil
}

// AssignProperties_To_ManagedClustersTrustedAccessRoleBinding populates the provided destination ManagedClustersTrustedAccessRoleBinding from our ManagedClustersTrustedAccessRoleBinding
func (binding *ManagedClustersTrustedAccessRoleBinding) AssignProperties_To_ManagedClustersTrustedAccessRoleBinding(destination *v1api20230202ps.ManagedClustersTrustedAccessRoleBinding) error {

	// ObjectMeta
	destination.ObjectMeta = *binding.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec
	err := binding.Spec.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS
	err = binding.Status.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (binding *ManagedClustersTrustedAccessRoleBinding) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: binding.Spec.OriginalVersion(),
		Kind:    "ManagedClustersTrustedAccessRoleBinding",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-02-02-preview/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}
type ManagedClustersTrustedAccessRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedClustersTrustedAccessRoleBinding `json:"items"`
}

type ManagedClusters_TrustedAccessRoleBinding_Spec struct {
	// +kubebuilder:validation:MaxLength=24
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="^([A-Za-z0-9-])+$"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a containerservice.azure.com/ManagedCluster resource
	Owner *genruntime.KnownResourceReference `group:"containerservice.azure.com" json:"owner,omitempty" kind:"ManagedCluster"`

	// +kubebuilder:validation:Required
	// Roles: A list of roles to bind, each item is a resource type qualified role name. For example:
	// 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles []string `json:"roles,omitempty"`

	// +kubebuilder:validation:Required
	// SourceResourceReference: The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceReference *genruntime.ResourceReference `armReference:"SourceResourceId" json:"sourceResourceReference,omitempty"`
}

var _ genruntime.ARMTransformer = &ManagedClusters_TrustedAccessRoleBinding_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if binding == nil {
		return nil, nil
	}
	result := &ManagedClusters_TrustedAccessRoleBinding_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if binding.Roles != nil || binding.SourceResourceReference != nil {
		result.Properties = &TrustedAccessRoleBindingProperties_ARM{}
	}
	for _, item := range binding.Roles {
		result.Properties.Roles = append(result.Properties.Roles, item)
	}
	if binding.SourceResourceReference != nil {
		sourceResourceIdARMID, err := resolved.ResolvedReferences.Lookup(*binding.SourceResourceReference)
		if err != nil {
			return nil, err
		}
		sourceResourceId := sourceResourceIdARMID
		result.Properties.SourceResourceId = &sourceResourceId
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ManagedClusters_TrustedAccessRoleBinding_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ManagedClusters_TrustedAccessRoleBinding_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ManagedClusters_TrustedAccessRoleBinding_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	binding.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Owner’:
	binding.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘Roles’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Roles {
			binding.Roles = append(binding.Roles, item)
		}
	}

	// no assignment for property ‘SourceResourceReference’

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &ManagedClusters_TrustedAccessRoleBinding_Spec{}

// ConvertSpecFrom populates our ManagedClusters_TrustedAccessRoleBinding_Spec from the provided source
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec)
	if ok {
		// Populate our instance from source
		return binding.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = binding.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our ManagedClusters_TrustedAccessRoleBinding_Spec
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec)
	if ok {
		// Populate destination from our instance
		return binding.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec{}
	err := binding.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec populates our ManagedClusters_TrustedAccessRoleBinding_Spec from the provided source ManagedClusters_TrustedAccessRoleBinding_Spec
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_Spec(source *v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec) error {

	// AzureName
	binding.AzureName = source.AzureName

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		binding.Owner = &owner
	} else {
		binding.Owner = nil
	}

	// Roles
	binding.Roles = genruntime.CloneSliceOfString(source.Roles)

	// SourceResourceReference
	if source.SourceResourceReference != nil {
		sourceResourceReference := source.SourceResourceReference.Copy()
		binding.SourceResourceReference = &sourceResourceReference
	} else {
		binding.SourceResourceReference = nil
	}

	// No error
	return nil
}

// AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec populates the provided destination ManagedClusters_TrustedAccessRoleBinding_Spec from our ManagedClusters_TrustedAccessRoleBinding_Spec
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_Spec(destination *v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = binding.AzureName

	// OriginalVersion
	destination.OriginalVersion = binding.OriginalVersion()

	// Owner
	if binding.Owner != nil {
		owner := binding.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Roles
	destination.Roles = genruntime.CloneSliceOfString(binding.Roles)

	// SourceResourceReference
	if binding.SourceResourceReference != nil {
		sourceResourceReference := binding.SourceResourceReference.Copy()
		destination.SourceResourceReference = &sourceResourceReference
	} else {
		destination.SourceResourceReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_ManagedClusters_TrustedAccessRoleBinding_STATUS populates our ManagedClusters_TrustedAccessRoleBinding_Spec from the provided source ManagedClusters_TrustedAccessRoleBinding_STATUS
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) Initialize_From_ManagedClusters_TrustedAccessRoleBinding_STATUS(source *ManagedClusters_TrustedAccessRoleBinding_STATUS) error {

	// Roles
	binding.Roles = genruntime.CloneSliceOfString(source.Roles)

	// SourceResourceReference
	if source.SourceResourceId != nil {
		sourceResourceReference := genruntime.CreateResourceReferenceFromARMID(*source.SourceResourceId)
		binding.SourceResourceReference = &sourceResourceReference
	} else {
		binding.SourceResourceReference = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (binding *ManagedClusters_TrustedAccessRoleBinding_Spec) SetAzureName(azureName string) {
	binding.AzureName = azureName
}

type ManagedClusters_TrustedAccessRoleBinding_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// ProvisioningState: The current provisioning state of trusted access role binding.
	ProvisioningState *TrustedAccessRoleBindingProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Roles: A list of roles to bind, each item is a resource type qualified role name. For example:
	// 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles []string `json:"roles,omitempty"`

	// SourceResourceId: The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceId *string `json:"sourceResourceId,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ManagedClusters_TrustedAccessRoleBinding_STATUS{}

// ConvertStatusFrom populates our ManagedClusters_TrustedAccessRoleBinding_STATUS from the provided source
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS)
	if ok {
		// Populate our instance from source
		return binding.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = binding.AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our ManagedClusters_TrustedAccessRoleBinding_STATUS
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS)
	if ok {
		// Populate destination from our instance
		return binding.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS{}
	err := binding.AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &ManagedClusters_TrustedAccessRoleBinding_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		binding.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		binding.Name = &name
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			binding.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘Roles’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Roles {
			binding.Roles = append(binding.Roles, item)
		}
	}

	// Set property ‘SourceResourceId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.SourceResourceId != nil {
			sourceResourceId := *typedInput.Properties.SourceResourceId
			binding.SourceResourceId = &sourceResourceId
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		binding.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		binding.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS populates our ManagedClusters_TrustedAccessRoleBinding_STATUS from the provided source ManagedClusters_TrustedAccessRoleBinding_STATUS
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) AssignProperties_From_ManagedClusters_TrustedAccessRoleBinding_STATUS(source *v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS) error {

	// Conditions
	binding.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	binding.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	binding.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := TrustedAccessRoleBindingProperties_ProvisioningState_STATUS(*source.ProvisioningState)
		binding.ProvisioningState = &provisioningState
	} else {
		binding.ProvisioningState = nil
	}

	// Roles
	binding.Roles = genruntime.CloneSliceOfString(source.Roles)

	// SourceResourceId
	binding.SourceResourceId = genruntime.ClonePointerToString(source.SourceResourceId)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		binding.SystemData = &systemDatum
	} else {
		binding.SystemData = nil
	}

	// Type
	binding.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS populates the provided destination ManagedClusters_TrustedAccessRoleBinding_STATUS from our ManagedClusters_TrustedAccessRoleBinding_STATUS
func (binding *ManagedClusters_TrustedAccessRoleBinding_STATUS) AssignProperties_To_ManagedClusters_TrustedAccessRoleBinding_STATUS(destination *v1api20230202ps.ManagedClusters_TrustedAccessRoleBinding_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(binding.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(binding.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(binding.Name)

	// ProvisioningState
	if binding.ProvisioningState != nil {
		provisioningState := string(*binding.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// Roles
	destination.Roles = genruntime.CloneSliceOfString(binding.Roles)

	// SourceResourceId
	destination.SourceResourceId = genruntime.ClonePointerToString(binding.SourceResourceId)

	// SystemData
	if binding.SystemData != nil {
		var systemDatum v1api20230202ps.SystemData_STATUS
		err := binding.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(binding.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type TrustedAccessRoleBindingProperties_ProvisioningState_STATUS string

const (
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Canceled  = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Canceled")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Deleting  = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Deleting")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Failed    = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Failed")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Succeeded = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Succeeded")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Updating  = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Updating")
)

func init() {
	SchemeBuilder.Register(&ManagedClustersTrustedAccessRoleBinding{}, &ManagedClustersTrustedAccessRoleBindingList{})
}
