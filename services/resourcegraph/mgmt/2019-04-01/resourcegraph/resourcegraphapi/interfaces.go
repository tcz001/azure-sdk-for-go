package resourcegraphapi

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
	"github.com/tcz001/azure-sdk-for-go/services/resourcegraph/mgmt/2019-04-01/resourcegraph"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	Resources(ctx context.Context, query resourcegraph.QueryRequest) (result resourcegraph.QueryResponse, err error)
}

var _ BaseClientAPI = (*resourcegraph.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result resourcegraph.OperationListResult, err error)
}

var _ OperationsClientAPI = (*resourcegraph.OperationsClient)(nil)
