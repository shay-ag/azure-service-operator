// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=mongodbdatabasecollections,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={mongodbdatabasecollections/status,mongodbdatabasecollections/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210515.MongodbDatabaseCollection
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}
type MongodbDatabaseCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_MongodbDatabases_Collection_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccounts_MongodbDatabases_Collection_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseCollection{}

// GetConditions returns the conditions of the resource
func (collection *MongodbDatabaseCollection) GetConditions() conditions.Conditions {
	return collection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (collection *MongodbDatabaseCollection) SetConditions(conditions conditions.Conditions) {
	collection.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseCollection{}

// AzureName returns the Azure name of the resource
func (collection *MongodbDatabaseCollection) AzureName() string {
	return collection.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (collection MongodbDatabaseCollection) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (collection *MongodbDatabaseCollection) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (collection *MongodbDatabaseCollection) GetSpec() genruntime.ConvertibleSpec {
	return &collection.Spec
}

// GetStatus returns the status of this resource
func (collection *MongodbDatabaseCollection) GetStatus() genruntime.ConvertibleStatus {
	return &collection.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
func (collection *MongodbDatabaseCollection) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
}

// NewEmptyStatus returns a new empty (blank) status
func (collection *MongodbDatabaseCollection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccounts_MongodbDatabases_Collection_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (collection *MongodbDatabaseCollection) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(collection.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  collection.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (collection *MongodbDatabaseCollection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccounts_MongodbDatabases_Collection_STATUS); ok {
		collection.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccounts_MongodbDatabases_Collection_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	collection.Status = st
	return nil
}

// Hub marks that this MongodbDatabaseCollection is the hub type for conversion
func (collection *MongodbDatabaseCollection) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (collection *MongodbDatabaseCollection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: collection.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseCollection",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210515.MongodbDatabaseCollection
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}
type MongodbDatabaseCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseCollection `json:"items"`
}

// Storage version of v1api20210515.DatabaseAccounts_MongodbDatabases_Collection_Spec
type DatabaseAccounts_MongodbDatabases_Collection_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/MongodbDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"MongodbDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionResource         `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_MongodbDatabases_Collection_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_MongodbDatabases_Collection_Spec from the provided source
func (collection *DatabaseAccounts_MongodbDatabases_Collection_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == collection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(collection)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_MongodbDatabases_Collection_Spec
func (collection *DatabaseAccounts_MongodbDatabases_Collection_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == collection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(collection)
}

// Storage version of v1api20210515.DatabaseAccounts_MongodbDatabases_Collection_STATUS
type DatabaseAccounts_MongodbDatabases_Collection_STATUS struct {
	Conditions  []conditions.Condition                          `json:"conditions,omitempty"`
	Id          *string                                         `json:"id,omitempty"`
	Location    *string                                         `json:"location,omitempty"`
	Name        *string                                         `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                         `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                               `json:"tags,omitempty"`
	Type        *string                                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccounts_MongodbDatabases_Collection_STATUS{}

// ConvertStatusFrom populates our DatabaseAccounts_MongodbDatabases_Collection_STATUS from the provided source
func (collection *DatabaseAccounts_MongodbDatabases_Collection_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == collection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(collection)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccounts_MongodbDatabases_Collection_STATUS
func (collection *DatabaseAccounts_MongodbDatabases_Collection_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == collection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(collection)
}

// Storage version of v1api20210515.MongoDBCollectionGetProperties_Resource_STATUS
type MongoDBCollectionGetProperties_Resource_STATUS struct {
	AnalyticalStorageTtl *int                   `json:"analyticalStorageTtl,omitempty"`
	Etag                 *string                `json:"_etag,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Indexes              []MongoIndex_STATUS    `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid                  *string                `json:"_rid,omitempty"`
	ShardKey             map[string]string      `json:"shardKey,omitempty"`
	Ts                   *float64               `json:"_ts,omitempty"`
}

// Storage version of v1api20210515.MongoDBCollectionResource
// Cosmos DB MongoDB collection resource object
type MongoDBCollectionResource struct {
	AnalyticalStorageTtl *int                   `json:"analyticalStorageTtl,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Indexes              []MongoIndex           `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ShardKey             map[string]string      `json:"shardKey,omitempty"`
}

// Storage version of v1api20210515.MongoIndex
// Cosmos DB MongoDB collection index key
type MongoIndex struct {
	Key         *MongoIndexKeys        `json:"key,omitempty"`
	Options     *MongoIndexOptions     `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.MongoIndex_STATUS
// Cosmos DB MongoDB collection index key
type MongoIndex_STATUS struct {
	Key         *MongoIndexKeys_STATUS    `json:"key,omitempty"`
	Options     *MongoIndexOptions_STATUS `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.MongoIndexKeys
// Cosmos DB MongoDB collection resource object
type MongoIndexKeys struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.MongoIndexKeys_STATUS
// Cosmos DB MongoDB collection resource object
type MongoIndexKeys_STATUS struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.MongoIndexOptions
// Cosmos DB MongoDB collection index options
type MongoIndexOptions struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

// Storage version of v1api20210515.MongoIndexOptions_STATUS
// Cosmos DB MongoDB collection index options
type MongoIndexOptions_STATUS struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseCollection{}, &MongodbDatabaseCollectionList{})
}
