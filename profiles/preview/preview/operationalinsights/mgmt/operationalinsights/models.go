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

package operationalinsights

import (
	"context"

	original "github.com/tcz001/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DataSourceKind = original.DataSourceKind

const (
	AzureActivityLog              DataSourceKind = original.AzureActivityLog
	ChangeTrackingCustomRegistry  DataSourceKind = original.ChangeTrackingCustomRegistry
	ChangeTrackingDefaultPath     DataSourceKind = original.ChangeTrackingDefaultPath
	ChangeTrackingDefaultRegistry DataSourceKind = original.ChangeTrackingDefaultRegistry
	ChangeTrackingPath            DataSourceKind = original.ChangeTrackingPath
	CustomLog                     DataSourceKind = original.CustomLog
	CustomLogCollection           DataSourceKind = original.CustomLogCollection
	GenericDataSource             DataSourceKind = original.GenericDataSource
	IISLogs                       DataSourceKind = original.IISLogs
	LinuxPerformanceCollection    DataSourceKind = original.LinuxPerformanceCollection
	LinuxPerformanceObject        DataSourceKind = original.LinuxPerformanceObject
	LinuxSyslog                   DataSourceKind = original.LinuxSyslog
	LinuxSyslogCollection         DataSourceKind = original.LinuxSyslogCollection
	WindowsEvent                  DataSourceKind = original.WindowsEvent
	WindowsPerformanceCounter     DataSourceKind = original.WindowsPerformanceCounter
)

type EntityStatus = original.EntityStatus

const (
	Canceled            EntityStatus = original.Canceled
	Creating            EntityStatus = original.Creating
	Deleting            EntityStatus = original.Deleting
	Failed              EntityStatus = original.Failed
	ProvisioningAccount EntityStatus = original.ProvisioningAccount
	Succeeded           EntityStatus = original.Succeeded
)

type SkuNameEnum = original.SkuNameEnum

const (
	Free       SkuNameEnum = original.Free
	PerGB2018  SkuNameEnum = original.PerGB2018
	PerNode    SkuNameEnum = original.PerNode
	Premium    SkuNameEnum = original.Premium
	Standalone SkuNameEnum = original.Standalone
	Standard   SkuNameEnum = original.Standard
)

type BaseClient = original.BaseClient
type DataSource = original.DataSource
type DataSourceFilter = original.DataSourceFilter
type DataSourceListResult = original.DataSourceListResult
type DataSourceListResultIterator = original.DataSourceListResultIterator
type DataSourceListResultPage = original.DataSourceListResultPage
type DataSourcesClient = original.DataSourcesClient
type IntelligencePack = original.IntelligencePack
type LinkedService = original.LinkedService
type LinkedServiceListResult = original.LinkedServiceListResult
type LinkedServiceProperties = original.LinkedServiceProperties
type LinkedServicesClient = original.LinkedServicesClient
type ListIntelligencePack = original.ListIntelligencePack
type ManagementGroup = original.ManagementGroup
type ManagementGroupProperties = original.ManagementGroupProperties
type MetricName = original.MetricName
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SharedKeys = original.SharedKeys
type Sku = original.Sku
type UsageMetric = original.UsageMetric
type Workspace = original.Workspace
type WorkspaceListManagementGroupsResult = original.WorkspaceListManagementGroupsResult
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListUsagesResult = original.WorkspaceListUsagesResult
type WorkspaceProperties = original.WorkspaceProperties
type WorkspacesClient = original.WorkspacesClient
type WorkspacesCreateOrUpdateFuture = original.WorkspacesCreateOrUpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDataSourceListResultIterator(page DataSourceListResultPage) DataSourceListResultIterator {
	return original.NewDataSourceListResultIterator(page)
}
func NewDataSourceListResultPage(getNextPage func(context.Context, DataSourceListResult) (DataSourceListResult, error)) DataSourceListResultPage {
	return original.NewDataSourceListResultPage(getNextPage)
}
func NewDataSourcesClient(subscriptionID string) DataSourcesClient {
	return original.NewDataSourcesClient(subscriptionID)
}
func NewDataSourcesClientWithBaseURI(baseURI string, subscriptionID string) DataSourcesClient {
	return original.NewDataSourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedServicesClient(subscriptionID string) LinkedServicesClient {
	return original.NewLinkedServicesClient(subscriptionID)
}
func NewLinkedServicesClientWithBaseURI(baseURI string, subscriptionID string) LinkedServicesClient {
	return original.NewLinkedServicesClientWithBaseURI(baseURI, subscriptionID)
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
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleDataSourceKindValues() []DataSourceKind {
	return original.PossibleDataSourceKindValues()
}
func PossibleEntityStatusValues() []EntityStatus {
	return original.PossibleEntityStatusValues()
}
func PossibleSkuNameEnumValues() []SkuNameEnum {
	return original.PossibleSkuNameEnumValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
