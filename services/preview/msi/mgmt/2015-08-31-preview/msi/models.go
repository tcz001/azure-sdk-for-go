package msi

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/tcz001/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"

// UserAssignedIdentities enumerates the values for user assigned identities.
type UserAssignedIdentities string

const (
	// MicrosoftManagedIdentityuserAssignedIdentities ...
	MicrosoftManagedIdentityuserAssignedIdentities UserAssignedIdentities = "Microsoft.ManagedIdentity/userAssignedIdentities"
)

// PossibleUserAssignedIdentitiesValues returns an array of possible values for the UserAssignedIdentities const type.
func PossibleUserAssignedIdentitiesValues() []UserAssignedIdentities {
	return []UserAssignedIdentities{MicrosoftManagedIdentityuserAssignedIdentities}
}

// CloudError an error response from the ManagedServiceIdentity service.
type CloudError struct {
	// Error - A list of additional details about the error.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody an error response from the ManagedServiceIdentity service.
type CloudErrorBody struct {
	// Code - An identifier for the error.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// Identity describes an identity resource.
type Identity struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The id of the created identity.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the created identity.
	Name *string `json:"name,omitempty"`
	// Location - The Azure region where the identity lives.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// IdentityProperties - The properties associated with the identity.
	*IdentityProperties `json:"properties,omitempty"`
	// Type - READ-ONLY; The type of resource i.e. Microsoft.ManagedIdentity/userAssignedIdentities. Possible values include: 'MicrosoftManagedIdentityuserAssignedIdentities'
	Type UserAssignedIdentities `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.Location != nil {
		objectMap["location"] = i.Location
	}
	if i.Tags != nil {
		objectMap["tags"] = i.Tags
	}
	if i.IdentityProperties != nil {
		objectMap["properties"] = i.IdentityProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Identity struct.
func (i *Identity) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				i.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				i.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				i.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				i.Tags = tags
			}
		case "properties":
			if v != nil {
				var identityProperties IdentityProperties
				err = json.Unmarshal(*v, &identityProperties)
				if err != nil {
					return err
				}
				i.IdentityProperties = &identityProperties
			}
		case "type":
			if v != nil {
				var typeVar UserAssignedIdentities
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				i.Type = typeVar
			}
		}
	}

	return nil
}

// IdentityProperties the properties associated with the identity.
type IdentityProperties struct {
	// TenantID - READ-ONLY; The id of the tenant which the identity belongs to.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// PrincipalID - READ-ONLY; The id of the service principal object associated with the created identity.
	PrincipalID *uuid.UUID `json:"principalId,omitempty"`
	// ClientID - READ-ONLY; The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientID *uuid.UUID `json:"clientId,omitempty"`
	// ClientSecretURL - READ-ONLY;  The ManagedServiceIdentity DataPlane URL that can be queried to obtain the identity credentials.
	ClientSecretURL *string `json:"clientSecretUrl,omitempty"`
}

// Operation operation supported by the Microsoft.ManagedIdentity REST API.
type Operation struct {
	// Name - The name of the REST Operation. This is of the format {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that describes the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that describes the operation.
type OperationDisplay struct {
	// Provider - Friendly name of the resource provider.
	Provider *string `json:"provider,omitempty"`
	// Operation - The type of operation. For example: read, write, delete.
	Operation *string `json:"operation,omitempty"`
	// Resource - The resource type on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Description - A description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult a list of operations supported by Microsoft.ManagedIdentity Resource Provider.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - A list of operations supported by Microsoft.ManagedIdentity Resource Provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The url to get the next page of results, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// UserAssignedIdentitiesListResult values returned by the List operation.
type UserAssignedIdentitiesListResult struct {
	autorest.Response `json:"-"`
	// Value - The collection of userAssignedIdentities returned by the listing operation.
	Value *[]Identity `json:"value,omitempty"`
	// NextLink - The url to get the next page of results, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// UserAssignedIdentitiesListResultIterator provides access to a complete listing of Identity values.
type UserAssignedIdentitiesListResultIterator struct {
	i    int
	page UserAssignedIdentitiesListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *UserAssignedIdentitiesListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserAssignedIdentitiesListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *UserAssignedIdentitiesListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter UserAssignedIdentitiesListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter UserAssignedIdentitiesListResultIterator) Response() UserAssignedIdentitiesListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter UserAssignedIdentitiesListResultIterator) Value() Identity {
	if !iter.page.NotDone() {
		return Identity{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the UserAssignedIdentitiesListResultIterator type.
func NewUserAssignedIdentitiesListResultIterator(page UserAssignedIdentitiesListResultPage) UserAssignedIdentitiesListResultIterator {
	return UserAssignedIdentitiesListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (uailr UserAssignedIdentitiesListResult) IsEmpty() bool {
	return uailr.Value == nil || len(*uailr.Value) == 0
}

// userAssignedIdentitiesListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (uailr UserAssignedIdentitiesListResult) userAssignedIdentitiesListResultPreparer(ctx context.Context) (*http.Request, error) {
	if uailr.NextLink == nil || len(to.String(uailr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(uailr.NextLink)))
}

// UserAssignedIdentitiesListResultPage contains a page of Identity values.
type UserAssignedIdentitiesListResultPage struct {
	fn    func(context.Context, UserAssignedIdentitiesListResult) (UserAssignedIdentitiesListResult, error)
	uailr UserAssignedIdentitiesListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *UserAssignedIdentitiesListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserAssignedIdentitiesListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.uailr)
	if err != nil {
		return err
	}
	page.uailr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *UserAssignedIdentitiesListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page UserAssignedIdentitiesListResultPage) NotDone() bool {
	return !page.uailr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page UserAssignedIdentitiesListResultPage) Response() UserAssignedIdentitiesListResult {
	return page.uailr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page UserAssignedIdentitiesListResultPage) Values() []Identity {
	if page.uailr.IsEmpty() {
		return nil
	}
	return *page.uailr.Value
}

// Creates a new instance of the UserAssignedIdentitiesListResultPage type.
func NewUserAssignedIdentitiesListResultPage(getNextPage func(context.Context, UserAssignedIdentitiesListResult) (UserAssignedIdentitiesListResult, error)) UserAssignedIdentitiesListResultPage {
	return UserAssignedIdentitiesListResultPage{fn: getNextPage}
}
