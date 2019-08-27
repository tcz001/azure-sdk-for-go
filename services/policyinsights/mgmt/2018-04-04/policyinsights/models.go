package policyinsights

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "github.com/tcz001/azure-sdk-for-go/services/policyinsights/mgmt/2018-04-04/policyinsights"

// PolicyStatesResource enumerates the values for policy states resource.
type PolicyStatesResource string

const (
	// Default ...
	Default PolicyStatesResource = "default"
	// Latest ...
	Latest PolicyStatesResource = "latest"
)

// PossiblePolicyStatesResourceValues returns an array of possible values for the PolicyStatesResource const type.
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return []PolicyStatesResource{Default, Latest}
}

// Operation operation definition.
type Operation struct {
	// Name - Operation name.
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Resource provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource name on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation name.
	Operation *string `json:"operation,omitempty"`
	// Description - Operation description.
	Description *string `json:"description,omitempty"`
}

// OperationsListResults list of available operations.
type OperationsListResults struct {
	autorest.Response `json:"-"`
	// OdataCount - OData entity count; represents the number of operations returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - List of available operations.
	Value *[]Operation `json:"value,omitempty"`
}

// PolicyAssignmentSummary policy assignment summary.
type PolicyAssignmentSummary struct {
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string `json:"policySetDefinitionId,omitempty"`
	// Results - Non-compliance summary for the policy assignment.
	Results *SummaryResults `json:"results,omitempty"`
	// PolicyDefinitions - Policy definitions summary.
	PolicyDefinitions *[]PolicyDefinitionSummary `json:"policyDefinitions,omitempty"`
}

// PolicyDefinitionSummary policy definition summary.
type PolicyDefinitionSummary struct {
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// PolicyDefinitionReferenceID - Policy definition reference ID.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty"`
	// Effect - Policy effect, i.e. policy definition action.
	Effect *string `json:"effect,omitempty"`
	// Results - Non-compliance summary for the policy definition.
	Results *SummaryResults `json:"results,omitempty"`
}

// PolicyEvent policy event record.
type PolicyEvent struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// OdataID - OData entity ID; always set to null since policy event records do not have an entity ID.
	OdataID *string `json:"@odata.id,omitempty"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// Timestamp - Timestamp for the policy event record.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *string `json:"effectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *string `json:"resourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *string `json:"policyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *string `json:"policyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *string `json:"policyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *string `json:"policyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *string `json:"policyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *string `json:"policyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *string `json:"policyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string `json:"policySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string `json:"policySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string `json:"policySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string `json:"policySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string `json:"policySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma separated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *string `json:"managementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty"`
	// TenantID - Tenant ID for the policy event record.
	TenantID *string `json:"tenantId,omitempty"`
	// PrincipalOid - Principal object ID for the user who initiated the resource operation that triggered the policy event.
	PrincipalOid *string `json:"principalOid,omitempty"`
}

