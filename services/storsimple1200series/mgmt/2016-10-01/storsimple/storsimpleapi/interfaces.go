package storsimpleapi

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
	"github.com/tcz001/azure-sdk-for-go/services/storsimple1200series/mgmt/2016-10-01/storsimple"
	"github.com/Azure/go-autorest/autorest"
)

// ManagersClientAPI contains the set of methods on the ManagersClient type.
type ManagersClientAPI interface {
	CreateExtendedInfo(ctx context.Context, managerExtendedInfo storsimple.ManagerExtendedInfo, resourceGroupName string, managerName string) (result storsimple.ManagerExtendedInfo, err error)
	CreateOrUpdate(ctx context.Context, manager storsimple.Manager, resourceGroupName string, managerName string) (result storsimple.Manager, err error)
	Delete(ctx context.Context, resourceGroupName string, managerName string) (result autorest.Response, err error)
	DeleteExtendedInfo(ctx context.Context, resourceGroupName string, managerName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.Manager, err error)
	GetEncryptionKey(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.SymmetricEncryptedSecret, err error)
	GetEncryptionSettings(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.EncryptionSettings, err error)
	GetExtendedInfo(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.ManagerExtendedInfo, err error)
	List(ctx context.Context) (result storsimple.ManagerList, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storsimple.ManagerList, err error)
	ListMetricDefinition(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.MetricDefinitionList, err error)
	ListMetrics(ctx context.Context, resourceGroupName string, managerName string, filter string) (result storsimple.MetricList, err error)
	Update(ctx context.Context, parameters storsimple.ManagerPatch, resourceGroupName string, managerName string) (result storsimple.Manager, err error)
	UpdateExtendedInfo(ctx context.Context, managerExtendedInfo storsimple.ManagerExtendedInfo, resourceGroupName string, managerName string, ifMatch string) (result storsimple.ManagerExtendedInfo, err error)
	UploadRegistrationCertificate(ctx context.Context, certificateName string, uploadCertificateRequestrequest storsimple.UploadCertificateRequest, resourceGroupName string, managerName string) (result storsimple.UploadCertificateResponse, err error)
}

var _ ManagersClientAPI = (*storsimple.ManagersClient)(nil)

// AvailableProviderOperationsClientAPI contains the set of methods on the AvailableProviderOperationsClient type.
type AvailableProviderOperationsClientAPI interface {
	List(ctx context.Context) (result storsimple.AvailableProviderOperationsPage, err error)
}

var _ AvailableProviderOperationsClientAPI = (*storsimple.AvailableProviderOperationsClient)(nil)

// AccessControlRecordsClientAPI contains the set of methods on the AccessControlRecordsClient type.
type AccessControlRecordsClientAPI interface {
	CreateOrUpdate(ctx context.Context, accessControlRecordName string, accessControlRecord storsimple.AccessControlRecord, resourceGroupName string, managerName string) (result storsimple.AccessControlRecordsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string) (result storsimple.AccessControlRecordsDeleteFuture, err error)
	Get(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string) (result storsimple.AccessControlRecord, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.AccessControlRecordList, err error)
}

var _ AccessControlRecordsClientAPI = (*storsimple.AccessControlRecordsClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	Clear(ctx context.Context, request storsimple.ClearAlertRequest, resourceGroupName string, managerName string) (result autorest.Response, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string, filter string) (result storsimple.AlertListPage, err error)
	SendTestEmail(ctx context.Context, deviceName string, request storsimple.SendTestAlertEmailRequest, resourceGroupName string, managerName string) (result autorest.Response, err error)
}

var _ AlertsClientAPI = (*storsimple.AlertsClient)(nil)

// BackupsClientAPI contains the set of methods on the BackupsClient type.
type BackupsClientAPI interface {
	Clone(ctx context.Context, deviceName string, backupName string, elementName string, cloneRequest storsimple.CloneRequest, resourceGroupName string, managerName string) (result storsimple.BackupsCloneFuture, err error)
	Delete(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string) (result storsimple.BackupsDeleteFuture, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string, forFailover *bool, filter string) (result storsimple.BackupListPage, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string, filter string) (result storsimple.BackupListPage, err error)
}

var _ BackupsClientAPI = (*storsimple.BackupsClient)(nil)

