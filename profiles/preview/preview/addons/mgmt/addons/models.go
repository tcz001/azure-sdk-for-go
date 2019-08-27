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

package addons

import original "github.com/tcz001/azure-sdk-for-go/services/preview/addons/mgmt/2018-03-01/addons"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type OneTimeCharge = original.OneTimeCharge

const (
	No          OneTimeCharge = original.No
	OnEnabled   OneTimeCharge = original.OnEnabled
	OnReenabled OneTimeCharge = original.OnReenabled
)

type PlanTypeName = original.PlanTypeName

const (
	Advanced  PlanTypeName = original.Advanced
	Essential PlanTypeName = original.Essential
	Standard  PlanTypeName = original.Standard
)

type ProvisioningState = original.ProvisioningState

const (
	Cancelled   ProvisioningState = original.Cancelled
	Cancelling  ProvisioningState = original.Cancelling
	Downgrading ProvisioningState = original.Downgrading
	Failed      ProvisioningState = original.Failed
	Purchasing  ProvisioningState = original.Purchasing
	Succeeded   ProvisioningState = original.Succeeded
	Upgrading   ProvisioningState = original.Upgrading
)

type SupportPlanType = original.SupportPlanType

const (
	SupportPlanTypeAdvanced  SupportPlanType = original.SupportPlanTypeAdvanced
	SupportPlanTypeEssential SupportPlanType = original.SupportPlanTypeEssential
	SupportPlanTypeStandard  SupportPlanType = original.SupportPlanTypeStandard
)

type BaseClient = original.BaseClient
type CanonicalSupportPlanInfoDefinition = original.CanonicalSupportPlanInfoDefinition
type CanonicalSupportPlanProperties = original.CanonicalSupportPlanProperties
type CanonicalSupportPlanResponseEnvelope = original.CanonicalSupportPlanResponseEnvelope
type ErrorDefinition = original.ErrorDefinition
type ListCanonicalSupportPlanInfoDefinition = original.ListCanonicalSupportPlanInfoDefinition
type ListOperationsDefinition = original.ListOperationsDefinition
type OperationsClient = original.OperationsClient
type OperationsDefinition = original.OperationsDefinition
type OperationsDisplayDefinition = original.OperationsDisplayDefinition
type SupportPlanTypesClient = original.SupportPlanTypesClient
type SupportPlanTypesCreateOrUpdateFuture = original.SupportPlanTypesCreateOrUpdateFuture
type SupportPlanTypesDeleteFuture = original.SupportPlanTypesDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSupportPlanTypesClient(subscriptionID string) SupportPlanTypesClient {
	return original.NewSupportPlanTypesClient(subscriptionID)
}
func NewSupportPlanTypesClientWithBaseURI(baseURI string, subscriptionID string) SupportPlanTypesClient {
	return original.NewSupportPlanTypesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleOneTimeChargeValues() []OneTimeCharge {
	return original.PossibleOneTimeChargeValues()
}
func PossiblePlanTypeNameValues() []PlanTypeName {
	return original.PossiblePlanTypeNameValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSupportPlanTypeValues() []SupportPlanType {
	return original.PossibleSupportPlanTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
