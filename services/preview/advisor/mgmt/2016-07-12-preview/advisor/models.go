package advisor

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/tcz001/azure-sdk-for-go/services/preview/advisor/mgmt/2016-07-12-preview/advisor"

// Category enumerates the values for category.
type Category string

const (
	// Cost ...
	Cost Category = "Cost"
	// HighAvailability ...
	HighAvailability Category = "HighAvailability"
	// Performance ...
	Performance Category = "Performance"
	// Security ...
	Security Category = "Security"
)

// PossibleCategoryValues returns an array of possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{Cost, HighAvailability, Performance, Security}
}

// Impact enumerates the values for impact.
type Impact string

const (
	// High ...
	High Impact = "High"
	// Low ...
	Low Impact = "Low"
	// Medium ...
	Medium Impact = "Medium"
)

// PossibleImpactValues returns an array of possible values for the Impact const type.
func PossibleImpactValues() []Impact {
	return []Impact{High, Low, Medium}
}

// Risk enumerates the values for risk.
type Risk string

const (
	// Error ...
	Error Risk = "Error"
	// None ...
	None Risk = "None"
	// Warning ...
	Warning Risk = "Warning"
)

// PossibleRiskValues returns an array of possible values for the Risk const type.
func PossibleRiskValues() []Risk {
	return []Risk{Error, None, Warning}
}

// ListSuppressionContract ...
type ListSuppressionContract struct {
	autorest.Response `json:"-"`
	Value             *[]SuppressionContract `json:"value,omitempty"`
}

// OperationDisplayInfo the operation supported by Advisor.
type OperationDisplayInfo struct {
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
	// Operation - The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Provider - Service provider: Microsoft Advisor.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity the operation supported by Advisor.
type OperationEntity struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The operation supported by Advisor.
	Display *OperationDisplayInfo `json:"display,omitempty"`
}

// OperationEntityListResult the list of Advisor operations.
type OperationEntityListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - The list of operations.
	Value *[]OperationEntity `json:"value,omitempty"`
}

// OperationEntityListResultIterator provides access to a complete listing of OperationEntity values.
type OperationEntityListResultIterator struct {
	i    int
	page OperationEntityListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationEntityListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationEntityListResultIterator.NextWithContext")
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
func (iter *OperationEntityListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationEntityListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationEntityListResultIterator) Response() OperationEntityListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationEntityListResultIterator) Value() OperationEntity {
	if !iter.page.NotDone() {
		return OperationEntity{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationEntityListResultIterator type.
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return OperationEntityListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (oelr OperationEntityListResult) IsEmpty() bool {
	return oelr.Value == nil || len(*oelr.Value) == 0
}

// operationEntityListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (oelr OperationEntityListResult) operationEntityListResultPreparer(ctx context.Context) (*http.Request, error) {
	if oelr.NextLink == nil || len(to.String(oelr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(oelr.NextLink)))
}

// OperationEntityListResultPage contains a page of OperationEntity values.
type OperationEntityListResultPage struct {
	fn   func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)
	oelr OperationEntityListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationEntityListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationEntityListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.oelr)
	if err != nil {
		return err
	}
	page.oelr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationEntityListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationEntityListResultPage) NotDone() bool {
	return !page.oelr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationEntityListResultPage) Response() OperationEntityListResult {
	return page.oelr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationEntityListResultPage) Values() []OperationEntity {
	if page.oelr.IsEmpty() {
		return nil
	}
	return *page.oelr.Value
}

// Creates a new instance of the OperationEntityListResultPage type.
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return OperationEntityListResultPage{fn: getNextPage}
}

