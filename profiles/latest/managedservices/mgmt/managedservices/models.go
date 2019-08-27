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

package managedservices

import (
	"context"

	original "github.com/tcz001/azure-sdk-for-go/services/managedservices/mgmt/2019-06-01/managedservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted     ProvisioningState = original.Accepted
	Canceled     ProvisioningState = original.Canceled
	Created      ProvisioningState = original.Created
	Creating     ProvisioningState = original.Creating
	Deleted      ProvisioningState = original.Deleted
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	NotSpecified ProvisioningState = original.NotSpecified
	Ready        ProvisioningState = original.Ready
	Running      ProvisioningState = original.Running
	Succeeded    ProvisioningState = original.Succeeded
	Updating     ProvisioningState = original.Updating
)

type Authorization = original.Authorization
type BaseClient = original.BaseClient
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationsClient = original.OperationsClient
type Plan = original.Plan
type RegistrationAssignment = original.RegistrationAssignment
type RegistrationAssignmentList = original.RegistrationAssignmentList
type RegistrationAssignmentListIterator = original.RegistrationAssignmentListIterator
type RegistrationAssignmentListPage = original.RegistrationAssignmentListPage
type RegistrationAssignmentProperties = original.RegistrationAssignmentProperties
type RegistrationAssignmentPropertiesRegistrationDefinition = original.RegistrationAssignmentPropertiesRegistrationDefinition
type RegistrationAssignmentPropertiesRegistrationDefinitionProperties = original.RegistrationAssignmentPropertiesRegistrationDefinitionProperties
type RegistrationAssignmentsClient = original.RegistrationAssignmentsClient
type RegistrationAssignmentsCreateOrUpdateFuture = original.RegistrationAssignmentsCreateOrUpdateFuture
type RegistrationAssignmentsDeleteFuture = original.RegistrationAssignmentsDeleteFuture
type RegistrationDefinition = original.RegistrationDefinition
type RegistrationDefinitionList = original.RegistrationDefinitionList
type RegistrationDefinitionListIterator = original.RegistrationDefinitionListIterator
type RegistrationDefinitionListPage = original.RegistrationDefinitionListPage
type RegistrationDefinitionProperties = original.RegistrationDefinitionProperties
type RegistrationDefinitionsClient = original.RegistrationDefinitionsClient
type RegistrationDefinitionsCreateOrUpdateFuture = original.RegistrationDefinitionsCreateOrUpdateFuture

func New() BaseClient {
	return original.New()
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewRegistrationAssignmentListIterator(page RegistrationAssignmentListPage) RegistrationAssignmentListIterator {
	return original.NewRegistrationAssignmentListIterator(page)
}
func NewRegistrationAssignmentListPage(getNextPage func(context.Context, RegistrationAssignmentList) (RegistrationAssignmentList, error)) RegistrationAssignmentListPage {
	return original.NewRegistrationAssignmentListPage(getNextPage)
}
func NewRegistrationAssignmentsClient() RegistrationAssignmentsClient {
	return original.NewRegistrationAssignmentsClient()
}
func NewRegistrationAssignmentsClientWithBaseURI(baseURI string) RegistrationAssignmentsClient {
	return original.NewRegistrationAssignmentsClientWithBaseURI(baseURI)
}
func NewRegistrationDefinitionListIterator(page RegistrationDefinitionListPage) RegistrationDefinitionListIterator {
	return original.NewRegistrationDefinitionListIterator(page)
}
func NewRegistrationDefinitionListPage(getNextPage func(context.Context, RegistrationDefinitionList) (RegistrationDefinitionList, error)) RegistrationDefinitionListPage {
	return original.NewRegistrationDefinitionListPage(getNextPage)
}
func NewRegistrationDefinitionsClient() RegistrationDefinitionsClient {
	return original.NewRegistrationDefinitionsClient()
}
func NewRegistrationDefinitionsClientWithBaseURI(baseURI string) RegistrationDefinitionsClient {
	return original.NewRegistrationDefinitionsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
