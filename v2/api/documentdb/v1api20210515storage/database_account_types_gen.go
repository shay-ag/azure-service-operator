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

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=databaseaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={databaseaccounts/status,databaseaccounts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210515.DatabaseAccount
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}
type DatabaseAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccount_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccount_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DatabaseAccount{}

// GetConditions returns the conditions of the resource
func (account *DatabaseAccount) GetConditions() conditions.Conditions {
	return account.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (account *DatabaseAccount) SetConditions(conditions conditions.Conditions) {
	account.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DatabaseAccount{}

// AzureName returns the Azure name of the resource
func (account *DatabaseAccount) AzureName() string {
	return account.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (account DatabaseAccount) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (account *DatabaseAccount) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (account *DatabaseAccount) GetSpec() genruntime.ConvertibleSpec {
	return &account.Spec
}

// GetStatus returns the status of this resource
func (account *DatabaseAccount) GetStatus() genruntime.ConvertibleStatus {
	return &account.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts"
func (account *DatabaseAccount) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts"
}

// NewEmptyStatus returns a new empty (blank) status
func (account *DatabaseAccount) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccount_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (account *DatabaseAccount) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(account.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  account.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (account *DatabaseAccount) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccount_STATUS); ok {
		account.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccount_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	account.Status = st
	return nil
}

// Hub marks that this DatabaseAccount is the hub type for conversion
func (account *DatabaseAccount) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (account *DatabaseAccount) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: account.Spec.OriginalVersion,
		Kind:    "DatabaseAccount",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210515.DatabaseAccount
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}
type DatabaseAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseAccount `json:"items"`
}

// Storage version of v1api20210515.APIVersion
// +kubebuilder:validation:Enum={"2021-05-15"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-05-15")

// Storage version of v1api20210515.DatabaseAccount_Spec
type DatabaseAccount_Spec struct {
	AnalyticalStorageConfiguration *AnalyticalStorageConfiguration `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                  *ApiProperties                  `json:"apiProperties,omitempty"`

	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:Pattern="^[a-z0-9]+(-[a-z0-9]+)*"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                          string                       `json:"azureName,omitempty"`
	BackupPolicy                       *BackupPolicy                `json:"backupPolicy,omitempty"`
	Capabilities                       []Capability                 `json:"capabilities,omitempty"`
	ConnectorOffer                     *string                      `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicy           `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicy                 `json:"cors,omitempty"`
	DatabaseAccountOfferType           *string                      `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                      `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                        `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	EnableAnalyticalStorage            *bool                        `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                        `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                        `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                        `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                        `json:"enableMultipleWriteLocations,omitempty"`
	Identity                           *ManagedServiceIdentity      `json:"identity,omitempty"`
	IpRules                            []IpAddressOrRange           `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                        `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                      `json:"keyVaultKeyUri,omitempty"`
	Kind                               *string                      `json:"kind,omitempty"`
	Location                           *string                      `json:"location,omitempty"`
	Locations                          []Location                   `json:"locations,omitempty"`
	NetworkAclBypass                   *string                      `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                     `json:"networkAclBypassResourceIds,omitempty"`
	OperatorSpec                       *DatabaseAccountOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion                    string                       `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
	VirtualNetworkRules []VirtualNetworkRule               `json:"virtualNetworkRules,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccount_Spec{}

// ConvertSpecFrom populates our DatabaseAccount_Spec from the provided source
func (account *DatabaseAccount_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(account)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccount_Spec
func (account *DatabaseAccount_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(account)
}

// Storage version of v1api20210515.DatabaseAccount_STATUS
type DatabaseAccount_STATUS struct {
	AnalyticalStorageConfiguration     *AnalyticalStorageConfiguration_STATUS `json:"analyticalStorageConfiguration,omitempty"`
	ApiProperties                      *ApiProperties_STATUS                  `json:"apiProperties,omitempty"`
	BackupPolicy                       *BackupPolicy_STATUS                   `json:"backupPolicy,omitempty"`
	Capabilities                       []Capability_STATUS                    `json:"capabilities,omitempty"`
	Conditions                         []conditions.Condition                 `json:"conditions,omitempty"`
	ConnectorOffer                     *string                                `json:"connectorOffer,omitempty"`
	ConsistencyPolicy                  *ConsistencyPolicy_STATUS              `json:"consistencyPolicy,omitempty"`
	Cors                               []CorsPolicy_STATUS                    `json:"cors,omitempty"`
	DatabaseAccountOfferType           *string                                `json:"databaseAccountOfferType,omitempty"`
	DefaultIdentity                    *string                                `json:"defaultIdentity,omitempty"`
	DisableKeyBasedMetadataWriteAccess *bool                                  `json:"disableKeyBasedMetadataWriteAccess,omitempty"`
	DocumentEndpoint                   *string                                `json:"documentEndpoint,omitempty"`
	EnableAnalyticalStorage            *bool                                  `json:"enableAnalyticalStorage,omitempty"`
	EnableAutomaticFailover            *bool                                  `json:"enableAutomaticFailover,omitempty"`
	EnableCassandraConnector           *bool                                  `json:"enableCassandraConnector,omitempty"`
	EnableFreeTier                     *bool                                  `json:"enableFreeTier,omitempty"`
	EnableMultipleWriteLocations       *bool                                  `json:"enableMultipleWriteLocations,omitempty"`
	FailoverPolicies                   []FailoverPolicy_STATUS                `json:"failoverPolicies,omitempty"`
	Id                                 *string                                `json:"id,omitempty"`
	Identity                           *ManagedServiceIdentity_STATUS         `json:"identity,omitempty"`
	IpRules                            []IpAddressOrRange_STATUS              `json:"ipRules,omitempty"`
	IsVirtualNetworkFilterEnabled      *bool                                  `json:"isVirtualNetworkFilterEnabled,omitempty"`
	KeyVaultKeyUri                     *string                                `json:"keyVaultKeyUri,omitempty"`
	Kind                               *string                                `json:"kind,omitempty"`
	Location                           *string                                `json:"location,omitempty"`
	Locations                          []Location_STATUS                      `json:"locations,omitempty"`
	Name                               *string                                `json:"name,omitempty"`
	NetworkAclBypass                   *string                                `json:"networkAclBypass,omitempty"`
	NetworkAclBypassResourceIds        []string                               `json:"networkAclBypassResourceIds,omitempty"`
	PrivateEndpointConnections         []PrivateEndpointConnection_STATUS     `json:"privateEndpointConnections,omitempty"`
	PropertyBag                        genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	ProvisioningState                  *string                                `json:"provisioningState,omitempty"`
	PublicNetworkAccess                *string                                `json:"publicNetworkAccess,omitempty"`
	ReadLocations                      []Location_STATUS                      `json:"readLocations,omitempty"`
	Tags                               map[string]string                      `json:"tags,omitempty"`
	Type                               *string                                `json:"type,omitempty"`
	VirtualNetworkRules                []VirtualNetworkRule_STATUS            `json:"virtualNetworkRules,omitempty"`
	WriteLocations                     []Location_STATUS                      `json:"writeLocations,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccount_STATUS{}

// ConvertStatusFrom populates our DatabaseAccount_STATUS from the provided source
func (account *DatabaseAccount_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(account)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccount_STATUS
func (account *DatabaseAccount_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(account)
}

// Storage version of v1api20210515.AnalyticalStorageConfiguration
// Analytical storage specific properties.
type AnalyticalStorageConfiguration struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SchemaType  *string                `json:"schemaType,omitempty"`
}

// Storage version of v1api20210515.AnalyticalStorageConfiguration_STATUS
// Analytical storage specific properties.
type AnalyticalStorageConfiguration_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SchemaType  *string                `json:"schemaType,omitempty"`
}

// Storage version of v1api20210515.ApiProperties
type ApiProperties struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerVersion *string                `json:"serverVersion,omitempty"`
}