// RecommendationProperties the properties of the recommendation.
type RecommendationProperties struct {
	// Category - The category of the recommendation. Possible values include: 'HighAvailability', 'Security', 'Performance', 'Cost'
	Category Category `json:"category,omitempty"`
	// Impact - The business impact of the recommendation. Possible values include: 'High', 'Medium', 'Low'
	Impact Impact `json:"impact,omitempty"`
	// ImpactedField - The resource type identified by Advisor.
	ImpactedField *string `json:"impactedField,omitempty"`
	// ImpactedValue - The resource identified by Advisor.
	ImpactedValue *string `json:"impactedValue,omitempty"`
	// LastUpdated - The most recent time that Advisor checked the validity of the recommendation.
	LastUpdated *date.Time `json:"lastUpdated,omitempty"`
	// Metadata - The recommendation metadata.
	Metadata map[string]interface{} `json:"metadata"`
	// RecommendationTypeID - The recommendation-type GUID.
	RecommendationTypeID *string `json:"recommendationTypeId,omitempty"`
	// Risk - The potential risk of not implementing the recommendation. Possible values include: 'Error', 'Warning', 'None'
	Risk Risk `json:"risk,omitempty"`
	// ShortDescription - A summary of the recommendation.
	ShortDescription *ShortDescription `json:"shortDescription,omitempty"`
}

// MarshalJSON is the custom marshaler for RecommendationProperties.
func (rp RecommendationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rp.Category != "" {
		objectMap["category"] = rp.Category
	}
	if rp.Impact != "" {
		objectMap["impact"] = rp.Impact
	}
	if rp.ImpactedField != nil {
		objectMap["impactedField"] = rp.ImpactedField
	}
	if rp.ImpactedValue != nil {
		objectMap["impactedValue"] = rp.ImpactedValue
	}
	if rp.LastUpdated != nil {
		objectMap["lastUpdated"] = rp.LastUpdated
	}
	if rp.Metadata != nil {
		objectMap["metadata"] = rp.Metadata
	}
	if rp.RecommendationTypeID != nil {
		objectMap["recommendationTypeId"] = rp.RecommendationTypeID
	}
	if rp.Risk != "" {
		objectMap["risk"] = rp.Risk
	}
	if rp.ShortDescription != nil {
		objectMap["shortDescription"] = rp.ShortDescription
	}
	return json.Marshal(objectMap)
}

// Resource an Azure resource.
type Resource struct {
	// ID - READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
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

// ResourceRecommendationBase advisor Recommendation.
type ResourceRecommendationBase struct {
	autorest.Response `json:"-"`
	// ID - The fully qualified recommendation ID, for example /subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.ClassicCompute/virtualMachines/vm1/providers/Microsoft.Advisor/recommendations/recommendationGUID.
	ID *string `json:"id,omitempty"`
	// Name - The name of recommendation.
	Name *string `json:"name,omitempty"`
	// RecommendationProperties - The properties of the recommendation.
	*RecommendationProperties `json:"properties,omitempty"`
	// SuppressionIds - The list of snoozed and dismissed rules for the recommendation.
	SuppressionIds *[]uuid.UUID `json:"suppressionIds,omitempty"`
	// Type - The recommendation type: Microsoft.Advisor/recommendations.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceRecommendationBase.
func (rrb ResourceRecommendationBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rrb.ID != nil {
		objectMap["id"] = rrb.ID
	}
	if rrb.Name != nil {
		objectMap["name"] = rrb.Name
	}
	if rrb.RecommendationProperties != nil {
		objectMap["properties"] = rrb.RecommendationProperties
	}
	if rrb.SuppressionIds != nil {
		objectMap["suppressionIds"] = rrb.SuppressionIds
	}
	if rrb.Type != nil {
		objectMap["type"] = rrb.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ResourceRecommendationBase struct.
func (rrb *ResourceRecommendationBase) UnmarshalJSON(body []byte) error {
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
				rrb.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rrb.Name = &name
			}
		case "properties":
			if v != nil {
				var recommendationProperties RecommendationProperties
				err = json.Unmarshal(*v, &recommendationProperties)
				if err != nil {
					return err
				}
				rrb.RecommendationProperties = &recommendationProperties
			}
		case "suppressionIds":
			if v != nil {
				var suppressionIds []uuid.UUID
				err = json.Unmarshal(*v, &suppressionIds)
				if err != nil {
					return err
				}
				rrb.SuppressionIds = &suppressionIds
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rrb.Type = &typeVar
			}
		}
	}

	return nil
}

// ResourceRecommendationBaseListResult the list of Advisor recommendations.
type ResourceRecommendationBaseListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of recommendations.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - The list of recommendations.
	Value *[]ResourceRecommendationBase `json:"value,omitempty"`
}

