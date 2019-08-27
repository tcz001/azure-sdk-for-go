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

package monitor

import original "github.com/tcz001/azure-sdk-for-go/services/preview/monitor/2018-09-01-preview/monitor"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type APIError = original.APIError
type APIFailureResponse = original.APIFailureResponse
type AzureMetricsBaseData = original.AzureMetricsBaseData
type AzureMetricsData = original.AzureMetricsData
type AzureMetricsDocument = original.AzureMetricsDocument
type AzureMetricsResult = original.AzureMetricsResult
type AzureTimeSeriesData = original.AzureTimeSeriesData
type BaseClient = original.BaseClient
type MetricsClient = original.MetricsClient

func New() BaseClient {
	return original.New()
}
func NewMetricsClient() MetricsClient {
	return original.NewMetricsClient()
}
func NewMetricsClientWithBaseURI(baseURI string) MetricsClient {
	return original.NewMetricsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