// DevicesClientAPI contains the set of methods on the DevicesClient type.
type DevicesClientAPI interface {
	CreateOrUpdateAlertSettings(ctx context.Context, deviceName string, alertSettings storsimple.AlertSettings, resourceGroupName string, managerName string) (result storsimple.DevicesCreateOrUpdateAlertSettingsFuture, err error)
	CreateOrUpdateSecuritySettings(ctx context.Context, deviceName string, securitySettings storsimple.SecuritySettings, resourceGroupName string, managerName string) (result storsimple.DevicesCreateOrUpdateSecuritySettingsFuture, err error)
	Deactivate(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.DevicesDeactivateFuture, err error)
	Delete(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.DevicesDeleteFuture, err error)
	DownloadUpdates(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.DevicesDownloadUpdatesFuture, err error)
	Failover(ctx context.Context, deviceName string, failoverRequest storsimple.FailoverRequest, resourceGroupName string, managerName string) (result storsimple.DevicesFailoverFuture, err error)
	Get(ctx context.Context, deviceName string, resourceGroupName string, managerName string, expand string) (result storsimple.Device, err error)
	GetAlertSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.AlertSettings, err error)
	GetNetworkSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.NetworkSettings, err error)
	GetTimeSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.TimeSettings, err error)
	GetUpdateSummary(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.Updates, err error)
	InstallUpdates(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.DevicesInstallUpdatesFuture, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string, expand string) (result storsimple.DeviceList, err error)
	ListFailoverTarget(ctx context.Context, deviceName string, resourceGroupName string, managerName string, expand string) (result storsimple.DeviceList, err error)
	ListMetricDefinition(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.MetricDefinitionList, err error)
	ListMetrics(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (result storsimple.MetricList, err error)
	Patch(ctx context.Context, deviceName string, devicePatch storsimple.DevicePatch, resourceGroupName string, managerName string) (result storsimple.DevicesPatchFuture, err error)
	ScanForUpdates(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.DevicesScanForUpdatesFuture, err error)
}

var _ DevicesClientAPI = (*storsimple.DevicesClient)(nil)

// BackupScheduleGroupsClientAPI contains the set of methods on the BackupScheduleGroupsClient type.
type BackupScheduleGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, scheduleGroupName string, scheduleGroup storsimple.BackupScheduleGroup, resourceGroupName string, managerName string) (result storsimple.BackupScheduleGroupsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, scheduleGroupName string, resourceGroupName string, managerName string) (result storsimple.BackupScheduleGroupsDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, scheduleGroupName string, resourceGroupName string, managerName string) (result storsimple.BackupScheduleGroup, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.BackupScheduleGroupList, err error)
}

var _ BackupScheduleGroupsClientAPI = (*storsimple.BackupScheduleGroupsClient)(nil)

// ChapSettingsClientAPI contains the set of methods on the ChapSettingsClient type.
type ChapSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, chapUserName string, chapSetting storsimple.ChapSettings, resourceGroupName string, managerName string) (result storsimple.ChapSettingsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string) (result storsimple.ChapSettingsDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, chapUserName string, resourceGroupName string, managerName string) (result storsimple.ChapSettings, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.ChapSettingsList, err error)
}

var _ ChapSettingsClientAPI = (*storsimple.ChapSettingsClient)(nil)

// IscsiDisksClientAPI contains the set of methods on the IscsiDisksClient type.
type IscsiDisksClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, iscsiServerName string, diskName string, iscsiDisk storsimple.ISCSIDisk, resourceGroupName string, managerName string) (result storsimple.IscsiDisksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string) (result storsimple.IscsiDisksDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string) (result storsimple.ISCSIDisk, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.ISCSIDiskList, err error)
	ListByIscsiServer(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string) (result storsimple.ISCSIDiskList, err error)
	ListMetricDefinition(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string) (result storsimple.MetricDefinitionList, err error)
	ListMetrics(ctx context.Context, deviceName string, iscsiServerName string, diskName string, resourceGroupName string, managerName string, filter string) (result storsimple.MetricList, err error)
}

var _ IscsiDisksClientAPI = (*storsimple.IscsiDisksClient)(nil)

