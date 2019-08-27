package powerbidedicatedapi

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
	"github.com/tcz001/azure-sdk-for-go/services/powerbidedicated/mgmt/2017-10-01/powerbidedicated"
)

// CapacitiesClientAPI contains the set of methods on the CapacitiesClient type.
type CapacitiesClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, capacityParameters powerbidedicated.CheckCapacityNameAvailabilityParameters) (result powerbidedicated.CheckCapacityNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityParameters powerbidedicated.DedicatedCapacity) (result powerbidedicated.CapacitiesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.CapacitiesDeleteFuture, err error)
	GetDetails(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.DedicatedCapacity, err error)
	List(ctx context.Context) (result powerbidedicated.DedicatedCapacities, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result powerbidedicated.DedicatedCapacities, err error)
	ListSkus(ctx context.Context) (result powerbidedicated.SkuEnumerationForNewResourceResult, err error)
	ListSkusForCapacity(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.SkuEnumerationForExistingResourceResult, err error)
	Resume(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.CapacitiesResumeFuture, err error)
	Suspend(ctx context.Context, resourceGroupName string, dedicatedCapacityName string) (result powerbidedicated.CapacitiesSuspendFuture, err error)
	Update(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityUpdateParameters powerbidedicated.DedicatedCapacityUpdateParameters) (result powerbidedicated.CapacitiesUpdateFuture, err error)
}

var _ CapacitiesClientAPI = (*powerbidedicated.CapacitiesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result powerbidedicated.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*powerbidedicated.OperationsClient)(nil)
