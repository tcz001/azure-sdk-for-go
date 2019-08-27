package devopsapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/devops/mgmt/2019-07-01-preview/devops"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result devops.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*devops.OperationsClient)(nil)

// PipelineTemplateDefinitionsClientAPI contains the set of methods on the PipelineTemplateDefinitionsClient type.
type PipelineTemplateDefinitionsClientAPI interface {
	List(ctx context.Context) (result devops.PipelineTemplateDefinitionListResultPage, err error)
}

var _ PipelineTemplateDefinitionsClientAPI = (*devops.PipelineTemplateDefinitionsClient)(nil)

// PipelinesClientAPI contains the set of methods on the PipelinesClient type.
type PipelinesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, pipelineName string, createOperationParameters devops.Pipeline) (result devops.PipelinesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, pipelineName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, pipelineName string) (result devops.Pipeline, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result devops.PipelineListResultPage, err error)
	ListBySubscription(ctx context.Context) (result devops.PipelineListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, pipelineName string, updateOperationParameters devops.PipelineUpdateParameters) (result devops.Pipeline, err error)
}

var _ PipelinesClientAPI = (*devops.PipelinesClient)(nil)
