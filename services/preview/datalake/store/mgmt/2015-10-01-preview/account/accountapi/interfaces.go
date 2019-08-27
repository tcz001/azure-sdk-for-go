package accountapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/datalake/store/mgmt/2015-10-01-preview/account"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeStoreAccount) (result account.CreateFuture, err error)
	CreateOrUpdateFirewallRule(ctx context.Context, resourceGroupName string, accountName string, name string, parameters account.FirewallRule) (result account.FirewallRule, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result account.DeleteFuture, err error)
	DeleteFirewallRule(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result autorest.Response, err error)
	EnableKeyVault(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeStoreAccount, err error)
	GetFirewallRule(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result account.FirewallRule, err error)
	List(ctx context.Context, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeStoreAccountListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeStoreAccountListResultPage, err error)
	ListFirewallRules(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeStoreFirewallRuleListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeStoreAccount) (result account.UpdateFuture, err error)
}

var _ ClientAPI = (*account.Client)(nil)
