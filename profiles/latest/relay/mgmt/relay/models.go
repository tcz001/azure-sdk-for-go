// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/tcz001/azure-sdk-for-go/tools/profileBuilder

package relay

import (
	"context"

	original "github.com/tcz001/azure-sdk-for-go/services/relay/mgmt/2017-04-01/relay"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessRights = original.AccessRights

const (
	Listen AccessRights = original.Listen
	Manage AccessRights = original.Manage
	Send   AccessRights = original.Send
)

type KeyType = original.KeyType

const (
	PrimaryKey   KeyType = original.PrimaryKey
	SecondaryKey KeyType = original.SecondaryKey
)

type ProvisioningStateEnum = original.ProvisioningStateEnum

const (
	Created   ProvisioningStateEnum = original.Created
	Deleted   ProvisioningStateEnum = original.Deleted
	Failed    ProvisioningStateEnum = original.Failed
	Succeeded ProvisioningStateEnum = original.Succeeded
	Unknown   ProvisioningStateEnum = original.Unknown
	Updating  ProvisioningStateEnum = original.Updating
)

type RelaytypeEnum = original.RelaytypeEnum

const (
	HTTP   RelaytypeEnum = original.HTTP
	NetTCP RelaytypeEnum = original.NetTCP
)

type SkuTier = original.SkuTier

const (
	Standard SkuTier = original.Standard
)

type UnavailableReason = original.UnavailableReason

const (
	InvalidName                           UnavailableReason = original.InvalidName
	NameInLockdown                        UnavailableReason = original.NameInLockdown
	NameInUse                             UnavailableReason = original.NameInUse
	None                                  UnavailableReason = original.None
	SubscriptionIsDisabled                UnavailableReason = original.SubscriptionIsDisabled
	TooManyNamespaceInCurrentSubscription UnavailableReason = original.TooManyNamespaceInCurrentSubscription
)

type AccessKeys = original.AccessKeys
type AuthorizationRule = original.AuthorizationRule
type AuthorizationRuleListResult = original.AuthorizationRuleListResult
type AuthorizationRuleListResultIterator = original.AuthorizationRuleListResultIterator
type AuthorizationRuleListResultPage = original.AuthorizationRuleListResultPage
type AuthorizationRuleProperties = original.AuthorizationRuleProperties
type BaseClient = original.BaseClient
type CheckNameAvailability = original.CheckNameAvailability
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ErrorResponse = original.ErrorResponse
type HybridConnection = original.HybridConnection
type HybridConnectionListResult = original.HybridConnectionListResult
type HybridConnectionListResultIterator = original.HybridConnectionListResultIterator
type HybridConnectionListResultPage = original.HybridConnectionListResultPage
type HybridConnectionProperties = original.HybridConnectionProperties
type HybridConnectionsClient = original.HybridConnectionsClient
type Namespace = original.Namespace
type NamespaceListResult = original.NamespaceListResult
type NamespaceListResultIterator = original.NamespaceListResultIterator
type NamespaceListResultPage = original.NamespaceListResultPage
type NamespaceProperties = original.NamespaceProperties
type NamespacesClient = original.NamespacesClient
type NamespacesCreateOrUpdateFuture = original.NamespacesCreateOrUpdateFuture
type NamespacesDeleteFuture = original.NamespacesDeleteFuture
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type RegenerateAccessKeyParameters = original.RegenerateAccessKeyParameters
type Resource = original.Resource
type ResourceNamespacePatch = original.ResourceNamespacePatch
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type UpdateParameters = original.UpdateParameters
type WCFRelaysClient = original.WCFRelaysClient
type WcfRelay = original.WcfRelay
type WcfRelayProperties = original.WcfRelayProperties
type WcfRelaysListResult = original.WcfRelaysListResult
type WcfRelaysListResultIterator = original.WcfRelaysListResultIterator
type WcfRelaysListResultPage = original.WcfRelaysListResultPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAuthorizationRuleListResultIterator(page AuthorizationRuleListResultPage) AuthorizationRuleListResultIterator {
	return original.NewAuthorizationRuleListResultIterator(page)
}
func NewAuthorizationRuleListResultPage(getNextPage func(context.Context, AuthorizationRuleListResult) (AuthorizationRuleListResult, error)) AuthorizationRuleListResultPage {
	return original.NewAuthorizationRuleListResultPage(getNextPage)
}
func NewHybridConnectionListResultIterator(page HybridConnectionListResultPage) HybridConnectionListResultIterator {
	return original.NewHybridConnectionListResultIterator(page)
}
func NewHybridConnectionListResultPage(getNextPage func(context.Context, HybridConnectionListResult) (HybridConnectionListResult, error)) HybridConnectionListResultPage {
	return original.NewHybridConnectionListResultPage(getNextPage)
}
func NewHybridConnectionsClient(subscriptionID string) HybridConnectionsClient {
	return original.NewHybridConnectionsClient(subscriptionID)
}
func NewHybridConnectionsClientWithBaseURI(baseURI string, subscriptionID string) HybridConnectionsClient {
	return original.NewHybridConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNamespaceListResultIterator(page NamespaceListResultPage) NamespaceListResultIterator {
	return original.NewNamespaceListResultIterator(page)
}
func NewNamespaceListResultPage(getNextPage func(context.Context, NamespaceListResult) (NamespaceListResult, error)) NamespaceListResultPage {
	return original.NewNamespaceListResultPage(getNextPage)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWCFRelaysClient(subscriptionID string) WCFRelaysClient {
	return original.NewWCFRelaysClient(subscriptionID)
}
func NewWCFRelaysClientWithBaseURI(baseURI string, subscriptionID string) WCFRelaysClient {
	return original.NewWCFRelaysClientWithBaseURI(baseURI, subscriptionID)
}
func NewWcfRelaysListResultIterator(page WcfRelaysListResultPage) WcfRelaysListResultIterator {
	return original.NewWcfRelaysListResultIterator(page)
}
func NewWcfRelaysListResultPage(getNextPage func(context.Context, WcfRelaysListResult) (WcfRelaysListResult, error)) WcfRelaysListResultPage {
	return original.NewWcfRelaysListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleProvisioningStateEnumValues() []ProvisioningStateEnum {
	return original.PossibleProvisioningStateEnumValues()
}
func PossibleRelaytypeEnumValues() []RelaytypeEnum {
	return original.PossibleRelaytypeEnumValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUnavailableReasonValues() []UnavailableReason {
	return original.PossibleUnavailableReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
