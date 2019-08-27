package customproviders

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
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/tcz001/azure-sdk-for-go/services/preview/customproviders/mgmt/2018-09-01-preview/customproviders"

// ActionRouting enumerates the values for action routing.
type ActionRouting string

const (
	// Proxy ...
	Proxy ActionRouting = "Proxy"
)

// PossibleActionRoutingValues returns an array of possible values for the ActionRouting const type.
func PossibleActionRoutingValues() []ActionRouting {
	return []ActionRouting{Proxy}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Deleting, Failed, Running, Succeeded}
}

// ResourceTypeRouting enumerates the values for resource type routing.
type ResourceTypeRouting string

const (
	// ResourceTypeRoutingProxy ...
	ResourceTypeRoutingProxy ResourceTypeRouting = "Proxy"
	// ResourceTypeRoutingProxyCache ...
	ResourceTypeRoutingProxyCache ResourceTypeRouting = "Proxy,Cache"
)

// PossibleResourceTypeRoutingValues returns an array of possible values for the ResourceTypeRouting const type.
func PossibleResourceTypeRoutingValues() []ResourceTypeRouting {
	return []ResourceTypeRouting{ResourceTypeRoutingProxy, ResourceTypeRoutingProxyCache}
}

// ValidationType enumerates the values for validation type.
type ValidationType string

const (
	// Swagger ...
	Swagger ValidationType = "Swagger"
)

// PossibleValidationTypeValues returns an array of possible values for the ValidationType const type.
func PossibleValidationTypeValues() []ValidationType {
	return []ValidationType{Swagger}
}

// Association the resource definition of this association.
type Association struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The association id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The association name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The association type.
	Type *string `json:"type,omitempty"`
	// AssociationProperties - The properties of the association.
	*AssociationProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Association.
func (a Association) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AssociationProperties != nil {
		objectMap["properties"] = a.AssociationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Association struct.
func (a *Association) UnmarshalJSON(body []byte) error {
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
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var associationProperties AssociationProperties
				err = json.Unmarshal(*v, &associationProperties)
				if err != nil {
					return err
				}
				a.AssociationProperties = &associationProperties
			}
		}
	}

	return nil
}

