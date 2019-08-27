package aadapi

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
	"github.com/tcz001/azure-sdk-for-go/services/aad/mgmt/2017-04-01/aad"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result aad.OperationsDiscoveryCollection, err error)
}

var _ OperationsClientAPI = (*aad.OperationsClient)(nil)

// DiagnosticSettingsClientAPI contains the set of methods on the DiagnosticSettingsClient type.
type DiagnosticSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters aad.DiagnosticSettingsResource, name string) (result aad.DiagnosticSettingsResource, err error)
	Delete(ctx context.Context, name string) (result autorest.Response, err error)
	Get(ctx context.Context, name string) (result aad.DiagnosticSettingsResource, err error)
	List(ctx context.Context) (result aad.DiagnosticSettingsResourceCollection, err error)
}

var _ DiagnosticSettingsClientAPI = (*aad.DiagnosticSettingsClient)(nil)

// DiagnosticSettingsCategoryClientAPI contains the set of methods on the DiagnosticSettingsCategoryClient type.
type DiagnosticSettingsCategoryClientAPI interface {
	List(ctx context.Context) (result aad.DiagnosticSettingsCategoryResourceCollection, err error)
}

var _ DiagnosticSettingsCategoryClientAPI = (*aad.DiagnosticSettingsCategoryClient)(nil)