// MarshalJSON is the custom marshaler for PolicyEvent.
func (peVar PolicyEvent) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if peVar.OdataID != nil {
		objectMap["@odata.id"] = peVar.OdataID
	}
	if peVar.OdataContext != nil {
		objectMap["@odata.context"] = peVar.OdataContext
	}
	if peVar.Timestamp != nil {
		objectMap["timestamp"] = peVar.Timestamp
	}
	if peVar.ResourceID != nil {
		objectMap["resourceId"] = peVar.ResourceID
	}
	if peVar.PolicyAssignmentID != nil {
		objectMap["policyAssignmentId"] = peVar.PolicyAssignmentID
	}
	if peVar.PolicyDefinitionID != nil {
		objectMap["policyDefinitionId"] = peVar.PolicyDefinitionID
	}
	if peVar.EffectiveParameters != nil {
		objectMap["effectiveParameters"] = peVar.EffectiveParameters
	}
	if peVar.IsCompliant != nil {
		objectMap["isCompliant"] = peVar.IsCompliant
	}
	if peVar.SubscriptionID != nil {
		objectMap["subscriptionId"] = peVar.SubscriptionID
	}
	if peVar.ResourceType != nil {
		objectMap["resourceType"] = peVar.ResourceType
	}
	if peVar.ResourceLocation != nil {
		objectMap["resourceLocation"] = peVar.ResourceLocation
	}
	if peVar.ResourceGroup != nil {
		objectMap["resourceGroup"] = peVar.ResourceGroup
	}
	if peVar.ResourceTags != nil {
		objectMap["resourceTags"] = peVar.ResourceTags
	}
	if peVar.PolicyAssignmentName != nil {
		objectMap["policyAssignmentName"] = peVar.PolicyAssignmentName
	}
	if peVar.PolicyAssignmentOwner != nil {
		objectMap["policyAssignmentOwner"] = peVar.PolicyAssignmentOwner
	}
	if peVar.PolicyAssignmentParameters != nil {
		objectMap["policyAssignmentParameters"] = peVar.PolicyAssignmentParameters
	}
	if peVar.PolicyAssignmentScope != nil {
		objectMap["policyAssignmentScope"] = peVar.PolicyAssignmentScope
	}
	if peVar.PolicyDefinitionName != nil {
		objectMap["policyDefinitionName"] = peVar.PolicyDefinitionName
	}
	if peVar.PolicyDefinitionAction != nil {
		objectMap["policyDefinitionAction"] = peVar.PolicyDefinitionAction
	}
	if peVar.PolicyDefinitionCategory != nil {
		objectMap["policyDefinitionCategory"] = peVar.PolicyDefinitionCategory
	}
	if peVar.PolicySetDefinitionID != nil {
		objectMap["policySetDefinitionId"] = peVar.PolicySetDefinitionID
	}
	if peVar.PolicySetDefinitionName != nil {
		objectMap["policySetDefinitionName"] = peVar.PolicySetDefinitionName
	}
	if peVar.PolicySetDefinitionOwner != nil {
		objectMap["policySetDefinitionOwner"] = peVar.PolicySetDefinitionOwner
	}
	if peVar.PolicySetDefinitionCategory != nil {
		objectMap["policySetDefinitionCategory"] = peVar.PolicySetDefinitionCategory
	}
	if peVar.PolicySetDefinitionParameters != nil {
		objectMap["policySetDefinitionParameters"] = peVar.PolicySetDefinitionParameters
	}
	if peVar.ManagementGroupIds != nil {
		objectMap["managementGroupIds"] = peVar.ManagementGroupIds
	}
	if peVar.PolicyDefinitionReferenceID != nil {
		objectMap["policyDefinitionReferenceId"] = peVar.PolicyDefinitionReferenceID
	}
	if peVar.TenantID != nil {
		objectMap["tenantId"] = peVar.TenantID
	}
	if peVar.PrincipalOid != nil {
		objectMap["principalOid"] = peVar.PrincipalOid
	}
	for k, v := range peVar.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PolicyEvent struct.
