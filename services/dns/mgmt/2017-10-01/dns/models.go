package dns

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
const fqdn = "github.com/tcz001/azure-sdk-for-go/services/dns/mgmt/2017-10-01/dns"

// RecordType enumerates the values for record type.
type RecordType string

const (
	// A ...
	A RecordType = "A"
	// AAAA ...
	AAAA RecordType = "AAAA"
	// CAA ...
	CAA RecordType = "CAA"
	// CNAME ...
	CNAME RecordType = "CNAME"
	// MX ...
	MX RecordType = "MX"
	// NS ...
	NS RecordType = "NS"
	// PTR ...
	PTR RecordType = "PTR"
	// SOA ...
	SOA RecordType = "SOA"
	// SRV ...
	SRV RecordType = "SRV"
	// TXT ...
	TXT RecordType = "TXT"
)

// PossibleRecordTypeValues returns an array of possible values for the RecordType const type.
func PossibleRecordTypeValues() []RecordType {
	return []RecordType{A, AAAA, CAA, CNAME, MX, NS, PTR, SOA, SRV, TXT}
}

// AaaaRecord an AAAA record.
type AaaaRecord struct {
	// Ipv6Address - The IPv6 address of this AAAA record.
	Ipv6Address *string `json:"ipv6Address,omitempty"`
}

// ARecord an A record.
type ARecord struct {
	// Ipv4Address - The IPv4 address of this A record.
	Ipv4Address *string `json:"ipv4Address,omitempty"`
}

// AzureEntityResource the resource model definition for a Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// CaaRecord a CAA record.
type CaaRecord struct {
	// Flags - The flags for this CAA record as an integer between 0 and 255.
	Flags *int32 `json:"flags,omitempty"`
	// Tag - The tag for this CAA record.
	Tag *string `json:"tag,omitempty"`
	// Value - The value for this CAA record.
	Value *string `json:"value,omitempty"`
}

