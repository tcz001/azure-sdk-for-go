package alertsmanagementapi

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
	"github.com/tcz001/azure-sdk-for-go/services/alertsmanagement/mgmt/2019-03-01/alertsmanagement"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result alertsmanagement.OperationsListPage, err error)
}

var _ OperationsClientAPI = (*alertsmanagement.OperationsClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	ChangeState(ctx context.Context, alertID string, newState alertsmanagement.AlertState) (result alertsmanagement.Alert, err error)
	GetAll(ctx context.Context, targetResource string, targetResourceType string, targetResourceGroup string, monitorService alertsmanagement.MonitorService, monitorCondition alertsmanagement.MonitorCondition, severity alertsmanagement.Severity, alertState alertsmanagement.AlertState, alertRule string, smartGroupID string, includeContext *bool, includeEgressConfig *bool, pageCount *int32, sortBy alertsmanagement.AlertsSortByFields, sortOrder string, selectParameter string, timeRange alertsmanagement.TimeRange, customTimeRange string) (result alertsmanagement.AlertsListPage, err error)
	GetByID(ctx context.Context, alertID string) (result alertsmanagement.Alert, err error)
	GetHistory(ctx context.Context, alertID string) (result alertsmanagement.AlertModification, err error)
	GetSummary(ctx context.Context, groupby alertsmanagement.AlertsSummaryGroupByFields, includeSmartGroupsCount *bool, targetResource string, targetResourceType string, targetResourceGroup string, monitorService alertsmanagement.MonitorService, monitorCondition alertsmanagement.MonitorCondition, severity alertsmanagement.Severity, alertState alertsmanagement.AlertState, alertRule string, timeRange alertsmanagement.TimeRange, customTimeRange string) (result alertsmanagement.AlertsSummary, err error)
}

var _ AlertsClientAPI = (*alertsmanagement.AlertsClient)(nil)

// SmartGroupsClientAPI contains the set of methods on the SmartGroupsClient type.
type SmartGroupsClientAPI interface {
	ChangeState(ctx context.Context, smartGroupID string, newState alertsmanagement.AlertState) (result alertsmanagement.SmartGroup, err error)
	GetAll(ctx context.Context, targetResource string, targetResourceGroup string, targetResourceType string, monitorService alertsmanagement.MonitorService, monitorCondition alertsmanagement.MonitorCondition, severity alertsmanagement.Severity, smartGroupState alertsmanagement.AlertState, timeRange alertsmanagement.TimeRange, pageCount *int32, sortBy alertsmanagement.SmartGroupsSortByFields, sortOrder string) (result alertsmanagement.SmartGroupsList, err error)
	GetByID(ctx context.Context, smartGroupID string) (result alertsmanagement.SmartGroup, err error)
	GetHistory(ctx context.Context, smartGroupID string) (result alertsmanagement.SmartGroupModification, err error)
}

var _ SmartGroupsClientAPI = (*alertsmanagement.SmartGroupsClient)(nil)

// SmartDetectorAlertRulesClientAPI contains the set of methods on the SmartDetectorAlertRulesClient type.
type SmartDetectorAlertRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, alertRuleName string, parameters alertsmanagement.AlertRule) (result alertsmanagement.AlertRule, err error)
	Delete(ctx context.Context, resourceGroupName string, alertRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, alertRuleName string, expandDetector *bool) (result alertsmanagement.AlertRule, err error)
	List(ctx context.Context) (result alertsmanagement.AlertRulesListPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result alertsmanagement.AlertRulesListPage, err error)
}

var _ SmartDetectorAlertRulesClientAPI = (*alertsmanagement.SmartDetectorAlertRulesClient)(nil)