func (peVar *PolicyEvent) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if peVar.AdditionalProperties == nil {
					peVar.AdditionalProperties = make(map[string]interface{})
				}
				peVar.AdditionalProperties[k] = additionalProperties
			}
		case "@odata.id":
			if v != nil {
				var odataID string
				err = json.Unmarshal(*v, &odataID)
				if err != nil {
					return err
				}
				peVar.OdataID = &odataID
			}
		case "@odata.context":
			if v != nil {
				var odataContext string
				err = json.Unmarshal(*v, &odataContext)
				if err != nil {
					return err
				}
				peVar.OdataContext = &odataContext
			}
		case "timestamp":
			if v != nil {
				var timestamp date.Time
				err = json.Unmarshal(*v, &timestamp)
				if err != nil {
					return err
				}
				peVar.Timestamp = &timestamp
			}
		case "resourceId":
			if v != nil {
				var resourceID string
				err = json.Unmarshal(*v, &resourceID)
				if err != nil {
					return err
				}
				peVar.ResourceID = &resourceID
			}
		case "policyAssignmentId":
			if v != nil {
				var policyAssignmentID string
				err = json.Unmarshal(*v, &policyAssignmentID)
				if err != nil {
					return err
				}
				peVar.PolicyAssignmentID = &policyAssignmentID
			}
		case "policyDefinitionId":
			if v != nil {
				var policyDefinitionID string
				err = json.Unmarshal(*v, &policyDefinitionID)
				if err != nil {
					return err
				}
				peVar.PolicyDefinitionID = &policyDefinitionID
			}
		case "effectiveParameters":
			if v != nil {
				var effectiveParameters string
				err = json.Unmarshal(*v, &effectiveParameters)
				if err != nil {
					return err
				}
				peVar.EffectiveParameters = &effectiveParameters
			}
		case "isCompliant":
			if v != nil {
				var isCompliant bool
				err = json.Unmarshal(*v, &isCompliant)
				if err != nil {
					return err
				}
				peVar.IsCompliant = &isCompliant
			}
		case "subscriptionId":
			if v != nil {
				var subscriptionID string
				err = json.Unmarshal(*v, &subscriptionID)
				if err != nil {
					return err
				}
				peVar.SubscriptionID = &subscriptionID
			}
		case "resourceType":
			if v != nil {
				var resourceType string
				err = json.Unmarshal(*v, &resourceType)
				if err != nil {
					return err
				}
				peVar.ResourceType = &resourceType
			}
		case "resourceLocation":
			if v != nil {
				var resourceLocation string
				err = json.Unmarshal(*v, &resourceLocation)
				if err != nil {
					return err
				}
				peVar.ResourceLocation = &resourceLocation
			}
		case "resourceGroup":
			if v != nil {
				var resourceGroup string
				err = json.Unmarshal(*v, &resourceGroup)
				if err != nil {
					return err
				}
				peVar.ResourceGroup = &resourceGroup
			}
		case "resourceTags":
			if v != nil {
				var resourceTags string
				err = json.Unmarshal(*v, &resourceTags)
				if err != nil {
					return err
				}
				peVar.ResourceTags = &resourceTags
			}
		case "policyAssignmentName":
			if v != nil {
				var policyAssignmentName string
				err = json.Unmarshal(*v, &policyAssignmentName)
				if err != nil {
					return err
				}
				peVar.PolicyAssignmentName = &policyAssignmentName
			}
		case "policyAssignmentOwner":
			if v != nil {
				var policyAssignmentOwner string
				err = json.Unmarshal(*v, &policyAssignmentOwner)
				if err != nil {
					return err
				}
				peVar.PolicyAssignmentOwner = &policyAssignmentOwner
			}
		case "policyAssignmentParameters":
			if v != nil {
				var policyAssignmentParameters string
				err = json.Unmarshal(*v, &policyAssignmentParameters)
				if err != nil {
					return err
				}
				peVar.PolicyAssignmentParameters = &policyAssignmentParameters
			}
		case "policyAssignmentScope":
			if v != nil {
				var policyAssignmentScope string
				err = json.Unmarshal(*v, &policyAssignmentScope)
				if err != nil {
					return err
				}
				peVar.PolicyAssignmentScope = &policyAssignmentScope
			}
		case "policyDefinitionName":
			if v != nil {
				var policyDefinitionName string
				err = json.Unmarshal(*v, &policyDefinitionName)
				if err != nil {
					return err
				}
				peVar.PolicyDefinitionName = &policyDefinitionName
			}
		case "policyDefinitionAction":
			if v != nil {
				var policyDefinitionAction string
				err = json.Unmarshal(*v, &policyDefinitionAction)
				if err != nil {
					return err
				}
				peVar.PolicyDefinitionAction = &policyDefinitionAction
			}
		case "policyDefinitionCategory":
			if v != nil {
				var policyDefinitionCategory string
				err = json.Unmarshal(*v, &policyDefinitionCategory)
				if err != nil {
					return err
				}
				peVar.PolicyDefinitionCategory = &policyDefinitionCategory
			}
		case "policySetDefinitionId":
			if v != nil {
				var policySetDefinitionID string
				err = json.Unmarshal(*v, &policySetDefinitionID)
				if err != nil {
					return err
				}
				peVar.PolicySetDefinitionID = &policySetDefinitionID
			}
		case "policySetDefinitionName":
			if v != nil {
				var policySetDefinitionName string
				err = json.Unmarshal(*v, &policySetDefinitionName)
				if err != nil {
					return err
				}
				peVar.PolicySetDefinitionName = &policySetDefinitionName
			}
		case "policySetDefinitionOwner":
			if v != nil {
				var policySetDefinitionOwner string
				err = json.Unmarshal(*v, &policySetDefinitionOwner)
				if err != nil {
					return err
				}
				peVar.PolicySetDefinitionOwner = &policySetDefinitionOwner
			}
		case "policySetDefinitionCategory":
			if v != nil {
				var policySetDefinitionCategory string
				err = json.Unmarshal(*v, &policySetDefinitionCategory)
				if err != nil {
					return err
				}
				peVar.PolicySetDefinitionCategory = &policySetDefinitionCategory
			}
		case "policySetDefinitionParameters":
			if v != nil {
				var policySetDefinitionParameters string
				err = json.Unmarshal(*v, &policySetDefinitionParameters)
				if err != nil {
					return err
				}
				peVar.PolicySetDefinitionParameters = &policySetDefinitionParameters
			}
		case "managementGroupIds":
			if v != nil {
				var managementGroupIds string
				err = json.Unmarshal(*v, &managementGroupIds)
				if err != nil {
					return err
				}
				peVar.ManagementGroupIds = &managementGroupIds
			}
		case "policyDefinitionReferenceId":
			if v != nil {
				var policyDefinitionReferenceID string
				err = json.Unmarshal(*v, &policyDefinitionReferenceID)
				if err != nil {
					return err
				}
				peVar.PolicyDefinitionReferenceID = &policyDefinitionReferenceID
			}
		case "tenantId":
			if v != nil {
				var tenantID string
				err = json.Unmarshal(*v, &tenantID)
				if err != nil {
					return err
				}
				peVar.TenantID = &tenantID
			}
		case "principalOid":
			if v != nil {
				var principalOid string
				err = json.Unmarshal(*v, &principalOid)
				if err != nil {
					return err
				}
				peVar.PrincipalOid = &principalOid
			}
		}
	}

	return nil
}

