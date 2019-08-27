package mediaapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/mediaservices/mgmt/2018-06-01-preview/media"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result media.OperationCollectionPage, err error)
}

var _ OperationsClientAPI = (*media.OperationsClient)(nil)

// MediaservicesClientAPI contains the set of methods on the MediaservicesClient type.
type MediaservicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, parameters media.Service) (result media.Service, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result media.Service, err error)
	GetBySubscription(ctx context.Context, accountName string) (result media.SubscriptionMediaService, err error)
	List(ctx context.Context, resourceGroupName string) (result media.ServiceCollectionPage, err error)
	ListBySubscription(ctx context.Context) (result media.SubscriptionMediaServiceCollectionPage, err error)
	SyncStorageKeys(ctx context.Context, resourceGroupName string, accountName string, parameters media.SyncStorageKeysInput) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters media.Service) (result media.Service, err error)
}

var _ MediaservicesClientAPI = (*media.MediaservicesClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	CheckNameAvailability(ctx context.Context, locationName string, parameters media.CheckNameAvailabilityInput) (result media.EntityNameAvailabilityCheckOutput, err error)
}

var _ LocationsClientAPI = (*media.LocationsClient)(nil)

// AssetsClientAPI contains the set of methods on the AssetsClient type.
type AssetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters media.Asset) (result media.Asset, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, assetName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, assetName string) (result media.Asset, err error)
	GetEncryptionKey(ctx context.Context, resourceGroupName string, accountName string, assetName string) (result media.AssetStorageEncryptionKey, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result media.AssetCollectionPage, err error)
	ListContainerSas(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters media.ListContainerSasInput) (result media.AssetContainerSas, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters media.Asset) (result media.Asset, err error)
}

var _ AssetsClientAPI = (*media.AssetsClient)(nil)

// ContentKeyPoliciesClientAPI contains the set of methods on the ContentKeyPoliciesClient type.
type ContentKeyPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters media.ContentKeyPolicy) (result media.ContentKeyPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (result media.ContentKeyPolicy, err error)
	GetPolicyPropertiesWithSecrets(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (result media.ContentKeyPolicyProperties, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result media.ContentKeyPolicyCollectionPage, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters media.ContentKeyPolicy) (result media.ContentKeyPolicy, err error)
}

var _ ContentKeyPoliciesClientAPI = (*media.ContentKeyPoliciesClient)(nil)

// TransformsClientAPI contains the set of methods on the TransformsClient type.
type TransformsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, transformName string, parameters media.Transform) (result media.Transform, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, transformName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, transformName string) (result media.Transform, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32) (result media.TransformCollectionPage, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, transformName string, parameters media.Transform) (result media.Transform, err error)
}

var _ TransformsClientAPI = (*media.TransformsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	CancelJob(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string, parameters media.Job) (result media.Job, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string) (result media.Job, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, transformName string, filter string, top *int32, skip *int32) (result media.JobCollectionPage, err error)
}

var _ JobsClientAPI = (*media.JobsClient)(nil)

// StreamingPoliciesClientAPI contains the set of methods on the StreamingPoliciesClient type.
type StreamingPoliciesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string, parameters media.StreamingPolicy) (result media.StreamingPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, streamingPolicyName string) (result media.StreamingPolicy, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result media.StreamingPolicyCollectionPage, err error)
}

var _ StreamingPoliciesClientAPI = (*media.StreamingPoliciesClient)(nil)

// StreamingLocatorsClientAPI contains the set of methods on the StreamingLocatorsClient type.
type StreamingLocatorsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string, parameters media.StreamingLocator) (result media.StreamingLocator, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result media.StreamingLocator, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result media.StreamingLocatorCollectionPage, err error)
	ListContentKeys(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result media.ListContentKeysResponse, err error)
	ListPaths(ctx context.Context, resourceGroupName string, accountName string, streamingLocatorName string) (result media.ListPathsResponse, err error)
}

var _ StreamingLocatorsClientAPI = (*media.StreamingLocatorsClient)(nil)

// LiveEventsClientAPI contains the set of methods on the LiveEventsClient type.
type LiveEventsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters media.LiveEvent, autoStart *bool) (result media.LiveEventsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result media.LiveEventsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result media.LiveEvent, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result media.LiveEventListResultPage, err error)
	Reset(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result media.LiveEventsResetFuture, err error)
	Start(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result media.LiveEventsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters media.LiveEventActionInput) (result media.LiveEventsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters media.LiveEvent) (result media.LiveEventsUpdateFuture, err error)
}

var _ LiveEventsClientAPI = (*media.LiveEventsClient)(nil)

// LiveOutputsClientAPI contains the set of methods on the LiveOutputsClient type.
type LiveOutputsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters media.LiveOutput) (result media.LiveOutputsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string) (result media.LiveOutputsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string) (result media.LiveOutput, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, liveEventName string) (result media.LiveOutputListResultPage, err error)
}

var _ LiveOutputsClientAPI = (*media.LiveOutputsClient)(nil)

// StreamingEndpointsClientAPI contains the set of methods on the StreamingEndpointsClient type.
type StreamingEndpointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters media.StreamingEndpoint, autoStart *bool) (result media.StreamingEndpointsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result media.StreamingEndpointsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result media.StreamingEndpoint, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result media.StreamingEndpointListResultPage, err error)
	Scale(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters media.StreamingEntityScaleUnit) (result media.StreamingEndpointsScaleFuture, err error)
	Start(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result media.StreamingEndpointsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string) (result media.StreamingEndpointsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, streamingEndpointName string, parameters media.StreamingEndpoint) (result media.StreamingEndpointsUpdateFuture, err error)
}

var _ StreamingEndpointsClientAPI = (*media.StreamingEndpointsClient)(nil)
