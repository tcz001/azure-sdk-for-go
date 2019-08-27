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

package managementpartner

import (
	"context"

	original "github.com/tcz001/azure-sdk-for-go/services/preview/managementpartner/mgmt/2018-02-01/managementpartner"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Code = original.Code

const (
	BadRequest Code = original.BadRequest
	Conflict   Code = original.Conflict
	NotFound   Code = original.NotFound
)

type State = original.State

const (
	Active  State = original.Active
	Deleted State = original.Deleted
)

type BaseClient = original.BaseClient
type Error = original.Error
type ExtendedErrorInfo = original.ExtendedErrorInfo
type OperationClient = original.OperationClient
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationResponse = original.OperationResponse
type PartnerClient = original.PartnerClient
type PartnerProperties = original.PartnerProperties
type PartnerResponse = original.PartnerResponse
type PartnersClient = original.PartnersClient

func New() BaseClient {
	return original.New()
}
func NewOperationClient() OperationClient {
	return original.NewOperationClient()
}
func NewOperationClientWithBaseURI(baseURI string) OperationClient {
	return original.NewOperationClientWithBaseURI(baseURI)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(getNextPage)
}
func NewPartnerClient() PartnerClient {
	return original.NewPartnerClient()
}
func NewPartnerClientWithBaseURI(baseURI string) PartnerClient {
	return original.NewPartnerClientWithBaseURI(baseURI)
}
func NewPartnersClient() PartnersClient {
	return original.NewPartnersClient()
}
func NewPartnersClientWithBaseURI(baseURI string) PartnersClient {
	return original.NewPartnersClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleCodeValues() []Code {
	return original.PossibleCodeValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