// PolicyEventsQueryResults query results.
type PolicyEventsQueryResults struct {
	autorest.Response `json:"-"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// OdataCount - OData entity count; represents the number of policy event records returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - Query results.
	Value *[]PolicyEvent `json:"value,omitempty"`
}

// PolicyState policy state record.
type PolicyState struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// OdataID - OData entity ID; always set to null since policy state records do not have an entity ID.
	OdataID *string `json:"@odata.id,omitempty"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// Timestamp - Timestamp for the policy state record.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *string `json:"effectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *string `json:"resourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *string `json:"policyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *string `json:"policyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *string `json:"policyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *string `json:"policyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *string `json:"policyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *string `json:"policyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *string `json:"policyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string `json:"policySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string `json:"policySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string `json:"policySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string `json:"policySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string `json:"policySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma separated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *string `json:"managementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty"`
}

// MarshalJSON is the custom marshaler for PolicyState.
func (ps PolicyState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ps.OdataID != nil {
		objectMap["@odata.id"] = ps.OdataID
	}
	if ps.OdataContext != nil {
		objectMap["@odata.context"] = ps.OdataContext
	}
	if ps.Timestamp != nil {
		objectMap["timestamp"] = ps.Timestamp
	}
	if ps.ResourceID != nil {
		objectMap["resourceId"] = ps.ResourceID
	}
	if ps.PolicyAssignmentID != nil {
		objectMap["policyAssignmentId"] = ps.PolicyAssignmentID
	}
	if ps.PolicyDefinitionID != nil {
		objectMap["policyDefinitionId"] = ps.PolicyDefinitionID
	}
	if ps.EffectiveParameters != nil {
		objectMap["effectiveParameters"] = ps.EffectiveParameters
	}
	if ps.IsCompliant != nil {
		objectMap["isCompliant"] = ps.IsCompliant
	}
	if ps.SubscriptionID != nil {
		objectMap["subscriptionId"] = ps.SubscriptionID
	}
	if ps.ResourceType != nil {
		objectMap["resourceType"] = ps.ResourceType
	}
	if ps.ResourceLocation != nil {
		objectMap["resourceLocation"] = ps.ResourceLocation
	}
	if ps.ResourceGroup != nil {
		objectMap["resourceGroup"] = ps.ResourceGroup
	}
	if ps.ResourceTags != nil {
		objectMap["resourceTags"] = ps.ResourceTags
	}
	if ps.PolicyAssignmentName != nil {
		objectMap["policyAssignmentName"] = ps.PolicyAssignmentName
	}
	if ps.PolicyAssignmentOwner != nil {
		objectMap["policyAssignmentOwner"] = ps.PolicyAssignmentOwner
	}
	if ps.PolicyAssignmentParameters != nil {
		objectMap["policyAssignmentParameters"] = ps.PolicyAssignmentParameters
	}
	if ps.PolicyAssignmentScope != nil {
		objectMap["policyAssignmentScope"] = ps.PolicyAssignmentScope
	}
	if ps.PolicyDefinitionName != nil {
		objectMap["policyDefinitionName"] = ps.PolicyDefinitionName
	}
	if ps.PolicyDefinitionAction != nil {
		objectMap["policyDefinitionAction"] = ps.PolicyDefinitionAction
	}
	if ps.PolicyDefinitionCategory != nil {
		objectMap["policyDefinitionCategory"] = ps.PolicyDefinitionCategory
	}
	if ps.PolicySetDefinitionID != nil {
		objectMap["policySetDefinitionId"] = ps.PolicySetDefinitionID
	}
	if ps.PolicySetDefinitionName != nil {
		objectMap["policySetDefinitionName"] = ps.PolicySetDefinitionName
	}
	if ps.PolicySetDefinitionOwner != nil {
		objectMap["policySetDefinitionOwner"] = ps.PolicySetDefinitionOwner
	}
	if ps.PolicySetDefinitionCategory != nil {
		objectMap["policySetDefinitionCategory"] = ps.PolicySetDefinitionCategory
	}
	if ps.PolicySetDefinitionParameters != nil {
		objectMap["policySetDefinitionParameters"] = ps.PolicySetDefinitionParameters
	}
	if ps.ManagementGroupIds != nil {
		objectMap["managementGroupIds"] = ps.ManagementGroupIds
	}
	if ps.PolicyDefinitionReferenceID != nil {
		objectMap["policyDefinitionReferenceId"] = ps.PolicyDefinitionReferenceID
	}
	for k, v := range ps.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PolicyState struct.
func (ps *PolicyState) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if ps.AdditionalProperties == nil {
					ps.AdditionalProperties = make(map[string]interface{})
				}
				ps.AdditionalProperties[k] = additionalProperties
			}
		case "@odata.id":
			if v != nil {
				var odataID string
				err = json.Unmarshal(*v, &odataID)
				if err != nil {
					return err
				}
				ps.OdataID = &odataID
			}
		case "@odata.context":
			if v != nil {
				var odataContext string
				err = json.Unmarshal(*v, &odataContext)
				if err != nil {
					return err
				}
				ps.OdataContext = &odataContext
			}
		case "timestamp":
			if v != nil {
				var timestamp date.Time
				err = json.Unmarshal(*v, &timestamp)
				if err != nil {
					return err
				}
				ps.Timestamp = &timestamp
			}
		case "resourceId":
			if v != nil {
				var resourceID string
				err = json.Unmarshal(*v, &resourceID)
				if err != nil {
					return err
				}
				ps.ResourceID = &resourceID
			}
		case "policyAssignmentId":
			if v != nil {
				var policyAssignmentID string
				err = json.Unmarshal(*v, &policyAssignmentID)
				if err != nil {
					return err
				}
				ps.PolicyAssignmentID = &policyAssignmentID
			}
		case "policyDefinitionId":
			if v != nil {
				var policyDefinitionID string
				err = json.Unmarshal(*v, &policyDefinitionID)
				if err != nil {
					return err
				}
				ps.PolicyDefinitionID = &policyDefinitionID
			}
		case "effectiveParameters":
			if v != nil {
				var effectiveParameters string
				err = json.Unmarshal(*v, &effectiveParameters)
				if err != nil {
					return err
				}
				ps.EffectiveParameters = &effectiveParameters
			}
		case "isCompliant":
			if v != nil {
				var isCompliant bool
				err = json.Unmarshal(*v, &isCompliant)
				if err != nil {
					return err
				}
				ps.IsCompliant = &isCompliant
			}
		case "subscriptionId":
			if v != nil {
				var subscriptionID string
				err = json.Unmarshal(*v, &subscriptionID)
				if err != nil {
					return err
				}
				ps.SubscriptionID = &subscriptionID
			}
		case "resourceType":
			if v != nil {
				var resourceType string
				err = json.Unmarshal(*v, &resourceType)
				if err != nil {
					return err
				}
				ps.ResourceType = &resourceType
			}
		case "resourceLocation":
			if v != nil {
				var resourceLocation string
				err = json.Unmarshal(*v, &resourceLocation)
				if err != nil {
					return err
				}
				ps.ResourceLocation = &resourceLocation
			}
		case "resourceGroup":
			if v != nil {
				var resourceGroup string
				err = json.Unmarshal(*v, &resourceGroup)
				if err != nil {
					return err
				}
				ps.ResourceGroup = &resourceGroup
			}
		case "resourceTags":
			if v != nil {
				var resourceTags string
				err = json.Unmarshal(*v, &resourceTags)
				if err != nil {
					return err
				}
				ps.ResourceTags = &resourceTags
			}
		case "policyAssignmentName":
			if v != nil {
				var policyAssignmentName string
				err = json.Unmarshal(*v, &policyAssignmentName)
				if err != nil {
					return err
				}
				ps.PolicyAssignmentName = &policyAssignmentName
			}
		case "policyAssignmentOwner":
			if v != nil {
				var policyAssignmentOwner string
				err = json.Unmarshal(*v, &policyAssignmentOwner)
				if err != nil {
					return err
				}
				ps.PolicyAssignmentOwner = &policyAssignmentOwner
			}
		case "policyAssignmentParameters":
			if v != nil {
				var policyAssignmentParameters string
				err = json.Unmarshal(*v, &policyAssignmentParameters)
				if err != nil {
					return err
				}
				ps.PolicyAssignmentParameters = &policyAssignmentParameters
			}
		case "policyAssignmentScope":
			if v != nil {
				var policyAssignmentScope string
				err = json.Unmarshal(*v, &policyAssignmentScope)
				if err != nil {
					return err
				}
				ps.PolicyAssignmentScope = &policyAssignmentScope
			}
		case "policyDefinitionName":
			if v != nil {
				var policyDefinitionName string
				err = json.Unmarshal(*v, &policyDefinitionName)
				if err != nil {
					return err
				}
				ps.PolicyDefinitionName = &policyDefinitionName
			}
		case "policyDefinitionAction":
			if v != nil {
				var policyDefinitionAction string
				err = json.Unmarshal(*v, &policyDefinitionAction)
				if err != nil {
					return err
				}
				ps.PolicyDefinitionAction = &policyDefinitionAction
			}
		case "policyDefinitionCategory":
			if v != nil {
				var policyDefinitionCategory string
				err = json.Unmarshal(*v, &policyDefinitionCategory)
				if err != nil {
					return err
				}
				ps.PolicyDefinitionCategory = &policyDefinitionCategory
			}
		case "policySetDefinitionId":
			if v != nil {
				var policySetDefinitionID string
				err = json.Unmarshal(*v, &policySetDefinitionID)
				if err != nil {
					return err
				}
				ps.PolicySetDefinitionID = &policySetDefinitionID
			}
		case "policySetDefinitionName":
			if v != nil {
				var policySetDefinitionName string
				err = json.Unmarshal(*v, &policySetDefinitionName)
				if err != nil {
					return err
				}
				ps.PolicySetDefinitionName = &policySetDefinitionName
			}
		case "policySetDefinitionOwner":
			if v != nil {
				var policySetDefinitionOwner string
				err = json.Unmarshal(*v, &policySetDefinitionOwner)
				if err != nil {
					return err
				}
				ps.PolicySetDefinitionOwner = &policySetDefinitionOwner
			}
		case "policySetDefinitionCategory":
			if v != nil {
				var policySetDefinitionCategory string
				err = json.Unmarshal(*v, &policySetDefinitionCategory)
				if err != nil {
					return err
				}
				ps.PolicySetDefinitionCategory = &policySetDefinitionCategory
			}
		case "policySetDefinitionParameters":
			if v != nil {
				var policySetDefinitionParameters string
				err = json.Unmarshal(*v, &policySetDefinitionParameters)
				if err != nil {
					return err
				}
				ps.PolicySetDefinitionParameters = &policySetDefinitionParameters
			}
		case "managementGroupIds":
			if v != nil {
				var managementGroupIds string
				err = json.Unmarshal(*v, &managementGroupIds)
				if err != nil {
					return err
				}
				ps.ManagementGroupIds = &managementGroupIds
			}
		case "policyDefinitionReferenceId":
			if v != nil {
				var policyDefinitionReferenceID string
				err = json.Unmarshal(*v, &policyDefinitionReferenceID)
				if err != nil {
					return err
				}
				ps.PolicyDefinitionReferenceID = &policyDefinitionReferenceID
			}
		}
	}

	return nil
}

