package timeseriesinsightsapi

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
	"github.com/tcz001/azure-sdk-for-go/services/timeseriesinsights/mgmt/2017-11-15/timeseriesinsights"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result timeseriesinsights.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*timeseriesinsights.OperationsClient)(nil)

// EnvironmentsClientAPI contains the set of methods on the EnvironmentsClient type.
type EnvironmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, parameters timeseriesinsights.EnvironmentCreateOrUpdateParameters) (result timeseriesinsights.EnvironmentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, environmentName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, environmentName string, expand string) (result timeseriesinsights.EnvironmentResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result timeseriesinsights.EnvironmentListResponse, err error)
	ListBySubscription(ctx context.Context) (result timeseriesinsights.EnvironmentListResponse, err error)
	Update(ctx context.Context, resourceGroupName string, environmentName string, environmentUpdateParameters timeseriesinsights.EnvironmentUpdateParameters) (result timeseriesinsights.EnvironmentsUpdateFuture, err error)
}

var _ EnvironmentsClientAPI = (*timeseriesinsights.EnvironmentsClient)(nil)

// EventSourcesClientAPI contains the set of methods on the EventSourcesClient type.
type EventSourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, parameters timeseriesinsights.BasicEventSourceCreateOrUpdateParameters) (result timeseriesinsights.EventSourceResourceModel, err error)
	Delete(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string) (result timeseriesinsights.EventSourceResourceModel, err error)
	ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string) (result timeseriesinsights.EventSourceListResponse, err error)
	Update(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, eventSourceUpdateParameters timeseriesinsights.EventSourceUpdateParameters) (result timeseriesinsights.EventSourceResourceModel, err error)
}

var _ EventSourcesClientAPI = (*timeseriesinsights.EventSourcesClient)(nil)

// ReferenceDataSetsClientAPI contains the set of methods on the ReferenceDataSetsClient type.
type ReferenceDataSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, parameters timeseriesinsights.ReferenceDataSetCreateOrUpdateParameters) (result timeseriesinsights.ReferenceDataSetResource, err error)
	Delete(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string) (result timeseriesinsights.ReferenceDataSetResource, err error)
	ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string) (result timeseriesinsights.ReferenceDataSetListResponse, err error)
	Update(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, referenceDataSetUpdateParameters timeseriesinsights.ReferenceDataSetUpdateParameters) (result timeseriesinsights.ReferenceDataSetResource, err error)
}

var _ ReferenceDataSetsClientAPI = (*timeseriesinsights.ReferenceDataSetsClient)(nil)

// AccessPoliciesClientAPI contains the set of methods on the AccessPoliciesClient type.
type AccessPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, parameters timeseriesinsights.AccessPolicyCreateOrUpdateParameters) (result timeseriesinsights.AccessPolicyResource, err error)
	Delete(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string) (result timeseriesinsights.AccessPolicyResource, err error)
	ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string) (result timeseriesinsights.AccessPolicyListResponse, err error)
	Update(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, accessPolicyUpdateParameters timeseriesinsights.AccessPolicyUpdateParameters) (result timeseriesinsights.AccessPolicyResource, err error)
}

var _ AccessPoliciesClientAPI = (*timeseriesinsights.AccessPoliciesClient)(nil)
