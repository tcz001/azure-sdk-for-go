package prediction

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/tcz001/azure-sdk-for-go/services/cognitiveservices/v1.1/customvision/prediction"

// BoundingBox ...
type BoundingBox struct {
	Left   *float64 `json:"left,omitempty"`
	Top    *float64 `json:"top,omitempty"`
	Width  *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
}

// ImagePrediction ...
type ImagePrediction struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY
	Predictions *[]Model `json:"predictions,omitempty"`
}

// ImageURL ...
type ImageURL struct {
	URL *string `json:"url,omitempty"`
}

// Model ...
type Model struct {
	// Probability - READ-ONLY
	Probability *float64 `json:"probability,omitempty"`
	// TagID - READ-ONLY
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// BoundingBox - READ-ONLY
	BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
}
