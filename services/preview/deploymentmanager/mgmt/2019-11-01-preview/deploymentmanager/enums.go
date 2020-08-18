package deploymentmanager

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

// DeploymentMode enumerates the values for deployment mode.
type DeploymentMode string

const (
	// Complete ...
	Complete DeploymentMode = "Complete"
	// Incremental ...
	Incremental DeploymentMode = "Incremental"
)

// PossibleDeploymentModeValues returns an array of possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{Complete, Incremental}
}

// RestAuthLocation enumerates the values for rest auth location.
type RestAuthLocation string

const (
	// Header ...
	Header RestAuthLocation = "Header"
	// Query ...
	Query RestAuthLocation = "Query"
)

// PossibleRestAuthLocationValues returns an array of possible values for the RestAuthLocation const type.
func PossibleRestAuthLocationValues() []RestAuthLocation {
	return []RestAuthLocation{Header, Query}
}

// RestMatchQuantifier enumerates the values for rest match quantifier.
type RestMatchQuantifier string

const (
	// All ...
	All RestMatchQuantifier = "All"
	// Any ...
	Any RestMatchQuantifier = "Any"
)

// PossibleRestMatchQuantifierValues returns an array of possible values for the RestMatchQuantifier const type.
func PossibleRestMatchQuantifierValues() []RestMatchQuantifier {
	return []RestMatchQuantifier{All, Any}
}

// RestRequestMethod enumerates the values for rest request method.
type RestRequestMethod string

const (
	// GET ...
	GET RestRequestMethod = "GET"
	// POST ...
	POST RestRequestMethod = "POST"
)

// PossibleRestRequestMethodValues returns an array of possible values for the RestRequestMethod const type.
func PossibleRestRequestMethodValues() []RestRequestMethod {
	return []RestRequestMethod{GET, POST}
}

// StepType enumerates the values for step type.
type StepType string

const (
	// StepTypeHealthCheck ...
	StepTypeHealthCheck StepType = "HealthCheck"
	// StepTypeStepProperties ...
	StepTypeStepProperties StepType = "StepProperties"
	// StepTypeWait ...
	StepTypeWait StepType = "Wait"
)

// PossibleStepTypeValues returns an array of possible values for the StepType const type.
func PossibleStepTypeValues() []StepType {
	return []StepType{StepTypeHealthCheck, StepTypeStepProperties, StepTypeWait}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeAuthentication ...
	TypeAuthentication Type = "Authentication"
	// TypeSas ...
	TypeSas Type = "Sas"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeAuthentication, TypeSas}
}

// TypeBasicHealthCheckStepAttributes enumerates the values for type basic health check step attributes.
type TypeBasicHealthCheckStepAttributes string

const (
	// TypeHealthCheckStepAttributes ...
	TypeHealthCheckStepAttributes TypeBasicHealthCheckStepAttributes = "HealthCheckStepAttributes"
	// TypeREST ...
	TypeREST TypeBasicHealthCheckStepAttributes = "REST"
)

// PossibleTypeBasicHealthCheckStepAttributesValues returns an array of possible values for the TypeBasicHealthCheckStepAttributes const type.
func PossibleTypeBasicHealthCheckStepAttributesValues() []TypeBasicHealthCheckStepAttributes {
	return []TypeBasicHealthCheckStepAttributes{TypeHealthCheckStepAttributes, TypeREST}
}

// TypeBasicRestRequestAuthentication enumerates the values for type basic rest request authentication.
type TypeBasicRestRequestAuthentication string

const (
	// TypeAPIKey ...
	TypeAPIKey TypeBasicRestRequestAuthentication = "ApiKey"
	// TypeRestRequestAuthentication ...
	TypeRestRequestAuthentication TypeBasicRestRequestAuthentication = "RestRequestAuthentication"
	// TypeRolloutIdentity ...
	TypeRolloutIdentity TypeBasicRestRequestAuthentication = "RolloutIdentity"
)

// PossibleTypeBasicRestRequestAuthenticationValues returns an array of possible values for the TypeBasicRestRequestAuthentication const type.
func PossibleTypeBasicRestRequestAuthenticationValues() []TypeBasicRestRequestAuthentication {
	return []TypeBasicRestRequestAuthentication{TypeAPIKey, TypeRestRequestAuthentication, TypeRolloutIdentity}
}