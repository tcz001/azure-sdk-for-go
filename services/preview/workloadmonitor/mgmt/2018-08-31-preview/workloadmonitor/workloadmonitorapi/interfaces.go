package workloadmonitorapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/workloadmonitor/mgmt/2018-08-31-preview/workloadmonitor"
	"github.com/satori/go.uuid"
)

// MonitorsClientAPI contains the set of methods on the MonitorsClient type.
type MonitorsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorID string) (result workloadmonitor.Monitor, err error)
	ListByResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, filter string, skiptoken string) (result workloadmonitor.MonitorsCollectionPage, err error)
	Update(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorID string, body workloadmonitor.Monitor) (result workloadmonitor.Monitor, err error)
}

var _ MonitorsClientAPI = (*workloadmonitor.MonitorsClient)(nil)

// ComponentsClientAPI contains the set of methods on the ComponentsClient type.
type ComponentsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, componentID uuid.UUID, selectParameter string, expand string) (result workloadmonitor.Component, err error)
	ListByResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result workloadmonitor.ComponentsCollectionPage, err error)
}

var _ ComponentsClientAPI = (*workloadmonitor.ComponentsClient)(nil)

// MonitorInstancesClientAPI contains the set of methods on the MonitorInstancesClient type.
type MonitorInstancesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorInstanceID uuid.UUID, selectParameter string, expand string) (result workloadmonitor.MonitorInstance, err error)
	ListByResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result workloadmonitor.MonitorInstancesCollectionPage, err error)
}

var _ MonitorInstancesClientAPI = (*workloadmonitor.MonitorInstancesClient)(nil)

// NotificationSettingsClientAPI contains the set of methods on the NotificationSettingsClient type.
type NotificationSettingsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result workloadmonitor.NotificationSetting, err error)
	ListByResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, skiptoken string) (result workloadmonitor.NotificationSettingsCollectionPage, err error)
	Update(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, body workloadmonitor.NotificationSetting) (result workloadmonitor.NotificationSetting, err error)
}

var _ NotificationSettingsClientAPI = (*workloadmonitor.NotificationSettingsClient)(nil)

// ComponentsSummaryClientAPI contains the set of methods on the ComponentsSummaryClient type.
type ComponentsSummaryClientAPI interface {
	List(ctx context.Context, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result workloadmonitor.ComponentsCollectionPage, err error)
}

var _ ComponentsSummaryClientAPI = (*workloadmonitor.ComponentsSummaryClient)(nil)

// MonitorInstancesSummaryClientAPI contains the set of methods on the MonitorInstancesSummaryClient type.
type MonitorInstancesSummaryClientAPI interface {
	List(ctx context.Context, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result workloadmonitor.MonitorInstancesCollectionPage, err error)
}

var _ MonitorInstancesSummaryClientAPI = (*workloadmonitor.MonitorInstancesSummaryClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context, skiptoken string) (result workloadmonitor.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*workloadmonitor.OperationsClient)(nil)
