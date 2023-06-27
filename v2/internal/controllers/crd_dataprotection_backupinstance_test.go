/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	// resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	// trustedaccess "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"

	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// The to package includes utilities for converting values to pointers.
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	// The genruntime package includes utilities for working with the Azure API runtime.
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Create a AKS Cluster

// Creation of Trusted Access Resource

func Test_Dataprotection_Backupinstance_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a new backupvault resource
	backupVault := newBackupVault(tc, rg, "asotestbackupvault")

	// Create a BackupPolicy
	backupPolicy := newBackupPolicy(tc, backupVault, "asotestbackuppolicy")

	// tc.CreateResourceAndWaitWithoutCleanup(rg)
	tc.CreateResourcesAndWait(backupVault, backupPolicy)

	// tc.CreateResourceAndWaitWithoutCleanup(backupvault)

	// Note:
	// It is mandatory to create a backupvault, backuppolicy before creating a backupinstance

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	if err != nil {
		t.Fatalf("failed to generate ssh key: %v", err)
	}
	// tc.Expect(err).ToNot(HaveOccurred())

	identityKind := aks.ManagedClusterIdentity_Type_SystemAssigned
	osType := aks.OSType_Linux
	agentPoolMode := aks.AgentPoolMode_System

	cluster := &aks.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: aks.ManagedCluster_Spec{
			Location:  backupVault.Spec.Location,
			Owner:     testcommon.AsOwner(rg),
			DnsPrefix: to.Ptr("aso"),
			AgentPoolProfiles: []aks.ManagedClusterAgentPoolProfile{
				{
					Name:   to.Ptr("ap1"),
					Count:  to.Ptr(1),
					VmSize: to.Ptr("Standard_DS2_v2"),
					OsType: &osType,
					Mode:   &agentPoolMode,
				},
			},
			LinuxProfile: &aks.ContainerServiceLinuxProfile{
				AdminUsername: &adminUsername,
				Ssh: &aks.ContainerServiceSshConfiguration{
					PublicKeys: []aks.ContainerServiceSshPublicKey{
						{
							KeyData: sshPublicKey,
						},
					},
				},
			},
			Identity: &aks.ManagedClusterIdentity{
				Type: &identityKind,
			},
		},
	}

	tc.CreateResourceAndWait(cluster)

	// Create a TrustedAccess

	ResourceId := "/subscriptions/f0c630e0-2995-4853-b056-0b3c09cb673f/resourceGroups/t-agrawals/providers/Microsoft.ContainerService/managedClusters/shayAKScluster"

	trustedAccess := &aks.ManagedClustersTrustedAccessRoleBinding{
		ObjectMeta: tc.MakeObjectMeta("trustedaccess"),
		Spec: aks.ManagedClusters_TrustedAccessRoleBinding_Spec{
			Owner:     testcommon.AsOwner(cluster),
			AzureName: "asotesttrustedaccess",
			Roles:     []string{"backup-operator"},
			SourceResourceReference: &genruntime.ResourceReference{
				ARMID: ResourceId,
			},
		},
	}

	tc.CreateResourceAndWait(trustedAccess)

	tc.Expect(trustedAccess.Status.ProvisioningState).To(BeEquivalentTo(to.Ptr(aks.TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Succeeded)))

	// CONSTANTS: BACKUP_INSTANCE
	// consts for BackupInstance:PolicyInfo:DataStoreParameters
	PolicyId_Value := "/subscriptions/f0c630e0-2995-4853-b056-0b3c09cb673f/resourceGroups/t-agrawals/providers/Microsoft.DataProtection/BackupVaults/shayvault1/backupPolicies/testsbackuppolicy"

	DataStoreParameters_ResourceGroupId_Value := "/subscriptions/f0c630e0-2995-4853-b056-0b3c09cb673f/resourceGroups/t-agrawals"

	ResourceUri_Value := "/subscriptions/f0c630e0-2995-4853-b056-0b3c09cb673f/resourceGroups/t-agrawals/providers/Microsoft.ContainerService/managedClusters/shayAKScluster"

	// Create a BackupInstance
	backupinstance := &dataprotection.BackupVaultsBackupInstance{
		ObjectMeta: tc.MakeObjectMeta("asotestbackupinstance"),
		Spec: dataprotection.BackupVaults_BackupInstance_Spec{
			Owner: testcommon.AsOwner(backupVault),
			Tags:  map[string]string{"cheese": "blue"},
			Properties: &dataprotection.BackupInstance{
				FriendlyName: to.Ptr("asotestbackupinstance"),
				ObjectType:   to.Ptr("BackupInstance"),
				DataSourceInfo: &dataprotection.Datasource{
					ObjectType: to.Ptr("Datasource"),
					ResourceReference: &genruntime.ResourceReference{
						ARMID: ResourceUri_Value,
					},
					ResourceName:     to.Ptr("shayAKScluster"),
					ResourceType:     to.Ptr("Microsoft.ContainerService/managedClusters"),
					ResourceLocation: backupVault.Spec.Location,
					ResourceUri:      &ResourceUri_Value,
					DatasourceType:   to.Ptr("Microsoft.ContainerService/managedClusters"),
				},
				DataSourceSetInfo: &dataprotection.DatasourceSet{
					ObjectType: to.Ptr("DatasourceSet"),
					ResourceReference: &genruntime.ResourceReference{
						ARMID: ResourceUri_Value,
					},
					ResourceName:     to.Ptr("shayAKScluster"),
					ResourceType:     to.Ptr("Microsoft.ContainerService/managedClusters"),
					ResourceLocation: backupVault.Spec.Location,
					ResourceUri:      &ResourceUri_Value,
					DatasourceType:   to.Ptr("Microsoft.ContainerService/managedClusters"),
				},
				PolicyInfo: &dataprotection.PolicyInfo{
					PolicyId: &PolicyId_Value,
					PolicyParameters: &dataprotection.PolicyParameters{
						DataStoreParametersList: []dataprotection.DataStoreParameters{
							{
								AzureOperationalStoreParameters: &dataprotection.AzureOperationalStoreParameters{
									DataStoreType:   to.Ptr(dataprotection.AzureOperationalStoreParameters_DataStoreType_OperationalStore),
									ObjectType:      to.Ptr(dataprotection.AzureOperationalStoreParameters_ObjectType_AzureOperationalStoreParameters),
									ResourceGroupId: &DataStoreParameters_ResourceGroupId_Value,
								},
							},
						},
						BackupDatasourceParametersList: []dataprotection.BackupDatasourceParameters{
							{
								KubernetesCluster: &dataprotection.KubernetesClusterBackupDatasourceParameters{
									ObjectType:                   to.Ptr(dataprotection.KubernetesClusterBackupDatasourceParameters_ObjectType_KubernetesClusterBackupDatasourceParameters),
									ExcludedResourceTypes:        []string{"v1/Secret"},
									SnapshotVolumes:              to.Ptr(true),
									IncludeClusterScopeResources: to.Ptr(true),
								},
							},
						},
					},
				},
			},
		},
	}
	// tc.CreateResourceAndWait(backupinstance)
	tc.CreateResourceAndWaitWithoutCleanup(backupinstance)
}

