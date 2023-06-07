/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	// . "github.com/onsi/gomega"

	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	dataprotectionstorage "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101storage"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Dataprotection_Backuppolicy_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// region := tc.AzureRegion

	// This code is meant for Object Type
	val1 := "BackupPolicy"
	amons := &dataprotectionstorage.BackupPolicy{
		ObjectType: &val1,
	}
	amon := &dataprotection.BackupPolicy{}
	err := amon.AssignProperties_From_BackupPolicy(amons)
	if err != nil {
		t.Fatalf("failed to assign properties from BackupPolicy: %v", err)
	}

	// This code is meant for Policy Rule: Azure Backup
	val2 := "BackupHourly"
	val3 := "AzureBackupRule"
	amons2 := &dataprotectionstorage.AzureBackupRule{
		Name:       &val2,
		ObjectType: &val3,
	}
	amon2 := &dataprotection.AzureBackupRule{}
	err2 := amon2.AssignProperties_From_AzureBackupRule(amons2)
	if err2 != nil {
		t.Fatalf("failed to assign properties from AzureBackupRule: %v", err2)
	}

	val4 := "OperationalStore"
	val5 := "DataStoreInfoBase"
	amons3 := &dataprotectionstorage.DataStoreInfoBase{
		DataStoreType: &val4,
		ObjectType:    &val5,
	}
	amon3 := &dataprotection.DataStoreInfoBase{}
	err3 := amon3.AssignProperties_From_DataStoreInfoBase(amons3)
	if err3 != nil {
		t.Fatalf("failed to assign properties from DataStoreInfoBase: %v", err3)
	}

	val6 := "AdhocBasedTriggerContext"
	amons4 := &dataprotectionstorage.AdhocBasedTriggerContext{
		ObjectType: &val6,
	}
	amon4 := &dataprotection.AdhocBasedTriggerContext{}
	err4 := amon4.AssignProperties_From_AdhocBasedTriggerContext(amons4)
	if err4 != nil {
		t.Fatalf("failed to assign properties from AdhocBasedTriggerContext: %v", err4)
	}

	val7 := "Default"
	amons5 := &dataprotectionstorage.RetentionTag{
		TagName: &val7,
	}
	amon5 := &dataprotection.RetentionTag{}
	err5 := amon5.AssignProperties_From_RetentionTag(amons5)
	if err5 != nil {
		t.Fatalf("failed to assign properties from RetentionTag: %v", err5)
	}

	val8 := "ScheduleBasedTriggerContext"
	amons6 := &dataprotectionstorage.ScheduleBasedTriggerContext{
		ObjectType: &val8,
	}
	amon6 := &dataprotection.ScheduleBasedTriggerContext{}
	err6 := amon6.AssignProperties_From_ScheduleBasedTriggerContext(amons6)
	if err6 != nil {
		t.Fatalf("failed to assign properties from ScheduleBasedTriggerContext: %v", err6)
	}

	val9 := "UTC"
	amons7 := &dataprotectionstorage.BackupSchedule{
		TimeZone: &val9,
	}
	amon7 := &dataprotection.BackupSchedule{}
	err7 := amon7.AssignProperties_From_BackupSchedule(amons7)
	if err7 != nil {
		t.Fatalf("failed to assign properties from BackupSchedule: %v", err7)
	}

	val10 := "Incremental"
	val11 := "AzureBackupParams"
	amons8 := &dataprotectionstorage.AzureBackupParams{
		BackupType: &val10,
		ObjectType: &val11,
	}
	amon8 := &dataprotection.AzureBackupParams{}
	err8 := amon8.AssignProperties_From_AzureBackupParams(amons8)
	if err8 != nil {
		t.Fatalf("failed to assign properties from AzureBackupParams: %v", err8)
	}

	val12 := "Default"
	val13 := "AzureRetentionRule"
	amons9 := &dataprotectionstorage.AzureRetentionRule{
		Name:       &val12,
		ObjectType: &val13,
	}
	amon9 := &dataprotection.AzureRetentionRule{}
	err9 := amon9.AssignProperties_From_AzureRetentionRule(amons9)
	if err9 != nil {
		t.Fatalf("failed to assign properties from AzureRetentionRule: %v", err9)
	}

	val14 := "P7D"
	val15 := "AbsoluteDeleteOption"
	amons10 := &dataprotectionstorage.AbsoluteDeleteOption{
		Duration:   &val14,
		ObjectType: &val15,
	}
	amon10 := &dataprotection.AbsoluteDeleteOption{}
	err10 := amon10.AssignProperties_From_AbsoluteDeleteOption(amons10)
	if err10 != nil {
		t.Fatalf("failed to assign properties from AbsoluteDeleteOption: %v", err10)
	}

	val16 := "OperationalStore"
	val17 := "DataStoreInfoBase"
	amons11 := &dataprotectionstorage.DataStoreInfoBase{
		DataStoreType: &val16,
		ObjectType:    &val17,
	}
	amon11 := &dataprotection.DataStoreInfoBase{}
	err11 := amon11.AssignProperties_From_DataStoreInfoBase(amons11)
	if err11 != nil {
		t.Fatalf("failed to assign properties from DataStoreInfoBase: %v", err11)
	}

	val18 := true
	val19 := 99
	amons12 := []dataprotection.TaggingCriteria{
		{
			IsDefault:       &val18,
			TaggingPriority: &val19,
		},
	}
	amon12 := []dataprotection.TaggingCriteria{}
	err12 := amon12.AssignProperties_From_TaggingCriteria(amons12)
	if err12 != nil {
		t.Fatalf("failed to assign properties from TaggingCriteria: %v", err12)
	}

	val20 := "Default"
	amons13 := &dataprotectionstorage.RetentionTag{
		TagName: &val20,
	}
	amon13 := &dataprotection.RetentionTag{}
	err13 := amon13.AssignProperties_From_RetentionTag(amons13)
	if err13 != nil {
		t.Fatalf("failed to assign properties from RetentionTag: %v", err13)
	}

	// Create a BackupPolicy
	backuppolicy := &dataprotection.BackupVaultsBackupPolicy{
		ObjectMeta: tc.MakeObjectMetaWithName("testsbackuppolicy"),
		Spec: dataprotection.BackupVaults_BackupPolicy_Spec{
			Owner: testcommon.AsOwner(rg),
			Properties: &dataprotection.BaseBackupPolicy{
				BackupPolicy: &dataprotection.BackupPolicy{
					DatasourceTypes: []string{"Microsoft.ContainerService/managedClusters"},
					ObjectType:      amon.ObjectType,
					PolicyRules: []dataprotection.BasePolicyRule{
						{
							AzureBackup: &dataprotection.AzureBackupRule{
								Name:       amon2.Name,
								ObjectType: amon2.ObjectType,
								DataStore: &dataprotection.DataStoreInfoBase{
									DataStoreType: amon3.DataStoreType,
									ObjectType:    amon3.ObjectType,
								},
								Trigger: &dataprotection.TriggerContext{
									Adhoc: &dataprotection.AdhocBasedTriggerContext{
										ObjectType: amon4.ObjectType,
										TaggingCriteria: &dataprotection.AdhocBasedTaggingCriteria{
											TagInfo: &dataprotection.RetentionTag{
												TagName: amon5.TagName,
											},
										},
									},
									Schedule: &dataprotection.ScheduleBasedTriggerContext{
										ObjectType: amon6.ObjectType,
										Schedule: &dataprotection.BackupSchedule{
											RepeatingTimeIntervals: []string{"R/2023-06-07T10:26:32+00:00/PT4H"},
											TimeZone:               amon7.TimeZone,
										},
										TaggingCriteria: []dataprotection.TaggingCriteria{
											{
												IsDefault:       amon12[0].IsDefault,
												TaggingPriority: amon12[0].TaggingPriority,
												TagInfo: &dataprotection.RetentionTag{
													TagName: amon13.TagName,
												},
												// Criteria: []dataprotection.BackupCriteria{
												// 	{
												// 		ScheduleBasedBackupCriteria: &dataprotection.ScheduleBasedBackupCriteria{
												// 			ObjectType:
												// 		},
												// 	},
												// },
											},
										},
									},
								},
								BackupParameters: &dataprotection.BackupParameters{
									AzureBackupParams: &dataprotection.AzureBackupParams{
										BackupType: amon8.BackupType,
										ObjectType: amon8.ObjectType,
									},
								},
							},
							AzureRetention: &dataprotection.AzureRetentionRule{
								Name:       amon9.Name,
								ObjectType: amon9.ObjectType,
								Lifecycles: []dataprotection.SourceLifeCycle{
									{
										DeleteAfter: &dataprotection.DeleteOption{
											AbsoluteDeleteOption: &dataprotection.AbsoluteDeleteOption{
												Duration:   amon10.Duration,
												ObjectType: amon10.ObjectType,
											},
										},
										SourceDataStore: &dataprotection.DataStoreInfoBase{
											DataStoreType: amon11.DataStoreType,
											ObjectType:    amon11.ObjectType,
										},
									},
								},
							},
						},
					},

					// PolicyRules: []dataprotection.BasePolicyRule{
					// 	AzureBackupRule: &dataprotection.AzureBackupRule{
					// 		Name: "BackupHourly",
					// 		ObjectType: &dataprotection.AzureBackupRule_ObjectType{
					// 			objectType: "AzureBackupRule",
					// 		},
					// 		// Trigger: &dataprotection.TriggerContext{},
					// 		DataStore: &dataprotection.DataStoreInfoBase{
					// 			DataStoreType: &dataprotection.DataStoreInfoBase_DataStoreType{
					// 				dataStoreType: "OperationalStore",
					// 			},
					// 			ObjectType: "DataStoreInfoBase",
					// 		},
					// 		BackupParameters: &dataprotection.BackupParameters{
					// 			AzureBackupParams: &dataprotection.AzureBackupParams{
					// 				BackupType: "Incremental",
					// 				ObjectType: &dataprotection.AzureBackupParams_ObjectType{
					// 					objectType: "AzureBackupParams",
					// 				},
					// 			},
					// 		},
					// 	},
					// 	AzureRetentionRule: &dataprotection.AzureRetentionRule{
					// 		Name: "Default",
					// 		ObjectType: &dataprotection.AzureRetentionRule_ObjectType{
					// 			objectType: "AzureRetentionRule",
					// 		},
					// 		IsDefault: true,
					// 		Lifecycles: []dataprotection.SourceLifeCycle{
					// 			SourceDataStore: &dataprotection.DataStoreInfoBase{
					// 				DataStoreType: &dataprotection.DataStoreInfoBase_DataStoreType{
					// 					dataStoreType: "OperationalStore",
					// 				},
					// 				ObjectType: "DataStoreInfoBase",
					// 			},
					// 		},
					// 	},
					// },
				},
			},
		},
	}

	tc.CreateResourceAndWait(backuppolicy)

	// Asserts

}

// "taggingCriteria": [
//         {
//           "isDefault": true,
//           "taggingPriority": 99,
//           "tagInfo": {
//             "id": "Default_",
//             "tagName": "Default"
//           }
//         }
//       ]
