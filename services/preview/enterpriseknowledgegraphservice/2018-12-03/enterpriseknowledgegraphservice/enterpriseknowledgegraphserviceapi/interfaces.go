package enterpriseknowledgegraphserviceapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/enterpriseknowledgegraphservice/2018-12-03/enterpriseknowledgegraphservice"
	"github.com/Azure/go-autorest/autorest"
)

// EnterpriseKnowledgeGraphClientAPI contains the set of methods on the EnterpriseKnowledgeGraphClient type.
type EnterpriseKnowledgeGraphClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, parameters enterpriseknowledgegraphservice.EnterpriseKnowledgeGraph) (result enterpriseknowledgegraphservice.EnterpriseKnowledgeGraph, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result enterpriseknowledgegraphservice.EnterpriseKnowledgeGraph, err error)
	List(ctx context.Context) (result enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphResponseListPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphResponseListPage, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters enterpriseknowledgegraphservice.EnterpriseKnowledgeGraph) (result enterpriseknowledgegraphservice.EnterpriseKnowledgeGraph, err error)
}

var _ EnterpriseKnowledgeGraphClientAPI = (*enterpriseknowledgegraphservice.EnterpriseKnowledgeGraphClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result enterpriseknowledgegraphservice.OperationEntityListResultPage, err error)
}

var _ OperationsClientAPI = (*enterpriseknowledgegraphservice.OperationsClient)(nil)