// DISCARDED CODE

// region := "East Asia"
// rg := &resources.ResourceGroup{
// 	ObjectMeta: tc.MakeObjectMetaWithName("t-agrawals"),
// 	Spec: resources.ResourceGroup_Spec{
// 		Location: &region,
// 	},
// }

// CONSTANTS: BACKUP_VAULT
// region := tc.AzureRegion
// region_backupvault := "East US"
// identityType := "SystemAssigned"
// alertsForAllJobFailures_Status := dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled
// StorageSetting_DatastoreType_Value := dataprotection.StorageSetting_DatastoreType_VaultStore
// StorageSetting_Type_Value := dataprotection.StorageSetting_Type_LocallyRedundant

// Create a backupvault
// backupvault := &dataprotection.BackupVault{
// 	ObjectMeta: tc.MakeObjectMetaWithName("shayvault1"),
// 	Spec: dataprotection.BackupVault_Spec{
// 		Location: &region_backupvault,
// 		Tags:     map[string]string{"cheese": "blue"},
// 		Owner:    testcommon.AsOwner(rg),
// 		Identity: &dataprotection.DppIdentityDetails{
// 			Type: &identityType,
// 		},
// 		Properties: &dataprotection.BackupVaultSpec{
// 			MonitoringSettings: &dataprotection.MonitoringSettings{
// 				AzureMonitorAlertSettings: &dataprotection.AzureMonitorAlertSettings{
// 					AlertsForAllJobFailures: &alertsForAllJobFailures_Status,
// 				},
// 			},
// 			StorageSettings: []dataprotection.StorageSetting{
// 				{
// 					DatastoreType: &StorageSetting_DatastoreType_Value,
// 					Type:          &StorageSetting_Type_Value,
// 				},
// 			},
// 		},
// 	},
// }
// tc.CreateResourceAndWait(backupvault)

