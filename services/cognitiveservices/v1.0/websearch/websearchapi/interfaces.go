package websearchapi

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
	"github.com/tcz001/azure-sdk-for-go/services/cognitiveservices/v1.0/websearch"
)

// WebClientAPI contains the set of methods on the WebClient type.
type WebClientAPI interface {
	Search(ctx context.Context, query string, acceptLanguage string, pragma string, userAgent string, clientID string, clientIP string, location string, answerCount *int32, countryCode string, count *int32, freshness websearch.Freshness, market string, offset *int32, promote []websearch.AnswerType, responseFilter []websearch.AnswerType, safeSearch websearch.SafeSearch, setLang string, textDecorations *bool, textFormat websearch.TextFormat) (result websearch.SearchResponse, err error)
}

var _ WebClientAPI = (*websearch.WebClient)(nil)