// FileServersClientAPI contains the set of methods on the FileServersClient type.
type FileServersClientAPI interface {
	BackupNow(ctx context.Context, deviceName string, fileServerName string, resourceGroupName string, managerName string) (result storsimple.FileServersBackupNowFuture, err error)
	CreateOrUpdate(ctx context.Context, deviceName string, fileServerName string, fileServer storsimple.FileServer, resourceGroupName string, managerName string) (result storsimple.FileServersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, fileServerName string, resourceGroupName string, managerName string) (result storsimple.FileServersDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, fileServerName string, resourceGroupName string, managerName string) (result storsimple.FileServer, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.FileServerList, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.FileServerList, err error)
	ListMetricDefinition(ctx context.Context, deviceName string, fileServerName string, resourceGroupName string, managerName string) (result storsimple.MetricDefinitionList, err error)
	ListMetrics(ctx context.Context, deviceName string, fileServerName string, resourceGroupName string, managerName string, filter string) (result storsimple.MetricList, err error)
}

var _ FileServersClientAPI = (*storsimple.FileServersClient)(nil)

// FileSharesClientAPI contains the set of methods on the FileSharesClient type.
type FileSharesClientAPI interface {
	CreateOrUpdate(ctx context.Context, deviceName string, fileServerName string, shareName string, fileShare storsimple.FileShare, resourceGroupName string, managerName string) (result storsimple.FileSharesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string) (result storsimple.FileSharesDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string) (result storsimple.FileShare, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.FileShareList, err error)
	ListByFileServer(ctx context.Context, deviceName string, fileServerName string, resourceGroupName string, managerName string) (result storsimple.FileShareList, err error)
	ListMetricDefinition(ctx context.Context, deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string) (result storsimple.MetricDefinitionList, err error)
	ListMetrics(ctx context.Context, deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string, filter string) (result storsimple.MetricList, err error)
}

var _ FileSharesClientAPI = (*storsimple.FileSharesClient)(nil)

// IscsiServersClientAPI contains the set of methods on the IscsiServersClient type.
type IscsiServersClientAPI interface {
	BackupNow(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string) (result storsimple.IscsiServersBackupNowFuture, err error)
	CreateOrUpdate(ctx context.Context, deviceName string, iscsiServerName string, iscsiServer storsimple.ISCSIServer, resourceGroupName string, managerName string) (result storsimple.IscsiServersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string) (result storsimple.IscsiServersDeleteFuture, err error)
	Get(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string) (result storsimple.ISCSIServer, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result storsimple.ISCSIServerList, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.ISCSIServerList, err error)
	ListMetricDefinition(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string) (result storsimple.MetricDefinitionList, err error)
	ListMetrics(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, filter string) (result storsimple.MetricList, err error)
}

var _ IscsiServersClientAPI = (*storsimple.IscsiServersClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Get(ctx context.Context, deviceName string, jobName string, resourceGroupName string, managerName string) (result storsimple.Job, err error)
	ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string, filter string) (result storsimple.JobListPage, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string, filter string) (result storsimple.JobListPage, err error)
}

var _ JobsClientAPI = (*storsimple.JobsClient)(nil)

// StorageAccountCredentialsClientAPI contains the set of methods on the StorageAccountCredentialsClient type.
type StorageAccountCredentialsClientAPI interface {
	CreateOrUpdate(ctx context.Context, credentialName string, storageAccount storsimple.StorageAccountCredential, resourceGroupName string, managerName string) (result storsimple.StorageAccountCredentialsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, credentialName string, resourceGroupName string, managerName string) (result storsimple.StorageAccountCredentialsDeleteFuture, err error)
	Get(ctx context.Context, credentialName string, resourceGroupName string, managerName string) (result storsimple.StorageAccountCredential, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.StorageAccountCredentialList, err error)
}

var _ StorageAccountCredentialsClientAPI = (*storsimple.StorageAccountCredentialsClient)(nil)

// StorageDomainsClientAPI contains the set of methods on the StorageDomainsClient type.
type StorageDomainsClientAPI interface {
	CreateOrUpdate(ctx context.Context, storageDomainName string, storageDomain storsimple.StorageDomain, resourceGroupName string, managerName string) (result storsimple.StorageDomainsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, storageDomainName string, resourceGroupName string, managerName string) (result storsimple.StorageDomainsDeleteFuture, err error)
	Get(ctx context.Context, storageDomainName string, resourceGroupName string, managerName string) (result storsimple.StorageDomain, err error)
	ListByManager(ctx context.Context, resourceGroupName string, managerName string) (result storsimple.StorageDomainList, err error)
}

var _ StorageDomainsClientAPI = (*storsimple.StorageDomainsClient)(nil)
