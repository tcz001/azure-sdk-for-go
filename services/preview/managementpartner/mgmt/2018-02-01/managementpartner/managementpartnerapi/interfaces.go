package managementpartnerapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/managementpartner/mgmt/2018-02-01/managementpartner"
	"github.com/Azure/go-autorest/autorest"
)

// PartnerClientAPI contains the set of methods on the PartnerClient type.
type PartnerClientAPI interface {
	Create(ctx context.Context, partnerID string) (result managementpartner.PartnerResponse, err error)
	Delete(ctx context.Context, partnerID string) (result autorest.Response, err error)
	Get(ctx context.Context, partnerID string) (result managementpartner.PartnerResponse, err error)
	Update(ctx context.Context, partnerID string) (result managementpartner.PartnerResponse, err error)
}

var _ PartnerClientAPI = (*managementpartner.PartnerClient)(nil)

// OperationClientAPI contains the set of methods on the OperationClient type.
type OperationClientAPI interface {
	List(ctx context.Context) (result managementpartner.OperationListPage, err error)
}

var _ OperationClientAPI = (*managementpartner.OperationClient)(nil)

// PartnersClientAPI contains the set of methods on the PartnersClient type.
type PartnersClientAPI interface {
	Get(ctx context.Context) (result managementpartner.PartnerResponse, err error)
}

var _ PartnersClientAPI = (*managementpartner.PartnersClient)(nil)
