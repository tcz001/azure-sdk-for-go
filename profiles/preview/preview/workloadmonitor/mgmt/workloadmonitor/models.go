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

package workloadmonitor

import (
	"context"

	original "github.com/tcz001/azure-sdk-for-go/services/preview/workloadmonitor/mgmt/2018-08-31-preview/workloadmonitor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlertGeneration = original.AlertGeneration

const (
	Disabled AlertGeneration = original.Disabled
	Enabled  AlertGeneration = original.Enabled
)

type HealthState = original.HealthState

const (
	Error         HealthState = original.Error
	Success       HealthState = original.Success
	Uninitialized HealthState = original.Uninitialized
	Unknown       HealthState = original.Unknown
	Warning       HealthState = original.Warning
)

type HealthStateCategory = original.HealthStateCategory

const (
	CustomGroup HealthStateCategory = original.CustomGroup
	Identity    HealthStateCategory = original.Identity
)

type MonitorCategory = original.MonitorCategory

const (
	AvailabilityHealth MonitorCategory = original.AvailabilityHealth
	Configuration      MonitorCategory = original.Configuration
	EntityHealth       MonitorCategory = original.EntityHealth
	PerformanceHealth  MonitorCategory = original.PerformanceHealth
	Security           MonitorCategory = original.Security
)

type MonitorState = original.MonitorState

const (
	MonitorStateDisabled MonitorState = original.MonitorStateDisabled
	MonitorStateEnabled  MonitorState = original.MonitorStateEnabled
)

type MonitorType = original.MonitorType

const (
	Aggregate  MonitorType = original.Aggregate
	Dependency MonitorType = original.Dependency
	Unit       MonitorType = original.Unit
)

type Operator = original.Operator

const (
	Equals             Operator = original.Equals
	GreaterThan        Operator = original.GreaterThan
	GreaterThanOrEqual Operator = original.GreaterThanOrEqual
	LessThan           Operator = original.LessThan
	LessThanOrEqual    Operator = original.LessThanOrEqual
	NotEquals          Operator = original.NotEquals
)

type WorkloadType = original.WorkloadType

const (
	Apache WorkloadType = original.Apache
	BaseOS WorkloadType = original.BaseOS
	IIS    WorkloadType = original.IIS
	SQL    WorkloadType = original.SQL
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type Component = original.Component
type ComponentProperties = original.ComponentProperties
type ComponentsClient = original.ComponentsClient
type ComponentsCollection = original.ComponentsCollection
type ComponentsCollectionIterator = original.ComponentsCollectionIterator
type ComponentsCollectionPage = original.ComponentsCollectionPage
type ComponentsSummaryClient = original.ComponentsSummaryClient
type ErrorFieldContract = original.ErrorFieldContract
type ErrorResponse = original.ErrorResponse
type HealthStateChange = original.HealthStateChange
type Monitor = original.Monitor
type MonitorCriteria = original.MonitorCriteria
type MonitorInstance = original.MonitorInstance
type MonitorInstanceProperties = original.MonitorInstanceProperties
type MonitorInstancesClient = original.MonitorInstancesClient
type MonitorInstancesCollection = original.MonitorInstancesCollection
type MonitorInstancesCollectionIterator = original.MonitorInstancesCollectionIterator
type MonitorInstancesCollectionPage = original.MonitorInstancesCollectionPage
type MonitorInstancesSummaryClient = original.MonitorInstancesSummaryClient
type MonitorProperties = original.MonitorProperties
type MonitorsClient = original.MonitorsClient
type MonitorsCollection = original.MonitorsCollection
type MonitorsCollectionIterator = original.MonitorsCollectionIterator
type MonitorsCollectionPage = original.MonitorsCollectionPage
type NotificationSetting = original.NotificationSetting
type NotificationSettingProperties = original.NotificationSettingProperties
type NotificationSettingsClient = original.NotificationSettingsClient
type NotificationSettingsCollection = original.NotificationSettingsCollection
type NotificationSettingsCollectionIterator = original.NotificationSettingsCollectionIterator
type NotificationSettingsCollectionPage = original.NotificationSettingsCollectionPage
type Operation = original.Operation
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewComponentsClient(subscriptionID string) ComponentsClient {
	return original.NewComponentsClient(subscriptionID)
}
func NewComponentsClientWithBaseURI(baseURI string, subscriptionID string) ComponentsClient {
	return original.NewComponentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewComponentsCollectionIterator(page ComponentsCollectionPage) ComponentsCollectionIterator {
	return original.NewComponentsCollectionIterator(page)
}
func NewComponentsCollectionPage(getNextPage func(context.Context, ComponentsCollection) (ComponentsCollection, error)) ComponentsCollectionPage {
	return original.NewComponentsCollectionPage(getNextPage)
}
func NewComponentsSummaryClient(subscriptionID string) ComponentsSummaryClient {
	return original.NewComponentsSummaryClient(subscriptionID)
}
func NewComponentsSummaryClientWithBaseURI(baseURI string, subscriptionID string) ComponentsSummaryClient {
	return original.NewComponentsSummaryClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitorInstancesClient(subscriptionID string) MonitorInstancesClient {
	return original.NewMonitorInstancesClient(subscriptionID)
}
func NewMonitorInstancesClientWithBaseURI(baseURI string, subscriptionID string) MonitorInstancesClient {
	return original.NewMonitorInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitorInstancesCollectionIterator(page MonitorInstancesCollectionPage) MonitorInstancesCollectionIterator {
	return original.NewMonitorInstancesCollectionIterator(page)
}
func NewMonitorInstancesCollectionPage(getNextPage func(context.Context, MonitorInstancesCollection) (MonitorInstancesCollection, error)) MonitorInstancesCollectionPage {
	return original.NewMonitorInstancesCollectionPage(getNextPage)
}
func NewMonitorInstancesSummaryClient(subscriptionID string) MonitorInstancesSummaryClient {
	return original.NewMonitorInstancesSummaryClient(subscriptionID)
}
func NewMonitorInstancesSummaryClientWithBaseURI(baseURI string, subscriptionID string) MonitorInstancesSummaryClient {
	return original.NewMonitorInstancesSummaryClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitorsClient(subscriptionID string) MonitorsClient {
	return original.NewMonitorsClient(subscriptionID)
}
func NewMonitorsClientWithBaseURI(baseURI string, subscriptionID string) MonitorsClient {
	return original.NewMonitorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitorsCollectionIterator(page MonitorsCollectionPage) MonitorsCollectionIterator {
	return original.NewMonitorsCollectionIterator(page)
}
func NewMonitorsCollectionPage(getNextPage func(context.Context, MonitorsCollection) (MonitorsCollection, error)) MonitorsCollectionPage {
	return original.NewMonitorsCollectionPage(getNextPage)
}
func NewNotificationSettingsClient(subscriptionID string) NotificationSettingsClient {
	return original.NewNotificationSettingsClient(subscriptionID)
}
func NewNotificationSettingsClientWithBaseURI(baseURI string, subscriptionID string) NotificationSettingsClient {
	return original.NewNotificationSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNotificationSettingsCollectionIterator(page NotificationSettingsCollectionPage) NotificationSettingsCollectionIterator {
	return original.NewNotificationSettingsCollectionIterator(page)
}
func NewNotificationSettingsCollectionPage(getNextPage func(context.Context, NotificationSettingsCollection) (NotificationSettingsCollection, error)) NotificationSettingsCollectionPage {
	return original.NewNotificationSettingsCollectionPage(getNextPage)
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
func PossibleAlertGenerationValues() []AlertGeneration {
	return original.PossibleAlertGenerationValues()
}
func PossibleHealthStateCategoryValues() []HealthStateCategory {
	return original.PossibleHealthStateCategoryValues()
}
func PossibleHealthStateValues() []HealthState {
	return original.PossibleHealthStateValues()
}
func PossibleMonitorCategoryValues() []MonitorCategory {
	return original.PossibleMonitorCategoryValues()
}
func PossibleMonitorStateValues() []MonitorState {
	return original.PossibleMonitorStateValues()
}
func PossibleMonitorTypeValues() []MonitorType {
	return original.PossibleMonitorTypeValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossibleWorkloadTypeValues() []WorkloadType {
	return original.PossibleWorkloadTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
