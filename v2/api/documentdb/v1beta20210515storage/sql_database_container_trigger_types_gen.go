// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"fmt"
	v1api20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210515.SqlDatabaseContainerTrigger
// Deprecated version of SqlDatabaseContainerTrigger. Use v1api20210515.SqlDatabaseContainerTrigger instead
type SqlDatabaseContainerTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerTrigger{}

// GetConditions returns the conditions of the resource
func (trigger *SqlDatabaseContainerTrigger) GetConditions() conditions.Conditions {
	return trigger.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (trigger *SqlDatabaseContainerTrigger) SetConditions(conditions conditions.Conditions) {
	trigger.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerTrigger{}

// ConvertFrom populates our SqlDatabaseContainerTrigger from the provided hub SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20210515s.SqlDatabaseContainerTrigger)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20210515storage/SqlDatabaseContainerTrigger but received %T instead", hub)
	}

	return trigger.AssignProperties_From_SqlDatabaseContainerTrigger(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerTrigger from our SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20210515s.SqlDatabaseContainerTrigger)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20210515storage/SqlDatabaseContainerTrigger but received %T instead", hub)
	}

	return trigger.AssignProperties_To_SqlDatabaseContainerTrigger(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerTrigger{}

// AzureName returns the Azure name of the resource
func (trigger *SqlDatabaseContainerTrigger) AzureName() string {
	return trigger.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (trigger SqlDatabaseContainerTrigger) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (trigger *SqlDatabaseContainerTrigger) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (trigger *SqlDatabaseContainerTrigger) GetSpec() genruntime.ConvertibleSpec {
	return &trigger.Spec
}

// GetStatus returns the status of this resource
func (trigger *SqlDatabaseContainerTrigger) GetStatus() genruntime.ConvertibleStatus {
	return &trigger.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
func (trigger *SqlDatabaseContainerTrigger) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
}

// NewEmptyStatus returns a new empty (blank) status
func (trigger *SqlDatabaseContainerTrigger) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (trigger *SqlDatabaseContainerTrigger) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(trigger.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  trigger.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (trigger *SqlDatabaseContainerTrigger) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS); ok {
		trigger.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	trigger.Status = st
	return nil
}

// AssignProperties_From_SqlDatabaseContainerTrigger populates our SqlDatabaseContainerTrigger from the provided source SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) AssignProperties_From_SqlDatabaseContainerTrigger(source *v1api20210515s.SqlDatabaseContainerTrigger) error {

	// ObjectMeta
	trigger.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec() to populate field Spec")
	}
	trigger.Spec = spec

	// Status
	var status DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS
	err = status.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS() to populate field Status")
	}
	trigger.Status = status

	// Invoke the augmentConversionForSqlDatabaseContainerTrigger interface (if implemented) to customize the conversion
	var triggerAsAny any = trigger
	if augmentedTrigger, ok := triggerAsAny.(augmentConversionForSqlDatabaseContainerTrigger); ok {
		err := augmentedTrigger.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerTrigger populates the provided destination SqlDatabaseContainerTrigger from our SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) AssignProperties_To_SqlDatabaseContainerTrigger(destination *v1api20210515s.SqlDatabaseContainerTrigger) error {

	// ObjectMeta
	destination.ObjectMeta = *trigger.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec
	err := trigger.Spec.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS
	err = trigger.Status.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForSqlDatabaseContainerTrigger interface (if implemented) to customize the conversion
	var triggerAsAny any = trigger
	if augmentedTrigger, ok := triggerAsAny.(augmentConversionForSqlDatabaseContainerTrigger); ok {
		err := augmentedTrigger.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (trigger *SqlDatabaseContainerTrigger) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: trigger.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerTrigger",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210515.SqlDatabaseContainerTrigger
// Deprecated version of SqlDatabaseContainerTrigger. Use v1api20210515.SqlDatabaseContainerTrigger instead
type SqlDatabaseContainerTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerTrigger `json:"items"`
}

type augmentConversionForSqlDatabaseContainerTrigger interface {
	AssignPropertiesFrom(src *v1api20210515s.SqlDatabaseContainerTrigger) error
	AssignPropertiesTo(dst *v1api20210515s.SqlDatabaseContainerTrigger) error
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec
type DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerResource                `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec from the provided source
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec)
	if ok {
		// Populate our instance from source
		return trigger.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = trigger.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec)
	if ok {
		// Populate destination from our instance
		return trigger.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec{}
	err := trigger.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(dst)
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

// AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec populates our DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec from the provided source DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(source *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	trigger.AzureName = source.AzureName

	// Location
	trigger.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignProperties_From_CreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CreateUpdateOptions() to populate field Options")
		}
		trigger.Options = &option
	} else {
		trigger.Options = nil
	}

	// OriginalVersion
	trigger.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		trigger.Owner = &owner
	} else {
		trigger.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlTriggerResource
		err := resource.AssignProperties_From_SqlTriggerResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlTriggerResource() to populate field Resource")
		}
		trigger.Resource = &resource
	} else {
		trigger.Resource = nil
	}

	// Tags
	trigger.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		trigger.PropertyBag = propertyBag
	} else {
		trigger.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec interface (if implemented) to customize the conversion
	var triggerAsAny any = trigger
	if augmentedTrigger, ok := triggerAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec); ok {
		err := augmentedTrigger.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec populates the provided destination DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec from our DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec(destination *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(trigger.PropertyBag)

	// AzureName
	destination.AzureName = trigger.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(trigger.Location)

	// Options
	if trigger.Options != nil {
		var option v1api20210515s.CreateUpdateOptions
		err := trigger.Options.AssignProperties_To_CreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = trigger.OriginalVersion

	// Owner
	if trigger.Owner != nil {
		owner := trigger.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if trigger.Resource != nil {
		var resource v1api20210515s.SqlTriggerResource
		err := trigger.Resource.AssignProperties_To_SqlTriggerResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlTriggerResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(trigger.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec interface (if implemented) to customize the conversion
	var triggerAsAny any = trigger
	if augmentedTrigger, ok := triggerAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec); ok {
		err := augmentedTrigger.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS
// Deprecated version of DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS. Use v1api20210515.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS instead
type DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS struct {
	Conditions  []conditions.Condition                   `json:"conditions,omitempty"`
	Id          *string                                  `json:"id,omitempty"`
	Location    *string                                  `json:"location,omitempty"`
	Name        *string                                  `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                        `json:"tags,omitempty"`
	Type        *string                                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS{}

// ConvertStatusFrom populates our DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS from the provided source
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS)
	if ok {
		// Populate our instance from source
		return trigger.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = trigger.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS)
	if ok {
		// Populate destination from our instance
		return trigger.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS{}
	err := trigger.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(dst)
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

// AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS populates our DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS from the provided source DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(source *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	trigger.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	trigger.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	trigger.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	trigger.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlTriggerGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_SqlTriggerGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlTriggerGetProperties_Resource_STATUS() to populate field Resource")
		}
		trigger.Resource = &resource
	} else {
		trigger.Resource = nil
	}

	// Tags
	trigger.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	trigger.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		trigger.PropertyBag = propertyBag
	} else {
		trigger.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS interface (if implemented) to customize the conversion
	var triggerAsAny any = trigger
	if augmentedTrigger, ok := triggerAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS); ok {
		err := augmentedTrigger.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS populates the provided destination DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS from our DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS
func (trigger *DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS(destination *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(trigger.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(trigger.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(trigger.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(trigger.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(trigger.Name)

	// Resource
	if trigger.Resource != nil {
		var resource v1api20210515s.SqlTriggerGetProperties_Resource_STATUS
		err := trigger.Resource.AssignProperties_To_SqlTriggerGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlTriggerGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(trigger.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(trigger.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS interface (if implemented) to customize the conversion
	var triggerAsAny any = trigger
	if augmentedTrigger, ok := triggerAsAny.(augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS); ok {
		err := augmentedTrigger.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_Spec interface {
	AssignPropertiesFrom(src *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) error
	AssignPropertiesTo(dst *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec) error
}

type augmentConversionForDatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS interface {
	AssignPropertiesFrom(src *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) error
	AssignPropertiesTo(dst *v1api20210515s.DatabaseAccounts_SqlDatabases_Containers_Trigger_STATUS) error
}

// Storage version of v1beta20210515.SqlTriggerGetProperties_Resource_STATUS
// Deprecated version of SqlTriggerGetProperties_Resource_STATUS. Use v1api20210515.SqlTriggerGetProperties_Resource_STATUS instead
type SqlTriggerGetProperties_Resource_STATUS struct {
	Body             *string                `json:"body,omitempty"`
	Etag             *string                `json:"_etag,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid              *string                `json:"_rid,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
	Ts               *float64               `json:"_ts,omitempty"`
}

// AssignProperties_From_SqlTriggerGetProperties_Resource_STATUS populates our SqlTriggerGetProperties_Resource_STATUS from the provided source SqlTriggerGetProperties_Resource_STATUS
func (resource *SqlTriggerGetProperties_Resource_STATUS) AssignProperties_From_SqlTriggerGetProperties_Resource_STATUS(source *v1api20210515s.SqlTriggerGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Rid
	resource.Rid = genruntime.ClonePointerToString(source.Rid)

	// TriggerOperation
	resource.TriggerOperation = genruntime.ClonePointerToString(source.TriggerOperation)

	// TriggerType
	resource.TriggerType = genruntime.ClonePointerToString(source.TriggerType)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		resource.Ts = &t
	} else {
		resource.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlTriggerGetProperties_Resource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlTriggerGetProperties_Resource_STATUS); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlTriggerGetProperties_Resource_STATUS populates the provided destination SqlTriggerGetProperties_Resource_STATUS from our SqlTriggerGetProperties_Resource_STATUS
func (resource *SqlTriggerGetProperties_Resource_STATUS) AssignProperties_To_SqlTriggerGetProperties_Resource_STATUS(destination *v1api20210515s.SqlTriggerGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Rid
	destination.Rid = genruntime.ClonePointerToString(resource.Rid)

	// TriggerOperation
	destination.TriggerOperation = genruntime.ClonePointerToString(resource.TriggerOperation)

	// TriggerType
	destination.TriggerType = genruntime.ClonePointerToString(resource.TriggerType)

	// Ts
	if resource.Ts != nil {
		t := *resource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlTriggerGetProperties_Resource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlTriggerGetProperties_Resource_STATUS); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210515.SqlTriggerResource
// Deprecated version of SqlTriggerResource. Use v1api20210515.SqlTriggerResource instead
type SqlTriggerResource struct {
	Body             *string                `json:"body,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
}

// AssignProperties_From_SqlTriggerResource populates our SqlTriggerResource from the provided source SqlTriggerResource
func (resource *SqlTriggerResource) AssignProperties_From_SqlTriggerResource(source *v1api20210515s.SqlTriggerResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// TriggerOperation
	resource.TriggerOperation = genruntime.ClonePointerToString(source.TriggerOperation)

	// TriggerType
	resource.TriggerType = genruntime.ClonePointerToString(source.TriggerType)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlTriggerResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlTriggerResource); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlTriggerResource populates the provided destination SqlTriggerResource from our SqlTriggerResource
func (resource *SqlTriggerResource) AssignProperties_To_SqlTriggerResource(destination *v1api20210515s.SqlTriggerResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// TriggerOperation
	destination.TriggerOperation = genruntime.ClonePointerToString(resource.TriggerOperation)

	// TriggerType
	destination.TriggerType = genruntime.ClonePointerToString(resource.TriggerType)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlTriggerResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlTriggerResource); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlTriggerGetProperties_Resource_STATUS interface {
	AssignPropertiesFrom(src *v1api20210515s.SqlTriggerGetProperties_Resource_STATUS) error
	AssignPropertiesTo(dst *v1api20210515s.SqlTriggerGetProperties_Resource_STATUS) error
}

type augmentConversionForSqlTriggerResource interface {
	AssignPropertiesFrom(src *v1api20210515s.SqlTriggerResource) error
	AssignPropertiesTo(dst *v1api20210515s.SqlTriggerResource) error
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerTrigger{}, &SqlDatabaseContainerTriggerList{})
}