// backupPolicy_ObjectType := dataprotection.BackupPolicy_ObjectType_BackupPolicy

// // consts for AzureBackupRule
// AzureBackRule_Name := "BackupHourly"
// AzureBackupRule_ObjectType := dataprotection.AzureBackupRule_ObjectType_AzureBackupRule

// AzureBackupParams_BackupType_Value := "Incremental"
// AzureBackupParams_ObjectType_Value := dataprotection.AzureBackupParams_ObjectType_AzureBackupParams

// DataStore_DataStoreType_Value := dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore
// DataStore_ObjectType_Value := "DataStoreInfoBase"

// Schedule_ObjectType_Value := dataprotection.ScheduleBasedTriggerContext_ObjectType_ScheduleBasedTriggerContext
// Schedule_Timezone_Value := "UTC"

// TaggingCriteria_isDefault_Value := true
// TaggingCriteria_TaggingPriority_Value := 99
// TaggingCriteria_TagInfo_TagName_Value := "Default"

// // consts for AzureRetentionRule
// AzureRetentionRule_Name := "Default"
// AzureRetentionRule_ObjectType := dataprotection.AzureRetentionRule_ObjectType_AzureRetentionRule
// AzureRetentionRule_IsDefault := true

// AzureRetentionRule_Lifecycles_DeleteAfter_Duration := "P9D"
// AzureRetentionRule_Lifecycles_DeleteAfter_ObjectType := dataprotection.AbsoluteDeleteOption_ObjectType_AbsoluteDeleteOption
// AzureRetentionRule_Lifecycles_SourceDataStore_DataStoreType := dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore
// AzureRetentionRule_Lifecycles_SourceDataStore_ObjectType := "DataStoreInfoBase"