// Storage version of v1api20210515.ApiProperties_STATUS
type ApiProperties_STATUS struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerVersion *string                `json:"serverVersion,omitempty"`
}

// Storage version of v1api20210515.BackupPolicy
type BackupPolicy struct {
	Continuous  *ContinuousModeBackupPolicy `json:"continuous,omitempty"`
	Periodic    *PeriodicModeBackupPolicy   `json:"periodic,omitempty"`
	PropertyBag genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.BackupPolicy_STATUS
type BackupPolicy_STATUS struct {
	Continuous  *ContinuousModeBackupPolicy_STATUS `json:"continuous,omitempty"`
	Periodic    *PeriodicModeBackupPolicy_STATUS   `json:"periodic,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.Capability
// Cosmos DB capability object
type Capability struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.Capability_STATUS
// Cosmos DB capability object
type Capability_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.ConsistencyPolicy
// The consistency policy for the Cosmos DB database account.
type ConsistencyPolicy struct {
	DefaultConsistencyLevel *string                `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                   `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                   `json:"maxStalenessPrefix,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.ConsistencyPolicy_STATUS
// The consistency policy for the Cosmos DB database account.
type ConsistencyPolicy_STATUS struct {
	DefaultConsistencyLevel *string                `json:"defaultConsistencyLevel,omitempty"`
	MaxIntervalInSeconds    *int                   `json:"maxIntervalInSeconds,omitempty"`
	MaxStalenessPrefix      *int                   `json:"maxStalenessPrefix,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.CorsPolicy
// The CORS policy for the Cosmos DB database account.
type CorsPolicy struct {
	AllowedHeaders  *string                `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string                `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string                `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string                `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.CorsPolicy_STATUS
// The CORS policy for the Cosmos DB database account.
type CorsPolicy_STATUS struct {
	AllowedHeaders  *string                `json:"allowedHeaders,omitempty"`
	AllowedMethods  *string                `json:"allowedMethods,omitempty"`
	AllowedOrigins  *string                `json:"allowedOrigins,omitempty"`
	ExposedHeaders  *string                `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.DatabaseAccountOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DatabaseAccountOperatorSpec struct {
	PropertyBag genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	Secrets     *DatabaseAccountOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1api20210515.FailoverPolicy_STATUS
// The failover policy for a given region of a database account.
type FailoverPolicy_STATUS struct {
	FailoverPriority *int                   `json:"failoverPriority,omitempty"`
	Id               *string                `json:"id,omitempty"`
	LocationName     *string                `json:"locationName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.IpAddressOrRange
// IpAddressOrRange object
type IpAddressOrRange struct {
	IpAddressOrRange *string                `json:"ipAddressOrRange,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.IpAddressOrRange_STATUS
// IpAddressOrRange object
type IpAddressOrRange_STATUS struct {
	IpAddressOrRange *string                `json:"ipAddressOrRange,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.Location
// A region in which the Azure Cosmos DB database account is deployed.
type Location struct {
	FailoverPriority *int                   `json:"failoverPriority,omitempty"`
	IsZoneRedundant  *bool                  `json:"isZoneRedundant,omitempty"`
	LocationName     *string                `json:"locationName,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.Location_STATUS
// A region in which the Azure Cosmos DB database account is deployed.
type Location_STATUS struct {
	DocumentEndpoint  *string                `json:"documentEndpoint,omitempty"`
	FailoverPriority  *int                   `json:"failoverPriority,omitempty"`
	Id                *string                `json:"id,omitempty"`
	IsZoneRedundant   *bool                  `json:"isZoneRedundant,omitempty"`
	LocationName      *string                `json:"locationName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
}

// Storage version of v1api20210515.ManagedServiceIdentity
// Identity for the resource.
type ManagedServiceIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210515.ManagedServiceIdentity_STATUS
// Identity for the resource.
type ManagedServiceIdentity_STATUS struct {
	PrincipalId            *string                                                         `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                                          `json:"$propertyBag,omitempty"`
	TenantId               *string                                                         `json:"tenantId,omitempty"`
	Type                   *string                                                         `json:"type,omitempty"`
	UserAssignedIdentities map[string]ManagedServiceIdentity_UserAssignedIdentities_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20210515.PrivateEndpointConnection_STATUS
// A private endpoint connection
type PrivateEndpointConnection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.VirtualNetworkRule
// Virtual Network ACL Rule object
type VirtualNetworkRule struct {
	IgnoreMissingVNetServiceEndpoint *bool                  `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID of a subnet, for example:
	// /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20210515.VirtualNetworkRule_STATUS
// Virtual Network ACL Rule object
type VirtualNetworkRule_STATUS struct {
	Id                               *string                `json:"id,omitempty"`
	IgnoreMissingVNetServiceEndpoint *bool                  `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
	PropertyBag                      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.ContinuousModeBackupPolicy
type ContinuousModeBackupPolicy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20210515.ContinuousModeBackupPolicy_STATUS
type ContinuousModeBackupPolicy_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20210515.DatabaseAccountOperatorSecrets
type DatabaseAccountOperatorSecrets struct {
	DocumentEndpoint           *genruntime.SecretDestination `json:"documentEndpoint,omitempty"`
	PrimaryMasterKey           *genruntime.SecretDestination `json:"primaryMasterKey,omitempty"`
	PrimaryReadonlyMasterKey   *genruntime.SecretDestination `json:"primaryReadonlyMasterKey,omitempty"`
	PropertyBag                genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryMasterKey         *genruntime.SecretDestination `json:"secondaryMasterKey,omitempty"`
	SecondaryReadonlyMasterKey *genruntime.SecretDestination `json:"secondaryReadonlyMasterKey,omitempty"`
}

// Storage version of v1api20210515.ManagedServiceIdentity_UserAssignedIdentities_STATUS
type ManagedServiceIdentity_UserAssignedIdentities_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.PeriodicModeBackupPolicy
type PeriodicModeBackupPolicy struct {
	PeriodicModeProperties *PeriodicModeProperties `json:"periodicModeProperties,omitempty"`
	PropertyBag            genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	Type                   *string                 `json:"type,omitempty"`
}

// Storage version of v1api20210515.PeriodicModeBackupPolicy_STATUS
type PeriodicModeBackupPolicy_STATUS struct {
	PeriodicModeProperties *PeriodicModeProperties_STATUS `json:"periodicModeProperties,omitempty"`
	PropertyBag            genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	Type                   *string                        `json:"type,omitempty"`
}

// Storage version of v1api20210515.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

// Storage version of v1api20210515.PeriodicModeProperties
// Configuration values for periodic mode backup
type PeriodicModeProperties struct {
	BackupIntervalInMinutes        *int                   `json:"backupIntervalInMinutes,omitempty"`
	BackupRetentionIntervalInHours *int                   `json:"backupRetentionIntervalInHours,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20210515.PeriodicModeProperties_STATUS
// Configuration values for periodic mode backup
type PeriodicModeProperties_STATUS struct {
	BackupIntervalInMinutes        *int                   `json:"backupIntervalInMinutes,omitempty"`
	BackupRetentionIntervalInHours *int                   `json:"backupRetentionIntervalInHours,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DatabaseAccount{}, &DatabaseAccountList{})
}
