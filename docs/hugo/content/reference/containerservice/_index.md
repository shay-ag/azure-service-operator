---
title: ContainerService Supported Resources
linktitle: ContainerService
no_list: true
---
### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.


| Resource                                                                                                                                                                                             | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                                |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [TrustedAccessRoleBinding](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20230202preview/#containerservice.azure.com/v1api20230202preview.TrustedAccessRoleBinding) | 2023-02-02-preview | v1api20230202preview | v2.2.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230202preview/v1api20230202preview_trustedaccessrolebinding.yaml) |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.


| Resource                                                                                                                                                                                             | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                                |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [ManagedCluster](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20230202preview/#containerservice.azure.com/v1api20230202preview.ManagedCluster)                     | 2023-02-02-preview | v1api20230202preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230202preview/v1api20230202preview_managedcluster.yaml)           |
| [ManagedCluster](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20230201/#containerservice.azure.com/v1api20230201.ManagedCluster)                                   | 2023-02-01         | v1api20230201        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230201/v1api20230201_managedcluster.yaml)                         |
| [ManagedCluster](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20210501/#containerservice.azure.com/v1api20210501.ManagedCluster)                                   | 2021-05-01         | v1api20210501        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20210501/v1api20210501_managedcluster.yaml)                         |
| [ManagedClustersAgentPool](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20230202preview/#containerservice.azure.com/v1api20230202preview.ManagedClustersAgentPool) | 2023-02-02-preview | v1api20230202preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230202preview/v1api20230202preview_managedclustersagentpool.yaml) |
| [ManagedClustersAgentPool](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20230201/#containerservice.azure.com/v1api20230201.ManagedClustersAgentPool)               | 2023-02-01         | v1api20230201        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20230201/v1api20230201_managedclustersagentpool.yaml)               |
| [ManagedClustersAgentPool](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20210501/#containerservice.azure.com/v1api20210501.ManagedClustersAgentPool)               | 2021-05-01         | v1api20210501        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1api20210501/v1api20210501_managedclustersagentpool.yaml)               |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                                                                 | ARM Version | CRD Version    | Supported From | Sample                                                                                                                                                    |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|----------------|----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| [ManagedCluster](https://azure.github.io/azure-service-operator/reference/containerservice/v1beta20210501/#containerservice.azure.com/v1beta20210501.ManagedCluster)                     | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1beta20210501/v1beta20210501_managedcluster.yaml)           |
| [ManagedClustersAgentPool](https://azure.github.io/azure-service-operator/reference/containerservice/v1beta20210501/#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPool) | 2021-05-01  | v1beta20210501 | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/containerservice/v1beta20210501/v1beta20210501_managedclustersagentpool.yaml) |

