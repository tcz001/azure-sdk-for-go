package datacatalog

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
)

// The package's fully qualified name.
const fqdn = "github.com/tcz001/azure-sdk-for-go/services/datacatalog/mgmt/2016-03-30/datacatalog"

// SkuType enumerates the values for sku type.
type SkuType string

const (
	// Free ...
	Free SkuType = "Free"
	// Standard ...
	Standard SkuType = "Standard"
)

// PossibleSkuTypeValues returns an array of possible values for the SkuType const type.
func PossibleSkuTypeValues() []SkuType {
	return []SkuType{Free, Standard}
}

// ADCCatalog azure Data Catalog.
type ADCCatalog struct {
	autorest.Response `json:"-"`
	// ADCCatalogProperties - Azure Data Catalog properties.
	*ADCCatalogProperties `json:"properties,omitempty"`
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
	// Etag - Resource etag
	Etag *string `json:"etag,omitempty"`
}

// MarshalJSON is the custom marshaler for ADCCatalog.
func (ac ADCCatalog) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ac.ADCCatalogProperties != nil {
		objectMap["properties"] = ac.ADCCatalogProperties
	}
	if ac.Location != nil {
		objectMap["location"] = ac.Location
	}
	if ac.Tags != nil {
		objectMap["tags"] = ac.Tags
	}
	if ac.Etag != nil {
		objectMap["etag"] = ac.Etag
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ADCCatalog struct.
func (ac *ADCCatalog) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var aDCCatalogProperties ADCCatalogProperties
				err = json.Unmarshal(*v, &aDCCatalogProperties)
				if err != nil {
					return err
				}
				ac.ADCCatalogProperties = &aDCCatalogProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ac.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ac.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ac.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ac.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ac.Tags = tags
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				ac.Etag = &etag
			}
		}
	}

	return nil
}

// ADCCatalogProperties properties of the data catalog.
type ADCCatalogProperties struct {
	// Sku - Azure data catalog SKU. Possible values include: 'Free', 'Standard'
	Sku SkuType `json:"sku,omitempty"`
	// Units - Azure data catalog units.
	Units *int32 `json:"units,omitempty"`
	// Admins - Azure data catalog admin list.
	Admins *[]Principals `json:"admins,omitempty"`
	// Users - Azure data catalog user list.
	Users *[]Principals `json:"users,omitempty"`
	// SuccessfullyProvisioned - Azure data catalog provision status.
	SuccessfullyProvisioned *bool `json:"successfullyProvisioned,omitempty"`
	// EnableAutomaticUnitAdjustment - Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment *bool `json:"enableAutomaticUnitAdjustment,omitempty"`
}

// ADCCatalogsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ADCCatalogsDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ADCCatalogsDeleteFuture) Result(client ADCCatalogsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("datacatalog.ADCCatalogsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ADCCatalogsListResult the response from the List Azure Data Catalog operation.
type ADCCatalogsListResult struct {
	autorest.Response `json:"-"`
	// Value - the list of Azure Data Catalogs.
	Value *[]ADCCatalog `json:"value,omitempty"`
}

// OperationDisplayInfo the operation supported by Azure Data Catalog Service.
type OperationDisplayInfo struct {
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
	// Operation - The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Provider - Service provider: Azure Data Catalog Service.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity the operation supported by Azure Data Catalog Service.
type OperationEntity struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The operation supported by Azure Data Catalog Service.
	Display *OperationDisplayInfo `json:"display,omitempty"`
}

// OperationEntityListResult the list of Azure data catalog service operation response.
type OperationEntityListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of operations.
	Value *[]OperationEntity `json:"value,omitempty"`
}

// Principals user principals.
type Principals struct {
	// Upn - UPN of the user.
	Upn *string `json:"upn,omitempty"`
	// ObjectID - Object Id for the user
	ObjectID *string `json:"objectId,omitempty"`
}

// Resource the Resource model definition.
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
	// Etag - Resource etag
	Etag *string `json:"etag,omitempty"`
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
	if r.Etag != nil {
		objectMap["etag"] = r.Etag
	}
	return json.Marshal(objectMap)
}
