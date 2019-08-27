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

package aad

import (
	"context"

	original "github.com/tcz001/azure-sdk-for-go/services/domainservices/mgmt/2017-06-01/aad"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ExternalAccess = original.ExternalAccess

const (
	Disabled ExternalAccess = original.Disabled
	Enabled  ExternalAccess = original.Enabled
)

type FilteredSync = original.FilteredSync

const (
	FilteredSyncDisabled FilteredSync = original.FilteredSyncDisabled
	FilteredSyncEnabled  FilteredSync = original.FilteredSyncEnabled
)

type Ldaps = original.Ldaps

const (
	LdapsDisabled Ldaps = original.LdapsDisabled
	LdapsEnabled  Ldaps = original.LdapsEnabled
)

type NotifyDcAdmins = original.NotifyDcAdmins

const (
	NotifyDcAdminsDisabled NotifyDcAdmins = original.NotifyDcAdminsDisabled
	NotifyDcAdminsEnabled  NotifyDcAdmins = original.NotifyDcAdminsEnabled
)

type NotifyGlobalAdmins = original.NotifyGlobalAdmins

const (
	NotifyGlobalAdminsDisabled NotifyGlobalAdmins = original.NotifyGlobalAdminsDisabled
	NotifyGlobalAdminsEnabled  NotifyGlobalAdmins = original.NotifyGlobalAdminsEnabled
)

type NtlmV1 = original.NtlmV1

const (
	NtlmV1Disabled NtlmV1 = original.NtlmV1Disabled
	NtlmV1Enabled  NtlmV1 = original.NtlmV1Enabled
)

type SyncNtlmPasswords = original.SyncNtlmPasswords

const (
	SyncNtlmPasswordsDisabled SyncNtlmPasswords = original.SyncNtlmPasswordsDisabled
	SyncNtlmPasswordsEnabled  SyncNtlmPasswords = original.SyncNtlmPasswordsEnabled
)

type TLSV1 = original.TLSV1

const (
	TLSV1Disabled TLSV1 = original.TLSV1Disabled
	TLSV1Enabled  TLSV1 = original.TLSV1Enabled
)

type BaseClient = original.BaseClient
type DomainSecuritySettings = original.DomainSecuritySettings
type DomainService = original.DomainService
type DomainServiceListResult = original.DomainServiceListResult
type DomainServiceListResultIterator = original.DomainServiceListResultIterator
type DomainServiceListResultPage = original.DomainServiceListResultPage
type DomainServiceOperationsClient = original.DomainServiceOperationsClient
type DomainServiceProperties = original.DomainServiceProperties
type DomainServicesClient = original.DomainServicesClient
type DomainServicesCreateOrUpdateFuture = original.DomainServicesCreateOrUpdateFuture
type DomainServicesDeleteFuture = original.DomainServicesDeleteFuture
type DomainServicesUpdateFuture = original.DomainServicesUpdateFuture
type HealthAlert = original.HealthAlert
type HealthMonitor = original.HealthMonitor
type LdapsSettings = original.LdapsSettings
type NotificationSettings = original.NotificationSettings
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type ReplicaSet = original.ReplicaSet
type ReplicaSetListResult = original.ReplicaSetListResult
type ReplicaSetListResultIterator = original.ReplicaSetListResultIterator
type ReplicaSetListResultPage = original.ReplicaSetListResultPage
type ReplicaSetProperties = original.ReplicaSetProperties
type ReplicaSetsClient = original.ReplicaSetsClient
type ReplicaSetsCreateOrUpdateFuture = original.ReplicaSetsCreateOrUpdateFuture
type ReplicaSetsDeleteFuture = original.ReplicaSetsDeleteFuture
type ReplicaSetsUpdateFuture = original.ReplicaSetsUpdateFuture
type Resource = original.Resource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDomainServiceListResultIterator(page DomainServiceListResultPage) DomainServiceListResultIterator {
	return original.NewDomainServiceListResultIterator(page)
}
func NewDomainServiceListResultPage(getNextPage func(context.Context, DomainServiceListResult) (DomainServiceListResult, error)) DomainServiceListResultPage {
	return original.NewDomainServiceListResultPage(getNextPage)
}
func NewDomainServiceOperationsClient(subscriptionID string) DomainServiceOperationsClient {
	return original.NewDomainServiceOperationsClient(subscriptionID)
}
func NewDomainServiceOperationsClientWithBaseURI(baseURI string, subscriptionID string) DomainServiceOperationsClient {
	return original.NewDomainServiceOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDomainServicesClient(subscriptionID string) DomainServicesClient {
	return original.NewDomainServicesClient(subscriptionID)
}
func NewDomainServicesClientWithBaseURI(baseURI string, subscriptionID string) DomainServicesClient {
	return original.NewDomainServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(getNextPage)
}
func NewReplicaSetListResultIterator(page ReplicaSetListResultPage) ReplicaSetListResultIterator {
	return original.NewReplicaSetListResultIterator(page)
}
func NewReplicaSetListResultPage(getNextPage func(context.Context, ReplicaSetListResult) (ReplicaSetListResult, error)) ReplicaSetListResultPage {
	return original.NewReplicaSetListResultPage(getNextPage)
}
func NewReplicaSetsClient(subscriptionID string) ReplicaSetsClient {
	return original.NewReplicaSetsClient(subscriptionID)
}
func NewReplicaSetsClientWithBaseURI(baseURI string, subscriptionID string) ReplicaSetsClient {
	return original.NewReplicaSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleExternalAccessValues() []ExternalAccess {
	return original.PossibleExternalAccessValues()
}
func PossibleFilteredSyncValues() []FilteredSync {
	return original.PossibleFilteredSyncValues()
}
func PossibleLdapsValues() []Ldaps {
	return original.PossibleLdapsValues()
}
func PossibleNotifyDcAdminsValues() []NotifyDcAdmins {
	return original.PossibleNotifyDcAdminsValues()
}
func PossibleNotifyGlobalAdminsValues() []NotifyGlobalAdmins {
	return original.PossibleNotifyGlobalAdminsValues()
}
func PossibleNtlmV1Values() []NtlmV1 {
	return original.PossibleNtlmV1Values()
}
func PossibleSyncNtlmPasswordsValues() []SyncNtlmPasswords {
	return original.PossibleSyncNtlmPasswordsValues()
}
func PossibleTLSV1Values() []TLSV1 {
	return original.PossibleTLSV1Values()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