// ResourceRecommendationBaseListResultIterator provides access to a complete listing of
// ResourceRecommendationBase values.
type ResourceRecommendationBaseListResultIterator struct {
	i    int
	page ResourceRecommendationBaseListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceRecommendationBaseListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceRecommendationBaseListResultIterator.NextWithContext")
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
func (iter *ResourceRecommendationBaseListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceRecommendationBaseListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceRecommendationBaseListResultIterator) Response() ResourceRecommendationBaseListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceRecommendationBaseListResultIterator) Value() ResourceRecommendationBase {
	if !iter.page.NotDone() {
		return ResourceRecommendationBase{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceRecommendationBaseListResultIterator type.
func NewResourceRecommendationBaseListResultIterator(page ResourceRecommendationBaseListResultPage) ResourceRecommendationBaseListResultIterator {
	return ResourceRecommendationBaseListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rrblr ResourceRecommendationBaseListResult) IsEmpty() bool {
	return rrblr.Value == nil || len(*rrblr.Value) == 0
}

// resourceRecommendationBaseListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rrblr ResourceRecommendationBaseListResult) resourceRecommendationBaseListResultPreparer(ctx context.Context) (*http.Request, error) {
	if rrblr.NextLink == nil || len(to.String(rrblr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rrblr.NextLink)))
}

// ResourceRecommendationBaseListResultPage contains a page of ResourceRecommendationBase values.
type ResourceRecommendationBaseListResultPage struct {
	fn    func(context.Context, ResourceRecommendationBaseListResult) (ResourceRecommendationBaseListResult, error)
	rrblr ResourceRecommendationBaseListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceRecommendationBaseListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceRecommendationBaseListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rrblr)
	if err != nil {
		return err
	}
	page.rrblr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceRecommendationBaseListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceRecommendationBaseListResultPage) NotDone() bool {
	return !page.rrblr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceRecommendationBaseListResultPage) Response() ResourceRecommendationBaseListResult {
	return page.rrblr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceRecommendationBaseListResultPage) Values() []ResourceRecommendationBase {
	if page.rrblr.IsEmpty() {
		return nil
	}
	return *page.rrblr.Value
}

// Creates a new instance of the ResourceRecommendationBaseListResultPage type.
func NewResourceRecommendationBaseListResultPage(getNextPage func(context.Context, ResourceRecommendationBaseListResult) (ResourceRecommendationBaseListResult, error)) ResourceRecommendationBaseListResultPage {
	return ResourceRecommendationBaseListResultPage{fn: getNextPage}
}

// ShortDescription a summary of the recommendation.
type ShortDescription struct {
	// Problem - The issue or opportunity identified by the recommendation.
	Problem *string `json:"problem,omitempty"`
	// Solution - The remediation action suggested by the recommendation.
	Solution *string `json:"solution,omitempty"`
}

// SuppressionContract the details of the snoozed or dismissed rule; for example, the duration, name, and
// GUID associated with the rule.
type SuppressionContract struct {
	autorest.Response `json:"-"`
	// SuppressionID - The GUID of the suppression.
	SuppressionID *string `json:"suppressionId,omitempty"`
	// TTL - The duration for which the suppression is valid.
	TTL *string `json:"ttl,omitempty"`
	// ID - READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for SuppressionContract.
func (sc SuppressionContract) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sc.SuppressionID != nil {
		objectMap["suppressionId"] = sc.SuppressionID
	}
	if sc.TTL != nil {
		objectMap["ttl"] = sc.TTL
	}
	if sc.Location != nil {
		objectMap["location"] = sc.Location
	}
	if sc.Tags != nil {
		objectMap["tags"] = sc.Tags
	}
	return json.Marshal(objectMap)
}
