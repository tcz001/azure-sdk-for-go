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

package insights

import original "github.com/tcz001/azure-sdk-for-go/services/preview/appinsights/v1/insights"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type EventType = original.EventType

const (
	All                 EventType = original.All
	AvailabilityResults EventType = original.AvailabilityResults
	BrowserTimings      EventType = original.BrowserTimings
	CustomEvents        EventType = original.CustomEvents
	CustomMetrics       EventType = original.CustomMetrics
	Dependencies        EventType = original.Dependencies
	Exceptions          EventType = original.Exceptions
	PageViews           EventType = original.PageViews
	PerformanceCounters EventType = original.PerformanceCounters
	Requests            EventType = original.Requests
	Traces              EventType = original.Traces
)

type MetricID = original.MetricID

const (
	AvailabilityResultsavailabilityPercentage  MetricID = original.AvailabilityResultsavailabilityPercentage
	AvailabilityResultsduration                MetricID = original.AvailabilityResultsduration
	BillingtelemetryCount                      MetricID = original.BillingtelemetryCount
	ClientnetworkDuration                      MetricID = original.ClientnetworkDuration
	ClientprocessingDuration                   MetricID = original.ClientprocessingDuration
	ClientreceiveDuration                      MetricID = original.ClientreceiveDuration
	ClientsendDuration                         MetricID = original.ClientsendDuration
	ClienttotalDuration                        MetricID = original.ClienttotalDuration
	CustomEventscount                          MetricID = original.CustomEventscount
	Dependenciescount                          MetricID = original.Dependenciescount
	Dependenciesduration                       MetricID = original.Dependenciesduration
	Dependenciesfailed                         MetricID = original.Dependenciesfailed
	Exceptionsbrowser                          MetricID = original.Exceptionsbrowser
	Exceptionscount                            MetricID = original.Exceptionscount
	Exceptionsserver                           MetricID = original.Exceptionsserver
	PageViewscount                             MetricID = original.PageViewscount
	PageViewsduration                          MetricID = original.PageViewsduration
	PerformanceCountersexceptionsPerSecond     MetricID = original.PerformanceCountersexceptionsPerSecond
	PerformanceCountersmemoryAvailableBytes    MetricID = original.PerformanceCountersmemoryAvailableBytes
	PerformanceCountersprocessCPUPercentage    MetricID = original.PerformanceCountersprocessCPUPercentage
	PerformanceCountersprocessIOBytesPerSecond MetricID = original.PerformanceCountersprocessIOBytesPerSecond
	PerformanceCountersprocessorCPUPercentage  MetricID = original.PerformanceCountersprocessorCPUPercentage
	PerformanceCountersprocessPrivateBytes     MetricID = original.PerformanceCountersprocessPrivateBytes
	PerformanceCountersrequestExecutionTime    MetricID = original.PerformanceCountersrequestExecutionTime
	PerformanceCountersrequestsInQueue         MetricID = original.PerformanceCountersrequestsInQueue
	PerformanceCountersrequestsPerSecond       MetricID = original.PerformanceCountersrequestsPerSecond
	Requestscount                              MetricID = original.Requestscount
	Requestsduration                           MetricID = original.Requestsduration
	Requestsfailed                             MetricID = original.Requestsfailed
	Sessionscount                              MetricID = original.Sessionscount
	Usersauthenticated                         MetricID = original.Usersauthenticated
	Userscount                                 MetricID = original.Userscount
)

type MetricsAggregation = original.MetricsAggregation

const (
	Avg    MetricsAggregation = original.Avg
	Count  MetricsAggregation = original.Count
	Max    MetricsAggregation = original.Max
	Min    MetricsAggregation = original.Min
	Sum    MetricsAggregation = original.Sum
	Unique MetricsAggregation = original.Unique
)

type MetricsSegment = original.MetricsSegment

const (
	ApplicationBuild                MetricsSegment = original.ApplicationBuild
	ApplicationVersion              MetricsSegment = original.ApplicationVersion
	AuthenticatedOrAnonymousTraffic MetricsSegment = original.AuthenticatedOrAnonymousTraffic
	Browser                         MetricsSegment = original.Browser
	BrowserVersion                  MetricsSegment = original.BrowserVersion
	City                            MetricsSegment = original.City
	CloudRoleName                   MetricsSegment = original.CloudRoleName
	CloudServiceName                MetricsSegment = original.CloudServiceName
	Continent                       MetricsSegment = original.Continent
	CountryOrRegion                 MetricsSegment = original.CountryOrRegion
	DeploymentID                    MetricsSegment = original.DeploymentID
	DeploymentUnit                  MetricsSegment = original.DeploymentUnit
	DeviceType                      MetricsSegment = original.DeviceType
	Environment                     MetricsSegment = original.Environment
	HostingLocation                 MetricsSegment = original.HostingLocation
	InstanceName                    MetricsSegment = original.InstanceName
)

type Type = original.Type