// // backuppolicy generation
// backuppolicy := &dataprotection.BackupVaultsBackupPolicy{
// 	ObjectMeta: tc.MakeObjectMeta("asotestbackuppolicy"),
// 	Spec: dataprotection.BackupVaults_BackupPolicy_Spec{
// 		Owner: testcommon.AsOwner(backupVault),
// 		Properties: &dataprotection.BaseBackupPolicy{
// 			BackupPolicy: &dataprotection.BackupPolicy{
// 				DatasourceTypes: []string{"Microsoft.ContainerService/managedClusters"},
// 				ObjectType:      to.Ptr(),
// 				PolicyRules: []dataprotection.BasePolicyRule{
// 					{
// 						AzureBackup: &dataprotection.AzureBackupRule{
// 							Name:       &AzureBackRule_Name,
// 							ObjectType: &AzureBackupRule_ObjectType,
// 							BackupParameters: &dataprotection.BackupParameters{
// 								AzureBackupParams: &dataprotection.AzureBackupParams{
// 									BackupType: &AzureBackupParams_BackupType_Value,
// 									ObjectType: &AzureBackupParams_ObjectType_Value,
// 								},
// 							},
// 							DataStore: &dataprotection.DataStoreInfoBase{
// 								DataStoreType: &DataStore_DataStoreType_Value,
// 								ObjectType:    &DataStore_ObjectType_Value,
// 							},
// 							Trigger: &dataprotection.TriggerContext{
// 								Schedule: &dataprotection.ScheduleBasedTriggerContext{
// 									ObjectType: &Schedule_ObjectType_Value,
// 									Schedule: &dataprotection.BackupSchedule{
// 										RepeatingTimeIntervals: []string{"R/2023-06-07T10:26:32+00:00/PT4H"},
// 										TimeZone:               &Schedule_Timezone_Value,
// 									},
// 									TaggingCriteria: []dataprotection.TaggingCriteria{
// 										{
// 											IsDefault:       &TaggingCriteria_isDefault_Value,
// 											TaggingPriority: &TaggingCriteria_TaggingPriority_Value,
// 											TagInfo: &dataprotection.RetentionTag{
// 												TagName: &TaggingCriteria_TagInfo_TagName_Value,
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 					{
// 						AzureRetention: &dataprotection.AzureRetentionRule{
// 							Name:       &AzureRetentionRule_Name,
// 							ObjectType: &AzureRetentionRule_ObjectType,
// 							IsDefault:  &AzureRetentionRule_IsDefault,
// 							Lifecycles: []dataprotection.SourceLifeCycle{
// 								{
// 									DeleteAfter: &dataprotection.DeleteOption{
// 										AbsoluteDeleteOption: &dataprotection.AbsoluteDeleteOption{
// 											Duration:   &AzureRetentionRule_Lifecycles_DeleteAfter_Duration,
// 											ObjectType: &AzureRetentionRule_Lifecycles_DeleteAfter_ObjectType,
// 										},
// 									},
// 									SourceDataStore: &dataprotection.DataStoreInfoBase{
// 										DataStoreType: &AzureRetentionRule_Lifecycles_SourceDataStore_DataStoreType,
// 										ObjectType:    &AzureRetentionRule_Lifecycles_SourceDataStore_ObjectType,
// 									},
// 									TargetDataStoreCopySettings: []dataprotection.TargetCopySetting{},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	},
// }
// // tc.CreateResourceAndWait(backuppolicy)
// tc.CreateResourceAndWaitWithoutCleanup(backuppolicy)

// BackupInstance_FriendlyName_Value := "asotestbackupinstance"
// BackupInstance_ObjectType_Value := "BackupInstance"

// consts for BackupInstance:DataSourceInfo
// DataSourceInfo_ObjectType_Value := "Datasource"
// DataSourceInfo_ResourceType_Value := "Microsoft.ContainerService/managedClusters"
// DataSourceInfo_DatasourceType_Value := "Microsoft.ContainerService/managedClusters"

// consts for BackupInstance:DataSourceSetInfo
// DataSourceSetInfo_ObjectType_Value := "DatasourceSet"
// DataSourceSetInfo_ResourceType_Value := "Microsoft.ContainerService/managedClusters"
// DataSourceSetInfo_DatasourceType_Value := "Microsoft.ContainerService/managedClusters"

// consts for BackupInstance:PolicyInfo:BackupDatasourceParameters
// BackupDatasourceParameters_ObjectType_Value := dataprotection.KubernetesClusterBackupDatasourceParameters_ObjectType_KubernetesClusterBackupDatasourceParameters
// BackupDatasourceParameters_SnapshotVolumes_Value := true
// BackupDatasourceParameters_IncludeClusterScopeResources_Value := true

// common consts for BackupInstance
// ResourceName_Value := "shayAKScluster"

// DataStoreParameters_DataStoreType_Value := dataprotection.AzureOperationalStoreParameters_DataStoreType_OperationalStore
// DataStoreParameters_ObjectType_Value := dataprotection.AzureOperationalStoreParameters_ObjectType_AzureOperationalStoreParameters
