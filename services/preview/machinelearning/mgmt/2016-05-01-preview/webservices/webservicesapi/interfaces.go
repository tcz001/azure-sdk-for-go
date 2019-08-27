package webservicesapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/machinelearning/mgmt/2016-05-01-preview/webservices"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, webServiceName string, createOrUpdatePayload webservices.WebService) (result webservices.CreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, webServiceName string) (result webservices.WebService, err error)
	List(ctx context.Context, skiptoken string) (result webservices.PaginatedWebServicesListPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skiptoken string) (result webservices.PaginatedWebServicesListPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, webServiceName string) (result webservices.Keys, err error)
	Patch(ctx context.Context, resourceGroupName string, webServiceName string, patchPayload webservices.WebService) (result webservices.PatchFuture, err error)
	Remove(ctx context.Context, resourceGroupName string, webServiceName string) (result webservices.RemoveFuture, err error)
}

var _ ClientAPI = (*webservices.Client)(nil)