const (
	TypeAvailabilityResult Type = original.TypeAvailabilityResult
	TypeBrowserTiming      Type = original.TypeBrowserTiming
	TypeCustomEvent        Type = original.TypeCustomEvent
	TypeCustomMetric       Type = original.TypeCustomMetric
	TypeDependency         Type = original.TypeDependency
	TypeEventsResultData   Type = original.TypeEventsResultData
	TypeException          Type = original.TypeException
	TypePageView           Type = original.TypePageView
	TypePerformanceCounter Type = original.TypePerformanceCounter
	TypeRequest            Type = original.TypeRequest
	TypeTrace              Type = original.TypeTrace
)

type BaseClient = original.BaseClient
type BasicEventsResultData = original.BasicEventsResultData
type Column = original.Column
type ErrorDetail = original.ErrorDetail
type ErrorInfo = original.ErrorInfo
type ErrorResponse = original.ErrorResponse
type EventsAiInfo = original.EventsAiInfo
type EventsApplicationInfo = original.EventsApplicationInfo
type EventsAvailabilityResultInfo = original.EventsAvailabilityResultInfo
type EventsAvailabilityResultResult = original.EventsAvailabilityResultResult
type EventsBrowserTimingInfo = original.EventsBrowserTimingInfo
type EventsBrowserTimingResult = original.EventsBrowserTimingResult
type EventsClient = original.EventsClient
type EventsClientInfo = original.EventsClientInfo
type EventsClientPerformanceInfo = original.EventsClientPerformanceInfo
type EventsCloudInfo = original.EventsCloudInfo
type EventsCustomEventInfo = original.EventsCustomEventInfo
type EventsCustomEventResult = original.EventsCustomEventResult
type EventsCustomMetricInfo = original.EventsCustomMetricInfo
type EventsCustomMetricResult = original.EventsCustomMetricResult
type EventsDependencyInfo = original.EventsDependencyInfo
type EventsDependencyResult = original.EventsDependencyResult
type EventsExceptionDetail = original.EventsExceptionDetail
type EventsExceptionDetailsParsedStack = original.EventsExceptionDetailsParsedStack
type EventsExceptionInfo = original.EventsExceptionInfo
type EventsExceptionResult = original.EventsExceptionResult
type EventsOperationInfo = original.EventsOperationInfo
type EventsPageViewInfo = original.EventsPageViewInfo
type EventsPageViewResult = original.EventsPageViewResult
type EventsPerformanceCounterInfo = original.EventsPerformanceCounterInfo
type EventsPerformanceCounterResult = original.EventsPerformanceCounterResult
type EventsRequestInfo = original.EventsRequestInfo
type EventsRequestResult = original.EventsRequestResult
type EventsResult = original.EventsResult
type EventsResultData = original.EventsResultData
type EventsResultDataCustomDimensions = original.EventsResultDataCustomDimensions
type EventsResultDataCustomMeasurements = original.EventsResultDataCustomMeasurements
type EventsResults = original.EventsResults
type EventsSessionInfo = original.EventsSessionInfo
type EventsTraceInfo = original.EventsTraceInfo
type EventsTraceResult = original.EventsTraceResult
type EventsUserInfo = original.EventsUserInfo
type ListMetricsResultsItem = original.ListMetricsResultsItem
type MetricsClient = original.MetricsClient
type MetricsPostBodySchema = original.MetricsPostBodySchema
type MetricsPostBodySchemaParameters = original.MetricsPostBodySchemaParameters
type MetricsResult = original.MetricsResult
type MetricsResultInfo = original.MetricsResultInfo
type MetricsResultsItem = original.MetricsResultsItem
type MetricsSegmentInfo = original.MetricsSegmentInfo
type QueryBody = original.QueryBody
type QueryClient = original.QueryClient
type QueryResults = original.QueryResults
type SetObject = original.SetObject
type Table = original.Table

func New() BaseClient {
	return original.New()
}
func NewEventsClient() EventsClient {
	return original.NewEventsClient()
}
func NewEventsClientWithBaseURI(baseURI string) EventsClient {
	return original.NewEventsClientWithBaseURI(baseURI)
}
func NewMetricsClient() MetricsClient {
	return original.NewMetricsClient()
}
func NewMetricsClientWithBaseURI(baseURI string) MetricsClient {
	return original.NewMetricsClientWithBaseURI(baseURI)
}
func NewQueryClient() QueryClient {
	return original.NewQueryClient()
}
func NewQueryClientWithBaseURI(baseURI string) QueryClient {
	return original.NewQueryClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleEventTypeValues() []EventType {
	return original.PossibleEventTypeValues()
}
func PossibleMetricIDValues() []MetricID {
	return original.PossibleMetricIDValues()
}
func PossibleMetricsAggregationValues() []MetricsAggregation {
	return original.PossibleMetricsAggregationValues()
}
func PossibleMetricsSegmentValues() []MetricsSegment {
	return original.PossibleMetricsSegmentValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