// CloudError an error message
type CloudError struct {
	// Error - The error message body
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody the body of an error message
type CloudErrorBody struct {
	// Code - The error code
	Code *string `json:"code,omitempty"`
	// Message - A description of what caused the error
	Message *string `json:"message,omitempty"`
	// Target - The target resource of the error message
	Target *string `json:"target,omitempty"`
	// Details - Extra error information
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// CnameRecord a CNAME record.
type CnameRecord struct {
	// Cname - The canonical name for this CNAME record.
	Cname *string `json:"cname,omitempty"`
}

// MxRecord an MX record.
type MxRecord struct {
	// Preference - The preference value for this MX record.
	Preference *int32 `json:"preference,omitempty"`
	// Exchange - The domain name of the mail host for this MX record.
	Exchange *string `json:"exchange,omitempty"`
}

// NsRecord an NS record.
type NsRecord struct {
	// Nsdname - The name server name for this NS record.
	Nsdname *string `json:"nsdname,omitempty"`
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// PtrRecord a PTR record.
type PtrRecord struct {
	// Ptrdname - The PTR target domain name for this PTR record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

// RecordSet describes a DNS record set (a collection of DNS records with the same name and type).
type RecordSet struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The ID of the record set.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the record set.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the record set.
	Type *string `json:"type,omitempty"`
	// Etag - The etag of the record set.
	Etag *string `json:"etag,omitempty"`
	// RecordSetProperties - The properties of the record set.
	*RecordSetProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for RecordSet.
func (rs RecordSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rs.Etag != nil {
		objectMap["etag"] = rs.Etag
	}
	if rs.RecordSetProperties != nil {
		objectMap["properties"] = rs.RecordSetProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for RecordSet struct.
func (rs *RecordSet) UnmarshalJSON(body []byte) error {
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
				rs.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rs.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rs.Type = &typeVar
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				rs.Etag = &etag
			}
		case "properties":
			if v != nil {
				var recordSetProperties RecordSetProperties
				err = json.Unmarshal(*v, &recordSetProperties)
				if err != nil {
					return err
				}
				rs.RecordSetProperties = &recordSetProperties
			}
		}
	}

	return nil
}

// RecordSetListResult the response to a record set List operation.
type RecordSetListResult struct {
	autorest.Response `json:"-"`
	// Value - Information about the record sets in the response.
	Value *[]RecordSet `json:"value,omitempty"`
	// NextLink - READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// RecordSetListResultIterator provides access to a complete listing of RecordSet values.
type RecordSetListResultIterator struct {
	i    int
	page RecordSetListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RecordSetListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetListResultIterator.NextWithContext")
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
func (iter *RecordSetListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RecordSetListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RecordSetListResultIterator) Response() RecordSetListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RecordSetListResultIterator) Value() RecordSet {
	if !iter.page.NotDone() {
		return RecordSet{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the RecordSetListResultIterator type.
func NewRecordSetListResultIterator(page RecordSetListResultPage) RecordSetListResultIterator {
	return RecordSetListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rslr RecordSetListResult) IsEmpty() bool {
	return rslr.Value == nil || len(*rslr.Value) == 0
}

// recordSetListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rslr RecordSetListResult) recordSetListResultPreparer(ctx context.Context) (*http.Request, error) {
	if rslr.NextLink == nil || len(to.String(rslr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rslr.NextLink)))
}

// RecordSetListResultPage contains a page of RecordSet values.
type RecordSetListResultPage struct {
	fn   func(context.Context, RecordSetListResult) (RecordSetListResult, error)
	rslr RecordSetListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RecordSetListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecordSetListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rslr)
	if err != nil {
		return err
	}
	page.rslr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *RecordSetListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RecordSetListResultPage) NotDone() bool {
	return !page.rslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RecordSetListResultPage) Response() RecordSetListResult {
	return page.rslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RecordSetListResultPage) Values() []RecordSet {
	if page.rslr.IsEmpty() {
		return nil
	}
	return *page.rslr.Value
}

// Creates a new instance of the RecordSetListResultPage type.
func NewRecordSetListResultPage(getNextPage func(context.Context, RecordSetListResult) (RecordSetListResult, error)) RecordSetListResultPage {
	return RecordSetListResultPage{fn: getNextPage}
}

// RecordSetProperties represents the properties of the records in the record set.
type RecordSetProperties struct {
	// Metadata - The metadata attached to the record set.
	Metadata map[string]*string `json:"metadata"`
	// TTL - The TTL (time-to-live) of the records in the record set.
	TTL *int64 `json:"TTL,omitempty"`
	// Fqdn - READ-ONLY; Fully qualified domain name of the record set.
	Fqdn *string `json:"fqdn,omitempty"`
	// ARecords - The list of A records in the record set.
	ARecords *[]ARecord `json:"ARecords,omitempty"`
	// AaaaRecords - The list of AAAA records in the record set.
	AaaaRecords *[]AaaaRecord `json:"AAAARecords,omitempty"`
	// MxRecords - The list of MX records in the record set.
	MxRecords *[]MxRecord `json:"MXRecords,omitempty"`
	// NsRecords - The list of NS records in the record set.
	NsRecords *[]NsRecord `json:"NSRecords,omitempty"`
	// PtrRecords - The list of PTR records in the record set.
	PtrRecords *[]PtrRecord `json:"PTRRecords,omitempty"`
	// SrvRecords - The list of SRV records in the record set.
	SrvRecords *[]SrvRecord `json:"SRVRecords,omitempty"`
	// TxtRecords - The list of TXT records in the record set.
	TxtRecords *[]TxtRecord `json:"TXTRecords,omitempty"`
	// CnameRecord - The CNAME record in the  record set.
	CnameRecord *CnameRecord `json:"CNAMERecord,omitempty"`
	// SoaRecord - The SOA record in the record set.
	SoaRecord *SoaRecord `json:"SOARecord,omitempty"`
	// CaaRecords - The list of CAA records in the record set.
	CaaRecords *[]CaaRecord `json:"caaRecords,omitempty"`
}

// MarshalJSON is the custom marshaler for RecordSetProperties.
func (rsp RecordSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rsp.Metadata != nil {
		objectMap["metadata"] = rsp.Metadata
	}
	if rsp.TTL != nil {
		objectMap["TTL"] = rsp.TTL
	}
	if rsp.ARecords != nil {
		objectMap["ARecords"] = rsp.ARecords
	}
	if rsp.AaaaRecords != nil {
		objectMap["AAAARecords"] = rsp.AaaaRecords
	}
	if rsp.MxRecords != nil {
		objectMap["MXRecords"] = rsp.MxRecords
	}
	if rsp.NsRecords != nil {
		objectMap["NSRecords"] = rsp.NsRecords
	}
	if rsp.PtrRecords != nil {
		objectMap["PTRRecords"] = rsp.PtrRecords
	}
	if rsp.SrvRecords != nil {
		objectMap["SRVRecords"] = rsp.SrvRecords
	}
	if rsp.TxtRecords != nil {
		objectMap["TXTRecords"] = rsp.TxtRecords
	}
	if rsp.CnameRecord != nil {
		objectMap["CNAMERecord"] = rsp.CnameRecord
	}
	if rsp.SoaRecord != nil {
		objectMap["SOARecord"] = rsp.SoaRecord
	}
	if rsp.CaaRecords != nil {
		objectMap["caaRecords"] = rsp.CaaRecords
	}
	return json.Marshal(objectMap)
}

// RecordSetUpdateParameters parameters supplied to update a record set.
type RecordSetUpdateParameters struct {
	// RecordSet - Specifies information about the record set being updated.
	RecordSet *RecordSet `json:"RecordSet,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// SoaRecord an SOA record.
type SoaRecord struct {
	// Host - The domain name of the authoritative name server for this SOA record.
	Host *string `json:"host,omitempty"`
	// Email - The email contact for this SOA record.
	Email *string `json:"email,omitempty"`
	// SerialNumber - The serial number for this SOA record.
	SerialNumber *int64 `json:"serialNumber,omitempty"`
	// RefreshTime - The refresh value for this SOA record.
	RefreshTime *int64 `json:"refreshTime,omitempty"`
	// RetryTime - The retry time for this SOA record.
	RetryTime *int64 `json:"retryTime,omitempty"`
	// ExpireTime - The expire time for this SOA record.
	ExpireTime *int64 `json:"expireTime,omitempty"`
	// MinimumTTL - The minimum value for this SOA record. By convention this is used to determine the negative caching duration.
	MinimumTTL *int64 `json:"minimumTTL,omitempty"`
}

// SrvRecord an SRV record.
type SrvRecord struct {
	// Priority - The priority value for this SRV record.
	Priority *int32 `json:"priority,omitempty"`
	// Weight - The weight value for this SRV record.
	Weight *int32 `json:"weight,omitempty"`
	// Port - The port value for this SRV record.
	Port *int32 `json:"port,omitempty"`
	// Target - The target domain name for this SRV record.
	Target *string `json:"target,omitempty"`
}

// SubResource a reference to a another resource
type SubResource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}

// TxtRecord a TXT record.
type TxtRecord struct {
	// Value - The text value of this TXT record.
	Value *[]string `json:"value,omitempty"`
}

// Zone describes a DNS zone.
type Zone struct {
	autorest.Response `json:"-"`
	// Etag - The etag of the zone.
	Etag *string `json:"etag,omitempty"`
	// ZoneProperties - The properties of the zone.
	*ZoneProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Zone.
func (z Zone) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if z.Etag != nil {
		objectMap["etag"] = z.Etag
	}
	if z.ZoneProperties != nil {
		objectMap["properties"] = z.ZoneProperties
	}
	if z.Tags != nil {
		objectMap["tags"] = z.Tags
	}
	if z.Location != nil {
		objectMap["location"] = z.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Zone struct.
func (z *Zone) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				z.Etag = &etag
			}
		case "properties":
			if v != nil {
				var zoneProperties ZoneProperties
				err = json.Unmarshal(*v, &zoneProperties)
				if err != nil {
					return err
				}
				z.ZoneProperties = &zoneProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				z.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				z.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				z.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				z.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				z.Type = &typeVar
			}
		}
	}

	return nil
}

// ZoneListResult the response to a Zone List or ListAll operation.
type ZoneListResult struct {
	autorest.Response `json:"-"`
	// Value - Information about the DNS zones.
	Value *[]Zone `json:"value,omitempty"`
	// NextLink - READ-ONLY; The continuation token for the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ZoneListResultIterator provides access to a complete listing of Zone values.
type ZoneListResultIterator struct {
	i    int
	page ZoneListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ZoneListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ZoneListResultIterator.NextWithContext")
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
func (iter *ZoneListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ZoneListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ZoneListResultIterator) Response() ZoneListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ZoneListResultIterator) Value() Zone {
	if !iter.page.NotDone() {
		return Zone{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ZoneListResultIterator type.
func NewZoneListResultIterator(page ZoneListResultPage) ZoneListResultIterator {
	return ZoneListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (zlr ZoneListResult) IsEmpty() bool {
	return zlr.Value == nil || len(*zlr.Value) == 0
}

// zoneListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (zlr ZoneListResult) zoneListResultPreparer(ctx context.Context) (*http.Request, error) {
	if zlr.NextLink == nil || len(to.String(zlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(zlr.NextLink)))
}

// ZoneListResultPage contains a page of Zone values.
type ZoneListResultPage struct {
	fn  func(context.Context, ZoneListResult) (ZoneListResult, error)
	zlr ZoneListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ZoneListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ZoneListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.zlr)
	if err != nil {
		return err
	}
	page.zlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ZoneListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ZoneListResultPage) NotDone() bool {
	return !page.zlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ZoneListResultPage) Response() ZoneListResult {
	return page.zlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ZoneListResultPage) Values() []Zone {
	if page.zlr.IsEmpty() {
		return nil
	}
	return *page.zlr.Value
}

// Creates a new instance of the ZoneListResultPage type.
func NewZoneListResultPage(getNextPage func(context.Context, ZoneListResult) (ZoneListResult, error)) ZoneListResultPage {
	return ZoneListResultPage{fn: getNextPage}
}

// ZoneProperties represents the properties of the zone.
type ZoneProperties struct {
	// MaxNumberOfRecordSets - READ-ONLY; The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int64 `json:"maxNumberOfRecordSets,omitempty"`
	// NumberOfRecordSets - READ-ONLY; The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets *int64 `json:"numberOfRecordSets,omitempty"`
	// NameServers - READ-ONLY; The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers *[]string `json:"nameServers,omitempty"`
}

// ZonesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type ZonesDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ZonesDeleteFuture) Result(client ZonesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.ZonesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("dns.ZonesDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ZoneUpdate describes a request to update a DNS zone.
type ZoneUpdate struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ZoneUpdate.
func (zu ZoneUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if zu.Tags != nil {
		objectMap["tags"] = zu.Tags
	}
	return json.Marshal(objectMap)
}
