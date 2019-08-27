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

package datacatalog

import original "github.com/tcz001/azure-sdk-for-go/services/datacatalog/mgmt/2016-03-30/datacatalog"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type SkuType = original.SkuType

const (
	Free     SkuType = original.Free
	Standard SkuType = original.Standard
)

type ADCCatalog = original.ADCCatalog
type ADCCatalogProperties = original.ADCCatalogProperties
type ADCCatalogsClient = original.ADCCatalogsClient
type ADCCatalogsDeleteFuture = original.ADCCatalogsDeleteFuture
type ADCCatalogsListResult = original.ADCCatalogsListResult
type ADCOperationsClient = original.ADCOperationsClient
type BaseClient = original.BaseClient
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type Principals = original.Principals
type Resource = original.Resource

func New(subscriptionID string, catalogName string) BaseClient {
	return original.New(subscriptionID, catalogName)
}
func NewADCCatalogsClient(subscriptionID string, catalogName string) ADCCatalogsClient {
	return original.NewADCCatalogsClient(subscriptionID, catalogName)
}
func NewADCCatalogsClientWithBaseURI(baseURI string, subscriptionID string, catalogName string) ADCCatalogsClient {
	return original.NewADCCatalogsClientWithBaseURI(baseURI, subscriptionID, catalogName)
}
func NewADCOperationsClient(subscriptionID string, catalogName string) ADCOperationsClient {
	return original.NewADCOperationsClient(subscriptionID, catalogName)
}
func NewADCOperationsClientWithBaseURI(baseURI string, subscriptionID string, catalogName string) ADCOperationsClient {
	return original.NewADCOperationsClientWithBaseURI(baseURI, subscriptionID, catalogName)
}
func NewWithBaseURI(baseURI string, subscriptionID string, catalogName string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, catalogName)
}
func PossibleSkuTypeValues() []SkuType {
	return original.PossibleSkuTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