// PolicyStatesQueryResults query results.
type PolicyStatesQueryResults struct {
	autorest.Response `json:"-"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// OdataCount - OData entity count; represents the number of policy state records returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - Query results.
	Value *[]PolicyState `json:"value,omitempty"`
}

// QueryFailure error response.
type QueryFailure struct {
	// Error - Error definition.
	Error *QueryFailureError `json:"error,omitempty"`
}

// QueryFailureError error definition.
type QueryFailureError struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
}

// String ...
type String struct {
	autorest.Response `json:"-"`
	Value             *string `json:"value,omitempty"`
}

// SummarizeResults summarize action results.
type SummarizeResults struct {
	autorest.Response `json:"-"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// OdataCount - OData entity count; represents the number of summaries returned; always set to 1.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - Summarize action results.
	Value *[]Summary `json:"value,omitempty"`
}

// Summary summary results.
type Summary struct {
	// OdataID - OData entity ID; always set to null since summaries do not have an entity ID.
	OdataID *string `json:"@odata.id,omitempty"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// Results - Non-compliance summary for all policy assignments.
	Results *SummaryResults `json:"results,omitempty"`
	// PolicyAssignments - Policy assignments summary.
	PolicyAssignments *[]PolicyAssignmentSummary `json:"policyAssignments,omitempty"`
}

// SummaryResults non-compliance summary on a particular summary level.
type SummaryResults struct {
	// QueryResultsURI - HTTP POST URI for queryResults action on Microsoft.PolicyInsights to retrieve raw results for the non-compliance summary.
	QueryResultsURI *string `json:"queryResultsUri,omitempty"`
	// NonCompliantResources - Number of non-compliant resources.
	NonCompliantResources *int32 `json:"nonCompliantResources,omitempty"`
	// NonCompliantPolicies - Number of non-compliant policies.
	NonCompliantPolicies *int32 `json:"nonCompliantPolicies,omitempty"`
}
