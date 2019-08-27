package containerserviceapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/tcz001/azure-sdk-for-go/services/containerservice/mgmt/2018-03-31/containerservice"
)

// ContainerServicesClientAPI contains the set of methods on the ContainerServicesClient type.
type ContainerServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters containerservice.ContainerService) (result containerservice.ContainerServicesCreateOrUpdateFutureType, err error)
	Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerServicesDeleteFutureType, err error)
	Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerService, err error)
	List(ctx context.Context) (result containerservice.ListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ListResultPage, err error)
	ListOrchestrators(ctx context.Context, location string, resourceType string) (result containerservice.OrchestratorVersionProfileListResult, err error)
}

var _ ContainerServicesClientAPI = (*containerservice.ContainerServicesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerservice.OperationListResult, err error)
}

var _ OperationsClientAPI = (*containerservice.OperationsClient)(nil)

// ManagedClustersClientAPI contains the set of methods on the ManagedClustersClient type.
type ManagedClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedCluster) (result containerservice.ManagedClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedCluster, err error)
	GetAccessProfile(ctx context.Context, resourceGroupName string, resourceName string, roleName string) (result containerservice.ManagedClusterAccessProfile, err error)
	GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClusterUpgradeProfile, err error)
	List(ctx context.Context) (result containerservice.ManagedClusterListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ManagedClusterListResultPage, err error)
	ListClusterAdminCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.CredentialResults, err error)
	ListClusterUserCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.CredentialResults, err error)
	ResetAADProfile(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedClusterAADProfile) (result containerservice.ManagedClustersResetAADProfileFuture, err error)
	ResetServicePrincipalProfile(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedClusterServicePrincipalProfile) (result containerservice.ManagedClustersResetServicePrincipalProfileFuture, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.TagsObject) (result containerservice.ManagedClustersUpdateTagsFuture, err error)
}

var _ ManagedClustersClientAPI = (*containerservice.ManagedClustersClient)(nil)
