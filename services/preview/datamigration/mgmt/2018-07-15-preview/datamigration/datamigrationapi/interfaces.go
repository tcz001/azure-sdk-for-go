package datamigrationapi

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
	"github.com/tcz001/azure-sdk-for-go/services/preview/datamigration/mgmt/2018-07-15-preview/datamigration"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceSkusClientAPI contains the set of methods on the ResourceSkusClient type.
type ResourceSkusClientAPI interface {
	ListSkus(ctx context.Context) (result datamigration.ResourceSkusResultPage, err error)
}

var _ ResourceSkusClientAPI = (*datamigration.ResourceSkusClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CheckChildrenNameAvailability(ctx context.Context, groupName string, serviceName string, parameters datamigration.NameAvailabilityRequest) (result datamigration.NameAvailabilityResponse, err error)
	CheckNameAvailability(ctx context.Context, location string, parameters datamigration.NameAvailabilityRequest) (result datamigration.NameAvailabilityResponse, err error)
	CheckStatus(ctx context.Context, groupName string, serviceName string) (result datamigration.ServiceStatusResponse, err error)
	CreateOrUpdate(ctx context.Context, parameters datamigration.Service, groupName string, serviceName string) (result datamigration.ServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, groupName string, serviceName string, deleteRunningTasks *bool) (result datamigration.ServicesDeleteFuture, err error)
	Get(ctx context.Context, groupName string, serviceName string) (result datamigration.Service, err error)
	List(ctx context.Context) (result datamigration.ServiceListPage, err error)
	ListByResourceGroup(ctx context.Context, groupName string) (result datamigration.ServiceListPage, err error)
	ListSkus(ctx context.Context, groupName string, serviceName string) (result datamigration.ServiceSkuListPage, err error)
	Start(ctx context.Context, groupName string, serviceName string) (result datamigration.ServicesStartFuture, err error)
	Stop(ctx context.Context, groupName string, serviceName string) (result datamigration.ServicesStopFuture, err error)
	Update(ctx context.Context, parameters datamigration.Service, groupName string, serviceName string) (result datamigration.ServicesUpdateFuture, err error)
}

var _ ServicesClientAPI = (*datamigration.ServicesClient)(nil)

// TasksClientAPI contains the set of methods on the TasksClient type.
type TasksClientAPI interface {
	Cancel(ctx context.Context, groupName string, serviceName string, projectName string, taskName string) (result datamigration.ProjectTask, err error)
	Command(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters datamigration.BasicCommandProperties) (result datamigration.CommandPropertiesModel, err error)
	CreateOrUpdate(ctx context.Context, parameters datamigration.ProjectTask, groupName string, serviceName string, projectName string, taskName string) (result datamigration.ProjectTask, err error)
	Delete(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, deleteRunningTasks *bool) (result autorest.Response, err error)
	Get(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, expand string) (result datamigration.ProjectTask, err error)
	List(ctx context.Context, groupName string, serviceName string, projectName string, taskType string) (result datamigration.TaskListPage, err error)
	Update(ctx context.Context, parameters datamigration.ProjectTask, groupName string, serviceName string, projectName string, taskName string) (result datamigration.ProjectTask, err error)
}

var _ TasksClientAPI = (*datamigration.TasksClient)(nil)

// ServiceTasksClientAPI contains the set of methods on the ServiceTasksClient type.
type ServiceTasksClientAPI interface {
	Cancel(ctx context.Context, groupName string, serviceName string, taskName string) (result datamigration.ProjectTask, err error)
	CreateOrUpdate(ctx context.Context, parameters datamigration.ProjectTask, groupName string, serviceName string, taskName string) (result datamigration.ProjectTask, err error)
	Delete(ctx context.Context, groupName string, serviceName string, taskName string, deleteRunningTasks *bool) (result autorest.Response, err error)
	Get(ctx context.Context, groupName string, serviceName string, taskName string, expand string) (result datamigration.ProjectTask, err error)
	List(ctx context.Context, groupName string, serviceName string, taskType string) (result datamigration.TaskListPage, err error)
	Update(ctx context.Context, parameters datamigration.ProjectTask, groupName string, serviceName string, taskName string) (result datamigration.ProjectTask, err error)
}

var _ ServiceTasksClientAPI = (*datamigration.ServiceTasksClient)(nil)

// ProjectsClientAPI contains the set of methods on the ProjectsClient type.
type ProjectsClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters datamigration.Project, groupName string, serviceName string, projectName string) (result datamigration.Project, err error)
	Delete(ctx context.Context, groupName string, serviceName string, projectName string, deleteRunningTasks *bool) (result autorest.Response, err error)
	Get(ctx context.Context, groupName string, serviceName string, projectName string) (result datamigration.Project, err error)
	List(ctx context.Context, groupName string, serviceName string) (result datamigration.ProjectListPage, err error)
	Update(ctx context.Context, parameters datamigration.Project, groupName string, serviceName string, projectName string) (result datamigration.Project, err error)
}

var _ ProjectsClientAPI = (*datamigration.ProjectsClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result datamigration.QuotaListPage, err error)
}

var _ UsagesClientAPI = (*datamigration.UsagesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result datamigration.ServiceOperationListPage, err error)
}

var _ OperationsClientAPI = (*datamigration.OperationsClient)(nil)

// FilesClientAPI contains the set of methods on the FilesClient type.
type FilesClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters datamigration.ProjectFile, groupName string, serviceName string, projectName string, fileName string) (result datamigration.ProjectFile, err error)
	Delete(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result autorest.Response, err error)
	Get(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result datamigration.ProjectFile, err error)
	List(ctx context.Context, groupName string, serviceName string, projectName string) (result datamigration.FileListPage, err error)
	Read(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result datamigration.FileStorageInfo, err error)
	ReadWrite(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result datamigration.FileStorageInfo, err error)
	Update(ctx context.Context, parameters datamigration.ProjectFile, groupName string, serviceName string, projectName string, fileName string) (result datamigration.ProjectFile, err error)
}

var _ FilesClientAPI = (*datamigration.FilesClient)(nil)