// AssociationProperties the properties of the association.
type AssociationProperties struct {
	// TargetResourceID - The REST resource instance of the target resource for this association.
	TargetResourceID *string `json:"targetResourceId,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state of the association. Possible values include: 'Accepted', 'Deleting', 'Running', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// AssociationsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type AssociationsCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AssociationsCreateOrUpdateFuture) Result(client AssociationsClient) (a Association, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("customproviders.AssociationsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if a.Response.Response, err = future.GetResult(sender); err == nil && a.Response.Response.StatusCode != http.StatusNoContent {
		a, err = client.CreateOrUpdateResponder(a.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customproviders.AssociationsCreateOrUpdateFuture", "Result", a.Response.Response, "Failure responding to request")
		}
	}
	return
}

// AssociationsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type AssociationsDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AssociationsDeleteFuture) Result(client AssociationsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("customproviders.AssociationsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// AssociationsList list of associations.
type AssociationsList struct {
	autorest.Response `json:"-"`
	// Value - The array of associations.
	Value *[]Association `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// AssociationsListIterator provides access to a complete listing of Association values.
type AssociationsListIterator struct {
	i    int
	page AssociationsListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AssociationsListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssociationsListIterator.NextWithContext")
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
func (iter *AssociationsListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AssociationsListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AssociationsListIterator) Response() AssociationsList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AssociationsListIterator) Value() Association {
	if !iter.page.NotDone() {
		return Association{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AssociationsListIterator type.
func NewAssociationsListIterator(page AssociationsListPage) AssociationsListIterator {
	return AssociationsListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (al AssociationsList) IsEmpty() bool {
	return al.Value == nil || len(*al.Value) == 0
}

// associationsListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (al AssociationsList) associationsListPreparer(ctx context.Context) (*http.Request, error) {
	if al.NextLink == nil || len(to.String(al.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(al.NextLink)))
}

// AssociationsListPage contains a page of Association values.
type AssociationsListPage struct {
	fn func(context.Context, AssociationsList) (AssociationsList, error)
	al AssociationsList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AssociationsListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssociationsListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.al)
	if err != nil {
		return err
	}
	page.al = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AssociationsListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AssociationsListPage) NotDone() bool {
	return !page.al.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AssociationsListPage) Response() AssociationsList {
	return page.al
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AssociationsListPage) Values() []Association {
	if page.al.IsEmpty() {
		return nil
	}
	return *page.al.Value
}

// Creates a new instance of the AssociationsListPage type.
func NewAssociationsListPage(getNextPage func(context.Context, AssociationsList) (AssociationsList, error)) AssociationsListPage {
	return AssociationsListPage{fn: getNextPage}
}

// CustomResourceProviderCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type CustomResourceProviderCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CustomResourceProviderCreateOrUpdateFuture) Result(client CustomResourceProviderClient) (crm CustomRPManifest, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("customproviders.CustomResourceProviderCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if crm.Response.Response, err = future.GetResult(sender); err == nil && crm.Response.Response.StatusCode != http.StatusNoContent {
		crm, err = client.CreateOrUpdateResponder(crm.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderCreateOrUpdateFuture", "Result", crm.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CustomResourceProviderDeleteFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type CustomResourceProviderDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CustomResourceProviderDeleteFuture) Result(client CustomResourceProviderClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("customproviders.CustomResourceProviderDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// CustomRPActionRouteDefinition the route definition for an action implemented by the custom resource
// provider.
type CustomRPActionRouteDefinition struct {
	// RoutingType - The routing types that are supported for action requests. Possible values include: 'Proxy'
	RoutingType ActionRouting `json:"routingType,omitempty"`
	// Name - The name of the route definition. This becomes the name for the ARM extension (e.g. '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}/{name}')
	Name *string `json:"name,omitempty"`
	// Endpoint - The route definition endpoint URI that the custom resource provider will proxy requests to. This can be in the form of a flat URI (e.g. 'https://testendpoint/') or can specify to route via a path (e.g. 'https://testendpoint/{requestPath}')
	Endpoint *string `json:"endpoint,omitempty"`
}

// CustomRPManifest a manifest file that defines the custom resource provider resources.
type CustomRPManifest struct {
	autorest.Response `json:"-"`
	// CustomRPManifestProperties - The manifest for the custom resource provider
	*CustomRPManifestProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for CustomRPManifest.
func (crm CustomRPManifest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if crm.CustomRPManifestProperties != nil {
		objectMap["properties"] = crm.CustomRPManifestProperties
	}
	if crm.Location != nil {
		objectMap["location"] = crm.Location
	}
	if crm.Tags != nil {
		objectMap["tags"] = crm.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CustomRPManifest struct.
func (crm *CustomRPManifest) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var customRPManifestProperties CustomRPManifestProperties
				err = json.Unmarshal(*v, &customRPManifestProperties)
				if err != nil {
					return err
				}
				crm.CustomRPManifestProperties = &customRPManifestProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				crm.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				crm.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				crm.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				crm.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				crm.Tags = tags
			}
		}
	}

	return nil
}

// CustomRPManifestProperties the manifest for the custom resource provider
type CustomRPManifestProperties struct {
	// Actions - A list of actions that the custom resource provider implements.
	Actions *[]CustomRPActionRouteDefinition `json:"actions,omitempty"`
	// ResourceTypes - A list of resource types that the custom resource provider implements.
	ResourceTypes *[]CustomRPResourceTypeRouteDefinition `json:"resourceTypes,omitempty"`
	// Validations - A list of validations to run on the custom resource provider's requests.
	Validations *[]CustomRPValidations `json:"validations,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state of the resource provider. Possible values include: 'Accepted', 'Deleting', 'Running', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// CustomRPResourceTypeRouteDefinition the route definition for a resource implemented by the custom
// resource provider.
type CustomRPResourceTypeRouteDefinition struct {
	// RoutingType - The routing types that are supported for resource requests. Possible values include: 'ResourceTypeRoutingProxy', 'ResourceTypeRoutingProxyCache'
	RoutingType ResourceTypeRouting `json:"routingType,omitempty"`
	// Name - The name of the route definition. This becomes the name for the ARM extension (e.g. '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}/{name}')
	Name *string `json:"name,omitempty"`
	// Endpoint - The route definition endpoint URI that the custom resource provider will proxy requests to. This can be in the form of a flat URI (e.g. 'https://testendpoint/') or can specify to route via a path (e.g. 'https://testendpoint/{requestPath}')
	Endpoint *string `json:"endpoint,omitempty"`
}

// CustomRPRouteDefinition a route definition that defines an action or resource that can be interacted
// with through the custom resource provider.
type CustomRPRouteDefinition struct {
	// Name - The name of the route definition. This becomes the name for the ARM extension (e.g. '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}/{name}')
	Name *string `json:"name,omitempty"`
	// Endpoint - The route definition endpoint URI that the custom resource provider will proxy requests to. This can be in the form of a flat URI (e.g. 'https://testendpoint/') or can specify to route via a path (e.g. 'https://testendpoint/{requestPath}')
	Endpoint *string `json:"endpoint,omitempty"`
}

// CustomRPValidations a validation to apply on custom resource provider requests.
type CustomRPValidations struct {
	// ValidationType - The type of validation to run against a matching request. Possible values include: 'Swagger'
	ValidationType ValidationType `json:"validationType,omitempty"`
	// Specification - A link to the validation specification. The specification must be hosted on raw.githubusercontent.com.
	Specification *string `json:"specification,omitempty"`
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
	// Details - READ-ONLY; Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// ListByCustomRPManifest list of custom resource providers.
type ListByCustomRPManifest struct {
	autorest.Response `json:"-"`
	// Value - The array of custom resource provider manifests.
	Value *[]CustomRPManifest `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListByCustomRPManifestIterator provides access to a complete listing of CustomRPManifest values.
type ListByCustomRPManifestIterator struct {
	i    int
	page ListByCustomRPManifestPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListByCustomRPManifestIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListByCustomRPManifestIterator.NextWithContext")
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
func (iter *ListByCustomRPManifestIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListByCustomRPManifestIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListByCustomRPManifestIterator) Response() ListByCustomRPManifest {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListByCustomRPManifestIterator) Value() CustomRPManifest {
	if !iter.page.NotDone() {
		return CustomRPManifest{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListByCustomRPManifestIterator type.
func NewListByCustomRPManifestIterator(page ListByCustomRPManifestPage) ListByCustomRPManifestIterator {
	return ListByCustomRPManifestIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lbcrm ListByCustomRPManifest) IsEmpty() bool {
	return lbcrm.Value == nil || len(*lbcrm.Value) == 0
}

// listByCustomRPManifestPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lbcrm ListByCustomRPManifest) listByCustomRPManifestPreparer(ctx context.Context) (*http.Request, error) {
	if lbcrm.NextLink == nil || len(to.String(lbcrm.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lbcrm.NextLink)))
}

// ListByCustomRPManifestPage contains a page of CustomRPManifest values.
type ListByCustomRPManifestPage struct {
	fn    func(context.Context, ListByCustomRPManifest) (ListByCustomRPManifest, error)
	lbcrm ListByCustomRPManifest
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListByCustomRPManifestPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListByCustomRPManifestPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.lbcrm)
	if err != nil {
		return err
	}
	page.lbcrm = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListByCustomRPManifestPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListByCustomRPManifestPage) NotDone() bool {
	return !page.lbcrm.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListByCustomRPManifestPage) Response() ListByCustomRPManifest {
	return page.lbcrm
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListByCustomRPManifestPage) Values() []CustomRPManifest {
	if page.lbcrm.IsEmpty() {
		return nil
	}
	return *page.lbcrm.Value
}

// Creates a new instance of the ListByCustomRPManifestPage type.
func NewListByCustomRPManifestPage(getNextPage func(context.Context, ListByCustomRPManifest) (ListByCustomRPManifest, error)) ListByCustomRPManifestPage {
	return ListByCustomRPManifestPage{fn: getNextPage}
}

// Resource the resource definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceProviderOperation supported operations of this resource provider.
type ResourceProviderOperation struct {
	// Name - Operation name, in format of {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`
}

// ResourceProviderOperationDisplay display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Provider - Resource provider: Microsoft Custom Providers.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of this operation.
	Description *string `json:"description,omitempty"`
}

// ResourceProviderOperationList results of the request to list operations.
type ResourceProviderOperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by this resource provider.
	Value *[]ResourceProviderOperation `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceProviderOperationListIterator provides access to a complete listing of ResourceProviderOperation
// values.
type ResourceProviderOperationListIterator struct {
	i    int
	page ResourceProviderOperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceProviderOperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListIterator.NextWithContext")
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
func (iter *ResourceProviderOperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceProviderOperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceProviderOperationListIterator) Response() ResourceProviderOperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceProviderOperationListIterator) Value() ResourceProviderOperation {
	if !iter.page.NotDone() {
		return ResourceProviderOperation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceProviderOperationListIterator type.
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return ResourceProviderOperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rpol ResourceProviderOperationList) IsEmpty() bool {
	return rpol.Value == nil || len(*rpol.Value) == 0
}

// resourceProviderOperationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rpol ResourceProviderOperationList) resourceProviderOperationListPreparer(ctx context.Context) (*http.Request, error) {
	if rpol.NextLink == nil || len(to.String(rpol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rpol.NextLink)))
}

// ResourceProviderOperationListPage contains a page of ResourceProviderOperation values.
type ResourceProviderOperationListPage struct {
	fn   func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)
	rpol ResourceProviderOperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceProviderOperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rpol)
	if err != nil {
		return err
	}
	page.rpol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceProviderOperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceProviderOperationListPage) NotDone() bool {
	return !page.rpol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceProviderOperationListPage) Response() ResourceProviderOperationList {
	return page.rpol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceProviderOperationListPage) Values() []ResourceProviderOperation {
	if page.rpol.IsEmpty() {
		return nil
	}
	return *page.rpol.Value
}

// Creates a new instance of the ResourceProviderOperationListPage type.
func NewResourceProviderOperationListPage(getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return ResourceProviderOperationListPage{fn: getNextPage}
}

// ResourceProvidersUpdate custom resource provider update information.
type ResourceProvidersUpdate struct {
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ResourceProvidersUpdate.
func (rpu ResourceProvidersUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rpu.Tags != nil {
		objectMap["tags"] = rpu.Tags
	}
	return json.Marshal(objectMap)
}
