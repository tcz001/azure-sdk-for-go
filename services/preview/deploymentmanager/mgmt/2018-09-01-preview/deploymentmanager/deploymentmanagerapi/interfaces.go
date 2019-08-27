package deploymentmanagerapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/deploymentmanager/mgmt/2018-09-01-preview/deploymentmanager"
	"github.com/Azure/go-autorest/autorest"
)

// ServiceTopologiesClientAPI contains the set of methods on the ServiceTopologiesClient type.
type ServiceTopologiesClientAPI interface {
	CreateOrUpdate(ctx context.Context, serviceTopologyInfo deploymentmanager.ServiceTopologyResource, resourceGroupName string, serviceTopologyName string) (result deploymentmanager.ServiceTopologyResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceTopologyName string) (result deploymentmanager.ServiceTopologyResource, err error)
}

var _ ServiceTopologiesClientAPI = (*deploymentmanager.ServiceTopologiesClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceInfo deploymentmanager.ServiceResource) (result deploymentmanager.ServiceResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string) (result deploymentmanager.ServiceResource, err error)
}

var _ ServicesClientAPI = (*deploymentmanager.ServicesClient)(nil)

// ServiceUnitsClientAPI contains the set of methods on the ServiceUnitsClient type.
type ServiceUnitsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string, serviceUnitInfo deploymentmanager.ServiceUnitResource) (result deploymentmanager.ServiceUnitsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (result deploymentmanager.ServiceUnitResource, err error)
}

var _ ServiceUnitsClientAPI = (*deploymentmanager.ServiceUnitsClient)(nil)

// StepsClientAPI contains the set of methods on the StepsClient type.
type StepsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, stepName string, stepInfo *deploymentmanager.StepResource) (result deploymentmanager.StepResource, err error)
	Delete(ctx context.Context, resourceGroupName string, stepName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, stepName string) (result deploymentmanager.StepResource, err error)
}

var _ StepsClientAPI = (*deploymentmanager.StepsClient)(nil)

// RolloutsClientAPI contains the set of methods on the RolloutsClient type.
type RolloutsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, rolloutName string) (result deploymentmanager.Rollout, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, rolloutName string, rolloutRequest *deploymentmanager.RolloutRequest) (result deploymentmanager.RolloutsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, rolloutName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, rolloutName string, retryAttempt *int32) (result deploymentmanager.Rollout, err error)
	Restart(ctx context.Context, resourceGroupName string, rolloutName string, skipSucceeded *bool) (result deploymentmanager.Rollout, err error)
}

var _ RolloutsClientAPI = (*deploymentmanager.RolloutsClient)(nil)

// ArtifactSourcesClientAPI contains the set of methods on the ArtifactSourcesClient type.
type ArtifactSourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, artifactSourceName string, artifactSourceInfo *deploymentmanager.ArtifactSource) (result deploymentmanager.ArtifactSource, err error)
	Delete(ctx context.Context, resourceGroupName string, artifactSourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, artifactSourceName string) (result deploymentmanager.ArtifactSource, err error)
}

var _ ArtifactSourcesClientAPI = (*deploymentmanager.ArtifactSourcesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	Get(ctx context.Context) (result deploymentmanager.ListOperation, err error)
}

var _ OperationsClientAPI = (*deploymentmanager.OperationsClient)(nil)
